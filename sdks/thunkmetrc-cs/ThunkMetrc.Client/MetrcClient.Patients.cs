using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Adds new patients to a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of patients.</param>
        public async Task CreatePatients(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patients/v2", body, queryParams);
        }
        /// <summary>
        /// Removes a Patient, identified by an Id, from a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to delete the Patient.</param>
        public async Task DeletePatientsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patients/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of active patients for a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active patients.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetActivePatients(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/active", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Patient by Id.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Patient will be validated against the specified License Number. If not specified, the Patient will be validated against all of the User's current Facilities. Please note that if the Patient is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPatientsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Updates Patient information for a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of patients.</param>
        public async Task UpdatePatients(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/patients/v2", body, queryParams);
        }
    }
}

