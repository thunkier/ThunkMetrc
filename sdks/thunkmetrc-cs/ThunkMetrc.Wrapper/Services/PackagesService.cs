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
    public class PackagesService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Records a list of adjustments for packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of adjustments.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateAdjustPackages(List<PackagesCreateAdjustPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateAdjustPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePackagesPackages(List<PackagesCreatePackagesPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePackagesPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new plantings from packages for a specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new plantings.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePackagesPlantings(List<PackagesCreatePackagesPlantingsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePackagesPlantings(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new packages for testing for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new packages for testing.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePackagesTesting(List<PackagesCreateTestingRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePackagesTesting(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Discontinues a Package at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to discontinue.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeletePackagesById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeletePackagesById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates a list of packages as finished for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of finish packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task FinishPackagesPackages(List<PackagesFinishPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.FinishPackagesPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Flags one or more Packages at the specified Facility as Finished Goods.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to flag finished goods.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task FinishedgoodFlagPackages(List<PackagesFinishedgoodFlagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.FinishedgoodFlagPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Removes the Finished Good flag one or more Packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to flag finished goods.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task FinishedgoodUnflagPackages(List<PackagesFinishedgoodUnflagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.FinishedgoodUnflagPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PackagesPackage>> GetActivePackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActivePackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of adjustment reasons for packages at a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of adjustment reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<AdjustReason>> GetPackagesAdjustReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesAdjustReasons(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<AdjustReason>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves the Package Adjustments for a Facility
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of Package Adjustments.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Adjustment>> GetPackagesAdjustments(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesAdjustments(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Adjustment>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of packages in transit for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of packages in transit.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<InTransit>> GetInTransitPackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInTransitPackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<InTransit>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PackagesPackage>> GetInactivePackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactivePackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of lab sample packages that have been created/sent for testing.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PackagesPackage>> GetPackagesLabSamples(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesLabSamples(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of packages on hold for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of packages on hold.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PackagesPackage>> GetOnHoldPackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetOnHoldPackages(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Package by its Id.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PackagesPackage> GetPackagesById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PackagesPackage>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Package by its label.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="label">label</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<PackagesPackage>> GetPackagesByLabel(string label, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesByLabel(label, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of available Package types.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetPackagesTypes(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesTypes(), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves the source harvests for a Package by its Id.
        /// Permissions Required:
        /// - View Package Source Harvests
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<SourceHarvest> GetPackagesSourceHarvestById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesSourceHarvestById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<SourceHarvest>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of transferred packages for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of transferred packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PackagesPackage>> GetPackagesTransferred(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPackagesTransferred(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PackagesPackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates a list of packages as unfinished for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of unfinish packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UnfinishPackagesPackages(List<PackagesUnfinishPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UnfinishPackagesPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Set the final quantity for a Package.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of adjustments.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateAdjustPackages(List<PackagesUpdateAdjustPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateAdjustPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the Product decontaminate information for a list of packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of product decontaminations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesDecontaminate(List<PackagesUpdateDecontaminateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesDecontaminate(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Flags one or more packages for donation at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of flagged donations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesDonationFlag(List<PackagesUpdateDonationFlagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesDonationFlag(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Removes the donation flag from one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of unflaged donations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesDonationUnflag(List<PackagesUpdateDonationUnflagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesDonationUnflag(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the external identifiers for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Package Inventory
        /// - External Id Enabled
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of change external Ids.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesExternalid(List<PackagesUpdateExternalidRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesExternalid(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the associated Item for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of changed items.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesItem(List<PackagesUpdateItemRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesItem(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the list of required lab test batches for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of required lab test batches.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesLabtestsRequired(List<PackagesUpdateLabtestsRequiredRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesLabtestsRequired(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates notes associated with one or more packages for the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// - Manage Package Notes
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of change notes.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesNote(List<PackagesUpdateNoteRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesNote(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the Location and Sublocation for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of change locations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesLocation(List<PackagesUpdatePackagesLocationRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesLocation(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates a list of Product remediations for packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of product remediations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesRemediate(List<PackagesUpdateRemediateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesRemediate(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Flags or unflags one or more packages at the specified Facility as trade samples.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update trade sample flags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesTradeSampleFlag(List<PackagesUpdateTradeSampleFlagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesTradeSampleFlag(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Removes the trade sample flag from one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update trade sample unflag.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesTradeSampleUnflag(List<PackagesUpdateTradeSampleUnflagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesTradeSampleUnflag(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the use-by date for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of changed items.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePackagesUseByDate(List<PackagesUpdateUseByDateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePackagesUseByDate(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

