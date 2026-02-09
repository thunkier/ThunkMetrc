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
    public class AdditivesTemplatesService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates new additive templates for a specified Facility.
        /// Permissions Required:
        /// - Manage Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateAdditivesTemplates(List<AdditivesTemplatesCreateAdditivesTemplatesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateAdditivesTemplates(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of active additive templates for a specified Facility.
        /// Permissions Required:
        /// - Manage Additives
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active additives.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<AdditivesTemplate>> GetActiveAdditivesTemplates(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveAdditivesTemplates(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<AdditivesTemplate>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves an Additive Template by its Id.
        /// Permissions Required:
        /// - Manage Additives
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Additive will be validated against the specified License Number. If not specified, the Additive will be validated against all of the User's current Facilities. Please note that if the Additive is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<AdditivesTemplate> GetAdditivesTemplatesById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetAdditivesTemplatesById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<AdditivesTemplate>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of inactive additive templates for a specified Facility.
        /// Permissions Required:
        /// - Manage Additives
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive additives.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<AdditivesTemplate>> GetInactiveAdditivesTemplates(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveAdditivesTemplates(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<AdditivesTemplate>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates existing additive templates for a specified Facility.
        /// Permissions Required:
        /// - Manage Additives
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of additives.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateAdditivesTemplates(List<AdditivesTemplatesUpdateAdditivesTemplatesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateAdditivesTemplates(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

