using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Records additive usage for plants based on their location within a specified Facility.
        /// Permissions Required:
        /// - Manage Plants
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        public async Task CreatePlantsAdditivesByLocation(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/bylocation", body, queryParams);
        }
        /// <summary>
        /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        public async Task CreatePlantsAdditivesByLocationUsingTemplate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/bylocation/usingtemplate", body, queryParams);
        }
        /// <summary>
        /// Creates harvest product records from plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of plants manicured.</param>
        public async Task CreatePlantsManicure(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/manicure", body, queryParams);
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
        public async Task CreatePlantsPlantBatchPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/plantbatch/packages", body, queryParams);
        }
        /// <summary>
        /// Records additive usage details applied to specific plants at a Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        public async Task CreatePlantsAdditives(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives", body, queryParams);
        }
        /// <summary>
        /// Records additive usage for plants using predefined additive templates at a specified Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record Plant additives.</param>
        public async Task CreatePlantsAdditivesUsingTemplate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/additives/usingtemplate", body, queryParams);
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
        public async Task CreatePlantsPlantings(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/plantings", body, queryParams);
        }
        /// <summary>
        /// Records waste events for plants at a Facility, including method, reason, and location details.
        /// Permissions Required:
        /// - Manage Plants Waste
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record waste.</param>
        public async Task CreatePlantsWaste(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plants/v2/waste", body, queryParams);
        }
        /// <summary>
        /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Destroy Veg/Flower Plants
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility that intends to destroy the plants.</param>
        public async Task DeletePlants(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plants/v2", null, queryParams);
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
        public async Task<object> GetPlantsAdditives(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/additives", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of all plant additive types defined within a Facility.
        /// </summary>
        public async Task<object> GetPlantsAdditivesTypes()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/additives/types", null, queryParams);
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
        public async Task<object> GetFloweringPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/flowering", null, queryParams);
        }
        /// <summary>
        /// Retrieves the list of growth phases supported by a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of available Plant Growth Phases.</param>
        public async Task<object> GetPlantsGrowthPhases(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/growthphases", null, queryParams);
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
        public async Task<object> GetInactivePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/inactive", null, queryParams);
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
        public async Task<object> GetMotherInactivePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother/inactive", null, queryParams);
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
        public async Task<object> GetMotherOnHoldPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother/onhold", null, queryParams);
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
        public async Task<object> GetMotherPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/mother", null, queryParams);
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
        public async Task<object> GetOnHoldPlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/onhold", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Plant by Id.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPlantsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Plant by label.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// </summary>
        /// <param name="label">label</param>
        /// <param name="licenseNumber">If specified, the Plant will be validated against the specified License Number. If not specified, the Plant will be validated against all of the User's current Facilities. Please note that if the Plant is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPlantsByLabel(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/{Uri.EscapeDataString(label)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of recorded plant waste events for a specific Facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPlantsWaste(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of all available plant waste methods for use within a Facility.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPlantsWasteMethods(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/methods/all", null, queryParams);
        }
        /// <summary>
        /// Retriveves available reasons for recording mature plant waste at a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPlantsWasteReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/reasons", null, queryParams);
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
        public async Task<object> GetVegetativePlants(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/vegetative", null, queryParams);
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
        public async Task<object> GetPlantsWasteById(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/{Uri.EscapeDataString(id)}/plant", null, queryParams);
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
        public async Task<object> GetPlantsWastePackageById(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plants/v2/waste/{Uri.EscapeDataString(id)}/package", null, queryParams);
        }
        /// <summary>
        /// Adjusts the recorded count of plants at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the Plant count.</param>
        public async Task UpdateAdjustPlants(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/adjust", body, queryParams);
        }
        /// <summary>
        /// Changes the growth phases of plants within a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants growth phase.</param>
        public async Task UpdatePlantsGrowthPhase(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/growthphase", body, queryParams);
        }
        /// <summary>
        /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Plant harvests.</param>
        public async Task UpdatePlantsHarvest(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/harvest", body, queryParams);
        }
        /// <summary>
        /// Merges multiple plant groups into a single group within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants merged.</param>
        public async Task UpdatePlantsMerge(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/merge", body, queryParams);
        }
        /// <summary>
        /// Moves plant batches to new locations within a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of plants moved.</param>
        public async Task UpdatePlantsLocation(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/location", body, queryParams);
        }
        /// <summary>
        /// Updates the strain information for plants within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Plant strains.</param>
        public async Task UpdatePlantsStrain(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/strain", body, queryParams);
        }
        /// <summary>
        /// Replaces existing plant tags with new tags for plants within a Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to upate the list of Plant tags.</param>
        public async Task UpdatePlantsTag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/tag", body, queryParams);
        }
        /// <summary>
        /// Splits an existing plant group into multiple groups within a Facility.
        /// Permissions Required:
        /// - View Plant
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update Plant splits.</param>
        public async Task UpdatePlantsSplit(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plants/v2/split", body, queryParams);
        }
    }
}

