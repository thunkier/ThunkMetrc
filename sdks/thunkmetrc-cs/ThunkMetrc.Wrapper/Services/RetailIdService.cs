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
    public class RetailIdService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateRetailIdAssociate(List<RetailIdCreateAssociateRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateRetailIdAssociate(body, licenseNumber); return null; }, cancellationToken);
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateRetailIdGenerate(RetailIdCreateGenerateRequest? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateRetailIdGenerate(body, licenseNumber); return null; }, cancellationToken);
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateRetailIdMerge(RetailIdCreateMergeRequest? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateRetailIdMerge(body, licenseNumber); return null; }, cancellationToken);
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateRetailIdPackagesInfo(RetailIdCreatePackagesInfoRequest? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateRetailIdPackagesInfo(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves the available Retail Item ID quota for a facility.
        /// Permissions Required:
        /// - Download Product Labels
        /// - Manage Product Labels
        /// - Manage Tag Orders
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to retrieve available Retail Item ID quota.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Allotment>> GetRetailIdAllotment(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetRetailIdAllotment(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Allotment>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<Receive>> GetRetailIdReceiveByLabel(string label, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetRetailIdReceiveByLabel(label, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<Receive>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
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
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<ReceiveQrByShortCode>> GetRetailIdReceiveQrByShortCode(string shortCode, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetRetailIdReceiveQrByShortCode(shortCode, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<ReceiveQrByShortCode>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
    }
}

