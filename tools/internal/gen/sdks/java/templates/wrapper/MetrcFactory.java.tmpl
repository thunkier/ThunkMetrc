package io.github.thunkier.thunkmetrc.wrapper;

import io.github.thunkier.thunkmetrc.client.MetrcClient;

/**
 * Factory for creating MetrcWrapper instances with a shared rate limiter.
 * Use this when you need to create multiple wrappers that share the same rate limiting pool.
 */
public class MetrcFactory {
    private final MetrcRateLimiter sharedLimiter;

    /**
     * Create a factory with default max concurrent requests (150).
     */
    public MetrcFactory() {
        this(150);
    }

    /**
     * Create a factory with custom max concurrent requests.
     * @param maxConcurrentRequests Maximum concurrent GET requests for the integrator
     */
    public MetrcFactory(int maxConcurrentRequests) {
        MetrcRateLimiter.Config config = new MetrcRateLimiter.Config();
        config.maxConcurrentGetIntegrator = maxConcurrentRequests;
        this.sharedLimiter = new MetrcRateLimiter(config);
    }

    /**
     * Create a MetrcWrapper with a shared rate limiter.
     * @param baseUrl Metrc API base URL
     * @param vendorKey Vendor API key
     * @param userKey User API key
     * @return A MetrcWrapper sharing this factory's rate limiter
     */
    public MetrcWrapper create(String baseUrl, String vendorKey, String userKey) {
        MetrcClient client = new MetrcClient(baseUrl, vendorKey, userKey);
        return new MetrcWrapper(client, this.sharedLimiter);
    }

    /**
     * Get the shared rate limiter for this factory.
     * @return The shared MetrcRateLimiter instance
     */
    public MetrcRateLimiter getRateLimiter() {
        return sharedLimiter;
    }
}
