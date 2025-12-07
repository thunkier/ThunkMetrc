using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesCreateDeliveryV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries", body, queryParams);
        }

        /// <summary>
        /// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// </summary>
        public async Task SalesCreateDeliveryV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries", body, queryParams);
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries/retailer", body, queryParams);
        }

        /// <summary>
        /// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerDepartV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries/retailer/depart", body, queryParams);
        }

        /// <summary>
        /// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerDepartV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/depart", body, queryParams);
        }

        /// <summary>
        /// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerEndV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries/retailer/end", body, queryParams);
        }

        /// <summary>
        /// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerEndV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/end", body, queryParams);
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerRestockV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries/retailer/restock", body, queryParams);
        }

        /// <summary>
        /// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerRestockV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/restock", body, queryParams);
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerSaleV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/deliveries/retailer/sale", body, queryParams);
        }

        /// <summary>
        /// Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// </summary>
        public async Task SalesCreateDeliveryRetailerSaleV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/deliveries/retailer/sale", body, queryParams);
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesCreateReceiptV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/receipts", body, queryParams);
        }

        /// <summary>
        /// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales (Write)
        ///   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
        ///   - Industry/Facility Type/Advanced Sales
        ///   - WebApi Sales Read Write State (All or WriteOnly)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        /// </summary>
        public async Task SalesCreateReceiptV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v2/receipts", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesCreateTransactionByDateV1(string date, object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/sales/v1/transactions/{Uri.EscapeDataString(date)}", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v1/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryRetailerV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v1/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryRetailerV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesDeleteReceiptV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v1/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesDeleteReceiptV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/sales/v2/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCountiesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/counties", null, queryParams);
        }

        /// <summary>
        /// Returns a list of counties available for sales deliveries.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCountiesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/counties", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCustomertypesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/customertypes", null, queryParams);
        }

        /// <summary>
        /// Returns a list of customer types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCustomertypesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/customertypes", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/active", null, queryParams);
        }

        /// <summary>
        /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        /// 
        ///   Permissions Required:
        ///   - View Sales Delivery
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/active", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/inactive", null, queryParams);
        }

        /// <summary>
        /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        /// 
        ///   Permissions Required:
        ///   - View Sales Delivery
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/inactive", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/retailer/active", null, queryParams);
        }

        /// <summary>
        /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
        /// 
        ///   Permissions Required:
        ///   - View Retailer Delivery
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/retailer/inactive", null, queryParams);
        }

        /// <summary>
        /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
        /// 
        ///   Permissions Required:
        ///   - View Retailer Delivery
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetDeliveriesReturnreasonsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/returnreasons", null, queryParams);
        }

        /// <summary>
        /// Returns a list of return reasons for sales deliveries based on the provided License Number.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesReturnreasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/returnreasons", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a sales delivery record by its Id, with an optional License Number.
        /// 
        ///   Permissions Required:
        ///   - View Sales Delivery
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryRetailerV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a retailer delivery record by its ID, with an optional License Number.
        /// 
        ///   Permissions Required:
        ///   - View Retailer Delivery
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryRetailerV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/deliveries/retailer/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetPatientRegistrationsLocationsV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/patientregistration/locations", null, queryParams);
        }

        /// <summary>
        /// Returns a list of valid Patient registration locations for sales.
        /// 
        ///   Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetPatientRegistrationsLocationsV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/patientregistration/locations", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetPaymenttypesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/paymenttypes", null, queryParams);
        }

        /// <summary>
        /// Returns a list of available payment types for the specified License Number.
        /// 
        ///   Permissions Required:
        ///   - View Sales Delivery
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task<object> SalesGetPaymenttypesV2(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/paymenttypes", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a sales receipt by its Id, with an optional License Number.
        /// 
        ///   Permissions Required:
        ///   - View Sales
        ///   - Manage Sales
        /// </summary>
        public async Task<object> SalesGetReceiptV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/receipts/active", null, queryParams);
        }

        /// <summary>
        /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        /// 
        ///   Permissions Required:
        ///   - View Sales
        ///   - Manage Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves a Sales Receipt by its external number, with an optional License Number.
        /// 
        ///   Permissions Required:
        ///   - View Sales
        ///   - Manage Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsExternalByExternalNumberV2(string externalNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/external/{Uri.EscapeDataString(externalNumber)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/receipts/inactive", null, queryParams);
        }

        /// <summary>
        /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        /// 
        ///   Permissions Required:
        ///   - View Sales
        ///   - Manage Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            if (!string.IsNullOrEmpty(salesDateEnd)) queryParams["salesDateEnd"] = salesDateEnd;
            if (!string.IsNullOrEmpty(salesDateStart)) queryParams["salesDateStart"] = salesDateStart;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v2/receipts/inactive", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetTransactionsV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/transactions", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(string salesDateStart, string salesDateEnd, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/sales/v1/transactions/{Uri.EscapeDataString(salesDateStart)}/{Uri.EscapeDataString(salesDateEnd)}", null, queryParams);
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries", body, queryParams);
        }

        /// <summary>
        /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryCompleteV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/complete", body, queryParams);
        }

        /// <summary>
        /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryCompleteV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/complete", body, queryParams);
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryHubV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/hub", body, queryParams);
        }

        /// <summary>
        /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery, Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubAcceptV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/hub/accept", body, queryParams);
        }

        /// <summary>
        /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubAcceptV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/accept", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubDepartV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/hub/depart", body, queryParams);
        }

        /// <summary>
        /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubDepartV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/depart", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubVerifyIdV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/hub/verifyID", body, queryParams);
        }

        /// <summary>
        /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubVerifyIdV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/hub/verifyID", body, queryParams);
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryRetailerV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/deliveries/retailer", body, queryParams);
        }

        /// <summary>
        /// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        ///   - Industry/Facility Type/Retailer Delivery
        ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        ///   - Manage Retailer Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryRetailerV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/deliveries/retailer", body, queryParams);
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateReceiptV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/receipts", body, queryParams);
        }

        /// <summary>
        /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts", body, queryParams);
        }

        /// <summary>
        /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptFinalizeV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts/finalize", body, queryParams);
        }

        /// <summary>
        /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptUnfinalizeV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v2/receipts/unfinalize", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateTransactionByDateV1(string date, object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/sales/v1/transactions/{Uri.EscapeDataString(date)}", body, queryParams);
        }

    }
}
