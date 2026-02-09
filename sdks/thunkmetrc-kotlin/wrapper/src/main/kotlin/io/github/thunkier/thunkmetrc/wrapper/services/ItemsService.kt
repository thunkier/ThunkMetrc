package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class ItemsService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

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
    suspend fun createBrand(licenseNumber: String? = null, body: List<ItemsCreateBrandRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.createItemsBrand(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createFile(licenseNumber: String? = null, body: List<ItemsCreateFileRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.createItemsFile(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createItems(licenseNumber: String? = null, body: List<ItemsCreateItemsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.createItems(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun createPhoto(licenseNumber: String? = null, body: List<ItemsCreatePhotoRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.createItemsPhoto(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun deleteBrandById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.deleteItemsBrandById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun deleteItemsById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.deleteItemsById(id, licenseNumber, ) 
        } as? Any
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
    suspend fun getActiveItems(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Item>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getActiveItems(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Item>
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
    suspend fun getBrands(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Brand>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getItemsBrands(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Brand>
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
    suspend fun getCategories(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Category>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getItemsCategories(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Category>
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
    suspend fun getFileById(id: String, licenseNumber: String? = null): File? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getItemsFileById(id, licenseNumber, ) 
        } as? File
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
    suspend fun getInactiveItems(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Item>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getInactiveItems(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Item>
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
    suspend fun getItemsById(id: String, licenseNumber: String? = null): Item? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getItemsById(id, licenseNumber, ) 
        } as? Item
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
    suspend fun getPhotoById(id: String, licenseNumber: String? = null): Photo? {
        return rateLimiter.execute(null,true,
        ) { 
            client.items.getItemsPhotoById(id, licenseNumber, ) 
        } as? Photo
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
    suspend fun updateBrand(licenseNumber: String? = null, body: List<ItemsUpdateBrandRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.updateItemsBrand(licenseNumber, body) 
        } as? WriteResponse
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
    suspend fun updateItems(licenseNumber: String? = null, body: List<ItemsUpdateItemsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.items.updateItems(licenseNumber, body) 
        } as? WriteResponse
    }
}

