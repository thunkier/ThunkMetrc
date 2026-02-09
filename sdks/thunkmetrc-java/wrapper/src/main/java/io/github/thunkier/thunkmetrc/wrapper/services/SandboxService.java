package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;

public class SandboxService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public SandboxService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * POST /sandbox/v2/integrator/setup
     * @param userKey Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public Object createSandboxIntegratorSetup(
        String userKey, Object body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.sandbox.createSandboxIntegratorSetup(
                userKey, body
            )
        );
    }

}

