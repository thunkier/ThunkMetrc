using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Records a list of adjustments for packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of adjustments.</param>
        public async Task CreateAdjustPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/adjust", body, queryParams);
        }
        /// <summary>
        /// Creates new packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new packages.</param>
        public async Task CreatePackagesPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2", body, queryParams);
        }
        /// <summary>
        /// Creates new plantings from packages for a specified Facility.
        /// Permissions Required:
        /// - View Immature Plants
        /// - Manage Immature Plants
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new plantings.</param>
        public async Task CreatePackagesPlantings(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/plantings", body, queryParams);
        }
        /// <summary>
        /// Creates new packages for testing for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of new packages for testing.</param>
        public async Task CreatePackagesTesting(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/testing", body, queryParams);
        }
        /// <summary>
        /// Discontinues a Package at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to discontinue.</param>
        public async Task DeletePackagesById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/packages/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Updates a list of packages as finished for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of finish packages.</param>
        public async Task FinishPackagesPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/finish", body, queryParams);
        }
        /// <summary>
        /// Flags one or more Packages at the specified Facility as Finished Goods.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to flag finished goods.</param>
        public async Task FinishedgoodFlagPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/finishedgood/flag", body, queryParams);
        }
        /// <summary>
        /// Removes the Finished Good flag one or more Packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to flag finished goods.</param>
        public async Task FinishedgoodUnflagPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/finishedgood/unflag", body, queryParams);
        }
        /// <summary>
        /// Retrieves a list of active packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetActivePackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/active", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of adjustment reasons for packages at a specified Facility.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of adjustment reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPackagesAdjustReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/adjust/reasons", null, queryParams);
        }
        /// <summary>
        /// Retrieves the Package Adjustments for a Facility
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of Package Adjustments.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPackagesAdjustments(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/adjustments", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of packages in transit for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of packages in transit.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetInTransitPackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/intransit", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of inactive packages for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetInactivePackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/inactive", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of lab sample packages that have been created/sent for testing.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPackagesLabSamples(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/labsamples", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of packages on hold for a specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of packages on hold.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetOnHoldPackages(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/onhold", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Package by its Id.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPackagesById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Package by its label.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="label">label</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPackagesByLabel(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(label)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of available Package types.
        /// </summary>
        public async Task<object> GetPackagesTypes()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/types", null, queryParams);
        }
        /// <summary>
        /// Retrieves the source harvests for a Package by its Id.
        /// Permissions Required:
        /// - View Package Source Harvests
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities. Please note that if the Package is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetPackagesSourceHarvestById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(id)}/source/harvests", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transferred packages for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of transferred packages.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetPackagesTransferred(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/transferred", null, queryParams);
        }
        /// <summary>
        /// Updates a list of packages as unfinished for a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of unfinish packages.</param>
        public async Task UnfinishPackagesPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/unfinish", body, queryParams);
        }
        /// <summary>
        /// Set the final quantity for a Package.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record the list of adjustments.</param>
        public async Task UpdateAdjustPackages(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/adjust", body, queryParams);
        }
        /// <summary>
        /// Updates the Product decontaminate information for a list of packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of product decontaminations.</param>
        public async Task UpdatePackagesDecontaminate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/decontaminate", body, queryParams);
        }
        /// <summary>
        /// Flags one or more packages for donation at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of flagged donations.</param>
        public async Task UpdatePackagesDonationFlag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/donation/flag", body, queryParams);
        }
        /// <summary>
        /// Removes the donation flag from one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of unflaged donations.</param>
        public async Task UpdatePackagesDonationUnflag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/donation/unflag", body, queryParams);
        }
        /// <summary>
        /// Updates the external identifiers for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Package Inventory
        /// - External Id Enabled
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of change external Ids.</param>
        public async Task UpdatePackagesExternalid(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/externalid", body, queryParams);
        }
        /// <summary>
        /// Updates the associated Item for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of changed items.</param>
        public async Task UpdatePackagesItem(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/item", body, queryParams);
        }
        /// <summary>
        /// Updates the list of required lab test batches for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of required lab test batches.</param>
        public async Task UpdatePackagesLabtestsRequired(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/labtests/required", body, queryParams);
        }
        /// <summary>
        /// Updates notes associated with one or more packages for the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// - Manage Package Notes
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of change notes.</param>
        public async Task UpdatePackagesNote(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/note", body, queryParams);
        }
        /// <summary>
        /// Updates the Location and Sublocation for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of change locations.</param>
        public async Task UpdatePackagesLocation(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/location", body, queryParams);
        }
        /// <summary>
        /// Updates a list of Product remediations for packages at a specific Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of product remediations.</param>
        public async Task UpdatePackagesRemediate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/remediate", body, queryParams);
        }
        /// <summary>
        /// Flags or unflags one or more packages at the specified Facility as trade samples.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update trade sample flags.</param>
        public async Task UpdatePackagesTradeSampleFlag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/tradesample/flag", body, queryParams);
        }
        /// <summary>
        /// Removes the trade sample flag from one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Manage Packages Inventory
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update trade sample unflag.</param>
        public async Task UpdatePackagesTradeSampleUnflag(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/tradesample/unflag", body, queryParams);
        }
        /// <summary>
        /// Updates the use-by date for one or more packages at the specified Facility.
        /// Permissions Required:
        /// - View Packages
        /// - Create/Submit/Discontinue Packages
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update list of changed items.</param>
        public async Task UpdatePackagesUseByDate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/usebydate", body, queryParams);
        }
    }
}

