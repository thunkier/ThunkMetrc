using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreateAdjustV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v1/adjust", body, queryParams);
        }

        /// <summary>
        /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateAdjustV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/adjust", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateJobtypesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v1/jobtypes", body, queryParams);
        }

        /// <summary>
        /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateJobtypesV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/jobtypes", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreateStartV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v1/start", body, queryParams);
        }

        /// <summary>
        /// Initiates new processing jobs at a Facility, including job details and associated packages.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateStartV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/start", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreatepackagesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v1/createpackages", body, queryParams);
        }

        /// <summary>
        /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreatepackagesV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/processing/v2/createpackages", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteJobtypesV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v1/jobtypes/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a Processing Job Type at a Facility, making it inactive for future use.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteJobtypesV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/processing/v2/jobtypes/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a ProcessingJob by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves active processing jobs at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves inactive processing jobs at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/jobtypes/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of all active processing job types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesAttributesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/jobtypes/attributes", null, queryParams);
        }

        /// <summary>
        /// Retrieves all processing job attributes available for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesAttributesV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/attributes", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesCategoriesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/jobtypes/categories", null, queryParams);
        }

        /// <summary>
        /// Retrieves all processing job categories available for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesCategoriesV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v2/jobtypes/categories", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/processing/v1/jobtypes/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of all inactive processing job types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateFinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v1/finish", body, queryParams);
        }

        /// <summary>
        /// Completes processing jobs at a Facility by recording final notes and waste measurements.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateFinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/finish", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateJobtypesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v1/jobtypes", body, queryParams);
        }

        /// <summary>
        /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateJobtypesV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/jobtypes", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateUnfinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v1/unfinish", body, queryParams);
        }

        /// <summary>
        /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateUnfinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/processing/v2/unfinish", body, queryParams);
        }

    }
}
