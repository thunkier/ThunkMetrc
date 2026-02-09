package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class TransfersService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public TransfersService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createTransfersHubArrive(
        String licenseNumber, List<TransfersCreateHubArriveRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersHubArrive(
                licenseNumber, body
            )
        );
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
    public WriteResponse createTransfersHubCheckin(
        String licenseNumber, List<TransfersCreateHubCheckinRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersHubCheckin(
                licenseNumber, body
            )
        );
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
    public WriteResponse createTransfersHubCheckout(
        String licenseNumber, List<TransfersCreateHubCheckoutRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersHubCheckout(
                licenseNumber, body
            )
        );
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
    public WriteResponse createTransfersHubDepart(
        String licenseNumber, List<TransfersCreateHubDepartRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersHubDepart(
                licenseNumber, body
            )
        );
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
    public WriteResponse createTransfersIncomingExternal(
        String licenseNumber, List<TransfersCreateIncomingExternalRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersIncomingExternal(
                licenseNumber, body
            )
        );
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
    public WriteResponse createTransfersOutgoingTemplates(
        String licenseNumber, List<TransfersCreateOutgoingTemplatesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.createTransfersOutgoingTemplates(
                licenseNumber, body
            )
        );
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
    public Object deleteTransfersIncomingExternalById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.transfers.deleteTransfersIncomingExternalById(
                id, licenseNumber
            )
        );
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
    public Object deleteTransfersOutgoingTemplateById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.transfers.deleteTransfersOutgoingTemplateById(
                id, licenseNumber
            )
        );
    }

    /**
     * Returns a list of available shipment Package states.
     * GET /transfers/v2/deliveries/packages/states
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getTransfersDeliveriesPackagesStates(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.transfers.getTransfersDeliveriesPackagesStates(
                
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveryPackage> getTransfersDeliveryPackageById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveryPackage>) client.transfers.getTransfersDeliveryPackageById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveryPackageRequiredlabtestbatch> getTransfersDeliveryPackageRequiredlabtestbatchById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveryPackageRequiredlabtestbatch>) client.transfers.getTransfersDeliveryPackageRequiredlabtestbatchById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveryPackageWholesale> getTransfersDeliveryPackageWholesaleById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveryPackageWholesale>) client.transfers.getTransfersDeliveryPackageWholesaleById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveryTransporter> getTransfersDeliveryTransporterById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveryTransporter>) client.transfers.getTransfersDeliveryTransporterById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveryTransporterDetail> getTransfersDeliveryTransporterDetailById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveryTransporterDetail>) client.transfers.getTransfersDeliveryTransporterDetailById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Hub> getTransfersHub(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Hub>) client.transfers.getTransfersHub(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Transfer> getIncomingTransfers(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Transfer>) client.transfers.getIncomingTransfers(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    public Object getTransfersManifestHtmlById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.transfers.getTransfersManifestHtmlById(
                id, licenseNumber
            )
        );
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
    public ManifestPdf getTransfersManifestPdfById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (ManifestPdf) client.transfers.getTransfersManifestPdfById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TemplateDelivery> getTransfersOutgoingTemplateDeliveryById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TemplateDelivery>) client.transfers.getTransfersOutgoingTemplateDeliveryById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TemplateDeliveryPackage> getTransfersOutgoingTemplateDeliveryPackageById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TemplateDeliveryPackage>) client.transfers.getTransfersOutgoingTemplateDeliveryPackageById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TemplateDeliveryTransporter> getTransfersOutgoingTemplateDeliveryTransporterById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TemplateDeliveryTransporter>) client.transfers.getTransfersOutgoingTemplateDeliveryTransporterById(
                id, pageNumber, pageSize
            )
        );
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
    public TemplateDeliveryTransporterDetail getTransfersOutgoingTemplateDeliveryTransporterDetailById(
        String id
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (TemplateDeliveryTransporterDetail) client.transfers.getTransfersOutgoingTemplateDeliveryTransporterDetailById(
                id
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Template> getTransfersOutgoingTemplates(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Template>) client.transfers.getTransfersOutgoingTemplates(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Transfer> getOutgoingTransfers(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Transfer>) client.transfers.getOutgoingTransfers(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<Transfer> getRejectedTransfers(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<Transfer>) client.transfers.getRejectedTransfers(
                licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TransfersDelivery> getTransfersDeliveryById(
        String id, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TransfersDelivery>) client.transfers.getTransfersDeliveryById(
                id, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TransfersType> getTransfersTypes(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TransfersType>) client.transfers.getTransfersTypes(
                licenseNumber, pageNumber, pageSize
            )
        );
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
    public WriteResponse updateTransfersIncomingExternal(
        String licenseNumber, List<TransfersUpdateIncomingExternalRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.updateTransfersIncomingExternal(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateTransfersOutgoingTemplates(
        String licenseNumber, List<TransfersUpdateOutgoingTemplatesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transfers.updateTransfersOutgoingTemplates(
                licenseNumber, body
            )
        );
    }

}

