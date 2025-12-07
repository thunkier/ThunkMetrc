package com.thunkmetrc.wrapper

import kotlinx.coroutines.sync.Semaphore
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import kotlinx.coroutines.delay
import java.util.concurrent.ConcurrentHashMap
import java.time.Instant

data class RateLimiterConfig(
    val enabled: Boolean = false,
    val maxGetPerSecondPerFacility: Double = 50.0,
    val maxGetPerSecondIntegrator: Double = 150.0,
    val maxConcurrentGetPerFacility: Int = 10,
    val maxConcurrentGetIntegrator: Int = 30
)

class MetrcRateLimiter(private val config: RateLimiterConfig = RateLimiterConfig()) {
    private val integratorRate = TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator)
    private val integratorSem = Semaphore(config.maxConcurrentGetIntegrator)
    
    private val facilityRates = ConcurrentHashMap<String, TokenBucket>()
    private val facilitySems = ConcurrentHashMap<String, Semaphore>()
    private val semMutex = Mutex()

    suspend fun <T> execute(facility: String?, isGet: Boolean, op: suspend () -> T): T {
        if (!config.enabled || !isGet) {
            return op()
        }

        integratorSem.acquire()
        try {
            var facSem: Semaphore? = null
            if (facility != null) {
                // Double-checked locking pattern equivalent for suspension-safe map access if needed
                // But ConcurrentHashMap is thread-safe. ComputeIfAbsent is atomic.
                // However, creating Semaphore is cheap.
                facSem = facilitySems.computeIfAbsent(facility) { Semaphore(config.maxConcurrentGetPerFacility) }
                facSem.acquire()
            }

            try {
                integratorRate.acquire(1.0)
                if (facility != null) {
                    val facRate = facilityRates.computeIfAbsent(facility) { TokenBucket(config.maxGetPerSecondPerFacility, config.maxGetPerSecondPerFacility) }
                    facRate.acquire(1.0)
                }

                while (true) {
                    try {
                        return op()
                    } catch (e: Exception) {
                        if (e.message?.contains("429") == true) {
                            delay(1000)
                            continue
                        }
                        throw e
                    }
                }
            } finally {
                facSem?.release()
            }
        } finally {
            integratorSem.release()
        }
    }

    class TokenBucket(private val rate: Double, private val capacity: Double) {
        private var tokens = capacity
        private var lastRefill = System.nanoTime()
        private val mutex = Mutex()

        suspend fun acquire(amount: Double) {
            while (true) {
                val waitMs = mutex.withLock {
                    refill()
                    if (tokens >= amount) {
                        tokens -= amount
                        return
                    }
                    val missing = amount - tokens
                    (missing / rate * 1000).toLong()
                }
                if (waitMs > 0) delay(waitMs)
            }
        }

        private fun refill() {
            val now = System.nanoTime()
            val elapsed = (now - lastRefill) / 1_000_000_000.0
            tokens = Math.min(capacity, tokens + elapsed * rate)
            lastRefill = now
        }
    }
}
