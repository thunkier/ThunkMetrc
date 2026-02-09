package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class SublocationsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public SublocationsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates new sublocation records for a Facility.
     * Permissions Required:
     * - Manage Locations
     * POST /sublocations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createSublocations(
        String licenseNumber, List<SublocationsCreateSublocationsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sublocations.createSublocations(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * Permissions Required:
     * - Manage Locations
     * DELETE /sublocations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteSublocationsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.sublocations.deleteSublocationsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Sublocation> getActiveSublocations(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Sublocation>) client.sublocations.getActiveSublocations(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Sublocation> getInactiveSublocations(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Sublocation>) client.sublocations.getInactiveSublocations(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * Permissions Required:
     * - Manage Locations
     * GET /sublocations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Sublocation getSublocationsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Sublocation) client.sublocations.getSublocationsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * PUT /sublocations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateSublocations(
        String licenseNumber, List<SublocationsUpdateSublocationsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sublocations.updateSublocations(
                licenseNumber, body
            )
        );
    }

}

