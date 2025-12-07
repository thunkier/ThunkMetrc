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
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/create", body, queryParams);
        }

        /// <summary>
        /// Creates new packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateAdjustV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/adjust", body, queryParams);
        }

        /// <summary>
        /// Records a list of adjustments for packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateAdjustV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/adjust", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateChangeItemV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/change/item", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateChangeLocationV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/change/locations", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateFinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/finish", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreatePlantingsV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/create/plantings", body, queryParams);
        }

        /// <summary>
        /// Creates new plantings from packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreatePlantingsV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/plantings", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateRemediateV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/remediate", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateTestingV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/create/testing", body, queryParams);
        }

        /// <summary>
        /// Creates new packages for testing for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateTestingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v2/testing", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateUnfinishV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/packages/v1/unfinish", body, queryParams);
        }

        /// <summary>
        /// Discontinues a Package at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesDeleteV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/packages/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Package by its Id.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of active packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetAdjustReasonsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/adjust/reasons", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of adjustment reasons for packages at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetAdjustReasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/adjust/reasons", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetByLabelV1(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/{Uri.EscapeDataString(label)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Package by its label.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetByLabelV2(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(label)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/inactive", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of inactive packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of packages in transit for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetIntransitV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetLabsamplesV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/onhold", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of packages on hold for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves the source harvests for a Package by its Id.
        /// 
        ///   Permissions Required:
        ///   - View Package Source Harvests
        /// </summary>
        public async Task<object> PackagesGetSourceHarvestV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/{Uri.EscapeDataString(id)}/source/harvests", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transferred packages for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetTransferredV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetTypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v1/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of available Package types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetTypesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/packages/v2/types", null, queryParams);
        }

        /// <summary>
        /// Set the final quantity for a Package.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateAdjustV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/adjust", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        ///   - Manage Package Notes
        /// </summary>
        public async Task PackagesUpdateChangeNoteV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v1/change/note", body, queryParams);
        }

        /// <summary>
        /// Updates the Product decontaminate information for a list of packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDecontaminateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/decontaminate", body, queryParams);
        }

        /// <summary>
        /// Flags one or more packages for donation at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDonationFlagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/donation/flag", body, queryParams);
        }

        /// <summary>
        /// Removes the donation flag from one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDonationUnflagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/donation/unflag", body, queryParams);
        }

        /// <summary>
        /// Updates the external identifiers for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Package Inventory
        ///   - External Id Enabled
        /// </summary>
        public async Task PackagesUpdateExternalidV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/externalid", body, queryParams);
        }

        /// <summary>
        /// Updates a list of packages as finished for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateFinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/finish", body, queryParams);
        }

        /// <summary>
        /// Updates the associated Item for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateItemV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/item", body, queryParams);
        }

        /// <summary>
        /// Updates the list of required lab test batches for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateLabTestRequiredV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/labtests/required", body, queryParams);
        }

        /// <summary>
        /// Updates the Location and Sublocation for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateLocationV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/location", body, queryParams);
        }

        /// <summary>
        /// Updates notes associated with one or more packages for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        ///   - Manage Package Notes
        /// </summary>
        public async Task PackagesUpdateNoteV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/note", body, queryParams);
        }

        /// <summary>
        /// Updates a list of Product remediations for packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateRemediateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/remediate", body, queryParams);
        }

        /// <summary>
        /// Flags or unflags one or more packages at the specified Facility as trade samples.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateTradesampleFlagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/tradesample/flag", body, queryParams);
        }

        /// <summary>
        /// Removes the trade sample flag from one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateTradesampleUnflagV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/tradesample/unflag", body, queryParams);
        }

        /// <summary>
        /// Updates a list of packages as unfinished for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateUnfinishV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/unfinish", body, queryParams);
        }

        /// <summary>
        /// Updates the use-by date for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateUsebydateV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/packages/v2/usebydate", body, queryParams);
        }

    }
}
