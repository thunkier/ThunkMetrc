using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using ThunkMetrc.Client;
using Microsoft.Extensions.Logging;
using System.Text.Json;
using System.Threading;
using System.Runtime.CompilerServices;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;
using File = ThunkMetrc.Wrapper.Models.File;

namespace ThunkMetrc.Wrapper.Services
{
    public class PatientCheckInsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Records patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of Patient check-ins.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePatientCheckIns(List<PatientCheckInsCreatePatientCheckInsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePatientCheckIns(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive Patient check-in.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeletePatientCheckInsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeletePatientCheckInsById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of Patient Check-In locations.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<PatientCheckInsLocation>> GetPatientCheckInsLocations(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPatientCheckInsLocations(), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<PatientCheckInsLocation>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="checkinDateEnd">The last modified end timestamp</param>
        /// <param name="checkinDateStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of Patient check-ins.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<PatientCheckIn>> GetPatientCheckIns(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPatientCheckIns(checkinDateEnd, checkinDateStart, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<PatientCheckIn>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates patient check-ins for a specified Facility.
        /// Permissions Required:
        /// - ManagePatientsCheckIns
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of Patient check-ins.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePatientCheckIns(List<PatientCheckInsUpdatePatientCheckInsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePatientCheckIns(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

