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
    public class TagsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Returns a list of available package tags. NOTE: This is a premium endpoint.
        /// Permissions Required:
        /// - Manage Tags
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return available Package tags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Tag>> GetTagsAvailablePackage(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTagsAvailablePackage(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Tag>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
        /// Permissions Required:
        /// - Manage Tags
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return available tags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Tag>> GetTagsAvailablePlant(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTagsAvailablePlant(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Tag>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of staged tags. NOTE: This is a premium endpoint.
        /// Permissions Required:
        /// - Manage Tags
        /// - RetailId.AllowPackageStaging Key Value enabled
        /// </summary>
        /// <param name="licenseNumber">The License Number for which to return staged tags.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Staged>> GetStagedTags(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetStagedTags(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Staged>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
    }
}

