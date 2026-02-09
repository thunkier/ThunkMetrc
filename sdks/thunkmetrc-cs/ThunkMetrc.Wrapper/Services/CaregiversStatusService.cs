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
    public class CaregiversStatusService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// Permissions Required:
        /// - Lookup Caregivers
        /// </summary>
        /// <param name="caregiverLicenseNumber">caregiverLicenseNumber</param>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the Caregiver status.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<CaregiversStatus>> GetCaregiversStatusByCaregiverLicenseNumber(string caregiverLicenseNumber, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<CaregiversStatus>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
    }
}

