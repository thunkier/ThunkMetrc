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
    public class PlantBatchesService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record adjustments.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateAdjustPlantBatches(List<PlantBatchesCreateAdjustPlantBatchesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateAdjustPlantBatches(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the growth phase of plants at a specified Facility based on tracking information.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to change the growth phase.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesGrowthPhase(List<PlantBatchesCreateGrowthPhaseRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesGrowthPhase(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates packages from mother plants at the specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create packages from a mother plant.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesPackagesFromMotherPlant(List<PlantBatchesCreatePackagesFromMotherPlantRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesPackagesFromMotherPlant(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records Additive usage details for plant batches at a specific Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesAdditives(List<PlantBatchesCreatePlantBatchesAdditivesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesAdditives(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records Additive usage for plant batches at a Facility using predefined additive templates.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesAdditivesUsingTemplate(List<PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesAdditivesUsingTemplate(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a package from the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesPackages(List<PlantBatchesCreatePlantBatchesPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create the plantings.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesPlantings(List<PlantBatchesCreatePlantBatchesPlantingsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesPlantings(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records waste information for plant batches based on the submitted data for the specified Facility.
        /// Permissions Required:
        /// - Manage Plants Waste
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record waste of the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesWaste(List<PlantBatchesCreatePlantBatchesWasteRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesWaste(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Splits an existing Plant Batch into multiple groups at the specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to split the plant batches.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantBatchesSplit(List<PlantBatchesCreateSplitRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantBatchesSplit(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Completes the destruction of plant batches based on the provided input data.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Destroy Immature Plants
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to destroy the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeletePlantBatches(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeletePlantBatches(licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
        /// Permissions Required:
        /// - View Immature Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active plant batches.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantBatch>> GetActivePlantBatches(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActivePlantBatches(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantBatch>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
        /// Permissions Required:
        /// - View Immature Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive plant batches.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantBatch>> GetInactivePlantBatches(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactivePlantBatches(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantBatch>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Plant Batch by Id.
        /// Permissions Required:
        /// - View Immature Plants
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Plant Batch will be validated against the specified License Number. If not specified, the Plant Batch will be validated against all of the User's current Facilities. Please note that if the Plant Batch is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PlantBatch> GetPlantBatchesById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantBatchesById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PlantBatch>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of plant batch types.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantBatchesType>> GetPlantBatchesTypes(string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantBatchesTypes(pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantBatchesType>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves waste details associated with plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantBatchesWaste>> GetPlantBatchesWaste(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantBatchesWaste(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantBatchesWaste>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a list of waste reasons.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<PlantBatchesWasteReason>> GetPlantBatchesWasteReasons(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantBatchesWasteReasons(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<PlantBatchesWasteReason>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Renames plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to rename the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantBatchesName(List<PlantBatchesUpdateNameRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantBatchesName(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Moves one or more plant batches to new locations with in a specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to change the Location of the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantBatchesLocation(List<PlantBatchesUpdatePlantBatchesLocationRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantBatchesLocation(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Changes the strain of plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to change the Strain of the Plant Batch.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantBatchesStrain(List<PlantBatchesUpdatePlantBatchesStrainRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantBatchesStrain(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Replaces tags for plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to replace Plant Batch tags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantBatchesTag(List<PlantBatchesUpdatePlantBatchesTagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantBatchesTag(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

