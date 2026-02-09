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
    public class ProcessingJobService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to adjust the Processing Job.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateAdjustProcessingJob(List<ProcessingJobCreateAdjustProcessingJobRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateAdjustProcessingJob(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Job Type.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateProcessingJobJobTypes(List<ProcessingJobCreateJobTypesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateProcessingJobJobTypes(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create the packages.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateProcessingJobPackages(List<ProcessingJobCreateProcessingJobPackagesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateProcessingJobPackages(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a Processing Job Type at a Facility, making it inactive for future use.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Processing Job Type.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteProcessingJobJobTypeById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteProcessingJobJobTypeById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Processing Job.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteProcessingJobById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteProcessingJobById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Completes processing jobs at a Facility by recording final notes and waste measurements.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to finish the Processing Job.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task FinishProcessingJobProcessingJob(List<ProcessingJobFinishProcessingJobRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.FinishProcessingJobProcessingJob(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a list of all active processing job types defined within a Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active processing job types.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ActiveJobType>> GetProcessingJobActiveJobTypes(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetProcessingJobActiveJobTypes(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ActiveJobType>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves active processing jobs at a specified Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active processing jobs.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ProcessingJob>> GetActiveProcessingJob(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetActiveProcessingJob(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ProcessingJob>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of all inactive processing job types defined within a Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive processing job types.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<InactiveJobType>> GetProcessingJobInactiveJobTypes(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetProcessingJobInactiveJobTypes(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<InactiveJobType>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves inactive processing jobs at a specified Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive processing jobs.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ProcessingJob>> GetInactiveProcessingJob(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetInactiveProcessingJob(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ProcessingJob>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves all processing job attributes available for a Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active processing job type attributes.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<JobTypesAttribute>> GetProcessingJobJobTypesAttributes(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetProcessingJobJobTypesAttributes(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<JobTypesAttribute>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves all processing job categories available for a specified Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of processing job type categories.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<JobTypesCategory>> GetProcessingJobJobTypesCategories(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetProcessingJobJobTypesCategories(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<JobTypesCategory>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a ProcessingJob by Id.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the processing job will be validated against the specified License Number. If not specified, the processing job will be validated against all of the User's current Facilities. Please note that if the processing job is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<ProcessingJob> GetProcessingJobById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetProcessingJobById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<ProcessingJob>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Initiates new processing jobs at a Facility, including job details and associated packages.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to start the Processing Job.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task StartProcessingJobProcessingJob(List<ProcessingJobStartProcessingJobRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StartProcessingJobProcessingJob(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to unfinish the Processing Job.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UnfinishProcessingJobProcessingJob(List<ProcessingJobUnfinishProcessingJobRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UnfinishProcessingJobProcessingJob(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the processing job types.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateProcessingJobJobTypes(List<ProcessingJobUpdateJobTypesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateProcessingJobJobTypes(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

