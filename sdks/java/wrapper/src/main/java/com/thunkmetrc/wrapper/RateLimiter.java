package com.thunkmetrc.wrapper;

import java.util.concurrent.Semaphore;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.TimeUnit;
import java.util.Map;
import java.util.concurrent.Callable;

public class RateLimiter {
    public static class Config {
        public boolean enabled = false;
        public double maxGetPerSecondPerFacility = 50.0;
        public double maxGetPerSecondIntegrator = 150.0;
        public int maxConcurrentGetPerFacility = 10;
        public int maxConcurrentGetIntegrator = 30;
    }

    private final Config config;
    private final Semaphore integratorSem;
    private final TokenBucket integratorRate;
    private final Map<String, Semaphore> facilitySems = new ConcurrentHashMap<>();
    private final Map<String, TokenBucket> facilityRates = new ConcurrentHashMap<>();

    public RateLimiter(Config config) {
        if (config == null) config = new Config();
        this.config = config;
        this.integratorSem = new Semaphore(config.maxConcurrentGetIntegrator);
        this.integratorRate = new TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator);
    }

    public <T> T execute(String facility, boolean isGet, Callable<T> op) throws Exception {
        if (!config.enabled || !isGet) {
            return op.call();
        }

        integratorSem.acquire();
        try {
            Semaphore facSem = null;
            if (facility != null) {
                facSem = facilitySems.computeIfAbsent(facility, k -> new Semaphore(config.maxConcurrentGetPerFacility));
                facSem.acquire();
            }

            try {
                integratorRate.acquire(1);
                if (facility != null) {
                    TokenBucket facRate = facilityRates.computeIfAbsent(facility, k -> new TokenBucket(config.maxGetPerSecondPerFacility, config.maxGetPerSecondPerFacility));
                    facRate.acquire(1);
                }

                while (true) {
                    try {
                        return op.call();
                    } catch (Exception e) {
                        if (e.getMessage() != null && e.getMessage().contains("429")) {
                            Thread.sleep(1000);
                            continue;
                        }
                        throw e;
                    }
                }
            } finally {
                if (facSem != null) facSem.release();
            }
        } finally {
            integratorSem.release();
        }
    }

    static class TokenBucket {
        private final double rate;
        private final double capacity;
        private double tokens;
        private long lastRefill;

        public TokenBucket(double rate, double capacity) {
            this.rate = rate;
            this.capacity = capacity;
            this.tokens = capacity;
            this.lastRefill = System.nanoTime();
        }

        public synchronized void acquire(double amount) throws InterruptedException {
            while (true) {
                refill();
                if (tokens >= amount) {
                    tokens -= amount;
                    return;
                }
                double missing = amount - tokens;
                long waitNs = (long) (missing / rate * 1_000_000_000.0);
                Thread.sleep(waitNs / 1_000_000, (int) (waitNs % 1_000_000));
            }
        }

        private void refill() {
            long now = System.nanoTime();
            double elapsed = (now - lastRefill) / 1_000_000_000.0;
            tokens = Math.min(capacity, tokens + elapsed * rate);
            lastRefill = now;
        }
    }
}
