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
    public class PatientsStatusService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// Permissions Required:
        /// - Lookup Patients
        /// </summary>
        /// <param name="patientLicenseNumber">patientLicenseNumber</param>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the Patient Status.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<PatientsStatus>> GetPatientsStatusesByPatientLicenseNumber(string patientLicenseNumber, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetPatientsStatusesByPatientLicenseNumber(patientLicenseNumber, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<PatientsStatus>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
    }
}

