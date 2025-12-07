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
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/locations/v1/create", body, queryParams);
        }

        /// <summary>
        /// Creates new locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/locations/v2", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateUpdateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/locations/v1/update", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/locations/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a specified Location, identified by its Id, for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/locations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Location by its Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetActiveV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v2/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v2/inactive", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetTypesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v1/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active location types for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetTypesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/locations/v2/types", null, queryParams);
        }

        /// <summary>
        /// Updates existing locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/locations/v2", body, queryParams);
        }

    }
}
