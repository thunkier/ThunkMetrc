using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
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
        public async Task CreateSalesDeliveries(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries", body, queryParams);
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
        public async Task CreateSalesDeliveriesRetailerDepart(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/depart", body, queryParams);
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
        public async Task CreateSalesDeliveriesRetailerEnd(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/end", body, queryParams);
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
        public async Task CreateSalesDeliveriesRetailerRestock(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/restock", body, queryParams);
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
        public async Task CreateSalesReceipts(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/receipts", body, queryParams);
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
        public async Task CreateSalesDeliveriesRetailer(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer", body, queryParams);
        }
        /// <summary>
        /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to void delivery.</param>
        public async Task DeleteSalesDeliveryById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
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
        public async Task DeleteSalesDeliveryRetailerById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive receipt.</param>
        public async Task DeleteSalesReceiptById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
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
        public async Task<object> GetSalesActiveDeliveries(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/active", null, queryParams);
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
        public async Task<object> GetSalesActiveDeliveriesRetailer(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/retailer/active", null, queryParams);
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
        public async Task<object> GetSalesActiveReceipts(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/active", null, queryParams);
        }
        /// <summary>
        /// Returns a list of counties available for sales deliveries.
        /// </summary>
        public async Task<object> GetSalesCounties()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/counties", null, queryParams);
        }
        /// <summary>
        /// Returns a list of customer types.
        /// </summary>
        public async Task<object> GetSalesCustomerTypes()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/customertypes", null, queryParams);
        }
        /// <summary>
        /// Returns a list of return reasons for sales deliveries based on the provided License Number.
        /// Permissions Required:
        /// - Sales Delivery
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of delivery return reasons.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetSalesDeliveriesReturnReasons(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/returnreasons", null, queryParams);
        }
        /// <summary>
        /// Retrieves a retailer delivery record by its ID, with an optional License Number.
        /// Permissions Required:
        /// - View Retailer Delivery
        /// - Manage Retailer Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetSalesDeliveryRetailerById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
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
        public async Task<object> GetSalesInactiveDeliveries(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/inactive", null, queryParams);
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
        public async Task<object> GetSalesInactiveDeliveriesRetailer(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/retailer/inactive", null, queryParams);
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
        public async Task<object> GetSalesInactiveReceipts(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/inactive", null, queryParams);
        }
        /// <summary>
        /// Returns a list of valid Patient registration locations for sales.
        /// </summary>
        public async Task<object> GetSalesPatientRegistrationLocations()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/patientregistration/locations", null, queryParams);
        }
        /// <summary>
        /// Returns a list of available payment types for the specified License Number.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of payment types.</param>
        public async Task<object> GetSalesPaymentTypes(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/paymenttypes", null, queryParams);
        }
        /// <summary>
        /// Retrieves a sales receipt by its Id, with an optional License Number.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Receipt will be validated against the specified License Number. If not specified, the Sales Receipt will be validated against all of the User's current Facilities. Please note that if the Sales Receipt is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetSalesReceiptById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a Sales Receipt by its external number, with an optional License Number.
        /// Permissions Required:
        /// - View Sales
        /// - Manage Sales
        /// </summary>
        /// <param name="externalNumber">externalNumber</param>
        /// <param name="licenseNumber">If specified, the External Sales Receipt Number will be validated against the specified License Number. If not specified, the External Sales Receipt Number will be validated against all of the User's current Facilities. Please note that if the External Sales Receipt Number is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetSalesReceiptsExternalByExternalNumber(string externalNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/external/{Uri.EscapeDataString(externalNumber)}", null, queryParams);
        }
        /// <summary>
        /// Retrieves a sales delivery record by its Id, with an optional License Number.
        /// Permissions Required:
        /// - View Sales Delivery
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">If specified, the Sales Delivery will be validated against the specified License Number. If not specified, the Sales Delivery will be validated against all of the User's current Facilities. Please note that if the Sales Delivery is not valid for the specified License Number, a 401 Unauthorized status will be returned.</param>
        public async Task<object> GetSalesDeliveryById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update deliveries.</param>
        public async Task UpdateSalesDeliveries(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries", body, queryParams);
        }
        /// <summary>
        /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update deliveries completed.</param>
        public async Task UpdateSalesDeliveriesComplete(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/complete", body, queryParams);
        }
        /// <summary>
        /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales Delivery, Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub transporters.</param>
        public async Task UpdateSalesDeliveriesHub(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub", body, queryParams);
        }
        /// <summary>
        /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub deliveries accepted.</param>
        public async Task UpdateSalesDeliveriesHubAccept(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/accept", body, queryParams);
        }
        /// <summary>
        /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub delivery departures.</param>
        public async Task UpdateSalesDeliveriesHubDepart(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/depart", body, queryParams);
        }
        /// <summary>
        /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        /// Permissions Required:
        /// - Manage Sales Delivery Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update hub delivery Id's.</param>
        public async Task UpdateSalesDeliveriesHubVerifyID(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/verifyID", body, queryParams);
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
        public async Task UpdateSalesDeliveriesRetailer(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/retailer", body, queryParams);
        }
        /// <summary>
        /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update receipts.</param>
        public async Task UpdateSalesReceipts(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts", body, queryParams);
        }
        /// <summary>
        /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update finalized receipts.</param>
        public async Task UpdateSalesReceiptsFinalize(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts/finalize", body, queryParams);
        }
        /// <summary>
        /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// Permissions Required:
        /// - Manage Sales
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update unfinalized receipts.</param>
        public async Task UpdateSalesReceiptsUnfinalize(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts/unfinalize", body, queryParams);
        }
    }
}

