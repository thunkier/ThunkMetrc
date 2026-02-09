package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class TransfersService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Arrive a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST /transfers/v2/hub/arrive
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHubArrive(licenseNumber: String? = null, body: List<TransfersCreateHubArriveRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersHubArrive(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * CheckIn a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST /transfers/v2/hub/checkin
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHubCheckin(licenseNumber: String? = null, body: List<TransfersCreateHubCheckinRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersHubCheckin(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * CheckOut a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST /transfers/v2/hub/checkout
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHubCheckout(licenseNumber: String? = null, body: List<TransfersCreateHubCheckoutRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersHubCheckout(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Depart a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST /transfers/v2/hub/depart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createHubDepart(licenseNumber: String? = null, body: List<TransfersCreateHubDepartRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersHubDepart(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * POST /transfers/v2/external/incoming
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createIncomingExternal(licenseNumber: String? = null, body: List<TransfersCreateIncomingExternalRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersIncomingExternal(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates new transfer templates for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * POST /transfers/v2/templates/outgoing
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createOutgoingTemplates(licenseNumber: String? = null, body: List<TransfersCreateOutgoingTemplatesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.createTransfersOutgoingTemplates(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * DELETE /transfers/v2/external/incoming/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteIncomingExternalById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.deleteTransfersIncomingExternalById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Archives a transfer template for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * DELETE /transfers/v2/templates/outgoing/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteOutgoingTemplateById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.deleteTransfersOutgoingTemplateById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Returns a list of available shipment Package states.
     * GET /transfers/v2/deliveries/packages/states
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveriesPackagesStates(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveriesPackagesStates() 
        } as? Any
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/deliveries/{id}/packages
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryPackageById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveryPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryPackageById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveryPackage>
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/deliveries/package/{id}/requiredlabtestbatches
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryPackageRequiredlabtestbatchById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveryPackageRequiredlabtestbatch>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryPackageRequiredlabtestbatchById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveryPackageRequiredlabtestbatch>
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/deliveries/{id}/packages/wholesale
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryPackageWholesaleById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveryPackageWholesale>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryPackageWholesaleById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveryPackageWholesale>
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/deliveries/{id}/transporters
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryTransporterById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveryTransporter>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryTransporterById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveryTransporter>
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/deliveries/{id}/transporters/details
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryTransporterDetailById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveryTransporterDetail>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryTransporterDetailById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveryTransporterDetail>
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/hub
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getHub(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Hub>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersHub(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Hub>
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/incoming
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getIncomingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Transfer>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getIncomingTransfers(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Transfer>
    }

    /**
     * Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/manifest/{id}/html
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getManifestHtmlById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersManifestHtmlById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Get Transfer Manifest PDF for a given Transfer Id
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/manifest/{id}/pdf
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getManifestPdfById(id: String, licenseNumber: String? = null): ManifestPdf? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersManifestPdfById(id, licenseNumber, ) 
        } as? ManifestPdf
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/templates/outgoing/{id}/deliveries
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTemplateDeliveryById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TemplateDelivery>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersOutgoingTemplateDeliveryById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TemplateDelivery>
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/templates/outgoing/deliveries/{id}/packages
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTemplateDeliveryPackageById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TemplateDeliveryPackage>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersOutgoingTemplateDeliveryPackageById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TemplateDeliveryPackage>
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTemplateDeliveryTransporterById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TemplateDeliveryTransporter>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersOutgoingTemplateDeliveryTransporterById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TemplateDeliveryTransporter>
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
     * @param id Path parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTemplateDeliveryTransporterDetailById(id: String): TemplateDeliveryTransporterDetail? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersOutgoingTemplateDeliveryTransporterDetailById(id, ) 
        } as? TemplateDeliveryTransporterDetail
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET /transfers/v2/templates/outgoing
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTemplates(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Template>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersOutgoingTemplates(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Template>
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/outgoing
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getOutgoingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Transfer>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getOutgoingTransfers(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Transfer>
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/rejected
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getRejectedTransfers(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<Transfer>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getRejectedTransfers(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<Transfer>
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET /transfers/v2/{id}/deliveries
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getTransfersDeliveryById(id: String, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TransfersDelivery>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersDeliveryById(id, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TransfersDelivery>
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * GET /transfers/v2/types
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getTransfersTypes(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TransfersType>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transfers.getTransfersTypes(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TransfersType>
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * PUT /transfers/v2/external/incoming
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateIncomingExternal(licenseNumber: String? = null, body: List<TransfersUpdateIncomingExternalRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.updateTransfersIncomingExternal(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates existing transfer templates for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * PUT /transfers/v2/templates/outgoing
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateOutgoingTemplates(licenseNumber: String? = null, body: List<TransfersUpdateOutgoingTemplatesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transfers.updateTransfersOutgoingTemplates(licenseNumber, body) 
        } as? WriteResponse
    }
}

