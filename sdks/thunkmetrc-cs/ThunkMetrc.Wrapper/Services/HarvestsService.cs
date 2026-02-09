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
    public class HarvestsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates packages from harvested products for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateHarvestsPackages(List<HarvestsCreateHarvestsPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateHarvestsPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests waste.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateHarvestsWaste(List<HarvestsCreateHarvestsWasteRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateHarvestsWaste(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates packages for testing from harvested products for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateHarvestsPackagesTesting(List<HarvestsCreatePackagesTestingRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateHarvestsPackagesTesting(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Discontinues a specific harvest waste record by Id for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Discontinue Harvest Waste
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests waste.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteHarvestsWasteById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteHarvestsWasteById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Marks one or more harvests as finished for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of finished harvests.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task FinishHarvestsHarvests(List<HarvestsFinishHarvestsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.FinishHarvestsHarvests(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active harvests for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active harvests.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Harvest>> GetActiveHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Harvest>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Harvest will be validated against the specified License Number. If not specified, the Harvest will be validated against all of the User's current Facilities. Please note that if the Harvest is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Harvest> GetHarvestsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetHarvestsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Harvest>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="harvestId">The harvestId is the unique identifier for each harvest.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste harvests.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<HarvestsWaste>> GetHarvestsWaste(string? harvestId = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetHarvestsWaste(harvestId, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<HarvestsWaste>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive harvests for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive harvests.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Harvest>> GetInactiveHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Harvest>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of harvests on hold for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests on hold.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Harvest>> GetOnHoldHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetOnHoldHarvests(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Harvest>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of Waste types for harvests.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<WasteType>> GetHarvestsWasteTypes(string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetHarvestsWasteTypes(pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<WasteType>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Reopens one or more previously finished harvests for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of unfinished harvests.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UnfinishHarvestsHarvests(List<HarvestsUnfinishHarvestsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UnfinishHarvestsHarvests(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the Location of Harvest for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests locations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateHarvestsLocation(List<HarvestsUpdateHarvestsLocationRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateHarvestsLocation(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Renames one or more harvests for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of renamed harvests.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateHarvestsRename(List<HarvestsUpdateRenameRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateHarvestsRename(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Restores previously harvested plants to their original state for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests restored.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateHarvestsRestoreHarvestedPlants(List<HarvestsUpdateRestoreHarvestedPlantsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateHarvestsRestoreHarvestedPlants(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

