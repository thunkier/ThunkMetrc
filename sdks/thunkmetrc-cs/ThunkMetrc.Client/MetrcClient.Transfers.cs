using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Arrive a transfer for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to arrive the transfer.</param>
        public async Task CreateTransfersHubArrive(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/hub/arrive", body, queryParams);
        }
        /// <summary>
        /// CheckIn a transfer for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to arrive the transfer.</param>
        public async Task CreateTransfersHubCheckin(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/hub/checkin", body, queryParams);
        }
        /// <summary>
        /// CheckOut a transfer for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to arrive the transfer.</param>
        public async Task CreateTransfersHubCheckout(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/hub/checkout", body, queryParams);
        }
        /// <summary>
        /// Depart a transfer for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Hub
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to depart the transfer.</param>
        public async Task CreateTransfersHubDepart(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/hub/depart", body, queryParams);
        }
        /// <summary>
        /// Creates external incoming shipment plans for a Facility.
        /// Permissions Required:
        /// - Manage Transfers
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number for which to create a Shipment Plan.</param>
        public async Task CreateTransfersIncomingExternal(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/external/incoming", body, queryParams);
        }
        /// <summary>
        /// Creates new transfer templates for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a Shipment Template.</param>
        public async Task CreateTransfersOutgoingTemplates(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("POST"), $"/transfers/v2/templates/outgoing", body, queryParams);
        }
        /// <summary>
        /// Voids an external incoming shipment plan for a Facility.
        /// Permissions Required:
        /// - Manage Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to void the Transfer.</param>
        public async Task DeleteTransfersIncomingExternalById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v2/external/incoming/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Archives a transfer template for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive the template.</param>
        public async Task DeleteTransfersOutgoingTemplateById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("DELETE"), $"/transfers/v2/templates/outgoing/{Uri.EscapeDataString(id)}", null, queryParams);
        }
        /// <summary>
        /// Returns a list of available shipment Package states.
        /// </summary>
        public async Task<object> GetTransfersDeliveriesPackagesStates()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/packages/states", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryPackageById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryPackageRequiredlabtestbatchById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/package/{Uri.EscapeDataString(id)}/requiredlabtestbatches", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryPackageWholesaleById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/packages/wholesale", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryTransporterById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryTransporterDetailById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the estimated arrival date parameters will result in an error.</param>
        /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the estimated arrival date parameters will result in an error.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return shipment plans.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersHub(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(lastModifiedEnd)) queryParams["lastModifiedEnd"] = lastModifiedEnd;
            if (!string.IsNullOrEmpty(lastModifiedStart)) queryParams["lastModifiedStart"] = lastModifiedStart;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/hub", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return incoming transfers.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetIncomingTransfers(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">Optional query parameter.</param>
        public async Task<object> GetTransfersManifestHtmlById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/manifest/{Uri.EscapeDataString(id)}/html", null, queryParams);
        }
        /// <summary>
        /// Get Transfer Manifest PDF for a given Transfer Id
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">Optional query parameter.</param>
        public async Task<object> GetTransfersManifestPdfById(string id, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/manifest/{Uri.EscapeDataString(id)}/pdf", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of deliveries associated with a specific transfer template.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersOutgoingTemplateDeliveryById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersOutgoingTemplateDeliveryPackageById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/packages", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersOutgoingTemplateDeliveryTransporterById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/transporters", null, queryParams);
        }
        /// <summary>
        /// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="id">id</param>
        public async Task<object> GetTransfersOutgoingTemplateDeliveryTransporterDetailById(string id)
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/templates/outgoing/deliveries/{Uri.EscapeDataString(id)}/transporters/details", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// - View Transfer Templates
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return outgoing shipment templates.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersOutgoingTemplates(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="lastModifiedEnd">The last modified end timestamp</param>
        /// <param name="lastModifiedStart">The last modified start timestamp</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return outgoing transfers.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetOutgoingTransfers(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
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
        /// Retrieves a list of shipments with rejected packages for a Facility.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return rejected transfers.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetRejectedTransfers(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/rejected", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        /// Permissions Required:
        /// - Manage Transfers
        /// - View Transfers
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersDeliveryById(string id, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/{Uri.EscapeDataString(id)}/deliveries", null, queryParams);
        }
        /// <summary>
        /// Retrieves a list of available transfer types for a Facility based on its license number.
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return transfer types.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetTransfersTypes(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/transfers/v2/types", null, queryParams);
        }
        /// <summary>
        /// Updates external incoming shipment plans for a Facility.
        /// Permissions Required:
        /// - Manage Transfers
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update external incoming shipment plans.</param>
        public async Task UpdateTransfersIncomingExternal(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v2/external/incoming", body, queryParams);
        }
        /// <summary>
        /// Updates existing transfer templates for a Facility.
        /// Permissions Required:
        /// - Manage Transfer Templates
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update shipment plan templates.</param>
        public async Task UpdateTransfersOutgoingTemplates(object? body = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            await SendAsync(new HttpMethod("PUT"), $"/transfers/v2/templates/outgoing", body, queryParams);
        }
    }
}

