package io.github.thunkier.thunkmetrc.wrapper

import io.github.thunkier.thunkmetrc.client.MetrcClient

/**
 * Factory for creating MetrcWrapper instances with a shared rate limiter.
 * Use this when you need to create multiple wrappers that share the same rate limiting pool.
 *
 * @param maxConcurrentRequests Maximum concurrent GET requests for the integrator. Default: 150
 */
class MetrcFactory(maxConcurrentRequests: Int = 150) {
    private val sharedLimiter: MetrcRateLimiter

    init {
        val config = RateLimiterConfig(maxConcurrentGetIntegrator = maxConcurrentRequests)
        sharedLimiter = MetrcRateLimiter(config)
    }

    /**
     * Create a MetrcWrapper with a shared rate limiter.
     * @param baseUrl Metrc API base URL
     * @param vendorKey Vendor API key
     * @param userKey User API key
     * @return A MetrcWrapper sharing this factory's rate limiter
     */
    fun create(baseUrl: String, vendorKey: String, userKey: String): MetrcWrapper {
        val client = MetrcClient(baseUrl, vendorKey, userKey)
        return MetrcWrapper(client, sharedLimiter)
    }

    /**
     * Get the shared rate limiter for this factory.
     */
    fun getRateLimiter(): MetrcRateLimiter = sharedLimiter
}
