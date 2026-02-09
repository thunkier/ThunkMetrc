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
    public class LabTestsService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility recording the Lab Test.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateLabTestsRecord(List<LabTestsCreateRecordRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateLabTestsRecord(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of Lab Test batches.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Batch>> GetLabTestsBatches(string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetLabTestsBatches(pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Batch>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a specific Lab Test result document by its Id for a given Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">LabTestDocumentFileId</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetLabTestsLabTestDocumentById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetLabTestsLabTestDocumentById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of Lab Test types.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<LabTestsType>> GetLabTestsTypes(string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetLabTestsTypes(pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<LabTestsType>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves Lab Test results for a specified Package.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="licenseNumber">Optional query parameter.</param>
        /// <param name="packageId">Optional query parameter.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Result>> GetLabTestsResults(string? licenseNumber = null, string? packageId = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetLabTestsResults(licenseNumber, packageId, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Result>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of all lab testing states.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetLabTestsStates(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetLabTestsStates(), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates one or more documents for previously submitted lab tests.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility updating the Lab Test document.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateLabTestsLabTestDocument(List<LabTestsUpdateLabTestDocumentRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateLabTestsLabTestDocument(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Releases Lab Test results for one or more packages.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility releasing the Lab Test results.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateLabTestsResultsRelease(List<LabTestsUpdateResultsReleaseRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateLabTestsResultsRelease(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

