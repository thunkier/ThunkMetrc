package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class WasteMethodsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public WasteMethodsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Retrieves all available waste methods.
     * GET /wastemethods/v2
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<WasteMethod> getWasteMethodsWasteMethods(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<WasteMethod>) client.wasteMethods.getWasteMethodsWasteMethods(
                
            )
        );
    }

}

