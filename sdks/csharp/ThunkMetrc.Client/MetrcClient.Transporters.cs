using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Creates new driver records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersCreateDriverV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transporters/v2/drivers", body, queryParams);
        }

        /// <summary>
        /// Creates new vehicle records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersCreateVehicleV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transporters/v2/vehicles", body, queryParams);
        }

        /// <summary>
        /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersDeleteDriverV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transporters/v2/drivers/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersDeleteVehicleV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transporters/v2/vehicles/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetDriverV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/drivers/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of drivers for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetDriversV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/drivers", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetVehicleV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/vehicles/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of vehicles for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetVehiclesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/vehicles", null, queryParams);
        }

        /// <summary>
        /// Updates existing driver records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersUpdateDriverV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transporters/v2/drivers", body, queryParams);
        }

        /// <summary>
        /// Updates existing vehicle records for a facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersUpdateVehicleV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transporters/v2/vehicles", body, queryParams);
        }

    }
}
