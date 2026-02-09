using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to adjust the Processing Job.</param>
        public async Task CreateAdjustProcessingJob(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/adjust", body, queryParams);
        }
        /// <summary>
        /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Job Type.</param>
        public async Task CreateProcessingJobJobTypes(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/jobtypes", body, queryParams);
        }
        /// <summary>
        /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create the packages.</param>
        public async Task CreateProcessingJobPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/createpackages", body, queryParams);
        }
        /// <summary>
        /// Archives a Processing Job Type at a Facility, making it inactive for future use.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Processing Job Type.</param>
        public async Task DeleteProcessingJobJobTypeById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v2/jobtypes/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility of the Processing Job.</param>
        public async Task DeleteProcessingJobById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Completes processing jobs at a Facility by recording final notes and waste measurements.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to finish the Processing Job.</param>
        public async Task FinishProcessingJobProcessingJob(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/finish", body, queryParams);
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
        public async Task<object> GetProcessingJobActiveJobTypes(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/active", null, queryParams);
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
        public async Task<object> GetActiveProcessingJob(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/active", null, queryParams);
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
        public async Task<object> GetProcessingJobInactiveJobTypes(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/inactive", null, queryParams);
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
        public async Task<object> GetInactiveProcessingJob(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/inactive", null, queryParams);
        }
        /// <summary>
        /// Retrieves all processing job attributes available for a Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active processing job type attributes.</param>
        public async Task<object> GetProcessingJobJobTypesAttributes(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/attributes", null, queryParams);
        }
        /// <summary>
        /// Retrieves all processing job categories available for a specified Facility.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of processing job type categories.</param>
        public async Task<object> GetProcessingJobJobTypesCategories(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/categories", null, queryParams);
        }
        /// <summary>
        /// Retrieves a ProcessingJob by Id.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the processing job will be validated against the specified License Number. If not specified, the processing job will be validated against all of the User's current Facilities. Please note that if the processing job is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetProcessingJobById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Initiates new processing jobs at a Facility, including job details and associated packages.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to start the Processing Job.</param>
        public async Task StartProcessingJobProcessingJob(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/start", body, queryParams);
        }
        /// <summary>
        /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility to unfinish the Processing Job.</param>
        public async Task UnfinishProcessingJobProcessingJob(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/unfinish", body, queryParams);
        }
        /// <summary>
        /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        /// Permissions Required:
        /// - Manage Processing Job
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the processing job types.</param>
        public async Task UpdateProcessingJobJobTypes(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/jobtypes", body, queryParams);
        }
    }
}

