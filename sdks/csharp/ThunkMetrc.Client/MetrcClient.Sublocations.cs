using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Creates new sublocation records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sublocations/v2", body, queryParams);
        }

        /// <summary>
        /// Archives an existing Sublocation record for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sublocations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Sublocation by its Id, with an optional license number.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sublocations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sublocations/v2/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive sublocations for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sublocations/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Updates existing sublocation records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sublocations/v2", body, queryParams);
        }

    }
}
