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
    public class PlantsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Records additive usage for plants based on their location within a specified Facility.
        /// Permissions Required:
        /// - Manage Plants
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsAdditivesByLocation(List<PlantsCreateAdditivesByLocationRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsAdditivesByLocation(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsAdditivesByLocationUsingTemplate(List<PlantsCreateAdditivesByLocationUsingTemplateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsAdditivesByLocationUsingTemplate(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates harvest product records from plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of plants manicured.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsManicure(List<PlantsCreateManicureRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsManicure(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates packages from plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of Plant Batch packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsPlantBatchPackages(List<PlantsCreatePlantBatchPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsPlantBatchPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records additive usage details applied to specific plants at a Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsAdditives(List<PlantsCreatePlantsAdditivesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsAdditives(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records additive usage for plants using predefined additive templates at a specified Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsAdditivesUsingTemplate(List<PlantsCreatePlantsAdditivesUsingTemplateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsAdditivesUsingTemplate(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new plant batches at a specified Facility from existing plant data.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of plant batches.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsPlantings(List<PlantsCreatePlantsPlantingsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsPlantings(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records waste events for plants at a Facility, including method, reason, and location details.
        /// Permissions Required:
        /// - Manage Plants Waste
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record waste.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePlantsWaste(List<PlantsCreatePlantsWasteRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePlantsWaste(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Destroy Veg/Flower Plants
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility that intends to destroy the plants.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeletePlants(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeletePlants(licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves additive records applied to plants at a specified Facility.
        /// Permissions Required:
        /// - View/Manage Plants Additives
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to retrieve plant additives.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Additive>> GetPlantsAdditives(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsAdditives(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Additive>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of all plant additive types defined within a Facility.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetPlantsAdditivesTypes(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsAdditivesTypes(), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of flowering plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Plant>> GetFloweringPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetFloweringPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Plant>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves the list of growth phases supported by a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of available Plant Growth Phases.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetPlantsGrowthPhases(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsGrowthPhases(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves inactive plants at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Plant>> GetInactivePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactivePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Plant>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves inactive mother-phase plants at a specified Facility.
        /// Permissions Required:
        /// - View Mother Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive mother plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Mother>> GetMotherInactivePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetMotherInactivePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Mother>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
        /// Permissions Required:
        /// - View Mother Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of mother plants on hold.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Mother>> GetMotherOnHoldPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetMotherOnHoldPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Mother>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves mother-phase plants at a specified Facility.
        /// Permissions Required:
        /// - View Mother Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of mother plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Mother>> GetMotherPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetMotherPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Mother>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves plants that are currently on hold at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of plants on hold.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Plant>> GetOnHoldPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetOnHoldPlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Plant>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Plant by Id.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Plant> GetPlantsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Plant>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Plant by label.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="label">label</param>
        /// <param name="licenseNumber">If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Plant>> GetPlantsByLabel(string label, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsByLabel(label, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Plant>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of recorded plant waste events for a specific Facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantsWaste>> GetPlantsWaste(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsWaste(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantsWaste>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of all available plant waste methods for use within a Facility.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<WasteMethod>> GetPlantsWasteMethods(string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsWasteMethods(pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<WasteMethod>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retriveves available reasons for recording mature plant waste at a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<WasteReason>> GetPlantsWasteReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsWasteReasons(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<WasteReason>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of vegetating plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Plant>> GetVegetativePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetVegetativePlants(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Plant>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PlantsWaste>> GetPlantsWasteById(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsWasteById(id, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PlantsWaste>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<WastePackage>> GetPlantsWastePackageById(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPlantsWastePackageById(id, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<WastePackage>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Adjusts the recorded count of plants at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the Plant count.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateAdjustPlants(List<PlantsUpdateAdjustPlantsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateAdjustPlants(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Changes the growth phases of plants within a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants growth phase.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsGrowthPhase(List<PlantsUpdateGrowthPhaseRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsGrowthPhase(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Plant harvests.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsHarvest(List<PlantsUpdateHarvestRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsHarvest(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Merges multiple plant groups into a single group within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants merged.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsMerge(List<PlantsUpdateMergeRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsMerge(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Moves plant batches to new locations within a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants moved.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsLocation(List<PlantsUpdatePlantsLocationRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsLocation(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates the strain information for plants within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Plant strains.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsStrain(List<PlantsUpdatePlantsStrainRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsStrain(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Replaces existing plant tags with new tags for plants within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to upate the list of Plant tags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsTag(List<PlantsUpdatePlantsTagRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsTag(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Splits an existing plant group into multiple groups within a Facility.
        /// Permissions Required:
        /// - View Plant
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update Plant splits.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePlantsSplit(List<PlantsUpdateSplitRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePlantsSplit(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

