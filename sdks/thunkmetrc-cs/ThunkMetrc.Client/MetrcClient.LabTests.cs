using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility recording the Lab Test.</param>
        public async Task CreateLabTestsRecord(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/labtests/v2/record", body, queryParams);
        }
        /// <summary>
        /// Retrieves a list of Lab Test batches.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetLabTestsBatches(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/batches", null, queryParams);
        }
        /// <summary>
        /// Retrieves a specific Lab Test result document by its Id for a given Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">LabTestDocumentFileId</param>
        public async Task<object> GetLabTestsLabTestDocumentById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/labtestdocument/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Returns a list of Lab Test types.
        /// </summary>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetLabTestsTypes(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/types", null, queryParams);
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
        public async Task<object> GetLabTestsResults(string? licenseNumber = null, string? packageId = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(packageId)) queryParams["packageId"] = packageId;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/results", null, queryParams);
        }
        /// <summary>
        /// Returns a list of all lab testing states.
        /// </summary>
        public async Task<object> GetLabTestsStates()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/states", null, queryParams);
        }
        /// <summary>
        /// Updates one or more documents for previously submitted lab tests.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility updating the Lab Test document.</param>
        public async Task UpdateLabTestsLabTestDocument(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v2/labtestdocument", body, queryParams);
        }
        /// <summary>
        /// Releases Lab Test results for one or more packages.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility releasing the Lab Test results.</param>
        public async Task UpdateLabTestsResultsRelease(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v2/results/release", body, queryParams);
        }
    }
}

