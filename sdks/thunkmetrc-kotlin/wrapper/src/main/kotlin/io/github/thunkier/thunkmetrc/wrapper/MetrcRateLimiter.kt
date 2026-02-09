package io.github.thunkier.thunkmetrc.wrapper

import kotlinx.coroutines.sync.Semaphore
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import kotlinx.coroutines.delay
import java.util.concurrent.ConcurrentHashMap

/**
 * Configuration for rate limiting behavior.
 * @property enabled Enable or disable rate limiting. Default: false
 * @property maxGetPerSecondPerFacility Maximum GET requests per second per facility. Default: 50
 * @property maxGetPerSecondIntegrator Maximum GET requests per second for the integrator. Default: 150
 * @property maxConcurrentGetPerFacility Maximum concurrent GET requests per facility. Default: 10
 * @property maxConcurrentGetIntegrator Maximum concurrent GET requests for the integrator. Default: 30
 * @property maxRetries Maximum retry attempts for rate-limited requests. Default: 5
 * @property initialBackoffMs Initial backoff in milliseconds. Default: 300
 * @property maxBackoffMs Maximum backoff in milliseconds. Default: 20000
 * @property backoffFactor Backoff multiplier factor. Default: 2.0
 */
data class RateLimiterConfig(
    val enabled: Boolean = false,
    val maxGetPerSecondPerFacility: Double = 50.0,
    val maxGetPerSecondIntegrator: Double = 150.0,
    val maxConcurrentGetPerFacility: Int = 10,
    val maxConcurrentGetIntegrator: Int = 30,
    val maxRetries: Int = 5,
    val initialBackoffMs: Long = 300,
    val maxBackoffMs: Long = 20000,
    val backoffFactor: Double = 2.0
)

/**
 * Rate limiter for Metrc API requests.
 * Implements token bucket algorithm with per-facility and integrator-level limits.
 * Includes exponential backoff with jitter for 429 and 5xx error handling.
 */
class MetrcRateLimiter(private val config: RateLimiterConfig = RateLimiterConfig()) {
    private val integratorRate = TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator)
    private val integratorSem = Semaphore(config.maxConcurrentGetIntegrator)

    private val facilityRates = ConcurrentHashMap<String, TokenBucket>()
    private val facilitySems = ConcurrentHashMap<String, Semaphore>()

    private val retryAfterRegex = Regex("(?i)retry-after[:\\s=]+(\\d+)")

    /**
     * Execute an operation with rate limiting.
     * @param facility The facility license number (for per-facility limits)
     * @param isGet Whether this is a GET request (rate limiting only applies to GETs)
     * @param op The suspend operation to execute
     * @return The result of the operation
     * @throws Exception If the operation fails after all retries
     */
    suspend fun <T> execute(facility: String?, isGet: Boolean, op: suspend () -> T): T {
        if (!config.enabled || !isGet) {
            return op()
        }

        integratorSem.acquire()
        try {
            var facSem: Semaphore? = null
            if (facility != null) {
                facSem = facilitySems.computeIfAbsent(facility) { Semaphore(config.maxConcurrentGetPerFacility) }
                facSem.acquire()
            }

            try {
                integratorRate.acquire(1.0)
                if (facility != null) {
                    val facRate = facilityRates.computeIfAbsent(facility) { TokenBucket(config.maxGetPerSecondPerFacility, config.maxGetPerSecondPerFacility) }
                    facRate.acquire(1.0)
                }

                var attempt = 0
                while (true) {
                    try {
                        return op()
                    } catch (e: Exception) {
                        attempt++
                        if (attempt > config.maxRetries) throw e

                        val msg = e.message ?: ""
                        val is429 = msg.contains("429") || msg.contains("Too Many Requests", ignoreCase = true)
                        val is5xx = msg.contains("500") || msg.contains("502") || msg.contains("503") || msg.contains("504")

                        if (is429 || is5xx) {
                            var waitMs = 0L

                            val match = retryAfterRegex.find(msg)
                            if (match != null) {
                                val seconds = match.groupValues[1].toLongOrNull()
                                if (seconds != null) {
                                    waitMs = seconds * 1000
                                }
                            }

                            if (waitMs == 0L) {
                                val backoff = config.initialBackoffMs * Math.pow(config.backoffFactor, (attempt - 1).toDouble())
                                val capped = Math.min(backoff, config.maxBackoffMs.toDouble())
                                val jitter = capped * (0.8 + 0.4 * Math.random())
                                waitMs = jitter.toLong()
                            }

                            delay(waitMs)
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

    private class TokenBucket(private val rate: Double, private val capacity: Double) {
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
