using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v1/create", body, queryParams);
        }

        /// <summary>
        /// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2", body, queryParams);
        }

        /// <summary>
        /// Creates one or more new item brands for the specified Facility identified by the License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateBrandV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/brand", body, queryParams);
        }

        /// <summary>
        /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateFileV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/file", body, queryParams);
        }

        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreatePhotoV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v1/photo", body, queryParams);
        }

        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreatePhotoV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/photo", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateUpdateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v1/update", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/items/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives the specified Product by Id for the given Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/items/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives the specified Item Brand by Id for the given Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteBrandV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/items/v2/brand/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves detailed information about a specific Item by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetActiveV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/active", null, queryParams);
        }

        /// <summary>
        /// Returns a list of active items for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/active", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetBrandsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/brands", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active item brands for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetBrandsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/brands", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> ItemsGetCategoriesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/categories", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of item categories.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> ItemsGetCategoriesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/categories", null, queryParams);
        }

        /// <summary>
        /// Retrieves a file by its Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetFileV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/file/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetInactiveV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive items for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetPhotoV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v1/photo/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves an image by its Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetPhotoV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/photo/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Updates one or more existing products for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/items/v2", body, queryParams);
        }

        /// <summary>
        /// Updates one or more existing item brands for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsUpdateBrandV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/items/v2/brand", body, queryParams);
        }

    }
}
