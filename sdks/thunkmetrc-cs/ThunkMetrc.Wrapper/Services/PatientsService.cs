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
    public class PatientsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Adds new patients to a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of patients.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreatePatients(List<PatientsCreatePatientsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreatePatients(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Removes a Patient, identified by an Id, from a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to delete the Patient.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeletePatientsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeletePatientsById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active patients for a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active patients.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Patient>> GetActivePatients(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActivePatients(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Patient>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Patient by Id.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Patient will be validated against the specified License Number. If not specified, the Patient will be validated against all of the User's current Facilities. Please note that if the Patient is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Patient> GetPatientsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPatientsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Patient>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates Patient information for a specified Facility.
        /// Permissions Required:
        /// - Manage Patients
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of patients.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdatePatients(List<PatientsUpdatePatientsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdatePatients(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

