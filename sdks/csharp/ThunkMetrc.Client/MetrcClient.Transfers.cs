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
        ///   - Transfers
        /// </summary>
        public async Task TransfersCreateExternalIncomingV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v1/external/incoming", body, queryParams);
        }

        /// <summary>
        /// Creates external incoming shipment plans for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersCreateExternalIncomingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/external/incoming", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersCreateTemplatesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v1/templates", body, queryParams);
        }

        /// <summary>
        /// Creates new transfer templates for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersCreateTemplatesOutgoingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/templates/outgoing", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task TransfersDeleteExternalIncomingV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v1/external/incoming/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Voids an external incoming shipment plan for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersDeleteExternalIncomingV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v2/external/incoming/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersDeleteTemplatesV1(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v1/templates/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Archives a transfer template for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersDeleteTemplatesOutgoingV2(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v2/templates/outgoing/{Uri.EscapeDataString(id)}", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetDeliveriesPackagesStatesV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/packages/states", null, queryParams);
        }

        /// <summary>
        /// Returns a list of available shipment Package states.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetDeliveriesPackagesStatesV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/packages/states", null, queryParams);
        }

        /// <summary>
        /// Please note: that the {id} parameter above represents a Shipment Plan ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageRequiredlabtestbatchesV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/package/{Uri.EscapeDataString(id)}/requiredlabtestbatches", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageRequiredlabtestbatchesV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/package/{Uri.EscapeDataString(id)}/requiredlabtestbatches", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageWholesaleV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/{Uri.EscapeDataString(id)}/packages/wholesale", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageWholesaleV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/packages/wholesale", null, queryParams);
        }

        /// <summary>
        /// Please note: that the {id} parameter above represents a Shipment Delivery ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Shipment Delivery ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersDetailsV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersDetailsV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetHubV2(string? estimatedArrivalEnd = null, string? estimatedArrivalStart = null, string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(estimatedArrivalEnd)) queryParams["estimatedArrivalEnd"] = estimatedArrivalEnd;
            if (!string.IsNullOrEmpty(estimatedArrivalStart)) queryParams["estimatedArrivalStart"] = estimatedArrivalStart;
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/hub", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetIncomingV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/incoming", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetIncomingV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/incoming", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetOutgoingV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/outgoing", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetOutgoingV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/outgoing", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetRejectedV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/rejected", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of shipments with rejected packages for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        ///   - View Transfers
        /// </summary>
        public async Task<object> TransfersGetRejectedV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/rejected", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/templates", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/templates/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryPackageV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/templates/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryTransportersV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/templates/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryTransportersDetailsV1(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/templates/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        ///   - View Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesOutgoingV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of deliveries associated with a specific transfer template.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        ///   - View Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesOutgoingDeliveryV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        ///   - View Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesOutgoingDeliveryPackageV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        ///   - View Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesOutgoingDeliveryTransportersV2(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }

        /// <summary>
        /// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        ///   - View Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(string id, string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetTypesV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v1/types", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of available transfer types for a Facility based on its license number.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetTypesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/types", null, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task TransfersUpdateExternalIncomingV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v1/external/incoming", body, queryParams);
        }

        /// <summary>
        /// Updates external incoming shipment plans for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersUpdateExternalIncomingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v2/external/incoming", body, queryParams);
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersUpdateTemplatesV1(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v1/templates", body, queryParams);
        }

        /// <summary>
        /// Updates existing transfer templates for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersUpdateTemplatesOutgoingV2(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v2/templates/outgoing", body, queryParams);
        }

    }
}
