package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;

public class UnitsOfMeasureService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public UnitsOfMeasureService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Retrieves all active units of measure.
     * GET /unitsofmeasure/v2/active
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<UnitOfMeasure> getActiveUnitsOfMeasure(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<UnitOfMeasure>) client.unitsOfMeasure.getActiveUnitsOfMeasure(
                
            )
        );
    }

    /**
     * Retrieves all inactive units of measure.
     * GET /unitsofmeasure/v2/inactive
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<UnitOfMeasure> getInactiveUnitsOfMeasure(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<UnitOfMeasure>) client.unitsOfMeasure.getInactiveUnitsOfMeasure(
                
            )
        );
    }

}

