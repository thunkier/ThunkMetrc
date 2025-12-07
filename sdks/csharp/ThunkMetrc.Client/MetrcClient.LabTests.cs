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
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsCreateRecordV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/labtests/v1/record", body, queryParams);
        }

        /// <summary>
        /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsCreateRecordV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/labtests/v2/record", body, queryParams);
        }

        /// <summary>
        /// Retrieves a list of Lab Test batches.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetBatchesV2(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/batches", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task<object> LabTestsGetLabtestdocumentV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v1/labtestdocument/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a specific Lab Test result document by its Id for a given Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task<object> LabTestsGetLabtestdocumentV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/labtestdocument/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> LabTestsGetResultsV1(string? licenseNumber = null, string? packageId = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(packageId)) queryParams["packageId"] = packageId;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v1/results", null, queryParams);
        }

        /// <summary>
        /// Retrieves Lab Test results for a specified Package.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task<object> LabTestsGetResultsV2(string? licenseNumber = null, string? packageId = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(packageId)) queryParams["packageId"] = packageId;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/results", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetStatesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v1/states", null, queryParams);
        }

        /// <summary>
        /// Returns a list of all lab testing states.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetStatesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/states", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetTypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v1/types", null, queryParams);
        }

        /// <summary>
        /// Returns a list of Lab Test types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/labtests/v2/types", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateLabtestdocumentV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v1/labtestdocument", body, queryParams);
        }

        /// <summary>
        /// Updates one or more documents for previously submitted lab tests.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateLabtestdocumentV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v2/labtestdocument", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateResultReleaseV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v1/results/release", body, queryParams);
        }

        /// <summary>
        /// Releases Lab Test results for one or more packages.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateResultReleaseV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/labtests/v2/results/release", body, queryParams);
        }

    }
}
