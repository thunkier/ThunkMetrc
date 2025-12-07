using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsCreateFinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v1/finish", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v1/create/packages", body, queryParams);
        }

        /// <summary>
        /// Creates packages from harvested products for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/packages", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageTestingV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v1/create/packages/testing", body, queryParams);
        }

        /// <summary>
        /// Creates packages for testing from harvested products for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageTestingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/packages/testing", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsCreateRemoveWasteV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v1/removewaste", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsCreateUnfinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v1/unfinish", body, queryParams);
        }

        /// <summary>
        /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsCreateWasteV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/harvests/v2/waste", body, queryParams);
        }

        /// <summary>
        /// Discontinues a specific harvest waste record by Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Discontinue Harvest Waste
        /// </summary>
        public async Task HarvestsDeleteWasteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/harvests/v2/waste/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active harvests for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive harvests for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v1/onhold", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of harvests on hold for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetWasteV2(string? harvestId = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(harvestId)) queryParams["harvestId"] = harvestId;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/waste", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> HarvestsGetWasteTypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v1/waste/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of Waste types for harvests.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> HarvestsGetWasteTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/harvests/v2/waste/types", null, queryParams);
        }

        /// <summary>
        /// Marks one or more harvests as finished for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateFinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/finish", body, queryParams);
        }

        /// <summary>
        /// Updates the Location of Harvest for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateLocationV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/location", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateMoveV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v1/move", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateRenameV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v1/rename", body, queryParams);
        }

        /// <summary>
        /// Renames one or more harvests for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateRenameV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/rename", body, queryParams);
        }

        /// <summary>
        /// Restores previously harvested plants to their original state for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateRestoreHarvestedPlantsV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/restore/harvestedplants", body, queryParams);
        }

        /// <summary>
        /// Reopens one or more previously finished harvests for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateUnfinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/harvests/v2/unfinish", body, queryParams);
        }

    }
}
