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
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a Driver.</param>
        public async Task CreateTransportersDrivers(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transporters/v2/drivers", body, queryParams);
        }
        /// <summary>
        /// Creates new vehicle records for a Facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a Vehicle.</param>
        public async Task CreateTransportersVehicles(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transporters/v2/vehicles", body, queryParams);
        }
        /// <summary>
        /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive a Driver.</param>
        public async Task DeleteTransportersDriverById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transporters/v2/drivers/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number for which to archive a Vehicle</param>
        public async Task DeleteTransportersVehicleById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transporters/v2/vehicles/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a Driver.</param>
        public async Task<object> GetTransportersDriverById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/drivers/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of drivers for a Facility.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of drivers.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransportersDrivers(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/drivers", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a Vehicle.</param>
        public async Task<object> GetTransportersVehicleById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/vehicles/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of vehicles for a Facility.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of vehicles.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransportersVehicles(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transporters/v2/vehicles", null, queryParams);
        }
        /// <summary>
        /// Updates existing driver records for a Facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of drivers.</param>
        public async Task UpdateTransportersDrivers(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transporters/v2/drivers", body, queryParams);
        }
        /// <summary>
        /// Updates existing vehicle records for a facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of vehicles.</param>
        public async Task UpdateTransportersVehicles(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transporters/v2/vehicles", body, queryParams);
        }
    }
}

