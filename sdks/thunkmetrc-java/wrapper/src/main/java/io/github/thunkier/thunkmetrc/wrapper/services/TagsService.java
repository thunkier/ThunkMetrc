package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class TagsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public TagsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET /tags/v2/package/available
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<Tag> getTagsAvailablePackage(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Tag>) client.tags.getTagsAvailablePackage(
                licenseNumber
            )
        );
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET /tags/v2/plant/available
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<Tag> getTagsAvailablePlant(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Tag>) client.tags.getTagsAvailablePlant(
                licenseNumber
            )
        );
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * - RetailId.AllowPackageStaging Key Value enabled
     * GET /tags/v2/staged
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<Staged> getStagedTags(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<Staged>) client.tags.getStagedTags(
                licenseNumber
            )
        );
    }

}

