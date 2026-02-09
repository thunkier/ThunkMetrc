package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class AdditivesTemplatesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public AdditivesTemplatesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates new additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * POST /additivestemplates/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createAdditivesTemplates(
        String licenseNumber, List<AdditivesTemplatesCreateAdditivesTemplatesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.additivesTemplates.createAdditivesTemplates(
                licenseNumber, body
            )
        );
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<AdditivesTemplate> getActiveAdditivesTemplates(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<AdditivesTemplate>) client.additivesTemplates.getActiveAdditivesTemplates(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves an Additive Template by its Id.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public AdditivesTemplate getAdditivesTemplatesById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (AdditivesTemplate) client.additivesTemplates.getAdditivesTemplatesById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * GET /additivestemplates/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<AdditivesTemplate> getInactiveAdditivesTemplates(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<AdditivesTemplate>) client.additivesTemplates.getInactiveAdditivesTemplates(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * Permissions Required:
     * - Manage Additives
     * PUT /additivestemplates/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateAdditivesTemplates(
        String licenseNumber, List<AdditivesTemplatesUpdateAdditivesTemplatesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.additivesTemplates.updateAdditivesTemplates(
                licenseNumber, body
            )
        );
    }

}

