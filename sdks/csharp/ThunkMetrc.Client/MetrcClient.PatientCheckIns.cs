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
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsCreateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patient-checkins/v1", body, queryParams);
        }

        /// <summary>
        /// Records patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patient-checkins/v2", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patient-checkins/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patient-checkins/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task<object> PatientCheckInsGetAllV1(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(checkinDateEnd)) queryParams["checkinDateEnd"] = checkinDateEnd;
            if (!string.IsNullOrEmpty(checkinDateStart)) queryParams["checkinDateStart"] = checkinDateStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v1", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task<object> PatientCheckInsGetAllV2(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(checkinDateEnd)) queryParams["checkinDateEnd"] = checkinDateEnd;
            if (!string.IsNullOrEmpty(checkinDateStart)) queryParams["checkinDateStart"] = checkinDateStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v2", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PatientCheckInsGetLocationsV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v1/locations", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of Patient Check-In locations.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PatientCheckInsGetLocationsV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v2/locations", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsUpdateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/patient-checkins/v1", body, queryParams);
        }

        /// <summary>
        /// Updates patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsUpdateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/patient-checkins/v2", body, queryParams);
        }

    }
}
