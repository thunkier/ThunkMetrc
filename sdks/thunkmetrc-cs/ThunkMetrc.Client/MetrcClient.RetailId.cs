using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        /// - WebApi Retail ID Read Write State (All or WriteOnly)
        /// - Industry/View Packages
        /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to associate the packages and QR codes provided.</param>
        public async Task CreateRetailIdAssociate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/retailid/v2/associate", body, queryParams);
        }
        /// <summary>
        /// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        /// - WebApi Retail ID Read Write State (All or WriteOnly)
        /// - Industry/View Packages
        /// - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to associate the packages and QR codes provided.</param>
        public async Task CreateRetailIdGenerate(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/retailid/v2/generate", body, queryParams);
        }
        /// <summary>
        /// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        /// - WebApi Retail ID Read Write State (All or WriteOnly)
        /// - Key Value Settings/Retail ID Merge Packages Enabled
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility.</param>
        public async Task CreateRetailIdMerge(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/retailid/v2/merge", body, queryParams);
        }
        /// <summary>
        /// Retrieves Package information for given list of Package labels.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        /// - WebApi Retail ID Read Write State (All or WriteOnly)
        /// - Industry/View Packages
        /// - Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility</param>
        public async Task CreateRetailIdPackagesInfo(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/retailid/v2/packages/info", body, queryParams);
        }
        /// <summary>
        /// Retrieves the available Retail Item ID quota for a facility.
        /// Permissions Required:
        /// - Download Product Labels
        /// - Manage Product Labels
        /// - Manage Tag Orders
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to retrieve available Retail Item ID quota.</param>
        public async Task<object> GetRetailIdAllotment(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/retailid/v2/allotment", null, queryParams);
        }
        /// <summary>
        /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Manage RetailId
        /// - WebApi Retail ID Read Write State (All or ReadOnly)
        /// - Industry/View Packages
        /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        /// </summary>
        /// <param name="label">label</param>
        /// <param name="licenseNumber">If specified, the Package will be validated against the specified License Number. If not specified, the Package will be validated against all of the User's current Facilities.</param>
        public async Task<object> GetRetailIdReceiveByLabel(string label, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/retailid/v2/receive/{Uri.EscapeDataString(label)}", null, queryParams);
        }
        /// <summary>
        /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Manage RetailId
        /// - WebApi Retail ID Read Write State (All or ReadOnly)
        /// - Industry/View Packages
        /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        /// </summary>
        /// <param name="shortCode">shortCode</param>
        /// <param name="licenseNumber">The License Number of the Facility</param>
        public async Task<object> GetRetailIdReceiveQrByShortCode(string shortCode, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/retailid/v2/receive/qr/{Uri.EscapeDataString(shortCode)}", null, queryParams);
        }
    }
}

