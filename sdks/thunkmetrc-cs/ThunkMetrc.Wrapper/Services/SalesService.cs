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
    public class SalesService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesDeliveries(List<SalesCreateDeliveriesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesDeliveries(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record depart delivery.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesDeliveriesRetailerDepart(List<SalesCreateDeliveriesRetailerDepartRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesDeliveriesRetailerDepart(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record end retailer deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesDeliveriesRetailerEnd(List<SalesCreateDeliveriesRetailerEndRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesDeliveriesRetailerEnd(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record restock retailer deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesDeliveriesRetailerRestock(List<SalesCreateDeliveriesRetailerRestockRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesDeliveriesRetailerRestock(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales (Write)
        /// - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
        /// - Industry/Facility Type/Advanced Sales
        /// - WebApi Sales Read Write State (All or WriteOnly)
        /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record receipts.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesReceipts(List<SalesCreateReceiptsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesReceipts(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to record retailer deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateSalesDeliveriesRetailer(List<SalesCreateSalesDeliveriesRetailerRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateSalesDeliveriesRetailer(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to void delivery.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteSalesDeliveryById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteSalesDeliveryById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to void retailer delivery.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteSalesDeliveryRetailerById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteSalesDeliveryRetailerById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive receipt.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteSalesReceiptById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteSalesReceiptById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active deliveries.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ActiveDelivery>> GetSalesActiveDeliveries(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesActiveDeliveries(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ActiveDelivery>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
        /// Permissions Required:
        /// - View Retailer Delivery
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return active retailer deliveries.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ActiveDeliveriesRetailer>> GetSalesActiveDeliveriesRetailer(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesActiveDeliveriesRetailer(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ActiveDeliveriesRetailer>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active receipts.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<ActiveReceipt>> GetSalesActiveReceipts(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesActiveReceipts(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<ActiveReceipt>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of counties available for sales deliveries.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<County>> GetSalesCounties(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesCounties(), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<County>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of customer types.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetSalesCustomerTypes(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesCustomerTypes(), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of return reasons for sales deliveries based on the provided License Number.
        /// Permissions Required:
        /// - Sales Delivery
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of delivery return reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<DeliveriesReturnReason>> GetSalesDeliveriesReturnReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesDeliveriesReturnReasons(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<DeliveriesReturnReason>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a retailer delivery record by its ID, with an optional License Number.
        /// Permissions Required:
        /// - View Retailer Delivery
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<DeliveryRetailer> GetSalesDeliveryRetailerById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesDeliveryRetailerById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<DeliveryRetailer>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive deliveries.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<InactiveDelivery>> GetSalesInactiveDeliveries(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesInactiveDeliveries(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<InactiveDelivery>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
        /// Permissions Required:
        /// - View Retailer Delivery
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return inactive retailer deliveries.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<InactiveDeliveriesRetailer>> GetSalesInactiveDeliveriesRetailer(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesInactiveDeliveriesRetailer(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<InactiveDeliveriesRetailer>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive receipts.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<InactiveReceipt>> GetSalesInactiveReceipts(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesInactiveReceipts(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<InactiveReceipt>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of valid Patient registration locations for sales.
        /// </summary>
        
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<PatientRegistrationLocation>> GetSalesPatientRegistrationLocations(CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesPatientRegistrationLocations(), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<PatientRegistrationLocation>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Returns a list of available payment types for the specified License Number.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of payment types.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetSalesPaymentTypes(string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesPaymentTypes(licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a sales receipt by its Id, with an optional License Number.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Receipt will be validated against the specified License Number. If not specified, the Sales Receipt will be validated against all of the User's current Facilities. Please note that if the Sales Receipt is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<ActiveReceipt> GetSalesReceiptById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesReceiptById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<ActiveReceipt>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Sales Receipt by its external number, with an optional License Number.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="externalNumber">externalNumber</param>
        /// <param name="licenseNumber">If specified, the External Sales Receipt Number will be validated against the specified License Number. If not specified, the External Sales Receipt Number will be validated against all of the User's current Facilities. Please note that if the External Sales Receipt Number is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<List<ReceiptsExternalByExternalNumber>> GetSalesReceiptsExternalByExternalNumber(string externalNumber, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesReceiptsExternalByExternalNumber(externalNumber, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<List<ReceiptsExternalByExternalNumber>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a sales delivery record by its Id, with an optional License Number.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<SalesDelivery> GetSalesDeliveryById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetSalesDeliveryById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<SalesDelivery>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveries(List<SalesUpdateDeliveriesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveries(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update deliveries completed.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesComplete(List<SalesUpdateDeliveriesCompleteRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesComplete(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales Delivery, Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub transporters.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesHub(List<SalesUpdateDeliveriesHubRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesHub(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub deliveries accepted.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesHubAccept(List<SalesUpdateDeliveriesHubAcceptRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesHubAccept(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub delivery departures.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesHubDepart(List<SalesUpdateDeliveriesHubDepartRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesHubDepart(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub delivery Id's.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesHubVerifyID(List<SalesUpdateDeliveriesHubVerifyIDRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesHubVerifyID(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        /// - Industry/Facility Type/Retailer Delivery
        /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update retailer deliveries.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesDeliveriesRetailer(List<SalesUpdateDeliveriesRetailerRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesDeliveriesRetailer(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update receipts.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesReceipts(List<SalesUpdateReceiptsRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesReceipts(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update finalized receipts.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesReceiptsFinalize(List<SalesUpdateReceiptsFinalizeRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesReceiptsFinalize(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update unfinalized receipts.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateSalesReceiptsUnfinalize(List<SalesUpdateReceiptsUnfinalizeRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateSalesReceiptsUnfinalize(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

