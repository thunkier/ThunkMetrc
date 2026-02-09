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
    public class StrainsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates new strain records for a specified Facility.
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of strains.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateStrains(List<StrainsCreateStrainsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateStrains(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives an existing strain record for a Facility
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which Strain to delete.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteStrainsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteStrainsById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active strains.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Strain>> GetActiveStrains(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveStrains(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Strain>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive strains.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Strain>> GetInactiveStrains(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveStrains(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Strain>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Strain record by its Id, with an optional license number.
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Strain will be validated against the specified License Number. If not specified, the Strain will be validated against all of the User's current Facilities. Please note that if the Strain is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Strain> GetStrainsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetStrainsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Strain>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates existing strain records for a specified Facility.
        /// Permissions Required:
        /// - Manage Strains
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of strains.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateStrains(List<StrainsUpdateStrainsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateStrains(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

