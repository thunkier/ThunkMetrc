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
    public class SublocationsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates new sublocation records for a Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of sublocations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSublocations(List<SublocationsCreateSublocationsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSublocations(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives an existing Sublocation record for a Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to delete the list of sublocations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteSublocationsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteSublocationsById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="lastModifiedEnd">The Last Modified end timestamp</param>
        /// <param name="lastModifiedStart">The Last Modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active sublocations.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Sublocation>> GetActiveSublocations(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveSublocations(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Sublocation>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive sublocations for the specified Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive sublocations.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Sublocation>> GetInactiveSublocations(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveSublocations(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Sublocation>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Sublocation by its Id, with an optional license number.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sublocation will be validated against the specified License Number. If not specified, the Sublocation will be validated against all of the User's current Facilities. Please note that if the Sublocation is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<Sublocation> GetSublocationsById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSublocationsById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<Sublocation>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates existing sublocation records for a specified Facility.
        /// Permissions Required:
        /// - Manage Locations
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of sublocations.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSublocations(List<SublocationsUpdateSublocationsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSublocations(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

