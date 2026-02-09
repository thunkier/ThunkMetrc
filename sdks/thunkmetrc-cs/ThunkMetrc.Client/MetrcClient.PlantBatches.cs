using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record adjustments.</param>
        public async Task CreateAdjustPlantBatches(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/adjust", body, queryParams);
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
        public async Task CreatePlantBatchesGrowthPhase(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/growthphase", body, queryParams);
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
        public async Task CreatePlantBatchesPackagesFromMotherPlant(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/packages/frommotherplant", body, queryParams);
        }
        /// <summary>
        /// Records Additive usage details for plant batches at a specific Facility.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record plant additives.</param>
        public async Task CreatePlantBatchesAdditives(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/additives", body, queryParams);
        }
        /// <summary>
        /// Records Additive usage for plant batches at a Facility using predefined additive templates.
        /// Permissions Required:
        /// - Manage Plants Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record plant additives.</param>
        public async Task CreatePlantBatchesAdditivesUsingTemplate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/additives/usingtemplate", body, queryParams);
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
        public async Task CreatePlantBatchesPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/packages", body, queryParams);
        }
        /// <summary>
        /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create the plantings.</param>
        public async Task CreatePlantBatchesPlantings(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/plantings", body, queryParams);
        }
        /// <summary>
        /// Records waste information for plant batches based on the submitted data for the specified Facility.
        /// Permissions Required:
        /// - Manage Plants Waste
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record waste of the Plant Batch.</param>
        public async Task CreatePlantBatchesWaste(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/waste", body, queryParams);
        }
        /// <summary>
        /// Splits an existing Plant Batch into multiple groups at the specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to split the plant batches.</param>
        public async Task CreatePlantBatchesSplit(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/plantbatches/v2/split", body, queryParams);
        }
        /// <summary>
        /// Completes the destruction of plant batches based on the provided input data.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Destroy Immature Plants
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to destroy the Plant Batch.</param>
        public async Task DeletePlantBatches(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/plantbatches/v2", null, queryParams);
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
        public async Task<object> GetActivePlantBatches(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/active", null, queryParams);
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
        public async Task<object> GetInactivePlantBatches(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/inactive", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Plant Batch by Id.
        /// Permissions Required:
        /// - View Immature Plants
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Plant Batch will be validated against the specified License Number. If not specified, the Plant Batch will be validated against all of the User's current Facilities. Please note that if the Plant Batch is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPlantBatchesById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of plant batch types.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPlantBatchesTypes(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/types", null, queryParams);
        }
        /// <summary>
        /// Retrieves waste details associated with plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Plants Waste
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPlantBatchesWaste(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/waste", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a list of waste reasons.</param>
        public async Task<object> GetPlantBatchesWasteReasons(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/plantbatches/v2/waste/reasons", null, queryParams);
        }
        /// <summary>
        /// Renames plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to rename the Plant Batch.</param>
        public async Task UpdatePlantBatchesName(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/name", body, queryParams);
        }
        /// <summary>
        /// Moves one or more plant batches to new locations with in a specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to change the Location of the Plant Batch.</param>
        public async Task UpdatePlantBatchesLocation(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/location", body, queryParams);
        }
        /// <summary>
        /// Changes the strain of plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to change the Strain of the Plant Batch.</param>
        public async Task UpdatePlantBatchesStrain(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/strain", body, queryParams);
        }
        /// <summary>
        /// Replaces tags for plant batches at a specified Facility.
        /// Permissions Required:
        /// - View Veg/Flower Plants
        /// - Manage Veg/Flower Plants Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to replace Plant Batch tags.</param>
        public async Task UpdatePlantBatchesTag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/plantbatches/v2/tag", body, queryParams);
        }
    }
}

