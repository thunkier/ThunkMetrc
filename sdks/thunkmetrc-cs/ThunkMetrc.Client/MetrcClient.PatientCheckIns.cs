using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Records patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of Patient check-ins.</param>
        public async Task CreatePatientCheckIns(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/patient-checkins/v2", body, queryParams);
        }
        /// <summary>
        /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive Patient check-in.</param>
        public async Task DeletePatientCheckInsById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/patient-checkins/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of Patient Check-In locations.
        /// </summary>
        public async Task<object> GetPatientCheckInsLocations()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v2/locations", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="checkinDateEnd">The last modified end timestamp</param>
        /// <param name="checkinDateStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of Patient check-ins.</param>
        public async Task<object> GetPatientCheckIns(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(checkinDateEnd)) queryParams["checkinDateEnd"] = checkinDateEnd;
            if (!string.IsNullOrEmpty(checkinDateStart)) queryParams["checkinDateStart"] = checkinDateStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patient-checkins/v2", null, queryParams);
        }
        /// <summary>
        /// Updates patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Patient check-ins.</param>
        public async Task UpdatePatientCheckIns(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/patient-checkins/v2", body, queryParams);
        }
    }
}

