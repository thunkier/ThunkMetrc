package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class FacilitiesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public FacilitiesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * GET /facilities/v2
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<Facility> getFacilities(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Facility>) client.facilities.getFacilities(
                
            )
        );
    }

}

