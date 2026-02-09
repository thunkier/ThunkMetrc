using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using ThunkMetrc.Client;
using Microsoft.Extensions.Logging;
using System.Text.Json;
using System.Threading;
using System.Runtime.CompilerServices;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;
using File = ThunkMetrc.Wrapper.Models.File;

namespace ThunkMetrc.Wrapper.Services
{
    public class ItemsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates one or more new item brands for the specified Facility identified by the License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brands to create.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateItemsBrand(List<ItemsCreateBrandRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateItemsBrand(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">Optional query parameter.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateItemsFile(List<ItemsCreateFileRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateItemsFile(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create new items.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateItems(List<ItemsCreateItemsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateItems(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to add an image.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateItemsPhoto(List<ItemsCreatePhotoRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateItemsPhoto(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives the specified Item Brand by Id for the given Facility License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brand to delete.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteItemsBrandById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteItemsBrandById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives the specified Product by Id for the given Facility License Number.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item to delete.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteItemsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteItemsById(id, licenseNumber); return null; }, cancellationToken);
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Item>> GetActiveItems(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveItems(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Item>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of active item brands for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active item brands.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Brand>> GetItemsBrands(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetItemsBrands(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Brand>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of item categories.
        /// </summary>
        /// <param name="licenseNumber">If specified, the categories will be retrieved for the specified License Number. If not specified, all item categories will be returned.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Category>> GetItemsCategories(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetItemsCategories(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Category>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a file by its Id for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the file.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<File> GetItemsFileById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetItemsFileById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<File>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive items for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive items.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Item>> GetInactiveItems(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveItems(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Item>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves detailed information about a specific Item by Id.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Item will be validated against the specified License Number. If not specified, the Item will be validated against all of the User's current Facilities. Please note that if the Item is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Item> GetItemsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetItemsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Item>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves an image by its Id for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the image.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Photo> GetItemsPhotoById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetItemsPhotoById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Photo>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates one or more existing item brands for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item Brand updates.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateItemsBrand(List<ItemsUpdateBrandRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateItemsBrand(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates one or more existing products for the specified Facility.
        /// Permissions Required:
        /// - Manage Items
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for the Item updates.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateItems(List<ItemsUpdateItemsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateItems(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

