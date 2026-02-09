package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class TransfersClient(private val client: MetrcClient) {
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
    fun createTransfersHubArrive(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/hub/arrive")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createTransfersHubCheckin(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/hub/checkin")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createTransfersHubCheckout(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/hub/checkout")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createTransfersHubDepart(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/hub/depart")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createTransfersIncomingExternal(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun createTransfersOutgoingTemplates(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
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
    fun deleteTransfersIncomingExternalById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
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
    fun deleteTransfersOutgoingTemplateById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Returns a list of available shipment Package states.
     * GET GetDeliveriesPackagesStates
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTransfersDeliveriesPackagesStates(): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/packages/states")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryPackageById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryPackageRequiredlabtestbatchById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/package/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/requiredlabtestbatches")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryPackageWholesaleById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages/wholesale")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryTransporterById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryTransporterDetailById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersHub(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/hub")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getIncomingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/incoming")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersManifestHtmlById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/manifest/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/html")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersManifestPdfById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/manifest/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/pdf")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersOutgoingTemplateDeliveryById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersOutgoingTemplateDeliveryPackageById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/packages")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersOutgoingTemplateDeliveryTransporterById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersOutgoingTemplateDeliveryTransporterDetailById(id: String, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/transporters/details")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersOutgoingTemplates(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getOutgoingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/outgoing")
        val query = ArrayList<String>()
        if (lastModifiedEnd != null) {
            query.add("lastModifiedEnd=${URLEncoder.encode(lastModifiedEnd, StandardCharsets.UTF_8)}")
        }
        if (lastModifiedStart != null) {
            query.add("lastModifiedStart=${URLEncoder.encode(lastModifiedStart, StandardCharsets.UTF_8)}")
        }
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getRejectedTransfers(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/rejected")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersDeliveryById(id: String, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/${URLEncoder.encode(id, StandardCharsets.UTF_8)}/deliveries")
        val query = ArrayList<String>()
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getTransfersTypes(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transfers/v2/types")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun updateTransfersIncomingExternal(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/external/incoming")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
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
    fun updateTransfersOutgoingTemplates(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transfers/v2/templates/outgoing")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

}

