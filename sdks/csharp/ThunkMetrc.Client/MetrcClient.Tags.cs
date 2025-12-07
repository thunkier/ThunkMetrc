using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Returns a list of available package tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        /// </summary>
        public async Task<object> TagsGetPackageAvailableV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/package/available", null, queryParams);
        }

        /// <summary>
        /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        /// </summary>
        public async Task<object> TagsGetPlantAvailableV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/plant/available", null, queryParams);
        }

        /// <summary>
        /// Returns a list of staged tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        ///   - RetailId.AllowPackageStaging Key Value enabled
        /// </summary>
        public async Task<object> TagsGetStagedV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/staged", null, queryParams);
        }

    }
}
