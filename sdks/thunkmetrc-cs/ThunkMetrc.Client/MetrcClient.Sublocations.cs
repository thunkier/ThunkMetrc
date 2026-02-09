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
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of sublocations.</param>
        public async Task CreateSublocations(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sublocations/v2", body, queryParams);
        }
        /// <summary>
        /// Archives an existing Sublocation record for a Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to delete the list of sublocations.</param>
        public async Task DeleteSublocationsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sublocations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="lastModifiedEnd">The Last Modified end timestamp</param>
        /// <param name="lastModifiedStart">The Last Modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active sublocations.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetActiveSublocations(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive sublocations.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetInactiveSublocations(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sublocations/v2/inactive", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Sublocation by its Id, with an optional license number.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sublocation will be validated against the specified License Number. If not specified, the Sublocation will be validated against all of the User's current Facilities. Please note that if the Sublocation is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetSublocationsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sublocations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Updates existing sublocation records for a specified Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of sublocations.</param>
        public async Task UpdateSublocations(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sublocations/v2", body, queryParams);
        }
    }
}

