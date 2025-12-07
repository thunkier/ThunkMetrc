using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using ThunkMetrc.Client;

namespace ThunkMetrc.Wrapper
{
    public class MetrcWrapper
    {
        private readonly MetrcClient _client;
        private readonly MetrcRateLimiter _rateLimiter;

        public MetrcWrapper(MetrcClient client, RateLimiterConfig? config = null)
        {
            _client = client;
            _rateLimiter = new MetrcRateLimiter(config);
        }

        /// <summary>
        /// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task SandboxCreateIntegratorSetupV2(object? body = null, string? userKey = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SandboxCreateIntegratorSetupV2(body, userKey); return null; });
        }

        /// <summary>
        /// This endpoint provides a list of facilities for which the authenticated user has access.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> FacilitiesGetAllV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.FacilitiesGetAllV1(No));
        }

        /// <summary>
        /// This endpoint provides a list of facilities for which the authenticated user has access.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> FacilitiesGetAllV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.FacilitiesGetAllV2(No));
        }

        /// <summary>
        /// Adds new patients to a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateV2(List<PatientsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateAddV1(List<PatientsCreateAddV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsCreateAddV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsCreateUpdateV1(List<PatientsCreateUpdateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsCreateUpdateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Removes a Patient, identified by an Id, from a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Patient by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetActiveV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsGetActiveV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active patients for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task<object> PatientsGetActiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsGetActiveV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates Patient information for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Patients
        /// </summary>
        public async Task PatientsUpdateV2(List<PatientsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task TransfersCreateExternalIncomingV1(List<TransfersCreateExternalIncomingV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersCreateExternalIncomingV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates external incoming shipment plans for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersCreateExternalIncomingV2(List<TransfersCreateExternalIncomingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersCreateExternalIncomingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersCreateTemplatesV1(List<TransfersCreateTemplatesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersCreateTemplatesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new transfer templates for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersCreateTemplatesOutgoingV2(List<TransfersCreateTemplatesOutgoingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersCreateTemplatesOutgoingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task TransfersDeleteExternalIncomingV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersDeleteExternalIncomingV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Voids an external incoming shipment plan for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersDeleteExternalIncomingV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersDeleteExternalIncomingV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersDeleteTemplatesV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersDeleteTemplatesV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a transfer template for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersDeleteTemplatesOutgoingV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersDeleteTemplatesOutgoingV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetDeliveriesPackagesStatesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveriesPackagesStatesV1(No));
        }

        /// <summary>
        /// Returns a list of available shipment Package states.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetDeliveriesPackagesStatesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveriesPackagesStatesV2(No));
        }

        /// <summary>
        /// Please note: that the {id} parameter above represents a Shipment Plan ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryV2(id, pageNumber, pageSize));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageV2(id, pageNumber, pageSize));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageRequiredlabtestbatchesV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageRequiredlabtestbatchesV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageRequiredlabtestbatchesV2(id, pageNumber, pageSize));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryPackageWholesaleV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageWholesaleV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryPackageWholesaleV2(id, pageNumber, pageSize));
        }

        /// <summary>
        /// Please note: that the {id} parameter above represents a Shipment Delivery ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryTransportersV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryTransportersV2(id, pageNumber, pageSize));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Shipment Delivery ID.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetDeliveryTransportersDetailsV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryTransportersDetailsV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetDeliveryTransportersDetailsV2(id, pageNumber, pageSize));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetHubV2(estimatedArrivalEnd, estimatedArrivalStart, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetIncomingV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetIncomingV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetIncomingV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetOutgoingV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetOutgoingV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetOutgoingV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetRejectedV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetRejectedV1(licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetRejectedV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesDeliveryV1(id, No));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryPackageV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesDeliveryPackageV1(id, No));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryTransportersV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesDeliveryTransportersV1(id, No));
        }

        /// <summary>
        /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
        /// 
        ///   Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task<object> TransfersGetTemplatesDeliveryTransportersDetailsV1(string id, string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesDeliveryTransportersDetailsV1(id, No));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesOutgoingV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesOutgoingDeliveryV2(id, pageNumber, pageSize));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesOutgoingDeliveryPackageV2(id, pageNumber, pageSize));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesOutgoingDeliveryTransportersV2(id, pageNumber, pageSize));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetTypesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTypesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of available transfer types for a Facility based on its license number.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> TransfersGetTypesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransfersGetTypesV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfers
        /// </summary>
        public async Task TransfersUpdateExternalIncomingV1(List<TransfersUpdateExternalIncomingV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersUpdateExternalIncomingV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates external incoming shipment plans for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfers
        /// </summary>
        public async Task TransfersUpdateExternalIncomingV2(List<TransfersUpdateExternalIncomingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersUpdateExternalIncomingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Transfer Templates
        /// </summary>
        public async Task TransfersUpdateTemplatesV1(List<TransfersUpdateTemplatesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersUpdateTemplatesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates existing transfer templates for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transfer Templates
        /// </summary>
        public async Task TransfersUpdateTemplatesOutgoingV2(List<TransfersUpdateTemplatesOutgoingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransfersUpdateTemplatesOutgoingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new driver records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersCreateDriverV2(List<TransportersCreateDriverV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersCreateDriverV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new vehicle records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersCreateVehicleV2(List<TransportersCreateVehicleV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersCreateVehicleV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersDeleteDriverV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersDeleteDriverV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersDeleteVehicleV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersDeleteVehicleV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetDriverV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransportersGetDriverV2(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of drivers for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetDriversV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransportersGetDriversV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetVehicleV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransportersGetVehicleV2(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of vehicles for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Transporters
        /// </summary>
        public async Task<object> TransportersGetVehiclesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TransportersGetVehiclesV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates existing driver records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersUpdateDriverV2(List<TransportersUpdateDriverV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersUpdateDriverV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates existing vehicle records for a facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Transporters
        /// </summary>
        public async Task TransportersUpdateVehicleV2(List<TransportersUpdateVehicleV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.TransportersUpdateVehicleV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetActiveV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.UnitsOfMeasureGetActiveV1(No));
        }

        /// <summary>
        /// Retrieves all active units of measure.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetActiveV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.UnitsOfMeasureGetActiveV2(No));
        }

        /// <summary>
        /// Retrieves all inactive units of measure.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetInactiveV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.UnitsOfMeasureGetInactiveV2(No));
        }

        /// <summary>
        /// Retrieves all available waste methods.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> WasteMethodsGetAllV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.WasteMethodsGetAllV2(No));
        }

        /// <summary>
        /// Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Caregivers
        /// </summary>
        public async Task<object> CaregiversStatusGetByCaregiverLicenseNumberV1(string caregiverLicenseNumber, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.CaregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber, licenseNumber));
        }

        /// <summary>
        /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Caregivers
        /// </summary>
        public async Task<object> CaregiversStatusGetByCaregiverLicenseNumberV2(string caregiverLicenseNumber, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.CaregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Employees
        /// </summary>
        public async Task<object> EmployeesGetAllV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.EmployeesGetAllV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of employees for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Employees
        ///   - View Employees
        /// </summary>
        public async Task<object> EmployeesGetAllV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.EmployeesGetAllV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Employees
        /// </summary>
        public async Task<object> EmployeesGetPermissionsV2(string? employeeLicenseNumber = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.EmployeesGetPermissionsV2(employeeLicenseNumber, licenseNumber));
        }

        /// <summary>
        /// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateV1(List<ItemsCreateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateV2(List<ItemsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates one or more new item brands for the specified Facility identified by the License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateBrandV2(List<ItemsCreateBrandV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreateBrandV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateFileV2(List<ItemsCreateFileV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreateFileV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreatePhotoV1(List<ItemsCreatePhotoV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreatePhotoV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreatePhotoV2(List<ItemsCreatePhotoV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreatePhotoV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsCreateUpdateV1(List<ItemsCreateUpdateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsCreateUpdateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives the specified Product by Id for the given Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives the specified Item Brand by Id for the given Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsDeleteBrandV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsDeleteBrandV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves detailed information about a specific Item by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetActiveV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetActiveV1(licenseNumber));
        }

        /// <summary>
        /// Returns a list of active items for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetBrandsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetBrandsV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active item brands for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetBrandsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetBrandsV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> ItemsGetCategoriesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetCategoriesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of item categories.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> ItemsGetCategoriesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetCategoriesV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a file by its Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetFileV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetFileV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetInactiveV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetInactiveV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of inactive items for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetInactiveV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetPhotoV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetPhotoV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves an image by its Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task<object> ItemsGetPhotoV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ItemsGetPhotoV2(id, licenseNumber));
        }

        /// <summary>
        /// Updates one or more existing products for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsUpdateV2(List<ItemsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates one or more existing item brands for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Items
        /// </summary>
        public async Task ItemsUpdateBrandV2(List<ItemsUpdateBrandV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ItemsUpdateBrandV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateV1(List<PackagesCreateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateV2(List<PackagesCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateAdjustV1(List<PackagesCreateAdjustV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateAdjustV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records a list of adjustments for packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateAdjustV2(List<PackagesCreateAdjustV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateAdjustV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateChangeItemV1(List<PackagesCreateChangeItemV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateChangeItemV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateChangeLocationV1(List<PackagesCreateChangeLocationV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateChangeLocationV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateFinishV1(List<PackagesCreateFinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateFinishV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreatePlantingsV1(List<PackagesCreatePlantingsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreatePlantingsV1(body, licenseNumber); return null; });
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
        public async Task PackagesCreatePlantingsV2(List<PackagesCreatePlantingsV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreatePlantingsV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateRemediateV1(List<PackagesCreateRemediateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateRemediateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateTestingV1(List<PackagesCreateTestingV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateTestingV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new packages for testing for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesCreateTestingV2(List<PackagesCreateTestingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateTestingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesCreateUnfinishV1(List<PackagesCreateUnfinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesCreateUnfinishV1(body, licenseNumber); return null; });
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
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Package by its Id.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetAdjustReasonsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetAdjustReasonsV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of adjustment reasons for packages at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetAdjustReasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetAdjustReasonsV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetByLabelV1(string label, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetByLabelV1(label, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Package by its label.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetByLabelV2(string label, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetByLabelV2(label, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of inactive packages for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of packages in transit for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetIntransitV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetIntransitV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetLabsamplesV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetLabsamplesV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetOnholdV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of packages on hold for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetOnholdV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves the source harvests for a Package by its Id.
        /// 
        ///   Permissions Required:
        ///   - View Package Source Harvests
        /// </summary>
        public async Task<object> PackagesGetSourceHarvestV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetSourceHarvestV2(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of transferred packages for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> PackagesGetTransferredV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetTransferredV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetTypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetTypesV1(No));
        }

        /// <summary>
        /// Retrieves a list of available Package types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PackagesGetTypesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PackagesGetTypesV2(No));
        }

        /// <summary>
        /// Set the final quantity for a Package.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateAdjustV2(List<PackagesUpdateAdjustV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateAdjustV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        ///   - Manage Package Notes
        /// </summary>
        public async Task PackagesUpdateChangeNoteV1(List<PackagesUpdateChangeNoteV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateChangeNoteV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the Product decontaminate information for a list of packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDecontaminateV2(List<PackagesUpdateDecontaminateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateDecontaminateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Flags one or more packages for donation at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDonationFlagV2(List<PackagesUpdateDonationFlagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateDonationFlagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Removes the donation flag from one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateDonationUnflagV2(List<PackagesUpdateDonationUnflagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateDonationUnflagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the external identifiers for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Package Inventory
        ///   - External Id Enabled
        /// </summary>
        public async Task PackagesUpdateExternalidV2(List<PackagesUpdateExternalidV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateExternalidV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates a list of packages as finished for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateFinishV2(List<PackagesUpdateFinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateFinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the associated Item for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateItemV2(List<PackagesUpdateItemV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateItemV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the list of required lab test batches for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateLabTestRequiredV2(List<PackagesUpdateLabTestRequiredV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateLabTestRequiredV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the Location and Sublocation for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateLocationV2(List<PackagesUpdateLocationV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateLocationV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates notes associated with one or more packages for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        ///   - Manage Package Notes
        /// </summary>
        public async Task PackagesUpdateNoteV2(List<PackagesUpdateNoteV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateNoteV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates a list of Product remediations for packages at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateRemediateV2(List<PackagesUpdateRemediateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateRemediateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Flags or unflags one or more packages at the specified Facility as trade samples.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateTradesampleFlagV2(List<PackagesUpdateTradesampleFlagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateTradesampleFlagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Removes the trade sample flag from one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateTradesampleUnflagV2(List<PackagesUpdateTradesampleUnflagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateTradesampleUnflagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates a list of packages as unfinished for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task PackagesUpdateUnfinishV2(List<PackagesUpdateUnfinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateUnfinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the use-by date for one or more packages at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PackagesUpdateUsebydateV2(List<PackagesUpdateUsebydateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PackagesUpdateUsebydateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsCreateV1(List<PatientCheckInsCreateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsCreateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsCreateV2(List<PatientCheckInsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task<object> PatientCheckInsGetAllV1(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientCheckInsGetAllV1(checkinDateEnd, checkinDateStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task<object> PatientCheckInsGetAllV2(string? checkinDateEnd = null, string? checkinDateStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientCheckInsGetAllV2(checkinDateEnd, checkinDateStart, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PatientCheckInsGetLocationsV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientCheckInsGetLocationsV1(No));
        }

        /// <summary>
        /// Retrieves a list of Patient Check-In locations.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PatientCheckInsGetLocationsV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientCheckInsGetLocationsV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsUpdateV1(List<PatientCheckInsUpdateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsUpdateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates patient check-ins for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - ManagePatientsCheckIns
        /// </summary>
        public async Task PatientCheckInsUpdateV2(List<PatientCheckInsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PatientCheckInsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantBatchesCreateAdditivesV1(List<PlantBatchesCreateAdditivesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateAdditivesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records Additive usage details for plant batches at a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantBatchesCreateAdditivesV2(List<PlantBatchesCreateAdditivesV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateAdditivesV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records Additive usage for plant batches at a Facility using predefined additive templates.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantBatchesCreateAdditivesUsingtemplateV2(List<PlantBatchesCreateAdditivesUsingtemplateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateAdditivesUsingtemplateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateAdjustV1(List<PlantBatchesCreateAdjustV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateAdjustV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateAdjustV2(List<PlantBatchesCreateAdjustV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateAdjustV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateChangegrowthphaseV1(List<PlantBatchesCreateChangegrowthphaseV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateChangegrowthphaseV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the growth phase of plants at a specified Facility based on tracking information.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateGrowthphaseV2(List<PlantBatchesCreateGrowthphaseV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateGrowthphaseV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageV2(List<PlantBatchesCreatePackageV2RequestItem> body, string? isFromMotherPlant = null, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreatePackageV2(body, isFromMotherPlant, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageFrommotherplantV1(List<PlantBatchesCreatePackageFrommotherplantV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreatePackageFrommotherplantV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages from mother plants at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatePackageFrommotherplantV2(List<PlantBatchesCreatePackageFrommotherplantV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreatePackageFrommotherplantV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreatePlantingsV2(List<PlantBatchesCreatePlantingsV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreatePlantingsV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateSplitV1(List<PlantBatchesCreateSplitV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateSplitV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Splits an existing Plant Batch into multiple groups at the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateSplitV2(List<PlantBatchesCreateSplitV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateSplitV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantBatchesCreateWasteV1(List<PlantBatchesCreateWasteV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateWasteV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records waste information for plant batches based on the submitted data for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantBatchesCreateWasteV2(List<PlantBatchesCreateWasteV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateWasteV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantBatchesCreatepackagesV1(List<PlantBatchesCreatepackagesV1RequestItem> body, string? isFromMotherPlant = null, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreatepackagesV1(body, isFromMotherPlant, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        /// </summary>
        public async Task PlantBatchesCreateplantingsV1(List<PlantBatchesCreateplantingsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesCreateplantingsV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Destroy Immature Plants
        /// </summary>
        public async Task PlantBatchesDeleteV1(string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesDeleteV1(licenseNumber); return null; });
        }

        /// <summary>
        /// Completes the destruction of plant batches based on the provided input data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Destroy Immature Plants
        /// </summary>
        public async Task PlantBatchesDeleteV2(string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesDeleteV2(licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Plant Batch by Id.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task<object> PlantBatchesGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetTypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetTypesV1(No));
        }

        /// <summary>
        /// Retrieves a list of plant batch types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetTypesV2(pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves waste details associated with plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantBatchesGetWasteV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetWasteV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetWasteReasonsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetWasteReasonsV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantBatchesGetWasteReasonsV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantBatchesGetWasteReasonsV2(licenseNumber));
        }

        /// <summary>
        /// Moves one or more plant batches to new locations with in a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants
        /// </summary>
        public async Task PlantBatchesUpdateLocationV2(List<PlantBatchesUpdateLocationV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesUpdateLocationV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        /// </summary>
        public async Task PlantBatchesUpdateMoveplantbatchesV1(List<PlantBatchesUpdateMoveplantbatchesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesUpdateMoveplantbatchesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Renames plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateNameV2(List<PlantBatchesUpdateNameV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesUpdateNameV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Changes the strain of plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateStrainV2(List<PlantBatchesUpdateStrainV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesUpdateStrainV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Replaces tags for plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantBatchesUpdateTagV2(List<PlantBatchesUpdateTagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantBatchesUpdateTagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreateAdjustV1(List<ProcessingJobsCreateAdjustV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateAdjustV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateAdjustV2(List<ProcessingJobsCreateAdjustV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateAdjustV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateJobtypesV1(List<ProcessingJobsCreateJobtypesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateJobtypesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateJobtypesV2(List<ProcessingJobsCreateJobtypesV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateJobtypesV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreateStartV1(List<ProcessingJobsCreateStartV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateStartV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Initiates new processing jobs at a Facility, including job details and associated packages.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreateStartV2(List<ProcessingJobsCreateStartV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreateStartV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - ManageProcessingJobs
        /// </summary>
        public async Task ProcessingJobsCreatepackagesV1(List<ProcessingJobsCreatepackagesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreatepackagesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsCreatepackagesV2(List<ProcessingJobsCreatepackagesV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsCreatepackagesV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteJobtypesV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsDeleteJobtypesV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a Processing Job Type at a Facility, making it inactive for future use.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsDeleteJobtypesV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsDeleteJobtypesV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a ProcessingJob by Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves active processing jobs at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves inactive processing jobs at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of all active processing job types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesAttributesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesAttributesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves all processing job attributes available for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesAttributesV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesAttributesV2(licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesCategoriesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesCategoriesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves all processing job categories available for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesCategoriesV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesCategoriesV2(licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of all inactive processing job types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task<object> ProcessingJobsGetJobtypesInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.ProcessingJobsGetJobtypesInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateFinishV1(List<ProcessingJobsUpdateFinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateFinishV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Completes processing jobs at a Facility by recording final notes and waste measurements.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateFinishV2(List<ProcessingJobsUpdateFinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateFinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateJobtypesV1(List<ProcessingJobsUpdateJobtypesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateJobtypesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateJobtypesV2(List<ProcessingJobsUpdateJobtypesV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateJobtypesV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateUnfinishV1(List<ProcessingJobsUpdateUnfinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateUnfinishV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        /// 
        ///   Permissions Required:
        ///   - Manage Processing Job
        /// </summary>
        public async Task ProcessingJobsUpdateUnfinishV2(List<ProcessingJobsUpdateUnfinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.ProcessingJobsUpdateUnfinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new sublocation records for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsCreateV2(List<SublocationsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SublocationsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives an existing Sublocation record for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SublocationsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Retrieves a Sublocation by its Id, with an optional license number.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SublocationsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SublocationsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of inactive sublocations for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> SublocationsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SublocationsGetInactiveV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates existing sublocation records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task SublocationsUpdateV2(List<SublocationsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SublocationsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesCreateDeliveryV1(List<SalesCreateDeliveryV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryV2(List<SalesCreateDeliveryV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerV1(List<SalesCreateDeliveryRetailerV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryRetailerV2(List<SalesCreateDeliveryRetailerV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerDepartV1(List<SalesCreateDeliveryRetailerDepartV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerDepartV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryRetailerDepartV2(List<SalesCreateDeliveryRetailerDepartV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerDepartV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerEndV1(List<SalesCreateDeliveryRetailerEndV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerEndV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryRetailerEndV2(List<SalesCreateDeliveryRetailerEndV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerEndV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerRestockV1(List<SalesCreateDeliveryRetailerRestockV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerRestockV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryRetailerRestockV2(List<SalesCreateDeliveryRetailerRestockV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerRestockV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesCreateDeliveryRetailerSaleV1(List<SalesCreateDeliveryRetailerSaleV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerSaleV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateDeliveryRetailerSaleV2(List<SalesCreateDeliveryRetailerSaleV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateDeliveryRetailerSaleV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesCreateReceiptV1(List<SalesCreateReceiptV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateReceiptV1(body, licenseNumber); return null; });
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
        public async Task SalesCreateReceiptV2(List<SalesCreateReceiptV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateReceiptV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesCreateTransactionByDateV1(string date, List<SalesCreateTransactionByDateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesCreateTransactionByDateV1(date, body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteDeliveryV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteDeliveryV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesDeleteDeliveryRetailerV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteDeliveryRetailerV1(id, licenseNumber); return null; });
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
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteDeliveryRetailerV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesDeleteReceiptV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteReceiptV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesDeleteReceiptV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesDeleteReceiptV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCountiesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetCountiesV1(No));
        }

        /// <summary>
        /// Returns a list of counties available for sales deliveries.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCountiesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetCountiesV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCustomertypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetCustomertypesV1(No));
        }

        /// <summary>
        /// Returns a list of customer types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> SalesGetCustomertypesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetCustomertypesV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesRetailerActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesRetailerActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesRetailerInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesRetailerInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesRetailerInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetDeliveriesReturnreasonsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesReturnreasonsV1(licenseNumber));
        }

        /// <summary>
        /// Returns a list of return reasons for sales deliveries based on the provided License Number.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveriesReturnreasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveriesReturnreasonsV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveryV1(id, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveryV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task<object> SalesGetDeliveryRetailerV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveryRetailerV1(id, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetDeliveryRetailerV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetPatientRegistrationsLocationsV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetPatientRegistrationsLocationsV1(No));
        }

        /// <summary>
        /// Returns a list of valid Patient registration locations for sales.
        /// 
        ///   Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> SalesGetPatientRegistrationsLocationsV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetPatientRegistrationsLocationsV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task<object> SalesGetPaymenttypesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetPaymenttypesV1(licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetPaymenttypesV2(licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptV1(id, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptsActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptsActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptsExternalByExternalNumberV2(externalNumber, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetReceiptsInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? salesDateEnd = null, string? salesDateStart = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptsInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetReceiptsInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetTransactionsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetTransactionsV1(licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task<object> SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(string salesDateStart, string salesDateEnd, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart, salesDateEnd, licenseNumber));
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryV1(List<SalesUpdateDeliveryV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryV2(List<SalesUpdateDeliveryV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryCompleteV1(List<SalesUpdateDeliveryCompleteV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryCompleteV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryCompleteV2(List<SalesUpdateDeliveryCompleteV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryCompleteV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryHubV1(List<SalesUpdateDeliveryHubV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery, Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubV2(List<SalesUpdateDeliveryHubV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubAcceptV1(List<SalesUpdateDeliveryHubAcceptV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubAcceptV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubAcceptV2(List<SalesUpdateDeliveryHubAcceptV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubAcceptV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubDepartV1(List<SalesUpdateDeliveryHubDepartV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubDepartV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubDepartV2(List<SalesUpdateDeliveryHubDepartV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubDepartV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateDeliveryHubVerifyIdV1(List<SalesUpdateDeliveryHubVerifyIdV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubVerifyIdV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales Delivery Hub
        /// </summary>
        public async Task SalesUpdateDeliveryHubVerifyIdV2(List<SalesUpdateDeliveryHubVerifyIdV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryHubVerifyIdV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Retailer Delivery
        /// </summary>
        public async Task SalesUpdateDeliveryRetailerV1(List<SalesUpdateDeliveryRetailerV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryRetailerV1(body, licenseNumber); return null; });
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
        public async Task SalesUpdateDeliveryRetailerV2(List<SalesUpdateDeliveryRetailerV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateDeliveryRetailerV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateReceiptV1(List<SalesUpdateReceiptV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateReceiptV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptV2(List<SalesUpdateReceiptV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateReceiptV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptFinalizeV2(List<SalesUpdateReceiptFinalizeV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateReceiptFinalizeV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        /// 
        ///   Permissions Required:
        ///   - Manage Sales
        /// </summary>
        public async Task SalesUpdateReceiptUnfinalizeV2(List<SalesUpdateReceiptUnfinalizeV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateReceiptUnfinalizeV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Sales
        /// </summary>
        public async Task SalesUpdateTransactionByDateV1(string date, List<SalesUpdateTransactionByDateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.SalesUpdateTransactionByDateV1(date, body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateV1(List<StrainsCreateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsCreateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new strain records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateV2(List<StrainsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsCreateUpdateV1(List<StrainsCreateUpdateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsCreateUpdateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives an existing strain record for a Facility
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.StrainsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Strain record by its Id, with an optional license number.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.StrainsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetActiveV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.StrainsGetActiveV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.StrainsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task<object> StrainsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.StrainsGetInactiveV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates existing strain records for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Strains
        /// </summary>
        public async Task StrainsUpdateV2(List<StrainsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.StrainsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Returns a list of available package tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        /// </summary>
        public async Task<object> TagsGetPackageAvailableV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TagsGetPackageAvailableV2(licenseNumber));
        }

        /// <summary>
        /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        /// </summary>
        public async Task<object> TagsGetPlantAvailableV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TagsGetPlantAvailableV2(licenseNumber));
        }

        /// <summary>
        /// Returns a list of staged tags. NOTE: This is a premium endpoint.
        /// 
        ///   Permissions Required:
        ///   - Manage Tags
        ///   - RetailId.AllowPackageStaging Key Value enabled
        /// </summary>
        public async Task<object> TagsGetStagedV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.TagsGetStagedV2(licenseNumber));
        }

        /// <summary>
        /// Creates new additive templates for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Additives
        /// </summary>
        public async Task AdditivesTemplatesCreateV2(List<AdditivesTemplatesCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.AdditivesTemplatesCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Retrieves an Additive Template by its Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Additives
        /// </summary>
        public async Task<object> AdditivesTemplatesGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.AdditivesTemplatesGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active additive templates for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Additives
        /// </summary>
        public async Task<object> AdditivesTemplatesGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.AdditivesTemplatesGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of inactive additive templates for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Additives
        /// </summary>
        public async Task<object> AdditivesTemplatesGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.AdditivesTemplatesGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates existing additive templates for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Additives
        /// </summary>
        public async Task AdditivesTemplatesUpdateV2(List<AdditivesTemplatesUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.AdditivesTemplatesUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsCreateFinishV1(List<HarvestsCreateFinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreateFinishV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageV1(List<HarvestsCreatePackageV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreatePackageV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages from harvested products for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageV2(List<HarvestsCreatePackageV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreatePackageV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageTestingV1(List<HarvestsCreatePackageTestingV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreatePackageTestingV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages for testing from harvested products for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task HarvestsCreatePackageTestingV2(List<HarvestsCreatePackageTestingV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreatePackageTestingV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsCreateRemoveWasteV1(List<HarvestsCreateRemoveWasteV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreateRemoveWasteV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsCreateUnfinishV1(List<HarvestsCreateUnfinishV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreateUnfinishV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsCreateWasteV2(List<HarvestsCreateWasteV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsCreateWasteV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Discontinues a specific harvest waste record by Id for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Discontinue Harvest Waste
        /// </summary>
        public async Task HarvestsDeleteWasteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsDeleteWasteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetActiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetActiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active harvests for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of inactive harvests for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetOnholdV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of harvests on hold for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetOnholdV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        /// </summary>
        public async Task<object> HarvestsGetWasteV2(string? harvestId = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetWasteV2(harvestId, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> HarvestsGetWasteTypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetWasteTypesV1(No));
        }

        /// <summary>
        /// Retrieves a list of Waste types for harvests.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> HarvestsGetWasteTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.HarvestsGetWasteTypesV2(pageNumber, pageSize));
        }

        /// <summary>
        /// Marks one or more harvests as finished for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateFinishV2(List<HarvestsUpdateFinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateFinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the Location of Harvest for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateLocationV2(List<HarvestsUpdateLocationV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateLocationV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateMoveV1(List<HarvestsUpdateMoveV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateMoveV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateRenameV1(List<HarvestsUpdateRenameV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateRenameV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Renames one or more harvests for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Manage Harvests
        /// </summary>
        public async Task HarvestsUpdateRenameV2(List<HarvestsUpdateRenameV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateRenameV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Restores previously harvested plants to their original state for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateRestoreHarvestedPlantsV2(List<HarvestsUpdateRestoreHarvestedPlantsV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateRestoreHarvestedPlantsV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Reopens one or more previously finished harvests for the specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Harvests
        ///   - Finish/Discontinue Harvests
        /// </summary>
        public async Task HarvestsUpdateUnfinishV2(List<HarvestsUpdateUnfinishV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.HarvestsUpdateUnfinishV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsCreateRecordV1(List<LabTestsCreateRecordV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsCreateRecordV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsCreateRecordV2(List<LabTestsCreateRecordV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsCreateRecordV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Retrieves a list of Lab Test batches.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetBatchesV2(string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetBatchesV2(pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task<object> LabTestsGetLabtestdocumentV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetLabtestdocumentV1(id, licenseNumber));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetLabtestdocumentV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        /// </summary>
        public async Task<object> LabTestsGetResultsV1(string? licenseNumber = null, string? packageId = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetResultsV1(licenseNumber, packageId));
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
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetResultsV2(licenseNumber, packageId, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetStatesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetStatesV1(No));
        }

        /// <summary>
        /// Returns a list of all lab testing states.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetStatesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetStatesV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetTypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetTypesV1(No));
        }

        /// <summary>
        /// Returns a list of Lab Test types.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> LabTestsGetTypesV2(string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LabTestsGetTypesV2(pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateLabtestdocumentV1(List<LabTestsUpdateLabtestdocumentV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsUpdateLabtestdocumentV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates one or more documents for previously submitted lab tests.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateLabtestdocumentV2(List<LabTestsUpdateLabtestdocumentV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsUpdateLabtestdocumentV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateResultReleaseV1(List<LabTestsUpdateResultReleaseV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsUpdateResultReleaseV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Releases Lab Test results for one or more packages.
        /// 
        ///   Permissions Required:
        ///   - View Packages
        ///   - Manage Packages Inventory
        /// </summary>
        public async Task LabTestsUpdateResultReleaseV2(List<LabTestsUpdateResultReleaseV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LabTestsUpdateResultReleaseV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateV1(List<LocationsCreateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsCreateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateV2(List<LocationsCreateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsCreateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsCreateUpdateV1(List<LocationsCreateUpdateV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsCreateUpdateV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsDeleteV1(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsDeleteV1(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Archives a specified Location, identified by its Id, for a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsDeleteV2(string id, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsDeleteV2(id, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Location by its Id.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetActiveV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetActiveV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetActiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetActiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of inactive locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetInactiveV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetInactiveV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetTypesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetTypesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of active location types for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task<object> LocationsGetTypesV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.LocationsGetTypesV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Updates existing locations for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Locations
        /// </summary>
        public async Task LocationsUpdateV2(List<LocationsUpdateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.LocationsUpdateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Patients
        /// </summary>
        public async Task<object> PatientsStatusGetStatusesByPatientLicenseNumberV1(string patientLicenseNumber, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber, licenseNumber));
        }

        /// <summary>
        /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Patients
        /// </summary>
        public async Task<object> PatientsStatusGetStatusesByPatientLicenseNumberV2(string patientLicenseNumber, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PatientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesV1(List<PlantsCreateAdditivesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records additive usage details applied to specific plants at a Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesV2(List<PlantsCreateAdditivesV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationV1(List<PlantsCreateAdditivesBylocationV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesBylocationV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records additive usage for plants based on their location within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationV2(List<PlantsCreateAdditivesBylocationV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesBylocationV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesBylocationUsingtemplateV2(List<PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesBylocationUsingtemplateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records additive usage for plants using predefined additive templates at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Additives
        /// </summary>
        public async Task PlantsCreateAdditivesUsingtemplateV2(List<PlantsCreateAdditivesUsingtemplateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateAdditivesUsingtemplateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreateChangegrowthphasesV1(List<PlantsCreateChangegrowthphasesV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateChangegrowthphasesV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateHarvestplantsV1(List<PlantsCreateHarvestplantsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateHarvestplantsV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates harvest product records from plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateManicureV2(List<PlantsCreateManicureV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateManicureV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsCreateManicureplantsV1(List<PlantsCreateManicureplantsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateManicureplantsV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreateMoveplantsV1(List<PlantsCreateMoveplantsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateMoveplantsV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantsCreatePlantbatchPackageV1(List<PlantsCreatePlantbatchPackageV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreatePlantbatchPackageV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates packages from plant batches at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        ///   - View Packages
        ///   - Create/Submit/Discontinue Packages
        /// </summary>
        public async Task PlantsCreatePlantbatchPackageV2(List<PlantsCreatePlantbatchPackageV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreatePlantbatchPackageV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreatePlantingsV1(List<PlantsCreatePlantingsV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreatePlantingsV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Creates new plant batches at a specified Facility from existing plant data.
        /// 
        ///   Permissions Required:
        ///   - View Immature Plants
        ///   - Manage Immature Plants Inventory
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsCreatePlantingsV2(List<PlantsCreatePlantingsV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreatePlantingsV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantsCreateWasteV1(List<PlantsCreateWasteV1RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateWasteV1(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Records waste events for plants at a Facility, including method, reason, and location details.
        /// 
        ///   Permissions Required:
        ///   - Manage Plants Waste
        /// </summary>
        public async Task PlantsCreateWasteV2(List<PlantsCreateWasteV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsCreateWasteV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Destroy Veg/Flower Plants
        /// </summary>
        public async Task PlantsDeleteV1(string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsDeleteV1(licenseNumber); return null; });
        }

        /// <summary>
        /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Destroy Veg/Flower Plants
        /// </summary>
        public async Task PlantsDeleteV2(string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsDeleteV2(licenseNumber); return null; });
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetV1(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetV1(id, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Plant by Id.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetV2(string id, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetV2(id, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View/Manage Plants Additives
        /// </summary>
        public async Task<object> PlantsGetAdditivesV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetAdditivesV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves additive records applied to plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View/Manage Plants Additives
        /// </summary>
        public async Task<object> PlantsGetAdditivesV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetAdditivesV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   -
        /// </summary>
        public async Task<object> PlantsGetAdditivesTypesV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetAdditivesTypesV1(No));
        }

        /// <summary>
        /// Retrieves a list of all plant additive types defined within a Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetAdditivesTypesV2(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetAdditivesTypesV2(No));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetByLabelV1(string label, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetByLabelV1(label, licenseNumber));
        }

        /// <summary>
        /// Retrieves a Plant by label.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetByLabelV2(string label, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetByLabelV2(label, licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetFloweringV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetFloweringV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetFloweringV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetFloweringV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetGrowthPhasesV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetGrowthPhasesV1(licenseNumber));
        }

        /// <summary>
        /// Retrieves the list of growth phases supported by a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetGrowthPhasesV2(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetGrowthPhasesV2(licenseNumber));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetInactiveV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetInactiveV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves inactive plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves mother-phase plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetMotherV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves inactive mother-phase plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherInactiveV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetMotherInactiveV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Mother Plants
        /// </summary>
        public async Task<object> PlantsGetMotherOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetMotherOnholdV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetOnholdV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetOnholdV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves plants that are currently on hold at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetOnholdV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetOnholdV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetVegetativeV1(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetVegetativeV1(lastModifiedEnd, lastModifiedStart, licenseNumber));
        }

        /// <summary>
        /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        /// </summary>
        public async Task<object> PlantsGetVegetativeV2(string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetVegetativeV2(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of recorded plant waste events for a specific Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWasteV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWasteV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteMethodsAllV1(string? No = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWasteMethodsAllV1(No));
        }

        /// <summary>
        /// Retrieves a list of all available plant waste methods for use within a Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteMethodsAllV2(string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWasteMethodsAllV2(pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWastePackageV2(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWastePackageV2(id, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
        /// 
        ///   Permissions Required:
        ///   - View Plants Waste
        /// </summary>
        public async Task<object> PlantsGetWastePlantV2(string id, string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWastePlantV2(id, licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteReasonsV1(string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWasteReasonsV1(licenseNumber));
        }

        /// <summary>
        /// Retriveves available reasons for recording mature plant waste at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> PlantsGetWasteReasonsV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.PlantsGetWasteReasonsV2(licenseNumber, pageNumber, pageSize));
        }

        /// <summary>
        /// Adjusts the recorded count of plants at a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateAdjustV2(List<PlantsUpdateAdjustV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateAdjustV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Changes the growth phases of plants within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateGrowthphaseV2(List<PlantsUpdateGrowthphaseV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateGrowthphaseV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsUpdateHarvestV2(List<PlantsUpdateHarvestV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateHarvestV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Moves plant batches to new locations within a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateLocationV2(List<PlantsUpdateLocationV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateLocationV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Merges multiple plant groups into a single group within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manicure/Harvest Veg/Flower Plants
        /// </summary>
        public async Task PlantsUpdateMergeV2(List<PlantsUpdateMergeV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateMergeV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Splits an existing plant group into multiple groups within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Plant
        /// </summary>
        public async Task PlantsUpdateSplitV2(List<PlantsUpdateSplitV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateSplitV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Updates the strain information for plants within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateStrainV2(List<PlantsUpdateStrainV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateStrainV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Replaces existing plant tags with new tags for plants within a Facility.
        /// 
        ///   Permissions Required:
        ///   - View Veg/Flower Plants
        ///   - Manage Veg/Flower Plants Inventory
        /// </summary>
        public async Task PlantsUpdateTagV2(List<PlantsUpdateTagV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.PlantsUpdateTagV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly)
        ///   - Industry/View Packages
        ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        public async Task RetailIdCreateAssociateV2(List<RetailIdCreateAssociateV2RequestItem> body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.RetailIdCreateAssociateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly)
        ///   - Industry/View Packages
        ///   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        public async Task RetailIdCreateGenerateV2(RetailIdCreateGenerateV2Request body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.RetailIdCreateGenerateV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly)
        ///   - Key Value Settings/Retail ID Merge Packages Enabled
        /// </summary>
        public async Task RetailIdCreateMergeV2(RetailIdCreateMergeV2Request body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.RetailIdCreateMergeV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Retrieves Package information for given list of Package labels.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        ///   - WebApi Retail ID Read Write State (All or WriteOnly)
        ///   - Industry/View Packages
        ///   - Admin/Employees/Packages Page/Product Labels(Manage)
        /// </summary>
        public async Task RetailIdCreatePackageInfoV2(RetailIdCreatePackageInfoV2Request body, string? licenseNumber = null)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.RetailIdCreatePackageInfoV2(body, licenseNumber); return null; });
        }

        /// <summary>
        /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
        ///   - WebApi Retail ID Read Write State (All or ReadOnly)
        ///   - Industry/View Packages
        ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        /// </summary>
        public async Task<object> RetailIdGetReceiveByLabelV2(string label, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.RetailIdGetReceiveByLabelV2(label, licenseNumber));
        }

        /// <summary>
        /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
        /// 
        ///   Permissions Required:
        ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
        ///   - WebApi Retail ID Read Write State (All or ReadOnly)
        ///   - Industry/View Packages
        ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        /// </summary>
        public async Task<object> RetailIdGetReceiveQrByShortCodeV2(string shortCode, string? licenseNumber = null)
        {
            return await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.RetailIdGetReceiveQrByShortCodeV2(shortCode, licenseNumber));
        }

    }
}
