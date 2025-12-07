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
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patients/v2", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateAddV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patients/v1/add", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateUpdateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patients/v1/update", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patients/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Removes a Patient, identified by an Id, from a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patients/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Patient by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetActiveV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active patients for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetActiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/active", null, queryParams);
        }

        /// <summary>
        /// Updates Patient information for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/patients/v2", body, queryParams);
        }

    }
}
