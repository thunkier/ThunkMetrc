package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class ItemsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public ItemsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates one or more new item brands for the specified Facility identified by the License Number.
     * Permissions Required:
     * - Manage Items
     * POST /items/v2/brand
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createItemsBrand(
        String licenseNumber, List<ItemsCreateBrandRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.createItemsBrand(
                licenseNumber, body
            )
        );
    }

    /**
     * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
     * Permissions Required:
     * - Manage Items
     * POST /items/v2/file
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createItemsFile(
        String licenseNumber, List<ItemsCreateFileRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.createItemsFile(
                licenseNumber, body
            )
        );
    }

    /**
     * Permissions Required:
     * - Manage Items
     * POST /items/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createItems(
        String licenseNumber, List<ItemsCreateItemsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.createItems(
                licenseNumber, body
            )
        );
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * Permissions Required:
     * - Manage Items
     * POST /items/v2/photo
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createItemsPhoto(
        String licenseNumber, List<ItemsCreatePhotoRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.createItemsPhoto(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives the specified Item Brand by Id for the given Facility License Number.
     * Permissions Required:
     * - Manage Items
     * DELETE /items/v2/brand/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteItemsBrandById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.items.deleteItemsBrandById(
                id, licenseNumber
            )
        );
    }

    /**
     * Archives the specified Product by Id for the given Facility License Number.
     * Permissions Required:
     * - Manage Items
     * DELETE /items/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteItemsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.items.deleteItemsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Returns a list of active items for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Item> getActiveItems(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Item>) client.items.getActiveItems(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of active item brands for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/brands
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Brand> getItemsBrands(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Brand>) client.items.getItemsBrands(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of item categories.
     * GET /items/v2/categories
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Category> getItemsCategories(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Category>) client.items.getItemsCategories(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a file by its Id for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/file/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public File getItemsFileById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (File) client.items.getItemsFileById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of inactive items for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Item> getInactiveItems(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Item>) client.items.getInactiveItems(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves detailed information about a specific Item by Id.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Item getItemsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Item) client.items.getItemsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves an image by its Id for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * GET /items/v2/photo/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Photo getItemsPhotoById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Photo) client.items.getItemsPhotoById(
                id, licenseNumber
            )
        );
    }

    /**
     * Updates one or more existing item brands for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * PUT /items/v2/brand
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateItemsBrand(
        String licenseNumber, List<ItemsUpdateBrandRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.updateItemsBrand(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates one or more existing products for the specified Facility.
     * Permissions Required:
     * - Manage Items
     * PUT /items/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateItems(
        String licenseNumber, List<ItemsUpdateItemsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.items.updateItems(
                licenseNumber, body
            )
        );
    }

}

