package io.github.thunkier.thunkmetrc.wrapper;

import java.util.Map;
import java.util.Random;
import java.util.concurrent.Callable;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.Semaphore;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Rate limiter for Metrc API requests.
 * Implements token bucket algorithm with per-facility and integrator-level limits.
 * Includes exponential backoff with jitter for 429 and 5xx error handling.
 */
public class MetrcRateLimiter {

    /**
     * Configuration for rate limiting behavior.
     */
    public static class Config {
        /** Enable or disable rate limiting. Default: false */
        public boolean enabled = false;
        /** Maximum GET requests per second per facility. Default: 50 */
        public double maxGetPerSecondPerFacility = 50.0;
        /** Maximum GET requests per second for the integrator. Default: 150 */
        public double maxGetPerSecondIntegrator = 150.0;
        /** Maximum concurrent GET requests per facility. Default: 10 */
        public int maxConcurrentGetPerFacility = 10;
        /** Maximum concurrent GET requests for the integrator. Default: 30 */
        public int maxConcurrentGetIntegrator = 30;
        /** Maximum retry attempts for rate-limited requests. Default: 5 */
        public int maxRetries = 5;
        /** Initial backoff in milliseconds. Default: 300 */
        public long initialBackoffMs = 300;
        /** Maximum backoff in milliseconds. Default: 20000 */
        public long maxBackoffMs = 20000;
        /** Backoff multiplier factor. Default: 2.0 */
        public double backoffFactor = 2.0;
    }

    private static final Pattern RETRY_AFTER_PATTERN = Pattern.compile("(?i)retry-after[:\\s=]+(\\d+)");

    private final Config config;
    private final Semaphore integratorSem;
    private final TokenBucket integratorRate;
    private final Map<String, Semaphore> facilitySems = new ConcurrentHashMap<>();
    private final Map<String, TokenBucket> facilityRates = new ConcurrentHashMap<>();
    private final Random random = new Random();

    public MetrcRateLimiter() {
        this(new Config());
    }

    public MetrcRateLimiter(Config config) {
        if (config == null) config = new Config();
        this.config = config;
        this.integratorSem = new Semaphore(config.maxConcurrentGetIntegrator);
        this.integratorRate = new TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator);
    }

    /**
     * Execute an operation with rate limiting.
     * @param facility The facility license number (for per-facility limits)
     * @param isGet Whether this is a GET request (rate limiting only applies to GETs)
     * @param op The operation to execute
     * @param <T> Return type
     * @return The result of the operation
     * @throws Exception If the operation fails after all retries
     */
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

                int attempt = 0;
                while (true) {
                    try {
                        return op.call();
                    } catch (Exception e) {
                        attempt++;
                        if (attempt > config.maxRetries) throw e;

                        String msg = e.getMessage() != null ? e.getMessage() : "";
                        boolean is429 = msg.contains("429") || msg.toLowerCase().contains("too many requests");
                        boolean is5xx = msg.contains("500") || msg.contains("502") || msg.contains("503") || msg.contains("504");

                        if (is429 || is5xx) {
                            long waitMs = 0;

                            Matcher matcher = RETRY_AFTER_PATTERN.matcher(msg);
                            if (matcher.find()) {
                                try {
                                    waitMs = Long.parseLong(matcher.group(1)) * 1000;
                                } catch (NumberFormatException ignored) {}
                            }

                            if (waitMs == 0) {
                                double backoff = config.initialBackoffMs * Math.pow(config.backoffFactor, attempt - 1);
                                double capped = Math.min(backoff, config.maxBackoffMs);
                                double jitter = capped * (0.8 + 0.4 * random.nextDouble());
                                waitMs = (long) jitter;
                            }

                            try {
                                Thread.sleep(waitMs);
                            } catch (InterruptedException ie) {
                                Thread.currentThread().interrupt();
                                throw e;
                            }
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

    private static class TokenBucket {
        private final double rate;
        private final double capacity;
        private double tokens;
        private long lastRefill;

        TokenBucket(double rate, double capacity) {
            this.rate = rate;
            this.capacity = capacity;
            this.tokens = capacity;
            this.lastRefill = System.nanoTime();
        }

        synchronized void acquire(double amount) throws InterruptedException {
            while (true) {
                refill();
                if (tokens >= amount) {
                    tokens -= amount;
                    return;
                }
                double missing = amount - tokens;
                long waitMs = (long) (missing / rate * 1000);
                if (waitMs > 0) {
                    Thread.sleep(waitMs);
                }
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
