using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
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
        public async Task CreateHarvestsPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/packages", body, queryParams);
        }
        /// <summary>
        /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests waste.</param>
        public async Task CreateHarvestsWaste(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/waste", body, queryParams);
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
        public async Task CreateHarvestsPackagesTesting(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/packages/testing", body, queryParams);
        }
        /// <summary>
        /// Discontinues a specific harvest waste record by Id for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Discontinue Harvest Waste
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests waste.</param>
        public async Task DeleteHarvestsWasteById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/harvests/v2/waste/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Marks one or more harvests as finished for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of finished harvests.</param>
        public async Task FinishHarvestsHarvests(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/finish", body, queryParams);
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
        public async Task<object> GetActiveHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/active", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        /// Permissions Required:
        /// - View Harvests
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Harvest will be validated against the specified License Number. If not specified, the Harvest will be validated against all of the User's current Facilities. Please note that if the Harvest is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetHarvestsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/{Uri.EscapeDataString(id)}", null, queryParams);
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
        public async Task<object> GetHarvestsWaste(string? harvestId = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(harvestId)) queryParams["harvestId"] = harvestId;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/waste", null, queryParams);
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
        public async Task<object> GetInactiveHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/inactive", null, queryParams);
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
        public async Task<object> GetOnHoldHarvests(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/onhold", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of Waste types for harvests.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetHarvestsWasteTypes(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/waste/types", null, queryParams);
        }
        /// <summary>
        /// Reopens one or more previously finished harvests for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of unfinished harvests.</param>
        public async Task UnfinishHarvestsHarvests(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/unfinish", body, queryParams);
        }
        /// <summary>
        /// Updates the Location of Harvest for a specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests locations.</param>
        public async Task UpdateHarvestsLocation(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/location", body, queryParams);
        }
        /// <summary>
        /// Renames one or more harvests for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Manage Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of renamed harvests.</param>
        public async Task UpdateHarvestsRename(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/rename", body, queryParams);
        }
        /// <summary>
        /// Restores previously harvested plants to their original state for the specified Facility.
        /// Permissions Required:
        /// - View Harvests
        /// - Finish/Discontinue Harvests
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests restored.</param>
        public async Task UpdateHarvestsRestoreHarvestedPlants(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/restore/harvestedplants", body, queryParams);
        }
    }
}

