using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Creates one or more new item brands for the specified Facility identified by the License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brands to create.</param>
        public async Task CreateItemsBrand(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/brand", body, queryParams);
        }
        /// <summary>
        /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">Optional query parameter.</param>
        public async Task CreateItemsFile(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/file", body, queryParams);
        }
        /// <summary>
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create new items.</param>
        public async Task CreateItems(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2", body, queryParams);
        }
        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to add an image.</param>
        public async Task CreateItemsPhoto(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/items/v2/photo", body, queryParams);
        }
        /// <summary>
        /// Archives the specified Item Brand by Id for the given Facility License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brand to delete.</param>
        public async Task DeleteItemsBrandById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/items/v2/brand/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Archives the specified Product by Id for the given Facility License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item to delete.</param>
        public async Task DeleteItemsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/items/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Returns a list of active items for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active items.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetActiveItems(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of active item brands for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active item brands.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetItemsBrands(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/brands", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of item categories.
        /// </summary>
        /// <param name="licenseNumber">If specified, the categories will be retrieved for the specified License Number. If not specified, all item categories will be returned.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetItemsCategories(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/categories", null, queryParams);
        }
        /// <summary>
        /// Retrieves a file by its Id for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the file.</param>
        public async Task<object> GetItemsFileById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/file/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of inactive items for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive items.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetInactiveItems(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/inactive", null, queryParams);
        }
        /// <summary>
        /// Retrieves detailed information about a specific Item by Id.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Item will be validated against the specified License Number. If not specified, the Item will be validated against all of the User's current Facilities. Please note that if the Item is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetItemsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves an image by its Id for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the image.</param>
        public async Task<object> GetItemsPhotoById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/items/v2/photo/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Updates one or more existing item brands for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brand updates.</param>
        public async Task UpdateItemsBrand(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/items/v2/brand", body, queryParams);
        }
        /// <summary>
        /// Updates one or more existing products for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item updates.</param>
        public async Task UpdateItems(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/items/v2", body, queryParams);
        }
    }
}

