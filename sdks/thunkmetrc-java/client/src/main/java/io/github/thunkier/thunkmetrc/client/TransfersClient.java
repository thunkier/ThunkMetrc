package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class TransfersClient {
    private final MetrcClient client;

    public TransfersClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Arrive a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST CreateHubArrive
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersHubArrive(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub/arrive");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * CheckIn a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST CreateHubCheckin
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersHubCheckin(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub/checkin");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * CheckOut a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST CreateHubCheckout
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersHubCheckout(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub/checkout");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Depart a transfer for a Facility.
     * Permissions Required:
     * - Manage Transfer Hub
     * POST CreateHubDepart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersHubDepart(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub/depart");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * POST CreateIncomingExternal
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersIncomingExternal(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates new transfer templates for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * POST CreateOutgoingTemplates
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransfersOutgoingTemplates(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * DELETE DeleteIncomingExternalById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteTransfersIncomingExternalById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Archives a transfer template for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * DELETE DeleteOutgoingTemplateById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteTransfersOutgoingTemplateById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Returns a list of available shipment Package states.
     * GET GetDeliveriesPackagesStates
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveriesPackagesStates(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/packages/states");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetDeliveryPackageById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryPackageById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetDeliveryPackageRequiredlabtestbatchById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryPackageRequiredlabtestbatchById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/package/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/requiredlabtestbatches");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetDeliveryPackageWholesaleById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryPackageWholesaleById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages/wholesale");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetDeliveryTransporterById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryTransporterById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetDeliveryTransporterDetailById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryTransporterDetailById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetHub
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersHub(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/hub");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetIncomingTransfers
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getIncomingTransfers(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/incoming");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetManifestHtmlById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersManifestHtmlById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/manifest/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/html");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Get Transfer Manifest PDF for a given Transfer Id
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetManifestPdfById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersManifestPdfById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/manifest/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/pdf");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetOutgoingTemplateDeliveryById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersOutgoingTemplateDeliveryById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetOutgoingTemplateDeliveryPackageById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersOutgoingTemplateDeliveryPackageById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/packages");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetOutgoingTemplateDeliveryTransporterById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersOutgoingTemplateDeliveryTransporterById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetOutgoingTemplateDeliveryTransporterDetailById
     * @param id Path parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersOutgoingTemplateDeliveryTransporterDetailById(
        String id
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/transporters/details");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfer Templates
     * - View Transfer Templates
     * GET GetOutgoingTemplates
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersOutgoingTemplates(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetOutgoingTransfers
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getOutgoingTransfers(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/outgoing");
        StringBuilder query = new StringBuilder();
        if (lastModifiedEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedEnd=").append(URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8));
        }
        if (lastModifiedStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("lastModifiedStart=").append(URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetRejectedTransfers
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getRejectedTransfers(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/rejected");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * Permissions Required:
     * - Manage Transfers
     * - View Transfers
     * GET GetTransfersDeliveryById
     * @param id Path parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersDeliveryById(
        String id, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"/deliveries");
        StringBuilder query = new StringBuilder();
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * GET GetTransfersTypes
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransfersTypes(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/types");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * Permissions Required:
     * - Manage Transfers
     * PUT UpdateIncomingExternal
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateTransfersIncomingExternal(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/external/incoming");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates existing transfer templates for a Facility.
     * Permissions Required:
     * - Manage Transfer Templates
     * PUT UpdateOutgoingTemplates
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateTransfersOutgoingTemplates(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transfers/v2/templates/outgoing");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

