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
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/strains/v1/create", body, queryParams);
        }

        /// <summary>
        /// Creates new strain records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/strains/v2", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateUpdateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/strains/v1/update", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/strains/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives an existing strain record for a Facility
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/strains/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/strains/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Strain record by its Id, with an optional license number.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/strains/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetActiveV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/strains/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/strains/v2/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/strains/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Updates existing strain records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/strains/v2", body, queryParams);
        }

    }
}
