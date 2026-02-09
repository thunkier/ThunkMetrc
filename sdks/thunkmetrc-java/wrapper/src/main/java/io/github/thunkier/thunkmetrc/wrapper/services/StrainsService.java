package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class StrainsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public StrainsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates new strain records for a specified Facility.
     * Permissions Required:
     * - Manage Strains
     * POST /strains/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createStrains(
        String licenseNumber, List<StrainsCreateStrainsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.strains.createStrains(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives an existing strain record for a Facility
     * Permissions Required:
     * - Manage Strains
     * DELETE /strains/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteStrainsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.strains.deleteStrainsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Strain> getActiveStrains(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Strain>) client.strains.getActiveStrains(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Strain> getInactiveStrains(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Strain>) client.strains.getInactiveStrains(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * Permissions Required:
     * - Manage Strains
     * GET /strains/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Strain getStrainsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Strain) client.strains.getStrainsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Updates existing strain records for a specified Facility.
     * Permissions Required:
     * - Manage Strains
     * PUT /strains/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateStrains(
        String licenseNumber, List<StrainsUpdateStrainsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.strains.updateStrains(
                licenseNumber, body
            )
        );
    }

}

