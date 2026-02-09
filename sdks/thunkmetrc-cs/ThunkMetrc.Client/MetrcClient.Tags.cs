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
        /// Permissions Required:
        /// - Manage Tags
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return available Package tags.</param>
        public async Task<object> GetTagsAvailablePackage(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/package/available", null, queryParams);
        }
        /// <summary>
        /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
        /// Permissions Required:
        /// - Manage Tags
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return available tags.</param>
        public async Task<object> GetTagsAvailablePlant(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/plant/available", null, queryParams);
        }
        /// <summary>
        /// Returns a list of staged tags. NOTE: This is a premium endpoint.
        /// Permissions Required:
        /// - Manage Tags
        /// - RetailId.AllowPackageStaging Key Value enabled
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return staged tags.</param>
        public async Task<object> GetStagedTags(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/tags/v2/staged", null, queryParams);
        }
    }
}

