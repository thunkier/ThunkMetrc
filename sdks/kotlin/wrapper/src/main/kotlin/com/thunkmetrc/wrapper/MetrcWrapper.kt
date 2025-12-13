package com.thunkmetrc.wrapper

import com.thunkmetrc.client.MetrcClient

class MetrcWrapper(private val client: MetrcClient, private val config: RateLimiterConfig = RateLimiterConfig()) {

    private val rateLimiter = MetrcRateLimiter(config)

    constructor(baseUrl: String, vendorKey: String, userKey: String) : this(MetrcClient(baseUrl, vendorKey, userKey))

    /**
     * Creates new sublocation records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     */
    suspend fun sublocationsCreateV2(licensenumber: String? = null, body: List<SublocationsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.sublocationsCreateV2(licensenumber, body) }
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     */
    suspend fun sublocationsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.sublocationsDeleteV2(id, licensenumber) }
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     */
    suspend fun sublocationsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.sublocationsGetV2(id, licensenumber) }
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     */
    suspend fun sublocationsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.sublocationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     */
    suspend fun sublocationsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.sublocationsGetInactiveV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     */
    suspend fun sublocationsUpdateV2(licensenumber: String? = null, body: List<SublocationsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.sublocationsUpdateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * POST CreateExternalIncoming V1
     */
    suspend fun transfersCreateExternalIncomingV1(licensenumber: String? = null, body: List<TransfersCreateExternalIncomingV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersCreateExternalIncomingV1(licensenumber, body) }
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * POST CreateExternalIncoming V2
     */
    suspend fun transfersCreateExternalIncomingV2(licensenumber: String? = null, body: List<TransfersCreateExternalIncomingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersCreateExternalIncomingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * POST CreateTemplates V1
     */
    suspend fun transfersCreateTemplatesV1(licensenumber: String? = null, body: List<TransfersCreateTemplatesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersCreateTemplatesV1(licensenumber, body) }
    }

    /**
     * Creates new transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * POST CreateTemplatesOutgoing V2
     */
    suspend fun transfersCreateTemplatesOutgoingV2(licensenumber: String? = null, body: List<TransfersCreateTemplatesOutgoingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersCreateTemplatesOutgoingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * DELETE DeleteExternalIncoming V1
     */
    suspend fun transfersDeleteExternalIncomingV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transfersDeleteExternalIncomingV1(id, licensenumber) }
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * DELETE DeleteExternalIncoming V2
     */
    suspend fun transfersDeleteExternalIncomingV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transfersDeleteExternalIncomingV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * DELETE DeleteTemplates V1
     */
    suspend fun transfersDeleteTemplatesV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transfersDeleteTemplatesV1(id, licensenumber) }
    }

    /**
     * Archives a transfer template for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * DELETE DeleteTemplatesOutgoing V2
     */
    suspend fun transfersDeleteTemplatesOutgoingV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transfersDeleteTemplatesOutgoingV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V1
     */
    suspend fun transfersGetDeliveriesPackagesStatesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveriesPackagesStatesV1(no) }
    }

    /**
     * Returns a list of available shipment Package states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V2
     */
    suspend fun transfersGetDeliveriesPackagesStatesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveriesPackagesStatesV2(no) }
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Plan ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDelivery V1
     */
    suspend fun transfersGetDeliveryV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryV1(id, no) }
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDelivery V2
     */
    suspend fun transfersGetDeliveryV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryV2(id, pagenumber, pagesize) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackage V1
     */
    suspend fun transfersGetDeliveryPackageV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageV1(id, no) }
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackage V2
     */
    suspend fun transfersGetDeliveryPackageV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageV2(id, pagenumber, pagesize) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V1
     */
    suspend fun transfersGetDeliveryPackageRequiredlabtestbatchesV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, no) }
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V2
     */
    suspend fun transfersGetDeliveryPackageRequiredlabtestbatchesV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, pagenumber, pagesize) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageWholesale V1
     */
    suspend fun transfersGetDeliveryPackageWholesaleV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageWholesaleV1(id, no) }
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageWholesale V2
     */
    suspend fun transfersGetDeliveryPackageWholesaleV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryPackageWholesaleV2(id, pagenumber, pagesize) }
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransporters V1
     */
    suspend fun transfersGetDeliveryTransportersV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryTransportersV1(id, no) }
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransporters V2
     */
    suspend fun transfersGetDeliveryTransportersV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryTransportersV2(id, pagenumber, pagesize) }
    }

    /**
     * Please note: The {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransportersDetails V1
     */
    suspend fun transfersGetDeliveryTransportersDetailsV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryTransportersDetailsV1(id, no) }
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransportersDetails V2
     */
    suspend fun transfersGetDeliveryTransportersDetailsV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetDeliveryTransportersDetailsV2(id, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetHub V2
     */
    suspend fun transfersGetHubV2(estimatedarrivalend: String? = null, estimatedarrivalstart: String? = null, lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetHubV2(estimatedarrivalend, estimatedarrivalstart, lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetIncoming V1
     */
    suspend fun transfersGetIncomingV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetIncomingV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetIncoming V2
     */
    suspend fun transfersGetIncomingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetIncomingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetOutgoing V1
     */
    suspend fun transfersGetOutgoingV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetOutgoingV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetOutgoing V2
     */
    suspend fun transfersGetOutgoingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetRejected V1
     */
    suspend fun transfersGetRejectedV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetRejectedV1(licensenumber) }
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetRejected V2
     */
    suspend fun transfersGetRejectedV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetRejectedV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplates V1
     */
    suspend fun transfersGetTemplatesV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDelivery V1
     */
    suspend fun transfersGetTemplatesDeliveryV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesDeliveryV1(id, no) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetTemplatesDeliveryPackage V1
     */
    suspend fun transfersGetTemplatesDeliveryPackageV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesDeliveryPackageV1(id, no) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransporters V1
     */
    suspend fun transfersGetTemplatesDeliveryTransportersV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesDeliveryTransportersV1(id, no) }
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransportersDetails V1
     */
    suspend fun transfersGetTemplatesDeliveryTransportersDetailsV1(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesDeliveryTransportersDetailsV1(id, no) }
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoing V2
     */
    suspend fun transfersGetTemplatesOutgoingV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDelivery V2
     */
    suspend fun transfersGetTemplatesOutgoingDeliveryV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesOutgoingDeliveryV2(id, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryPackage V2
     */
    suspend fun transfersGetTemplatesOutgoingDeliveryPackageV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesOutgoingDeliveryPackageV2(id, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransporters V2
     */
    suspend fun transfersGetTemplatesOutgoingDeliveryTransportersV2(id: String, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesOutgoingDeliveryTransportersV2(id, pagenumber, pagesize) }
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransportersDetails V2
     */
    suspend fun transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id: String, no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, no) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    suspend fun transfersGetTypesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTypesV1(licensenumber) }
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    suspend fun transfersGetTypesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transfersGetTypesV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * PUT UpdateExternalIncoming V1
     */
    suspend fun transfersUpdateExternalIncomingV1(licensenumber: String? = null, body: List<TransfersUpdateExternalIncomingV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersUpdateExternalIncomingV1(licensenumber, body) }
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * PUT UpdateExternalIncoming V2
     */
    suspend fun transfersUpdateExternalIncomingV2(licensenumber: String? = null, body: List<TransfersUpdateExternalIncomingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersUpdateExternalIncomingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * PUT UpdateTemplates V1
     */
    suspend fun transfersUpdateTemplatesV1(licensenumber: String? = null, body: List<TransfersUpdateTemplatesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersUpdateTemplatesV1(licensenumber, body) }
    }

    /**
     * Updates existing transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * PUT UpdateTemplatesOutgoing V2
     */
    suspend fun transfersUpdateTemplatesOutgoingV2(licensenumber: String? = null, body: List<TransfersUpdateTemplatesOutgoingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transfersUpdateTemplatesOutgoingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     */
    suspend fun plantsCreateAdditivesV1(licensenumber: String? = null, body: List<PlantsCreateAdditivesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesV1(licensenumber, body) }
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     */
    suspend fun plantsCreateAdditivesV2(licensenumber: String? = null, body: List<PlantsCreateAdditivesV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V1
     */
    suspend fun plantsCreateAdditivesBylocationV1(licensenumber: String? = null, body: List<PlantsCreateAdditivesBylocationV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesBylocationV1(licensenumber, body) }
    }

    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V2
     */
    suspend fun plantsCreateAdditivesBylocationV2(licensenumber: String? = null, body: List<PlantsCreateAdditivesBylocationV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesBylocationV2(licensenumber, body) }
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocationUsingtemplate V2
     */
    suspend fun plantsCreateAdditivesBylocationUsingtemplateV2(licensenumber: String? = null, body: List<PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesBylocationUsingtemplateV2(licensenumber, body) }
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     */
    suspend fun plantsCreateAdditivesUsingtemplateV2(licensenumber: String? = null, body: List<PlantsCreateAdditivesUsingtemplateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateAdditivesUsingtemplateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphases V1
     */
    suspend fun plantsCreateChangegrowthphasesV1(licensenumber: String? = null, body: List<PlantsCreateChangegrowthphasesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateChangegrowthphasesV1(licensenumber, body) }
    }

    /**
     * NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateHarvestplants V1
     */
    suspend fun plantsCreateHarvestplantsV1(licensenumber: String? = null, body: List<PlantsCreateHarvestplantsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateHarvestplantsV1(licensenumber, body) }
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicure V2
     */
    suspend fun plantsCreateManicureV2(licensenumber: String? = null, body: List<PlantsCreateManicureV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateManicureV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicureplants V1
     */
    suspend fun plantsCreateManicureplantsV1(licensenumber: String? = null, body: List<PlantsCreateManicureplantsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateManicureplantsV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateMoveplants V1
     */
    suspend fun plantsCreateMoveplantsV1(licensenumber: String? = null, body: List<PlantsCreateMoveplantsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateMoveplantsV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePlantbatchPackage V1
     */
    suspend fun plantsCreatePlantbatchPackageV1(licensenumber: String? = null, body: List<PlantsCreatePlantbatchPackageV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreatePlantbatchPackageV1(licensenumber, body) }
    }

    /**
     * Creates packages from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePlantbatchPackage V2
     */
    suspend fun plantsCreatePlantbatchPackageV2(licensenumber: String? = null, body: List<PlantsCreatePlantbatchPackageV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreatePlantbatchPackageV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V1
     */
    suspend fun plantsCreatePlantingsV1(licensenumber: String? = null, body: List<PlantsCreatePlantingsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreatePlantingsV1(licensenumber, body) }
    }

    /**
     * Creates new plant batches at a specified Facility from existing plant data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V2
     */
    suspend fun plantsCreatePlantingsV2(licensenumber: String? = null, body: List<PlantsCreatePlantingsV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreatePlantingsV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     */
    suspend fun plantsCreateWasteV1(licensenumber: String? = null, body: List<PlantsCreateWasteV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateWasteV1(licensenumber, body) }
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     */
    suspend fun plantsCreateWasteV2(licensenumber: String? = null, body: List<PlantsCreateWasteV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsCreateWasteV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V1
     */
    suspend fun plantsDeleteV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.plantsDeleteV1(licensenumber) }
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V2
     */
    suspend fun plantsDeleteV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.plantsDeleteV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V1
     */
    suspend fun plantsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Plant by Id.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V2
     */
    suspend fun plantsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V1
     */
    suspend fun plantsGetAdditivesV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetAdditivesV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V2
     */
    suspend fun plantsGetAdditivesV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetAdditivesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetAdditivesTypes V1
     */
    suspend fun plantsGetAdditivesTypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetAdditivesTypesV1(no) }
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdditivesTypes V2
     */
    suspend fun plantsGetAdditivesTypesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetAdditivesTypesV2(no) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V1
     */
    suspend fun plantsGetByLabelV1(label: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetByLabelV1(label, licensenumber) }
    }

    /**
     * Retrieves a Plant by label.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V2
     */
    suspend fun plantsGetByLabelV2(label: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetByLabelV2(label, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V1
     */
    suspend fun plantsGetFloweringV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetFloweringV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V2
     */
    suspend fun plantsGetFloweringV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetFloweringV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V1
     */
    suspend fun plantsGetGrowthPhasesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetGrowthPhasesV1(licensenumber) }
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V2
     */
    suspend fun plantsGetGrowthPhasesV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetGrowthPhasesV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V1
     */
    suspend fun plantsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V2
     */
    suspend fun plantsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMother V2
     */
    suspend fun plantsGetMotherV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetMotherV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherInactive V2
     */
    suspend fun plantsGetMotherInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetMotherInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherOnhold V2
     */
    suspend fun plantsGetMotherOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetMotherOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V1
     */
    suspend fun plantsGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V2
     */
    suspend fun plantsGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V1
     */
    suspend fun plantsGetVegetativeV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetVegetativeV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V2
     */
    suspend fun plantsGetVegetativeV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetVegetativeV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     */
    suspend fun plantsGetWasteV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWasteV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V1
     */
    suspend fun plantsGetWasteMethodsAllV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWasteMethodsAllV1(no) }
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V2
     */
    suspend fun plantsGetWasteMethodsAllV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWasteMethodsAllV2(pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePackage V2
     */
    suspend fun plantsGetWastePackageV2(id: String, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWastePackageV2(id, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePlant V2
     */
    suspend fun plantsGetWastePlantV2(id: String, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWastePlantV2(id, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     */
    suspend fun plantsGetWasteReasonsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWasteReasonsV1(licensenumber) }
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     */
    suspend fun plantsGetWasteReasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantsGetWasteReasonsV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateAdjust V2
     */
    suspend fun plantsUpdateAdjustV2(licensenumber: String? = null, body: List<PlantsUpdateAdjustV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateAdjustV2(licensenumber, body) }
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateGrowthphase V2
     */
    suspend fun plantsUpdateGrowthphaseV2(licensenumber: String? = null, body: List<PlantsUpdateGrowthphaseV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateGrowthphaseV2(licensenumber, body) }
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateHarvest V2
     */
    suspend fun plantsUpdateHarvestV2(licensenumber: String? = null, body: List<PlantsUpdateHarvestV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateHarvestV2(licensenumber, body) }
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateLocation V2
     */
    suspend fun plantsUpdateLocationV2(licensenumber: String? = null, body: List<PlantsUpdateLocationV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateLocationV2(licensenumber, body) }
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateMerge V2
     */
    suspend fun plantsUpdateMergeV2(licensenumber: String? = null, body: List<PlantsUpdateMergeV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateMergeV2(licensenumber, body) }
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * 
     *   Permissions Required:
     *   - View Plant
     *
     * PUT UpdateSplit V2
     */
    suspend fun plantsUpdateSplitV2(licensenumber: String? = null, body: List<PlantsUpdateSplitV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateSplitV2(licensenumber, body) }
    }

    /**
     * Updates the strain information for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     */
    suspend fun plantsUpdateStrainV2(licensenumber: String? = null, body: List<PlantsUpdateStrainV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateStrainV2(licensenumber, body) }
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     */
    suspend fun plantsUpdateTagV2(licensenumber: String? = null, body: List<PlantsUpdateTagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantsUpdateTagV2(licensenumber, body) }
    }

    /**
     * Creates new additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * POST Create V2
     */
    suspend fun additivesTemplatesCreateV2(licensenumber: String? = null, body: List<AdditivestemplatesCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.additivesTemplatesCreateV2(licensenumber, body) }
    }

    /**
     * Retrieves an Additive Template by its Id.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET Get V2
     */
    suspend fun additivesTemplatesGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.additivesTemplatesGetV2(id, licensenumber) }
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetActive V2
     */
    suspend fun additivesTemplatesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.additivesTemplatesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetInactive V2
     */
    suspend fun additivesTemplatesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.additivesTemplatesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * PUT Update V2
     */
    suspend fun additivesTemplatesUpdateV2(licensenumber: String? = null, body: List<AdditivestemplatesUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.additivesTemplatesUpdateV2(licensenumber, body) }
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V1
     */
    suspend fun caregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.caregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber, licensenumber) }
    }

    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V2
     */
    suspend fun caregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.caregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateFinish V1
     */
    suspend fun harvestsCreateFinishV1(licensenumber: String? = null, body: List<HarvestsCreateFinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreateFinishV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V1
     */
    suspend fun harvestsCreatePackageV1(licensenumber: String? = null, body: List<HarvestsCreatePackageV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreatePackageV1(licensenumber, body) }
    }

    /**
     * Creates packages from harvested products for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V2
     */
    suspend fun harvestsCreatePackageV2(licensenumber: String? = null, body: List<HarvestsCreatePackageV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreatePackageV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V1
     */
    suspend fun harvestsCreatePackageTestingV1(licensenumber: String? = null, body: List<HarvestsCreatePackageTestingV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreatePackageTestingV1(licensenumber, body) }
    }

    /**
     * Creates packages for testing from harvested products for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V2
     */
    suspend fun harvestsCreatePackageTestingV2(licensenumber: String? = null, body: List<HarvestsCreatePackageTestingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreatePackageTestingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateRemoveWaste V1
     */
    suspend fun harvestsCreateRemoveWasteV1(licensenumber: String? = null, body: List<HarvestsCreateRemoveWasteV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreateRemoveWasteV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateUnfinish V1
     */
    suspend fun harvestsCreateUnfinishV1(licensenumber: String? = null, body: List<HarvestsCreateUnfinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreateUnfinishV1(licensenumber, body) }
    }

    /**
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateWaste V2
     */
    suspend fun harvestsCreateWasteV2(licensenumber: String? = null, body: List<HarvestsCreateWasteV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsCreateWasteV2(licensenumber, body) }
    }

    /**
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Discontinue Harvest Waste
     *
     * DELETE DeleteWaste V2
     */
    suspend fun harvestsDeleteWasteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.harvestsDeleteWasteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET Get V1
     */
    suspend fun harvestsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET Get V2
     */
    suspend fun harvestsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V1
     */
    suspend fun harvestsGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of active harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V2
     */
    suspend fun harvestsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V1
     */
    suspend fun harvestsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of inactive harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V2
     */
    suspend fun harvestsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V1
     */
    suspend fun harvestsGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of harvests on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V2
     */
    suspend fun harvestsGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetWaste V2
     */
    suspend fun harvestsGetWasteV2(harvestid: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetWasteV2(harvestid, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V1
     */
    suspend fun harvestsGetWasteTypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetWasteTypesV1(no) }
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V2
     */
    suspend fun harvestsGetWasteTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.harvestsGetWasteTypesV2(pagenumber, pagesize) }
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateFinish V2
     */
    suspend fun harvestsUpdateFinishV2(licensenumber: String? = null, body: List<HarvestsUpdateFinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateFinishV2(licensenumber, body) }
    }

    /**
     * Updates the Location of Harvest for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateLocation V2
     */
    suspend fun harvestsUpdateLocationV2(licensenumber: String? = null, body: List<HarvestsUpdateLocationV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateLocationV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateMove V1
     */
    suspend fun harvestsUpdateMoveV1(licensenumber: String? = null, body: List<HarvestsUpdateMoveV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateMoveV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V1
     */
    suspend fun harvestsUpdateRenameV1(licensenumber: String? = null, body: List<HarvestsUpdateRenameV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateRenameV1(licensenumber, body) }
    }

    /**
     * Renames one or more harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V2
     */
    suspend fun harvestsUpdateRenameV2(licensenumber: String? = null, body: List<HarvestsUpdateRenameV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateRenameV2(licensenumber, body) }
    }

    /**
     * Restores previously harvested plants to their original state for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateRestoreHarvestedPlants V2
     */
    suspend fun harvestsUpdateRestoreHarvestedPlantsV2(licensenumber: String? = null, body: List<HarvestsUpdateRestoreHarvestedPlantsV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateRestoreHarvestedPlantsV2(licensenumber, body) }
    }

    /**
     * Reopens one or more previously finished harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateUnfinish V2
     */
    suspend fun harvestsUpdateUnfinishV2(licensenumber: String? = null, body: List<HarvestsUpdateUnfinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.harvestsUpdateUnfinishV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V1
     */
    suspend fun labTestsCreateRecordV1(licensenumber: String? = null, body: List<LabtestsCreateRecordV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsCreateRecordV1(licensenumber, body) }
    }

    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V2
     */
    suspend fun labTestsCreateRecordV2(licensenumber: String? = null, body: List<LabtestsCreateRecordV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsCreateRecordV2(licensenumber, body) }
    }

    /**
     * Retrieves a list of Lab Test batches.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetBatches V2
     */
    suspend fun labTestsGetBatchesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetBatchesV2(pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V1
     */
    suspend fun labTestsGetLabtestdocumentV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetLabtestdocumentV1(id, licensenumber) }
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V2
     */
    suspend fun labTestsGetLabtestdocumentV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetLabtestdocumentV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetResults V1
     */
    suspend fun labTestsGetResultsV1(licensenumber: String? = null, packageid: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetResultsV1(licensenumber, packageid) }
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetResults V2
     */
    suspend fun labTestsGetResultsV2(licensenumber: String? = null, packageid: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetResultsV2(licensenumber, packageid, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetStates V1
     */
    suspend fun labTestsGetStatesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetStatesV1(no) }
    }

    /**
     * Returns a list of all lab testing states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetStates V2
     */
    suspend fun labTestsGetStatesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetStatesV2(no) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    suspend fun labTestsGetTypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetTypesV1(no) }
    }

    /**
     * Returns a list of Lab Test types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    suspend fun labTestsGetTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.labTestsGetTypesV2(pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V1
     */
    suspend fun labTestsUpdateLabtestdocumentV1(licensenumber: String? = null, body: List<LabtestsUpdateLabtestdocumentV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsUpdateLabtestdocumentV1(licensenumber, body) }
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V2
     */
    suspend fun labTestsUpdateLabtestdocumentV2(licensenumber: String? = null, body: List<LabtestsUpdateLabtestdocumentV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsUpdateLabtestdocumentV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V1
     */
    suspend fun labTestsUpdateResultReleaseV1(licensenumber: String? = null, body: List<LabtestsUpdateResultReleaseV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsUpdateResultReleaseV1(licensenumber, body) }
    }

    /**
     * Releases Lab Test results for one or more packages.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V2
     */
    suspend fun labTestsUpdateResultReleaseV2(licensenumber: String? = null, body: List<LabtestsUpdateResultReleaseV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.labTestsUpdateResultReleaseV2(licensenumber, body) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * POST CreateDelivery V1
     */
    suspend fun salesCreateDeliveryV1(licensenumber: String? = null, body: List<SalesCreateDeliveryV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryV1(licensenumber, body) }
    }

    /**
     * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateDelivery V2
     */
    suspend fun salesCreateDeliveryV2(licensenumber: String? = null, body: List<SalesCreateDeliveryV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryV2(licensenumber, body) }
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailer V1
     */
    suspend fun salesCreateDeliveryRetailerV1(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerV1(licensenumber, body) }
    }

    /**
     * Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailer V2
     */
    suspend fun salesCreateDeliveryRetailerV2(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V1
     */
    suspend fun salesCreateDeliveryRetailerDepartV1(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerDepartV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerDepartV1(licensenumber, body) }
    }

    /**
     * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V2
     */
    suspend fun salesCreateDeliveryRetailerDepartV2(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerDepartV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerDepartV2(licensenumber, body) }
    }

    /**
     * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V1
     */
    suspend fun salesCreateDeliveryRetailerEndV1(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerEndV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerEndV1(licensenumber, body) }
    }

    /**
     * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V2
     */
    suspend fun salesCreateDeliveryRetailerEndV2(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerEndV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerEndV2(licensenumber, body) }
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V1
     */
    suspend fun salesCreateDeliveryRetailerRestockV1(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerRestockV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerRestockV1(licensenumber, body) }
    }

    /**
     * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V2
     */
    suspend fun salesCreateDeliveryRetailerRestockV2(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerRestockV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerRestockV2(licensenumber, body) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerSale V1
     */
    suspend fun salesCreateDeliveryRetailerSaleV1(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerSaleV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerSaleV1(licensenumber, body) }
    }

    /**
     * Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateDeliveryRetailerSale V2
     */
    suspend fun salesCreateDeliveryRetailerSaleV2(licensenumber: String? = null, body: List<SalesCreateDeliveryRetailerSaleV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateDeliveryRetailerSaleV2(licensenumber, body) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * POST CreateReceipt V1
     */
    suspend fun salesCreateReceiptV1(licensenumber: String? = null, body: List<SalesCreateReceiptV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateReceiptV1(licensenumber, body) }
    }

    /**
     * Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales (Write)
     *   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
     *   - Industry/Facility Type/Advanced Sales
     *   - WebApi Sales Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *
     * POST CreateReceipt V2
     */
    suspend fun salesCreateReceiptV2(licensenumber: String? = null, body: List<SalesCreateReceiptV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateReceiptV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * POST CreateTransactionByDate V1
     */
    suspend fun salesCreateTransactionByDateV1(date: String, licensenumber: String? = null, body: List<SalesCreateTransactionByDateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesCreateTransactionByDateV1(date, licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * DELETE DeleteDelivery V1
     */
    suspend fun salesDeleteDeliveryV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteDeliveryV1(id, licensenumber) }
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * DELETE DeleteDelivery V2
     */
    suspend fun salesDeleteDeliveryV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteDeliveryV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V1
     */
    suspend fun salesDeleteDeliveryRetailerV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteDeliveryRetailerV1(id, licensenumber) }
    }

    /**
     * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - Manage Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V2
     */
    suspend fun salesDeleteDeliveryRetailerV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteDeliveryRetailerV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * DELETE DeleteReceipt V1
     */
    suspend fun salesDeleteReceiptV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteReceiptV1(id, licensenumber) }
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * DELETE DeleteReceipt V2
     */
    suspend fun salesDeleteReceiptV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.salesDeleteReceiptV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCounties V1
     */
    suspend fun salesGetCountiesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetCountiesV1(no) }
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCounties V2
     */
    suspend fun salesGetCountiesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetCountiesV2(no) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V1
     */
    suspend fun salesGetCustomertypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetCustomertypesV1(no) }
    }

    /**
     * Returns a list of customer types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V2
     */
    suspend fun salesGetCustomertypesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetCustomertypesV2(no) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesActive V1
     */
    suspend fun salesGetDeliveriesActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart) }
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesActive V2
     */
    suspend fun salesGetDeliveriesActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesInactive V1
     */
    suspend fun salesGetDeliveriesInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart) }
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesInactive V2
     */
    suspend fun salesGetDeliveriesInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart) }
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V1
     */
    suspend fun salesGetDeliveriesRetailerActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesRetailerActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V2
     */
    suspend fun salesGetDeliveriesRetailerActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesRetailerActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V1
     */
    suspend fun salesGetDeliveriesRetailerInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesRetailerInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V2
     */
    suspend fun salesGetDeliveriesRetailerInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesRetailerInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetDeliveriesReturnreasons V1
     */
    suspend fun salesGetDeliveriesReturnreasonsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesReturnreasonsV1(licensenumber) }
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesReturnreasons V2
     */
    suspend fun salesGetDeliveriesReturnreasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveriesReturnreasonsV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDelivery V1
     */
    suspend fun salesGetDeliveryV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveryV1(id, licensenumber) }
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDelivery V2
     */
    suspend fun salesGetDeliveryV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveryV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveryRetailer V1
     */
    suspend fun salesGetDeliveryRetailerV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveryRetailerV1(id, licensenumber) }
    }

    /**
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveryRetailer V2
     */
    suspend fun salesGetDeliveryRetailerV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetDeliveryRetailerV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V1
     */
    suspend fun salesGetPatientRegistrationsLocationsV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetPatientRegistrationsLocationsV1(no) }
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * 
     *   Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V2
     */
    suspend fun salesGetPatientRegistrationsLocationsV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetPatientRegistrationsLocationsV2(no) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetPaymenttypes V1
     */
    suspend fun salesGetPaymenttypesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetPaymenttypesV1(licensenumber) }
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetPaymenttypes V2
     */
    suspend fun salesGetPaymenttypesV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetPaymenttypesV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceipt V1
     */
    suspend fun salesGetReceiptV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptV1(id, licensenumber) }
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceipt V2
     */
    suspend fun salesGetReceiptV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsActive V1
     */
    suspend fun salesGetReceiptsActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptsActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart) }
    }

    /**
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsActive V2
     */
    suspend fun salesGetReceiptsActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptsActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart) }
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsExternalByExternalNumber V2
     */
    suspend fun salesGetReceiptsExternalByExternalNumberV2(externalnumber: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptsExternalByExternalNumberV2(externalnumber, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsInactive V1
     */
    suspend fun salesGetReceiptsInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptsInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart) }
    }

    /**
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsInactive V2
     */
    suspend fun salesGetReceiptsInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null, salesdateend: String? = null, salesdatestart: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetReceiptsInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactions V1
     */
    suspend fun salesGetTransactionsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetTransactionsV1(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
     */
    suspend fun salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart: String, salesdateend: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart, salesdateend, licensenumber) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDelivery V1
     */
    suspend fun salesUpdateDeliveryV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryV1(licensenumber, body) }
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDelivery V2
     */
    suspend fun salesUpdateDeliveryV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryComplete V1
     */
    suspend fun salesUpdateDeliveryCompleteV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryCompleteV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryCompleteV1(licensenumber, body) }
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDeliveryComplete V2
     */
    suspend fun salesUpdateDeliveryCompleteV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryCompleteV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryCompleteV2(licensenumber, body) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryHub V1
     */
    suspend fun salesUpdateDeliveryHubV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubV1(licensenumber, body) }
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery, Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHub V2
     */
    suspend fun salesUpdateDeliveryHubV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubAccept V1
     */
    suspend fun salesUpdateDeliveryHubAcceptV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubAcceptV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubAcceptV1(licensenumber, body) }
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubAccept V2
     */
    suspend fun salesUpdateDeliveryHubAcceptV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubAcceptV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubAcceptV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubDepart V1
     */
    suspend fun salesUpdateDeliveryHubDepartV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubDepartV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubDepartV1(licensenumber, body) }
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubDepart V2
     */
    suspend fun salesUpdateDeliveryHubDepartV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubDepartV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubDepartV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubVerifyID V1
     */
    suspend fun salesUpdateDeliveryHubVerifyIdV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubVerifyIdV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubVerifyIdV1(licensenumber, body) }
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubVerifyID V2
     */
    suspend fun salesUpdateDeliveryHubVerifyIdV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryHubVerifyIdV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryHubVerifyIdV2(licensenumber, body) }
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V1
     */
    suspend fun salesUpdateDeliveryRetailerV1(licensenumber: String? = null, body: List<SalesUpdateDeliveryRetailerV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryRetailerV1(licensenumber, body) }
    }

    /**
     * Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     *   - Industry/Facility Type/Retailer Delivery
     *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     *   - Manage Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V2
     */
    suspend fun salesUpdateDeliveryRetailerV2(licensenumber: String? = null, body: List<SalesUpdateDeliveryRetailerV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateDeliveryRetailerV2(licensenumber, body) }
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * PUT UpdateReceipt V1
     */
    suspend fun salesUpdateReceiptV1(licensenumber: String? = null, body: List<SalesUpdateReceiptV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateReceiptV1(licensenumber, body) }
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceipt V2
     */
    suspend fun salesUpdateReceiptV2(licensenumber: String? = null, body: List<SalesUpdateReceiptV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateReceiptV2(licensenumber, body) }
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptFinalize V2
     */
    suspend fun salesUpdateReceiptFinalizeV2(licensenumber: String? = null, body: List<SalesUpdateReceiptFinalizeV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateReceiptFinalizeV2(licensenumber, body) }
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptUnfinalize V2
     */
    suspend fun salesUpdateReceiptUnfinalizeV2(licensenumber: String? = null, body: List<SalesUpdateReceiptUnfinalizeV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateReceiptUnfinalizeV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateTransactionByDate V1
     */
    suspend fun salesUpdateTransactionByDateV1(date: String, licensenumber: String? = null, body: List<SalesUpdateTransactionByDateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.salesUpdateTransactionByDateV1(date, licensenumber, body) }
    }

    /**
     * Creates new driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateDriver V2
     */
    suspend fun transportersCreateDriverV2(licensenumber: String? = null, body: List<TransportersCreateDriverV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transportersCreateDriverV2(licensenumber, body) }
    }

    /**
     * Creates new vehicle records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateVehicle V2
     */
    suspend fun transportersCreateVehicleV2(licensenumber: String? = null, body: List<TransportersCreateVehicleV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transportersCreateVehicleV2(licensenumber, body) }
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteDriver V2
     */
    suspend fun transportersDeleteDriverV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transportersDeleteDriverV2(id, licensenumber) }
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteVehicle V2
     */
    suspend fun transportersDeleteVehicleV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.transportersDeleteVehicleV2(id, licensenumber) }
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDriver V2
     */
    suspend fun transportersGetDriverV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transportersGetDriverV2(id, licensenumber) }
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDrivers V2
     */
    suspend fun transportersGetDriversV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transportersGetDriversV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicle V2
     */
    suspend fun transportersGetVehicleV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transportersGetVehicleV2(id, licensenumber) }
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicles V2
     */
    suspend fun transportersGetVehiclesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.transportersGetVehiclesV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates existing driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateDriver V2
     */
    suspend fun transportersUpdateDriverV2(licensenumber: String? = null, body: List<TransportersUpdateDriverV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transportersUpdateDriverV2(licensenumber, body) }
    }

    /**
     * Updates existing vehicle records for a facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateVehicle V2
     */
    suspend fun transportersUpdateVehicleV2(licensenumber: String? = null, body: List<TransportersUpdateVehicleV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.transportersUpdateVehicleV2(licensenumber, body) }
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V1
     */
    suspend fun facilitiesGetAllV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.facilitiesGetAllV1(no) }
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     */
    suspend fun facilitiesGetAllV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.facilitiesGetAllV2(no) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V1
     */
    suspend fun packagesCreateV1(licensenumber: String? = null, body: List<PackagesCreateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateV1(licensenumber, body) }
    }

    /**
     * Creates new packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V2
     */
    suspend fun packagesCreateV2(licensenumber: String? = null, body: List<PackagesCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V1
     */
    suspend fun packagesCreateAdjustV1(licensenumber: String? = null, body: List<PackagesCreateAdjustV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateAdjustV1(licensenumber, body) }
    }

    /**
     * Records a list of adjustments for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V2
     */
    suspend fun packagesCreateAdjustV2(licensenumber: String? = null, body: List<PackagesCreateAdjustV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateAdjustV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeItem V1
     */
    suspend fun packagesCreateChangeItemV1(licensenumber: String? = null, body: List<PackagesCreateChangeItemV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateChangeItemV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeLocation V1
     */
    suspend fun packagesCreateChangeLocationV1(licensenumber: String? = null, body: List<PackagesCreateChangeLocationV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateChangeLocationV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateFinish V1
     */
    suspend fun packagesCreateFinishV1(licensenumber: String? = null, body: List<PackagesCreateFinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateFinishV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V1
     */
    suspend fun packagesCreatePlantingsV1(licensenumber: String? = null, body: List<PackagesCreatePlantingsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreatePlantingsV1(licensenumber, body) }
    }

    /**
     * Creates new plantings from packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V2
     */
    suspend fun packagesCreatePlantingsV2(licensenumber: String? = null, body: List<PackagesCreatePlantingsV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreatePlantingsV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRemediate V1
     */
    suspend fun packagesCreateRemediateV1(licensenumber: String? = null, body: List<PackagesCreateRemediateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateRemediateV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V1
     */
    suspend fun packagesCreateTestingV1(licensenumber: String? = null, body: List<PackagesCreateTestingV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateTestingV1(licensenumber, body) }
    }

    /**
     * Creates new packages for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V2
     */
    suspend fun packagesCreateTestingV2(licensenumber: String? = null, body: List<PackagesCreateTestingV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateTestingV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateUnfinish V1
     */
    suspend fun packagesCreateUnfinishV1(licensenumber: String? = null, body: List<PackagesCreateUnfinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesCreateUnfinishV1(licensenumber, body) }
    }

    /**
     * Discontinues a Package at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * DELETE Delete V2
     */
    suspend fun packagesDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.packagesDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET Get V1
     */
    suspend fun packagesGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET Get V2
     */
    suspend fun packagesGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetActive V1
     */
    suspend fun packagesGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of active packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetActive V2
     */
    suspend fun packagesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V1
     */
    suspend fun packagesGetAdjustReasonsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetAdjustReasonsV1(licensenumber) }
    }

    /**
     * Retrieves a list of adjustment reasons for packages at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V2
     */
    suspend fun packagesGetAdjustReasonsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetAdjustReasonsV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V1
     */
    suspend fun packagesGetByLabelV1(label: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetByLabelV1(label, licensenumber) }
    }

    /**
     * Retrieves a Package by its label.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V2
     */
    suspend fun packagesGetByLabelV2(label: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetByLabelV2(label, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V1
     */
    suspend fun packagesGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of inactive packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V2
     */
    suspend fun packagesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of packages in transit for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetIntransit V2
     */
    suspend fun packagesGetIntransitV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetIntransitV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetLabsamples V2
     */
    suspend fun packagesGetLabsamplesV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetLabsamplesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V1
     */
    suspend fun packagesGetOnholdV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of packages on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V2
     */
    suspend fun packagesGetOnholdV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves the source harvests for a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Package Source Harvests
     *
     * GET GetSourceHarvest V2
     */
    suspend fun packagesGetSourceHarvestV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetSourceHarvestV2(id, licensenumber) }
    }

    /**
     * Retrieves a list of transferred packages for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetTransferred V2
     */
    suspend fun packagesGetTransferredV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetTransferredV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    suspend fun packagesGetTypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetTypesV1(no) }
    }

    /**
     * Retrieves a list of available Package types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    suspend fun packagesGetTypesV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.packagesGetTypesV2(no) }
    }

    /**
     * Set the final quantity for a Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateAdjust V2
     */
    suspend fun packagesUpdateAdjustV2(licensenumber: String? = null, body: List<PackagesUpdateAdjustV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateAdjustV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateChangeNote V1
     */
    suspend fun packagesUpdateChangeNoteV1(licensenumber: String? = null, body: List<PackagesUpdateChangeNoteV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateChangeNoteV1(licensenumber, body) }
    }

    /**
     * Updates the Product decontaminate information for a list of packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDecontaminate V2
     */
    suspend fun packagesUpdateDecontaminateV2(licensenumber: String? = null, body: List<PackagesUpdateDecontaminateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateDecontaminateV2(licensenumber, body) }
    }

    /**
     * Flags one or more packages for donation at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationFlag V2
     */
    suspend fun packagesUpdateDonationFlagV2(licensenumber: String? = null, body: List<PackagesUpdateDonationFlagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateDonationFlagV2(licensenumber, body) }
    }

    /**
     * Removes the donation flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationUnflag V2
     */
    suspend fun packagesUpdateDonationUnflagV2(licensenumber: String? = null, body: List<PackagesUpdateDonationUnflagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateDonationUnflagV2(licensenumber, body) }
    }

    /**
     * Updates the external identifiers for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Package Inventory
     *   - External Id Enabled
     *
     * PUT UpdateExternalid V2
     */
    suspend fun packagesUpdateExternalidV2(licensenumber: String? = null, body: List<PackagesUpdateExternalidV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateExternalidV2(licensenumber, body) }
    }

    /**
     * Updates a list of packages as finished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateFinish V2
     */
    suspend fun packagesUpdateFinishV2(licensenumber: String? = null, body: List<PackagesUpdateFinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateFinishV2(licensenumber, body) }
    }

    /**
     * Updates the associated Item for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateItem V2
     */
    suspend fun packagesUpdateItemV2(licensenumber: String? = null, body: List<PackagesUpdateItemV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateItemV2(licensenumber, body) }
    }

    /**
     * Updates the list of required lab test batches for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLabTestRequired V2
     */
    suspend fun packagesUpdateLabTestRequiredV2(licensenumber: String? = null, body: List<PackagesUpdateLabTestRequiredV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateLabTestRequiredV2(licensenumber, body) }
    }

    /**
     * Updates the Location and Sublocation for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLocation V2
     */
    suspend fun packagesUpdateLocationV2(licensenumber: String? = null, body: List<PackagesUpdateLocationV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateLocationV2(licensenumber, body) }
    }

    /**
     * Updates notes associated with one or more packages for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateNote V2
     */
    suspend fun packagesUpdateNoteV2(licensenumber: String? = null, body: List<PackagesUpdateNoteV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateNoteV2(licensenumber, body) }
    }

    /**
     * Updates a list of Product remediations for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateRemediate V2
     */
    suspend fun packagesUpdateRemediateV2(licensenumber: String? = null, body: List<PackagesUpdateRemediateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateRemediateV2(licensenumber, body) }
    }

    /**
     * Flags or unflags one or more packages at the specified Facility as trade samples.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleFlag V2
     */
    suspend fun packagesUpdateTradesampleFlagV2(licensenumber: String? = null, body: List<PackagesUpdateTradesampleFlagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateTradesampleFlagV2(licensenumber, body) }
    }

    /**
     * Removes the trade sample flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleUnflag V2
     */
    suspend fun packagesUpdateTradesampleUnflagV2(licensenumber: String? = null, body: List<PackagesUpdateTradesampleUnflagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateTradesampleUnflagV2(licensenumber, body) }
    }

    /**
     * Updates a list of packages as unfinished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateUnfinish V2
     */
    suspend fun packagesUpdateUnfinishV2(licensenumber: String? = null, body: List<PackagesUpdateUnfinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateUnfinishV2(licensenumber, body) }
    }

    /**
     * Updates the use-by date for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateUsebydate V2
     */
    suspend fun packagesUpdateUsebydateV2(licensenumber: String? = null, body: List<PackagesUpdateUsebydateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.packagesUpdateUsebydateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V1
     */
    suspend fun patientCheckInsCreateV1(licensenumber: String? = null, body: List<PatientcheckinsCreateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsCreateV1(licensenumber, body) }
    }

    /**
     * Records patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V2
     */
    suspend fun patientCheckInsCreateV2(licensenumber: String? = null, body: List<PatientcheckinsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsCreateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V1
     */
    suspend fun patientCheckInsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsDeleteV1(id, licensenumber) }
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V2
     */
    suspend fun patientCheckInsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V1
     */
    suspend fun patientCheckInsGetAllV1(checkindateend: String? = null, checkindatestart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientCheckInsGetAllV1(checkindateend, checkindatestart, licensenumber) }
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V2
     */
    suspend fun patientCheckInsGetAllV2(checkindateend: String? = null, checkindatestart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientCheckInsGetAllV2(checkindateend, checkindatestart, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetLocations V1
     */
    suspend fun patientCheckInsGetLocationsV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientCheckInsGetLocationsV1(no) }
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetLocations V2
     */
    suspend fun patientCheckInsGetLocationsV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientCheckInsGetLocationsV2(no) }
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V1
     */
    suspend fun patientCheckInsUpdateV1(licensenumber: String? = null, body: List<PatientcheckinsUpdateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsUpdateV1(licensenumber, body) }
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V2
     */
    suspend fun patientCheckInsUpdateV2(licensenumber: String? = null, body: List<PatientcheckinsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientCheckInsUpdateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetActive V1
     */
    suspend fun unitsOfMeasureGetActiveV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.unitsOfMeasureGetActiveV1(no) }
    }

    /**
     * Retrieves all active units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetActive V2
     */
    suspend fun unitsOfMeasureGetActiveV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.unitsOfMeasureGetActiveV2(no) }
    }

    /**
     * Retrieves all inactive units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetInactive V2
     */
    suspend fun unitsOfMeasureGetInactiveV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.unitsOfMeasureGetInactiveV2(no) }
    }

    /**
     * Retrieves all available waste methods.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     */
    suspend fun wasteMethodsGetAllV2(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.wasteMethodsGetAllV2(no) }
    }

    /**
     * Permissions Required:
     *   - Manage Employees
     *
     * GET GetAll V1
     */
    suspend fun employeesGetAllV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.employeesGetAllV1(licensenumber) }
    }

    /**
     * Retrieves a list of employees for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *   - View Employees
     *
     * GET GetAll V2
     */
    suspend fun employeesGetAllV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.employeesGetAllV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *
     * GET GetPermissions V2
     */
    suspend fun employeesGetPermissionsV2(employeelicensenumber: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.employeesGetPermissionsV2(employeelicensenumber, licensenumber) }
    }

    /**
     * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V1
     */
    suspend fun itemsCreateV1(licensenumber: String? = null, body: List<ItemsCreateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreateV1(licensenumber, body) }
    }

    /**
     * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V2
     */
    suspend fun itemsCreateV2(licensenumber: String? = null, body: List<ItemsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreateV2(licensenumber, body) }
    }

    /**
     * Creates one or more new item brands for the specified Facility identified by the License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateBrand V2
     */
    suspend fun itemsCreateBrandV2(licensenumber: String? = null, body: List<ItemsCreateBrandV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreateBrandV2(licensenumber, body) }
    }

    /**
     * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateFile V2
     */
    suspend fun itemsCreateFileV2(licensenumber: String? = null, body: List<ItemsCreateFileV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreateFileV2(licensenumber, body) }
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V1
     */
    suspend fun itemsCreatePhotoV1(licensenumber: String? = null, body: List<ItemsCreatePhotoV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreatePhotoV1(licensenumber, body) }
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V2
     */
    suspend fun itemsCreatePhotoV2(licensenumber: String? = null, body: List<ItemsCreatePhotoV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreatePhotoV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * POST CreateUpdate V1
     */
    suspend fun itemsCreateUpdateV1(licensenumber: String? = null, body: List<ItemsCreateUpdateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsCreateUpdateV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V1
     */
    suspend fun itemsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.itemsDeleteV1(id, licensenumber) }
    }

    /**
     * Archives the specified Product by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V2
     */
    suspend fun itemsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.itemsDeleteV2(id, licensenumber) }
    }

    /**
     * Archives the specified Item Brand by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE DeleteBrand V2
     */
    suspend fun itemsDeleteBrandV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.itemsDeleteBrandV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET Get V1
     */
    suspend fun itemsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves detailed information about a specific Item by Id.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET Get V2
     */
    suspend fun itemsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V1
     */
    suspend fun itemsGetActiveV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetActiveV1(licensenumber) }
    }

    /**
     * Returns a list of active items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V2
     */
    suspend fun itemsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V1
     */
    suspend fun itemsGetBrandsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetBrandsV1(licensenumber) }
    }

    /**
     * Retrieves a list of active item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V2
     */
    suspend fun itemsGetBrandsV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetBrandsV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCategories V1
     */
    suspend fun itemsGetCategoriesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetCategoriesV1(licensenumber) }
    }

    /**
     * Retrieves a list of item categories.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCategories V2
     */
    suspend fun itemsGetCategoriesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetCategoriesV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a file by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetFile V2
     */
    suspend fun itemsGetFileV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetFileV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V1
     */
    suspend fun itemsGetInactiveV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetInactiveV1(licensenumber) }
    }

    /**
     * Retrieves a list of inactive items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V2
     */
    suspend fun itemsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetInactiveV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V1
     */
    suspend fun itemsGetPhotoV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetPhotoV1(id, licensenumber) }
    }

    /**
     * Retrieves an image by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V2
     */
    suspend fun itemsGetPhotoV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.itemsGetPhotoV2(id, licensenumber) }
    }

    /**
     * Updates one or more existing products for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT Update V2
     */
    suspend fun itemsUpdateV2(licensenumber: String? = null, body: List<ItemsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsUpdateV2(licensenumber, body) }
    }

    /**
     * Updates one or more existing item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT UpdateBrand V2
     */
    suspend fun itemsUpdateBrandV2(licensenumber: String? = null, body: List<ItemsUpdateBrandV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.itemsUpdateBrandV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST Create V1
     */
    suspend fun locationsCreateV1(licensenumber: String? = null, body: List<LocationsCreateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.locationsCreateV1(licensenumber, body) }
    }

    /**
     * Creates new locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     */
    suspend fun locationsCreateV2(licensenumber: String? = null, body: List<LocationsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.locationsCreateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST CreateUpdate V1
     */
    suspend fun locationsCreateUpdateV1(licensenumber: String? = null, body: List<LocationsCreateUpdateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.locationsCreateUpdateV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V1
     */
    suspend fun locationsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.locationsDeleteV1(id, licensenumber) }
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     */
    suspend fun locationsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.locationsDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET Get V1
     */
    suspend fun locationsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Location by its Id.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     */
    suspend fun locationsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V1
     */
    suspend fun locationsGetActiveV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetActiveV1(licensenumber) }
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     */
    suspend fun locationsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     */
    suspend fun locationsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetInactiveV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V1
     */
    suspend fun locationsGetTypesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetTypesV1(licensenumber) }
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V2
     */
    suspend fun locationsGetTypesV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.locationsGetTypesV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates existing locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     */
    suspend fun locationsUpdateV2(licensenumber: String? = null, body: List<LocationsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.locationsUpdateV2(licensenumber, body) }
    }

    /**
     * Adds new patients to a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * POST Create V2
     */
    suspend fun patientsCreateV2(licensenumber: String? = null, body: List<PatientsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientsCreateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateAdd V1
     */
    suspend fun patientsCreateAddV1(licensenumber: String? = null, body: List<PatientsCreateAddV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientsCreateAddV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateUpdate V1
     */
    suspend fun patientsCreateUpdateV1(licensenumber: String? = null, body: List<PatientsCreateUpdateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientsCreateUpdateV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V1
     */
    suspend fun patientsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.patientsDeleteV1(id, licensenumber) }
    }

    /**
     * Removes a Patient, identified by an Id, from a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V2
     */
    suspend fun patientsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.patientsDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET Get V1
     */
    suspend fun patientsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Patient by Id.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET Get V2
     */
    suspend fun patientsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V1
     */
    suspend fun patientsGetActiveV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsGetActiveV1(licensenumber) }
    }

    /**
     * Retrieves a list of active patients for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V2
     */
    suspend fun patientsGetActiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsGetActiveV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates Patient information for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * PUT Update V2
     */
    suspend fun patientsUpdateV2(licensenumber: String? = null, body: List<PatientsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.patientsUpdateV2(licensenumber, body) }
    }

    /**
     * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreateAssociate V2
     */
    suspend fun retailIdCreateAssociateV2(licensenumber: String? = null, body: List<RetailidCreateAssociateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.retailIdCreateAssociateV2(licensenumber, body) }
    }

    /**
     * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreateGenerate V2
     */
    suspend fun retailIdCreateGenerateV2(licensenumber: String? = null, body: RetailidCreateGenerateV2Request): Any? {
        return rateLimiter.execute(null, false) { client.retailIdCreateGenerateV2(licensenumber, body) }
    }

    /**
     * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Key Value Settings/Retail ID Merge Packages Enabled
     *
     * POST CreateMerge V2
     */
    suspend fun retailIdCreateMergeV2(licensenumber: String? = null, body: RetailidCreateMergeV2Request): Any? {
        return rateLimiter.execute(null, false) { client.retailIdCreateMergeV2(licensenumber, body) }
    }

    /**
     * Retrieves Package information for given list of Package labels.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
     *   - WebApi Retail ID Read Write State (All or WriteOnly)
     *   - Industry/View Packages
     *   - Admin/Employees/Packages Page/Product Labels(Manage)
     *
     * POST CreatePackageInfo V2
     */
    suspend fun retailIdCreatePackageInfoV2(licensenumber: String? = null, body: RetailidCreatePackageInfoV2Request): Any? {
        return rateLimiter.execute(null, false) { client.retailIdCreatePackageInfoV2(licensenumber, body) }
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
     *   - WebApi Retail ID Read Write State (All or ReadOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     *
     * GET GetReceiveByLabel V2
     */
    suspend fun retailIdGetReceiveByLabelV2(label: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.retailIdGetReceiveByLabelV2(label, licensenumber) }
    }

    /**
     * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
     * 
     *   Permissions Required:
     *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
     *   - WebApi Retail ID Read Write State (All or ReadOnly)
     *   - Industry/View Packages
     *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
     *
     * GET GetReceiveQrByShortCode V2
     */
    suspend fun retailIdGetReceiveQrByShortCodeV2(shortcode: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.retailIdGetReceiveQrByShortCodeV2(shortcode, licensenumber) }
    }

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * 
     *   Permissions Required:
     *   - None
     *
     * POST CreateIntegratorSetup V2
     */
    suspend fun sandboxCreateIntegratorSetupV2(userkey: String? = null, body: Any? = null): Any? {
        return rateLimiter.execute(null, false) { client.sandboxCreateIntegratorSetupV2(userkey, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST Create V1
     */
    suspend fun strainsCreateV1(licensenumber: String? = null, body: List<StrainsCreateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.strainsCreateV1(licensenumber, body) }
    }

    /**
     * Creates new strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * POST Create V2
     */
    suspend fun strainsCreateV2(licensenumber: String? = null, body: List<StrainsCreateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.strainsCreateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST CreateUpdate V1
     */
    suspend fun strainsCreateUpdateV1(licensenumber: String? = null, body: List<StrainsCreateUpdateV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.strainsCreateUpdateV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V1
     */
    suspend fun strainsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.strainsDeleteV1(id, licensenumber) }
    }

    /**
     * Archives an existing strain record for a Facility
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V2
     */
    suspend fun strainsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.strainsDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET Get V1
     */
    suspend fun strainsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.strainsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET Get V2
     */
    suspend fun strainsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.strainsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V1
     */
    suspend fun strainsGetActiveV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.strainsGetActiveV1(licensenumber) }
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V2
     */
    suspend fun strainsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.strainsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetInactive V2
     */
    suspend fun strainsGetInactiveV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.strainsGetInactiveV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Updates existing strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * PUT Update V2
     */
    suspend fun strainsUpdateV2(licensenumber: String? = null, body: List<StrainsUpdateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.strainsUpdateV2(licensenumber, body) }
    }

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPackageAvailable V2
     */
    suspend fun tagsGetPackageAvailableV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.tagsGetPackageAvailableV2(licensenumber) }
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPlantAvailable V2
     */
    suspend fun tagsGetPlantAvailableV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.tagsGetPlantAvailableV2(licensenumber) }
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *   - RetailId.AllowPackageStaging Key Value enabled
     *
     * GET GetStaged V2
     */
    suspend fun tagsGetStagedV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.tagsGetStagedV2(licensenumber) }
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V1
     */
    suspend fun patientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber, licensenumber) }
    }

    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V2
     */
    suspend fun patientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.patientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     */
    suspend fun plantBatchesCreateAdditivesV1(licensenumber: String? = null, body: List<PlantbatchesCreateAdditivesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateAdditivesV1(licensenumber, body) }
    }

    /**
     * Records Additive usage details for plant batches at a specific Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     */
    suspend fun plantBatchesCreateAdditivesV2(licensenumber: String? = null, body: List<PlantbatchesCreateAdditivesV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateAdditivesV2(licensenumber, body) }
    }

    /**
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     */
    suspend fun plantBatchesCreateAdditivesUsingtemplateV2(licensenumber: String? = null, body: List<PlantbatchesCreateAdditivesUsingtemplateV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateAdditivesUsingtemplateV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V1
     */
    suspend fun plantBatchesCreateAdjustV1(licensenumber: String? = null, body: List<PlantbatchesCreateAdjustV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateAdjustV1(licensenumber, body) }
    }

    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V2
     */
    suspend fun plantBatchesCreateAdjustV2(licensenumber: String? = null, body: List<PlantbatchesCreateAdjustV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateAdjustV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphase V1
     */
    suspend fun plantBatchesCreateChangegrowthphaseV1(licensenumber: String? = null, body: List<PlantbatchesCreateChangegrowthphaseV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateChangegrowthphaseV1(licensenumber, body) }
    }

    /**
     * Updates the growth phase of plants at a specified Facility based on tracking information.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateGrowthphase V2
     */
    suspend fun plantBatchesCreateGrowthphaseV2(licensenumber: String? = null, body: List<PlantbatchesCreateGrowthphaseV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateGrowthphaseV2(licensenumber, body) }
    }

    /**
     * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V2
     */
    suspend fun plantBatchesCreatePackageV2(isfrommotherplant: String? = null, licensenumber: String? = null, body: List<PlantbatchesCreatePackageV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreatePackageV2(isfrommotherplant, licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V1
     */
    suspend fun plantBatchesCreatePackageFrommotherplantV1(licensenumber: String? = null, body: List<PlantbatchesCreatePackageFrommotherplantV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreatePackageFrommotherplantV1(licensenumber, body) }
    }

    /**
     * Creates packages from mother plants at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V2
     */
    suspend fun plantBatchesCreatePackageFrommotherplantV2(licensenumber: String? = null, body: List<PlantbatchesCreatePackageFrommotherplantV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreatePackageFrommotherplantV2(licensenumber, body) }
    }

    /**
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreatePlantings V2
     */
    suspend fun plantBatchesCreatePlantingsV2(licensenumber: String? = null, body: List<PlantbatchesCreatePlantingsV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreatePlantingsV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V1
     */
    suspend fun plantBatchesCreateSplitV1(licensenumber: String? = null, body: List<PlantbatchesCreateSplitV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateSplitV1(licensenumber, body) }
    }

    /**
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V2
     */
    suspend fun plantBatchesCreateSplitV2(licensenumber: String? = null, body: List<PlantbatchesCreateSplitV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateSplitV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     */
    suspend fun plantBatchesCreateWasteV1(licensenumber: String? = null, body: List<PlantbatchesCreateWasteV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateWasteV1(licensenumber, body) }
    }

    /**
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     */
    suspend fun plantBatchesCreateWasteV2(licensenumber: String? = null, body: List<PlantbatchesCreateWasteV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateWasteV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Createpackages V1
     */
    suspend fun plantBatchesCreatepackagesV1(isfrommotherplant: String? = null, licensenumber: String? = null, body: List<PlantbatchesCreatepackagesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreatepackagesV1(isfrommotherplant, licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST Createplantings V1
     */
    suspend fun plantBatchesCreateplantingsV1(licensenumber: String? = null, body: List<PlantbatchesCreateplantingsV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesCreateplantingsV1(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V1
     */
    suspend fun plantBatchesDeleteV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesDeleteV1(licensenumber) }
    }

    /**
     * Completes the destruction of plant batches based on the provided input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V2
     */
    suspend fun plantBatchesDeleteV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesDeleteV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V1
     */
    suspend fun plantBatchesGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a Plant Batch by Id.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V2
     */
    suspend fun plantBatchesGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V1
     */
    suspend fun plantBatchesGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V2
     */
    suspend fun plantBatchesGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V1
     */
    suspend fun plantBatchesGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V2
     */
    suspend fun plantBatchesGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     */
    suspend fun plantBatchesGetTypesV1(no: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetTypesV1(no) }
    }

    /**
     * Retrieves a list of plant batch types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     */
    suspend fun plantBatchesGetTypesV2(pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetTypesV2(pagenumber, pagesize) }
    }

    /**
     * Retrieves waste details associated with plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     */
    suspend fun plantBatchesGetWasteV2(licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetWasteV2(licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     */
    suspend fun plantBatchesGetWasteReasonsV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetWasteReasonsV1(licensenumber) }
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     */
    suspend fun plantBatchesGetWasteReasonsV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.plantBatchesGetWasteReasonsV2(licensenumber) }
    }

    /**
     * Moves one or more plant batches to new locations with in a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *
     * PUT UpdateLocation V2
     */
    suspend fun plantBatchesUpdateLocationV2(licensenumber: String? = null, body: List<PlantbatchesUpdateLocationV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesUpdateLocationV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * PUT UpdateMoveplantbatches V1
     */
    suspend fun plantBatchesUpdateMoveplantbatchesV1(licensenumber: String? = null, body: List<PlantbatchesUpdateMoveplantbatchesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesUpdateMoveplantbatchesV1(licensenumber, body) }
    }

    /**
     * Renames plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateName V2
     */
    suspend fun plantBatchesUpdateNameV2(licensenumber: String? = null, body: List<PlantbatchesUpdateNameV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesUpdateNameV2(licensenumber, body) }
    }

    /**
     * Changes the strain of plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     */
    suspend fun plantBatchesUpdateStrainV2(licensenumber: String? = null, body: List<PlantbatchesUpdateStrainV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesUpdateStrainV2(licensenumber, body) }
    }

    /**
     * Replaces tags for plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     */
    suspend fun plantBatchesUpdateTagV2(licensenumber: String? = null, body: List<PlantbatchesUpdateTagV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.plantBatchesUpdateTagV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateAdjust V1
     */
    suspend fun processingJobsCreateAdjustV1(licensenumber: String? = null, body: List<ProcessingjobsCreateAdjustV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateAdjustV1(licensenumber, body) }
    }

    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateAdjust V2
     */
    suspend fun processingJobsCreateAdjustV2(licensenumber: String? = null, body: List<ProcessingjobsCreateAdjustV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateAdjustV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V1
     */
    suspend fun processingJobsCreateJobtypesV1(licensenumber: String? = null, body: List<ProcessingjobsCreateJobtypesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateJobtypesV1(licensenumber, body) }
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V2
     */
    suspend fun processingJobsCreateJobtypesV2(licensenumber: String? = null, body: List<ProcessingjobsCreateJobtypesV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateJobtypesV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateStart V1
     */
    suspend fun processingJobsCreateStartV1(licensenumber: String? = null, body: List<ProcessingjobsCreateStartV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateStartV1(licensenumber, body) }
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateStart V2
     */
    suspend fun processingJobsCreateStartV2(licensenumber: String? = null, body: List<ProcessingjobsCreateStartV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreateStartV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST Createpackages V1
     */
    suspend fun processingJobsCreatepackagesV1(licensenumber: String? = null, body: List<ProcessingjobsCreatepackagesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreatepackagesV1(licensenumber, body) }
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST Createpackages V2
     */
    suspend fun processingJobsCreatepackagesV2(licensenumber: String? = null, body: List<ProcessingjobsCreatepackagesV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsCreatepackagesV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V1
     */
    suspend fun processingJobsDeleteV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsDeleteV1(id, licensenumber) }
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V2
     */
    suspend fun processingJobsDeleteV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsDeleteV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V1
     */
    suspend fun processingJobsDeleteJobtypesV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsDeleteJobtypesV1(id, licensenumber) }
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V2
     */
    suspend fun processingJobsDeleteJobtypesV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsDeleteJobtypesV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V1
     */
    suspend fun processingJobsGetV1(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetV1(id, licensenumber) }
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V2
     */
    suspend fun processingJobsGetV2(id: String, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetV2(id, licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V1
     */
    suspend fun processingJobsGetActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves active processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V2
     */
    suspend fun processingJobsGetActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V1
     */
    suspend fun processingJobsGetInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves inactive processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V2
     */
    suspend fun processingJobsGetInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V1
     */
    suspend fun processingJobsGetJobtypesActiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V2
     */
    suspend fun processingJobsGetJobtypesActiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V1
     */
    suspend fun processingJobsGetJobtypesAttributesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesAttributesV1(licensenumber) }
    }

    /**
     * Retrieves all processing job attributes available for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V2
     */
    suspend fun processingJobsGetJobtypesAttributesV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesAttributesV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V1
     */
    suspend fun processingJobsGetJobtypesCategoriesV1(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesCategoriesV1(licensenumber) }
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V2
     */
    suspend fun processingJobsGetJobtypesCategoriesV2(licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesCategoriesV2(licensenumber) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V1
     */
    suspend fun processingJobsGetJobtypesInactiveV1(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber) }
    }

    /**
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V2
     */
    suspend fun processingJobsGetJobtypesInactiveV2(lastmodifiedend: String? = null, lastmodifiedstart: String? = null, licensenumber: String? = null, pagenumber: String? = null, pagesize: String? = null): Any? {
        return rateLimiter.execute(null, true) { client.processingJobsGetJobtypesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V1
     */
    suspend fun processingJobsUpdateFinishV1(licensenumber: String? = null, body: List<ProcessingjobsUpdateFinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateFinishV1(licensenumber, body) }
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V2
     */
    suspend fun processingJobsUpdateFinishV2(licensenumber: String? = null, body: List<ProcessingjobsUpdateFinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateFinishV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V1
     */
    suspend fun processingJobsUpdateJobtypesV1(licensenumber: String? = null, body: List<ProcessingjobsUpdateJobtypesV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateJobtypesV1(licensenumber, body) }
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V2
     */
    suspend fun processingJobsUpdateJobtypesV2(licensenumber: String? = null, body: List<ProcessingjobsUpdateJobtypesV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateJobtypesV2(licensenumber, body) }
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V1
     */
    suspend fun processingJobsUpdateUnfinishV1(licensenumber: String? = null, body: List<ProcessingjobsUpdateUnfinishV1RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateUnfinishV1(licensenumber, body) }
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V2
     */
    suspend fun processingJobsUpdateUnfinishV2(licensenumber: String? = null, body: List<ProcessingjobsUpdateUnfinishV2RequestItem>): Any? {
        return rateLimiter.execute(null, false) { client.processingJobsUpdateUnfinishV2(licensenumber, body) }
    }


    data class SublocationsCreateV2RequestItem(
        val Name: String? = null
    )

    data class SublocationsUpdateV2RequestItem(
        val Id: Int? = null,
        val Name: String? = null
    )

    data class TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages(
        val ExpirationDate: String? = null,
        val ExternalId: String? = null,
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val ItemName: String? = null,
        val PackagedDate: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val UnitOfMeasureName: String? = null,
        val UseByDate: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateExternalIncomingV1RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val GrossUnitOfWeightId: Int? = null,
        val GrossWeight: Double? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersCreateExternalIncomingV1RequestItem(
        val Destinations: List<TransfersCreateExternalIncomingV1RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val ShipperAddress1: String? = null,
        val ShipperAddress2: String? = null,
        val ShipperAddressCity: String? = null,
        val ShipperAddressPostalCode: String? = null,
        val ShipperAddressState: String? = null,
        val ShipperLicenseNumber: String? = null,
        val ShipperMainPhoneNumber: String? = null,
        val ShipperName: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages(
        val ExpirationDate: String? = null,
        val ExternalId: String? = null,
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val ItemName: String? = null,
        val PackagedDate: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val UnitOfMeasureName: String? = null,
        val UseByDate: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateExternalIncomingV2RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val GrossUnitOfWeightId: Int? = null,
        val GrossWeight: Double? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersCreateExternalIncomingV2RequestItem(
        val Destinations: List<TransfersCreateExternalIncomingV2RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val ShipperAddress1: String? = null,
        val ShipperAddress2: String? = null,
        val ShipperAddressCity: String? = null,
        val ShipperAddressPostalCode: String? = null,
        val ShipperAddressState: String? = null,
        val ShipperLicenseNumber: String? = null,
        val ShipperMainPhoneNumber: String? = null,
        val ShipperName: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateTemplatesV1RequestItem_Destinations_Packages(
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val PackageLabel: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersCreateTemplatesV1RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateTemplatesV1RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersCreateTemplatesV1RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersCreateTemplatesV1RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersCreateTemplatesV1RequestItem(
        val Destinations: List<TransfersCreateTemplatesV1RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val Name: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages(
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val PackageLabel: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersCreateTemplatesOutgoingV2RequestItem(
        val Destinations: List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val Name: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages(
        val ExpirationDate: String? = null,
        val ExternalId: String? = null,
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val ItemName: String? = null,
        val PackagedDate: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val UnitOfMeasureName: String? = null,
        val UseByDate: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateExternalIncomingV1RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val GrossUnitOfWeightId: Int? = null,
        val GrossWeight: Double? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferDestinationId: Int? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersUpdateExternalIncomingV1RequestItem(
        val Destinations: List<TransfersUpdateExternalIncomingV1RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val ShipperAddress1: String? = null,
        val ShipperAddress2: String? = null,
        val ShipperAddressCity: String? = null,
        val ShipperAddressPostalCode: String? = null,
        val ShipperAddressState: String? = null,
        val ShipperLicenseNumber: String? = null,
        val ShipperMainPhoneNumber: String? = null,
        val ShipperName: String? = null,
        val TransferId: Int? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages(
        val ExpirationDate: String? = null,
        val ExternalId: String? = null,
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val ItemName: String? = null,
        val PackagedDate: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val UnitOfMeasureName: String? = null,
        val UseByDate: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateExternalIncomingV2RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val GrossUnitOfWeightId: Int? = null,
        val GrossWeight: Double? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferDestinationId: Int? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersUpdateExternalIncomingV2RequestItem(
        val Destinations: List<TransfersUpdateExternalIncomingV2RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val ShipperAddress1: String? = null,
        val ShipperAddress2: String? = null,
        val ShipperAddressCity: String? = null,
        val ShipperAddressPostalCode: String? = null,
        val ShipperAddressState: String? = null,
        val ShipperLicenseNumber: String? = null,
        val ShipperMainPhoneNumber: String? = null,
        val ShipperName: String? = null,
        val TransferId: Int? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateTemplatesV1RequestItem_Destinations_Packages(
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val PackageLabel: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: String? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateTemplatesV1RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersUpdateTemplatesV1RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferDestinationId: Int? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersUpdateTemplatesV1RequestItem(
        val Destinations: List<TransfersUpdateTemplatesV1RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val Name: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransferTemplateId: Int? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages(
        val GrossUnitOfWeightName: String? = null,
        val GrossWeight: Double? = null,
        val PackageLabel: String? = null,
        val WholesalePrice: String? = null
    )

    data class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters(
        val DriverLayoverLeg: String? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val IsLayover: Boolean? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransporterDetails: List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations(
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val InvoiceNumber: String? = null,
        val Packages: List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages>? = null,
        val PlannedRoute: String? = null,
        val RecipientLicenseNumber: String? = null,
        val TransferDestinationId: Int? = null,
        val TransferTypeName: String? = null,
        val Transporters: List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters>? = null
    )

    data class TransfersUpdateTemplatesOutgoingV2RequestItem(
        val Destinations: List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations>? = null,
        val DriverLicenseNumber: String? = null,
        val DriverName: String? = null,
        val DriverOccupationalLicenseNumber: String? = null,
        val Name: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val TransferTemplateId: Int? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class PlantsCreateAdditivesV1RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantsCreateAdditivesV1RequestItem(
        val ActiveIngredients: List<PlantsCreateAdditivesV1RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val PlantLabels: List<String>? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantsCreateAdditivesV2RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantsCreateAdditivesV2RequestItem(
        val ActiveIngredients: List<PlantsCreateAdditivesV2RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val PlantLabels: List<String>? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantsCreateAdditivesBylocationV1RequestItem(
        val ActiveIngredients: List<PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val LocationName: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val SublocationName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantsCreateAdditivesBylocationV2RequestItem(
        val ActiveIngredients: List<PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val LocationName: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val SublocationName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem(
        val ActualDate: String? = null,
        val AdditivesTemplateName: String? = null,
        val LocationName: String? = null,
        val Rate: String? = null,
        val SublocationName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null,
        val Volume: String? = null
    )

    data class PlantsCreateAdditivesUsingtemplateV2RequestItem(
        val ActualDate: String? = null,
        val AdditivesTemplateName: String? = null,
        val PlantLabels: List<String>? = null,
        val Rate: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null,
        val Volume: String? = null
    )

    data class PlantsCreateChangegrowthphasesV1RequestItem(
        val GrowthDate: String? = null,
        val GrowthPhase: String? = null,
        val Id: Int? = null,
        val Label: String? = null,
        val NewLocation: String? = null,
        val NewSublocation: String? = null,
        val NewTag: String? = null
    )

    data class PlantsCreateHarvestplantsV1RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val PatientLicenseNumber: String? = null,
        val Plant: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class PlantsCreateManicureV2RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val PatientLicenseNumber: String? = null,
        val Plant: String? = null,
        val PlantCount: Int? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class PlantsCreateManicureplantsV1RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val PatientLicenseNumber: String? = null,
        val Plant: String? = null,
        val PlantCount: Int? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class PlantsCreateMoveplantsV1RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null,
        val Label: String? = null,
        val Location: String? = null,
        val Sublocation: String? = null
    )

    data class PlantsCreatePlantbatchPackageV1RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PackageTag: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchType: String? = null,
        val PlantLabel: String? = null,
        val Sublocation: String? = null
    )

    data class PlantsCreatePlantbatchPackageV2RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PackageTag: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchType: String? = null,
        val PlantLabel: String? = null,
        val Sublocation: String? = null
    )

    data class PlantsCreatePlantingsV1RequestItem(
        val ActualDate: String? = null,
        val LocationName: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchName: String? = null,
        val PlantBatchType: String? = null,
        val PlantCount: Int? = null,
        val PlantLabel: String? = null,
        val StrainName: String? = null,
        val SublocationName: String? = null
    )

    data class PlantsCreatePlantingsV2RequestItem(
        val ActualDate: String? = null,
        val LocationName: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchName: String? = null,
        val PlantBatchType: String? = null,
        val PlantCount: Int? = null,
        val PlantLabel: String? = null,
        val StrainName: String? = null,
        val SublocationName: String? = null
    )

    data class PlantsCreateWasteV1RequestItem(
        val LocationName: String? = null,
        val MixedMaterial: String? = null,
        val Note: String? = null,
        val PlantLabels: List<Any>? = null,
        val ReasonName: String? = null,
        val SublocationName: String? = null,
        val UnitOfMeasureName: String? = null,
        val WasteDate: String? = null,
        val WasteMethodName: String? = null,
        val WasteWeight: Double? = null
    )

    data class PlantsCreateWasteV2RequestItem(
        val LocationName: String? = null,
        val MixedMaterial: String? = null,
        val Note: String? = null,
        val PlantLabels: List<Any>? = null,
        val ReasonName: String? = null,
        val SublocationName: String? = null,
        val UnitOfMeasureName: String? = null,
        val WasteDate: String? = null,
        val WasteMethodName: String? = null,
        val WasteWeight: Double? = null
    )

    data class PlantsUpdateAdjustV2RequestItem(
        val AdjustCount: Int? = null,
        val AdjustReason: String? = null,
        val AdjustmentDate: String? = null,
        val Id: Int? = null,
        val Label: String? = null,
        val ReasonNote: String? = null
    )

    data class PlantsUpdateGrowthphaseV2RequestItem(
        val GrowthDate: String? = null,
        val GrowthPhase: String? = null,
        val Id: Int? = null,
        val Label: String? = null,
        val NewLocation: String? = null,
        val NewSublocation: String? = null,
        val NewTag: String? = null
    )

    data class PlantsUpdateHarvestV2RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val PatientLicenseNumber: String? = null,
        val Plant: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class PlantsUpdateLocationV2RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null,
        val Label: String? = null,
        val Location: String? = null,
        val Sublocation: String? = null
    )

    data class PlantsUpdateMergeV2RequestItem(
        val MergeDate: String? = null,
        val SourcePlantGroupLabel: String? = null,
        val TargetPlantGroupLabel: String? = null
    )

    data class PlantsUpdateSplitV2RequestItem(
        val PlantCount: Int? = null,
        val SourcePlantLabel: String? = null,
        val SplitDate: String? = null,
        val StrainLabel: String? = null,
        val TagLabel: String? = null
    )

    data class PlantsUpdateStrainV2RequestItem(
        val Id: Int? = null,
        val Label: String? = null,
        val StrainId: Int? = null,
        val StrainName: String? = null
    )

    data class PlantsUpdateTagV2RequestItem(
        val Id: Int? = null,
        val Label: String? = null,
        val NewTag: String? = null,
        val ReplaceDate: String? = null,
        val TagId: Int? = null
    )

    data class AdditivestemplatesCreateV2RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class AdditivestemplatesCreateV2RequestItem(
        val ActiveIngredients: List<AdditivestemplatesCreateV2RequestItem_ActiveIngredients>? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val Name: String? = null,
        val Note: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val RestrictiveEntryIntervalQuantityDescription: String? = null,
        val RestrictiveEntryIntervalTimeDescription: String? = null
    )

    data class AdditivestemplatesUpdateV2RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class AdditivestemplatesUpdateV2RequestItem(
        val ActiveIngredients: List<AdditivestemplatesUpdateV2RequestItem_ActiveIngredients>? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val Id: Int? = null,
        val Name: String? = null,
        val Note: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val RestrictiveEntryIntervalQuantityDescription: String? = null,
        val RestrictiveEntryIntervalTimeDescription: String? = null
    )

    data class HarvestsCreateFinishV1RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null
    )

    data class HarvestsCreatePackageV1RequestItem_Ingredients(
        val HarvestId: Int? = null,
        val HarvestName: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class HarvestsCreatePackageV1RequestItem(
        val ActualDate: String? = null,
        val DecontaminateProduct: Boolean? = null,
        val DecontaminationDate: String? = null,
        val DecontaminationSteps: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<HarvestsCreatePackageV1RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresDecontamination: Boolean? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val RemediateProduct: Boolean? = null,
        val RemediationDate: String? = null,
        val RemediationMethodId: Int? = null,
        val RemediationSteps: String? = null,
        val RequiredLabTestBatches: List<Any>? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfWeight: String? = null,
        val UseByDate: String? = null
    )

    data class HarvestsCreatePackageV2RequestItem_Ingredients(
        val HarvestId: Int? = null,
        val HarvestName: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class HarvestsCreatePackageV2RequestItem(
        val ActualDate: String? = null,
        val DecontaminateProduct: Boolean? = null,
        val DecontaminationDate: String? = null,
        val DecontaminationSteps: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<HarvestsCreatePackageV2RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresDecontamination: Boolean? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val RemediateProduct: Boolean? = null,
        val RemediationDate: String? = null,
        val RemediationMethodId: Int? = null,
        val RemediationSteps: String? = null,
        val RequiredLabTestBatches: List<Any>? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfWeight: String? = null,
        val UseByDate: String? = null
    )

    data class HarvestsCreatePackageTestingV1RequestItem_Ingredients(
        val HarvestId: Int? = null,
        val HarvestName: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class HarvestsCreatePackageTestingV1RequestItem(
        val ActualDate: String? = null,
        val DecontaminateProduct: Boolean? = null,
        val DecontaminationDate: String? = null,
        val DecontaminationSteps: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<HarvestsCreatePackageTestingV1RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresDecontamination: Boolean? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val RemediateProduct: Boolean? = null,
        val RemediationDate: String? = null,
        val RemediationMethodId: Int? = null,
        val RemediationSteps: String? = null,
        val RequiredLabTestBatches: List<Any>? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfWeight: String? = null,
        val UseByDate: String? = null
    )

    data class HarvestsCreatePackageTestingV2RequestItem_Ingredients(
        val HarvestId: Int? = null,
        val HarvestName: String? = null,
        val UnitOfWeight: String? = null,
        val Weight: Double? = null
    )

    data class HarvestsCreatePackageTestingV2RequestItem(
        val ActualDate: String? = null,
        val DecontaminateProduct: Boolean? = null,
        val DecontaminationDate: String? = null,
        val DecontaminationSteps: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<HarvestsCreatePackageTestingV2RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresDecontamination: Boolean? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val RemediateProduct: Boolean? = null,
        val RemediationDate: String? = null,
        val RemediationMethodId: Int? = null,
        val RemediationSteps: String? = null,
        val RequiredLabTestBatches: List<Any>? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfWeight: String? = null,
        val UseByDate: String? = null
    )

    data class HarvestsCreateRemoveWasteV1RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null,
        val UnitOfWeight: String? = null,
        val WasteType: String? = null,
        val WasteWeight: Double? = null
    )

    data class HarvestsCreateUnfinishV1RequestItem(
        val Id: Int? = null
    )

    data class HarvestsCreateWasteV2RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null,
        val UnitOfWeight: String? = null,
        val WasteType: String? = null,
        val WasteWeight: Double? = null
    )

    data class HarvestsUpdateFinishV2RequestItem(
        val ActualDate: String? = null,
        val Id: Int? = null
    )

    data class HarvestsUpdateLocationV2RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val Id: Int? = null
    )

    data class HarvestsUpdateMoveV1RequestItem(
        val ActualDate: String? = null,
        val DryingLocation: String? = null,
        val DryingSublocation: String? = null,
        val HarvestName: String? = null,
        val Id: Int? = null
    )

    data class HarvestsUpdateRenameV1RequestItem(
        val Id: Int? = null,
        val NewName: String? = null,
        val OldName: String? = null
    )

    data class HarvestsUpdateRenameV2RequestItem(
        val Id: Int? = null,
        val NewName: String? = null,
        val OldName: String? = null
    )

    data class HarvestsUpdateRestoreHarvestedPlantsV2RequestItem(
        val HarvestId: Int? = null,
        val PlantTags: List<String>? = null
    )

    data class HarvestsUpdateUnfinishV2RequestItem(
        val Id: Int? = null
    )

    data class LabtestsCreateRecordV1RequestItem_Results(
        val LabTestTypeName: String? = null,
        val Notes: String? = null,
        val Passed: Boolean? = null,
        val Quantity: Double? = null
    )

    data class LabtestsCreateRecordV1RequestItem(
        val DocumentFileBase64: String? = null,
        val DocumentFileName: String? = null,
        val Label: String? = null,
        val ResultDate: String? = null,
        val Results: List<LabtestsCreateRecordV1RequestItem_Results>? = null
    )

    data class LabtestsCreateRecordV2RequestItem_Results(
        val LabTestTypeName: String? = null,
        val Notes: String? = null,
        val Passed: Boolean? = null,
        val Quantity: Double? = null
    )

    data class LabtestsCreateRecordV2RequestItem(
        val DocumentFileBase64: String? = null,
        val DocumentFileName: String? = null,
        val Label: String? = null,
        val ResultDate: String? = null,
        val Results: List<LabtestsCreateRecordV2RequestItem_Results>? = null
    )

    data class LabtestsUpdateLabtestdocumentV1RequestItem(
        val DocumentFileBase64: String? = null,
        val DocumentFileName: String? = null,
        val LabTestResultId: Int? = null
    )

    data class LabtestsUpdateLabtestdocumentV2RequestItem(
        val DocumentFileBase64: String? = null,
        val DocumentFileName: String? = null,
        val LabTestResultId: Int? = null
    )

    data class LabtestsUpdateResultReleaseV1RequestItem(
        val PackageLabel: String? = null
    )

    data class LabtestsUpdateResultReleaseV2RequestItem(
        val PackageLabel: String? = null
    )

    data class SalesCreateDeliveryV1RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryV1RequestItem(
        val ConsumerId: Int? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateDeliveryV1RequestItem_Transactions>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesCreateDeliveryV2RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryV2RequestItem(
        val ConsumerId: Int? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateDeliveryV2RequestItem_Transactions>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerV1RequestItem_Destinations(
        val ConsumerId: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val Transactions: List<SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions>? = null
    )

    data class SalesCreateDeliveryRetailerV1RequestItem_Packages(
        val DateTime: String? = null,
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerV1RequestItem(
        val DateTime: String? = null,
        val Destinations: List<SalesCreateDeliveryRetailerV1RequestItem_Destinations>? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerV1RequestItem_Packages>? = null,
        val PhoneNumberForQuestions: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerV2RequestItem_Destinations(
        val ConsumerId: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val Transactions: List<SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions>? = null
    )

    data class SalesCreateDeliveryRetailerV2RequestItem_Packages(
        val DateTime: String? = null,
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerV2RequestItem(
        val DateTime: String? = null,
        val Destinations: List<SalesCreateDeliveryRetailerV2RequestItem_Destinations>? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerV2RequestItem_Packages>? = null,
        val PhoneNumberForQuestions: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesCreateDeliveryRetailerDepartV1RequestItem(
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerDepartV2RequestItem(
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerEndV1RequestItem_Packages(
        val EndQuantity: Int? = null,
        val EndUnitOfMeasure: String? = null,
        val Label: String? = null
    )

    data class SalesCreateDeliveryRetailerEndV1RequestItem(
        val ActualArrivalDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerEndV1RequestItem_Packages>? = null,
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerEndV2RequestItem_Packages(
        val EndQuantity: Int? = null,
        val EndUnitOfMeasure: String? = null,
        val Label: String? = null
    )

    data class SalesCreateDeliveryRetailerEndV2RequestItem(
        val ActualArrivalDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerEndV2RequestItem_Packages>? = null,
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerRestockV1RequestItem_Packages(
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val RemoveCurrentPackage: Boolean? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerRestockV1RequestItem(
        val DateTime: String? = null,
        val Destinations: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerRestockV1RequestItem_Packages>? = null,
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerRestockV2RequestItem_Packages(
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val RemoveCurrentPackage: Boolean? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerRestockV2RequestItem(
        val DateTime: String? = null,
        val Destinations: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Packages: List<SalesCreateDeliveryRetailerRestockV2RequestItem_Packages>? = null,
        val RetailerDeliveryId: Int? = null
    )

    data class SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerSaleV1RequestItem(
        val ConsumerId: Int? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: Int? = null,
        val RetailerDeliveryId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions>? = null
    )

    data class SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateDeliveryRetailerSaleV2RequestItem(
        val ConsumerId: Int? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: Int? = null,
        val RetailerDeliveryId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions>? = null
    )

    data class SalesCreateReceiptV1RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateReceiptV1RequestItem(
        val CaregiverLicenseNumber: String? = null,
        val ExternalReceiptNumber: String? = null,
        val IdentificationMethod: String? = null,
        val PatientLicenseNumber: String? = null,
        val PatientRegistrationLocationId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateReceiptV1RequestItem_Transactions>? = null
    )

    data class SalesCreateReceiptV2RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesCreateReceiptV2RequestItem(
        val CaregiverLicenseNumber: String? = null,
        val ExternalReceiptNumber: String? = null,
        val IdentificationMethod: String? = null,
        val PatientLicenseNumber: String? = null,
        val PatientRegistrationLocationId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesCreateReceiptV2RequestItem_Transactions>? = null
    )

    data class SalesCreateTransactionByDateV1RequestItem(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryV1RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryV1RequestItem(
        val ConsumerId: Int? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateDeliveryV1RequestItem_Transactions>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryV2RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryV2RequestItem(
        val ConsumerId: Int? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateDeliveryV2RequestItem_Transactions>? = null,
        val TransporterFacilityLicenseNumber: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages(
        val Label: String? = null,
        val ReturnQuantityVerified: Int? = null,
        val ReturnReason: String? = null,
        val ReturnReasonNote: String? = null,
        val ReturnUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryCompleteV1RequestItem(
        val AcceptedPackages: List<String>? = null,
        val ActualArrivalDateTime: String? = null,
        val Id: Int? = null,
        val PaymentType: String? = null,
        val ReturnedPackages: List<SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages>? = null
    )

    data class SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages(
        val Label: String? = null,
        val ReturnQuantityVerified: Int? = null,
        val ReturnReason: String? = null,
        val ReturnReasonNote: String? = null,
        val ReturnUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryCompleteV2RequestItem(
        val AcceptedPackages: List<String>? = null,
        val ActualArrivalDateTime: String? = null,
        val Id: Int? = null,
        val PaymentType: String? = null,
        val ReturnedPackages: List<SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages>? = null
    )

    data class SalesUpdateDeliveryHubV1RequestItem(
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val TransporterFacilityId: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryHubV2RequestItem(
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val TransporterFacilityId: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryHubAcceptV1RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateDeliveryHubAcceptV2RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateDeliveryHubDepartV1RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateDeliveryHubDepartV2RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateDeliveryHubVerifyIdV1RequestItem(
        val Id: Int? = null,
        val PaymentType: String? = null
    )

    data class SalesUpdateDeliveryHubVerifyIdV2RequestItem(
        val Id: Int? = null,
        val PaymentType: String? = null
    )

    data class SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryRetailerV1RequestItem_Destinations(
        val ConsumerId: String? = null,
        val DriverEmployeeId: Int? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions>? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryRetailerV1RequestItem_Packages(
        val DateTime: String? = null,
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryRetailerV1RequestItem(
        val DateTime: String? = null,
        val Destinations: List<SalesUpdateDeliveryRetailerV1RequestItem_Destinations>? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val Packages: List<SalesUpdateDeliveryRetailerV1RequestItem_Packages>? = null,
        val PhoneNumberForQuestions: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryRetailerV2RequestItem_Destinations(
        val ConsumerId: String? = null,
        val DriverEmployeeId: Int? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedArrivalDateTime: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val PhoneNumberForQuestions: String? = null,
        val PlannedRoute: String? = null,
        val RecipientAddressCity: String? = null,
        val RecipientAddressCounty: String? = null,
        val RecipientAddressPostalCode: String? = null,
        val RecipientAddressState: String? = null,
        val RecipientAddressStreet1: String? = null,
        val RecipientAddressStreet2: String? = null,
        val RecipientName: String? = null,
        val RecipientZoneId: String? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions>? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateDeliveryRetailerV2RequestItem_Packages(
        val DateTime: String? = null,
        val PackageLabel: String? = null,
        val Quantity: Int? = null,
        val TotalPrice: Double? = null,
        val UnitOfMeasure: String? = null
    )

    data class SalesUpdateDeliveryRetailerV2RequestItem(
        val DateTime: String? = null,
        val Destinations: List<SalesUpdateDeliveryRetailerV2RequestItem_Destinations>? = null,
        val DriverEmployeeId: String? = null,
        val DriverName: String? = null,
        val DriversLicenseNumber: String? = null,
        val EstimatedDepartureDateTime: String? = null,
        val Id: Int? = null,
        val Packages: List<SalesUpdateDeliveryRetailerV2RequestItem_Packages>? = null,
        val PhoneNumberForQuestions: String? = null,
        val VehicleLicensePlateNumber: String? = null,
        val VehicleMake: String? = null,
        val VehicleModel: String? = null
    )

    data class SalesUpdateReceiptV1RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateReceiptV1RequestItem(
        val CaregiverLicenseNumber: String? = null,
        val ExternalReceiptNumber: String? = null,
        val Id: Int? = null,
        val IdentificationMethod: String? = null,
        val PatientLicenseNumber: String? = null,
        val PatientRegistrationLocationId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateReceiptV1RequestItem_Transactions>? = null
    )

    data class SalesUpdateReceiptV2RequestItem_Transactions(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class SalesUpdateReceiptV2RequestItem(
        val CaregiverLicenseNumber: String? = null,
        val ExternalReceiptNumber: String? = null,
        val Id: Int? = null,
        val IdentificationMethod: String? = null,
        val PatientLicenseNumber: String? = null,
        val PatientRegistrationLocationId: Int? = null,
        val SalesCustomerType: String? = null,
        val SalesDateTime: String? = null,
        val Transactions: List<SalesUpdateReceiptV2RequestItem_Transactions>? = null
    )

    data class SalesUpdateReceiptFinalizeV2RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateReceiptUnfinalizeV2RequestItem(
        val Id: Int? = null
    )

    data class SalesUpdateTransactionByDateV1RequestItem(
        val CityTax: String? = null,
        val CountyTax: String? = null,
        val DiscountAmount: String? = null,
        val ExciseTax: String? = null,
        val InvoiceNumber: String? = null,
        val MunicipalTax: String? = null,
        val PackageLabel: String? = null,
        val Price: String? = null,
        val QrCodes: String? = null,
        val Quantity: Int? = null,
        val SalesTax: String? = null,
        val SubTotal: String? = null,
        val TotalAmount: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class TransportersCreateDriverV2RequestItem(
        val DriversLicenseNumber: String? = null,
        val EmployeeId: String? = null,
        val Name: String? = null
    )

    data class TransportersCreateVehicleV2RequestItem(
        val LicensePlateNumber: String? = null,
        val Make: String? = null,
        val Model: String? = null
    )

    data class TransportersUpdateDriverV2RequestItem(
        val DriversLicenseNumber: String? = null,
        val EmployeeId: String? = null,
        val Id: String? = null,
        val Name: String? = null
    )

    data class TransportersUpdateVehicleV2RequestItem(
        val Id: String? = null,
        val LicensePlateNumber: String? = null,
        val Make: String? = null,
        val Model: String? = null
    )

    data class PackagesCreateV1RequestItem_Ingredients(
        val Package: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateV1RequestItem(
        val ActualDate: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<PackagesCreateV1RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val RequiredLabTestBatches: Boolean? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val UseSameItem: Boolean? = null
    )

    data class PackagesCreateV2RequestItem_Ingredients(
        val Package: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateV2RequestItem(
        val ActualDate: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<PackagesCreateV2RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val RequiredLabTestBatches: Boolean? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val UseSameItem: Boolean? = null
    )

    data class PackagesCreateAdjustV1RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentReason: String? = null,
        val Label: String? = null,
        val Quantity: Int? = null,
        val ReasonNote: String? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateAdjustV2RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentReason: String? = null,
        val Label: String? = null,
        val Quantity: Int? = null,
        val ReasonNote: String? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateChangeItemV1RequestItem(
        val Item: String? = null,
        val Label: String? = null
    )

    data class PackagesCreateChangeLocationV1RequestItem(
        val Label: String? = null,
        val Location: String? = null,
        val MoveDate: String? = null,
        val Sublocation: String? = null
    )

    data class PackagesCreateFinishV1RequestItem(
        val ActualDate: String? = null,
        val Label: String? = null
    )

    data class PackagesCreatePlantingsV1RequestItem(
        val LocationName: String? = null,
        val PackageAdjustmentAmount: Int? = null,
        val PackageAdjustmentUnitOfMeasureName: String? = null,
        val PackageLabel: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchName: String? = null,
        val PlantBatchType: String? = null,
        val PlantCount: Int? = null,
        val PlantedDate: String? = null,
        val StrainName: String? = null,
        val SublocationName: String? = null,
        val UnpackagedDate: String? = null
    )

    data class PackagesCreatePlantingsV2RequestItem(
        val LocationName: String? = null,
        val PackageAdjustmentAmount: Int? = null,
        val PackageAdjustmentUnitOfMeasureName: String? = null,
        val PackageLabel: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatchName: String? = null,
        val PlantBatchType: String? = null,
        val PlantCount: Int? = null,
        val PlantedDate: String? = null,
        val StrainName: String? = null,
        val SublocationName: String? = null,
        val UnpackagedDate: String? = null
    )

    data class PackagesCreateRemediateV1RequestItem(
        val PackageLabel: String? = null,
        val RemediationDate: String? = null,
        val RemediationMethodName: String? = null,
        val RemediationSteps: String? = null
    )

    data class PackagesCreateTestingV1RequestItem_Ingredients(
        val Package: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateTestingV1RequestItem(
        val ActualDate: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<PackagesCreateTestingV1RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val RequiredLabTestBatches: Boolean? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val UseSameItem: Boolean? = null
    )

    data class PackagesCreateTestingV2RequestItem_Ingredients(
        val Package: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesCreateTestingV2RequestItem(
        val ActualDate: String? = null,
        val ExpirationDate: String? = null,
        val Ingredients: List<PackagesCreateTestingV2RequestItem_Ingredients>? = null,
        val IsDonation: Boolean? = null,
        val IsProductionBatch: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val LabTestStageId: Int? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProcessingJobTypeId: Int? = null,
        val ProductRequiresRemediation: Boolean? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val RequiredLabTestBatches: List<String>? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val UseSameItem: Boolean? = null
    )

    data class PackagesCreateUnfinishV1RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateAdjustV2RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentReason: String? = null,
        val Label: String? = null,
        val Quantity: Int? = null,
        val ReasonNote: String? = null,
        val UnitOfMeasure: String? = null
    )

    data class PackagesUpdateChangeNoteV1RequestItem(
        val Note: String? = null,
        val PackageLabel: String? = null
    )

    data class PackagesUpdateDecontaminateV2RequestItem(
        val DecontaminationDate: String? = null,
        val DecontaminationMethodName: String? = null,
        val DecontaminationSteps: String? = null,
        val PackageLabel: String? = null
    )

    data class PackagesUpdateDonationFlagV2RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateDonationUnflagV2RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateExternalidV2RequestItem(
        val ExternalId: String? = null,
        val PackageLabel: String? = null
    )

    data class PackagesUpdateFinishV2RequestItem(
        val ActualDate: String? = null,
        val Label: String? = null
    )

    data class PackagesUpdateItemV2RequestItem(
        val Item: String? = null,
        val Label: String? = null
    )

    data class PackagesUpdateLabTestRequiredV2RequestItem(
        val Label: String? = null,
        val RequiredLabTestBatches: List<String>? = null
    )

    data class PackagesUpdateLocationV2RequestItem(
        val Label: String? = null,
        val Location: String? = null,
        val MoveDate: String? = null,
        val Sublocation: String? = null
    )

    data class PackagesUpdateNoteV2RequestItem(
        val Note: String? = null,
        val PackageLabel: String? = null
    )

    data class PackagesUpdateRemediateV2RequestItem(
        val PackageLabel: String? = null,
        val RemediationDate: String? = null,
        val RemediationMethodName: String? = null,
        val RemediationSteps: String? = null
    )

    data class PackagesUpdateTradesampleFlagV2RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateTradesampleUnflagV2RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateUnfinishV2RequestItem(
        val Label: String? = null
    )

    data class PackagesUpdateUsebydateV2RequestItem(
        val ExpirationDate: String? = null,
        val Label: String? = null,
        val SellByDate: String? = null,
        val UseByDate: String? = null
    )

    data class PatientcheckinsCreateV1RequestItem(
        val CheckInDate: String? = null,
        val CheckInLocationId: Int? = null,
        val PatientLicenseNumber: String? = null,
        val RegistrationExpiryDate: String? = null,
        val RegistrationStartDate: String? = null
    )

    data class PatientcheckinsCreateV2RequestItem(
        val CheckInDate: String? = null,
        val CheckInLocationId: Int? = null,
        val PatientLicenseNumber: String? = null,
        val RegistrationExpiryDate: String? = null,
        val RegistrationStartDate: String? = null
    )

    data class PatientcheckinsUpdateV1RequestItem(
        val CheckInDate: String? = null,
        val CheckInLocationId: Int? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val RegistrationExpiryDate: String? = null,
        val RegistrationStartDate: String? = null
    )

    data class PatientcheckinsUpdateV2RequestItem(
        val CheckInDate: String? = null,
        val CheckInLocationId: Int? = null,
        val Id: Int? = null,
        val PatientLicenseNumber: String? = null,
        val RegistrationExpiryDate: String? = null,
        val RegistrationStartDate: String? = null
    )

    data class ItemsCreateV1RequestItem(
        val AdministrationMethod: String? = null,
        val Allergens: String? = null,
        val Description: String? = null,
        val GlobalProductName: String? = null,
        val ItemBrand: String? = null,
        val ItemCategory: String? = null,
        val ItemIngredients: String? = null,
        val LabelImageFileSystemIds: String? = null,
        val LabelPhotoDescription: String? = null,
        val Name: String? = null,
        val NumberOfDoses: String? = null,
        val PackagingImageFileSystemIds: String? = null,
        val PackagingPhotoDescription: String? = null,
        val ProcessingJobCategoryName: String? = null,
        val ProcessingJobTypeName: String? = null,
        val ProductImageFileSystemIds: String? = null,
        val ProductPDFFileSystemIds: String? = null,
        val ProductPhotoDescription: String? = null,
        val PublicIngredients: String? = null,
        val ServingSize: String? = null,
        val Strain: String? = null,
        val SupplyDurationDays: Int? = null,
        val UnitCbdAContent: Double? = null,
        val UnitCbdAContentDose: Double? = null,
        val UnitCbdAContentDoseUnitOfMeasure: String? = null,
        val UnitCbdAContentUnitOfMeasure: String? = null,
        val UnitCbdAPercent: Double? = null,
        val UnitCbdContent: Double? = null,
        val UnitCbdContentDose: Double? = null,
        val UnitCbdContentDoseUnitOfMeasure: String? = null,
        val UnitCbdContentUnitOfMeasure: String? = null,
        val UnitCbdPercent: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcAContent: Double? = null,
        val UnitThcAContentDose: Double? = null,
        val UnitThcAContentDoseUnitOfMeasure: String? = null,
        val UnitThcAContentUnitOfMeasure: String? = null,
        val UnitThcAPercent: Double? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentDose: Double? = null,
        val UnitThcContentDoseUnitOfMeasure: String? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitVolume: Double? = null,
        val UnitVolumeUnitOfMeasure: String? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class ItemsCreateV2RequestItem(
        val AdministrationMethod: String? = null,
        val Allergens: String? = null,
        val Description: String? = null,
        val GlobalProductName: String? = null,
        val ItemBrand: String? = null,
        val ItemCategory: String? = null,
        val ItemIngredients: String? = null,
        val LabelImageFileSystemIds: String? = null,
        val LabelPhotoDescription: String? = null,
        val Name: String? = null,
        val NumberOfDoses: String? = null,
        val PackagingImageFileSystemIds: String? = null,
        val PackagingPhotoDescription: String? = null,
        val ProcessingJobCategoryName: String? = null,
        val ProcessingJobTypeName: String? = null,
        val ProductImageFileSystemIds: String? = null,
        val ProductPDFFileSystemIds: String? = null,
        val ProductPhotoDescription: String? = null,
        val PublicIngredients: String? = null,
        val ServingSize: String? = null,
        val Strain: String? = null,
        val SupplyDurationDays: Int? = null,
        val UnitCbdAContent: Double? = null,
        val UnitCbdAContentDose: Double? = null,
        val UnitCbdAContentDoseUnitOfMeasure: String? = null,
        val UnitCbdAContentUnitOfMeasure: String? = null,
        val UnitCbdAPercent: Double? = null,
        val UnitCbdContent: Double? = null,
        val UnitCbdContentDose: Double? = null,
        val UnitCbdContentDoseUnitOfMeasure: String? = null,
        val UnitCbdContentUnitOfMeasure: String? = null,
        val UnitCbdPercent: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcAContent: Double? = null,
        val UnitThcAContentDose: Double? = null,
        val UnitThcAContentDoseUnitOfMeasure: String? = null,
        val UnitThcAContentUnitOfMeasure: String? = null,
        val UnitThcAPercent: Double? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentDose: Double? = null,
        val UnitThcContentDoseUnitOfMeasure: String? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitVolume: Double? = null,
        val UnitVolumeUnitOfMeasure: String? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class ItemsCreateBrandV2RequestItem(
        val Name: String? = null
    )

    data class ItemsCreateFileV2RequestItem(
        val EncodedImageBase64: String? = null,
        val FileName: String? = null
    )

    data class ItemsCreatePhotoV1RequestItem(
        val EncodedImageBase64: String? = null,
        val FileName: String? = null
    )

    data class ItemsCreatePhotoV2RequestItem(
        val EncodedImageBase64: String? = null,
        val FileName: String? = null
    )

    data class ItemsCreateUpdateV1RequestItem(
        val AdministrationMethod: String? = null,
        val Allergens: String? = null,
        val Description: String? = null,
        val GlobalProductName: String? = null,
        val Id: Int? = null,
        val ItemBrand: String? = null,
        val ItemCategory: String? = null,
        val ItemIngredients: String? = null,
        val LabelImageFileSystemIds: String? = null,
        val LabelPhotoDescription: String? = null,
        val Name: String? = null,
        val NumberOfDoses: String? = null,
        val PackagingImageFileSystemIds: String? = null,
        val PackagingPhotoDescription: String? = null,
        val ProcessingJobCategoryName: String? = null,
        val ProcessingJobTypeName: String? = null,
        val ProductImageFileSystemIds: String? = null,
        val ProductPDFFileSystemIds: String? = null,
        val ProductPhotoDescription: String? = null,
        val PublicIngredients: String? = null,
        val ServingSize: String? = null,
        val Strain: String? = null,
        val SupplyDurationDays: Int? = null,
        val UnitCbdAContent: Double? = null,
        val UnitCbdAContentDose: Double? = null,
        val UnitCbdAContentDoseUnitOfMeasure: String? = null,
        val UnitCbdAContentUnitOfMeasure: String? = null,
        val UnitCbdAPercent: Double? = null,
        val UnitCbdContent: Double? = null,
        val UnitCbdContentDose: Double? = null,
        val UnitCbdContentDoseUnitOfMeasure: String? = null,
        val UnitCbdContentUnitOfMeasure: String? = null,
        val UnitCbdPercent: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcAContent: Double? = null,
        val UnitThcAContentDose: Double? = null,
        val UnitThcAContentDoseUnitOfMeasure: String? = null,
        val UnitThcAContentUnitOfMeasure: String? = null,
        val UnitThcAPercent: Double? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentDose: Double? = null,
        val UnitThcContentDoseUnitOfMeasure: String? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitVolume: Double? = null,
        val UnitVolumeUnitOfMeasure: String? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class ItemsUpdateV2RequestItem(
        val AdministrationMethod: String? = null,
        val Allergens: String? = null,
        val Description: String? = null,
        val GlobalProductName: String? = null,
        val Id: Int? = null,
        val ItemBrand: String? = null,
        val ItemCategory: String? = null,
        val ItemIngredients: String? = null,
        val LabelImageFileSystemIds: String? = null,
        val LabelPhotoDescription: String? = null,
        val Name: String? = null,
        val NumberOfDoses: String? = null,
        val PackagingImageFileSystemIds: String? = null,
        val PackagingPhotoDescription: String? = null,
        val ProcessingJobCategoryName: String? = null,
        val ProcessingJobTypeName: String? = null,
        val ProductImageFileSystemIds: String? = null,
        val ProductPDFFileSystemIds: String? = null,
        val ProductPhotoDescription: String? = null,
        val PublicIngredients: String? = null,
        val ServingSize: String? = null,
        val Strain: String? = null,
        val SupplyDurationDays: Int? = null,
        val UnitCbdAContent: Double? = null,
        val UnitCbdAContentDose: Double? = null,
        val UnitCbdAContentDoseUnitOfMeasure: String? = null,
        val UnitCbdAContentUnitOfMeasure: String? = null,
        val UnitCbdAPercent: Double? = null,
        val UnitCbdContent: Double? = null,
        val UnitCbdContentDose: Double? = null,
        val UnitCbdContentDoseUnitOfMeasure: String? = null,
        val UnitCbdContentUnitOfMeasure: String? = null,
        val UnitCbdPercent: Double? = null,
        val UnitOfMeasure: String? = null,
        val UnitThcAContent: Double? = null,
        val UnitThcAContentDose: Double? = null,
        val UnitThcAContentDoseUnitOfMeasure: String? = null,
        val UnitThcAContentUnitOfMeasure: String? = null,
        val UnitThcAPercent: Double? = null,
        val UnitThcContent: Double? = null,
        val UnitThcContentDose: Double? = null,
        val UnitThcContentDoseUnitOfMeasure: String? = null,
        val UnitThcContentUnitOfMeasure: String? = null,
        val UnitThcPercent: Double? = null,
        val UnitVolume: Double? = null,
        val UnitVolumeUnitOfMeasure: String? = null,
        val UnitWeight: Double? = null,
        val UnitWeightUnitOfMeasure: String? = null
    )

    data class ItemsUpdateBrandV2RequestItem(
        val Id: Int? = null,
        val Name: String? = null
    )

    data class LocationsCreateV1RequestItem(
        val LocationTypeName: String? = null,
        val Name: String? = null
    )

    data class LocationsCreateV2RequestItem(
        val LocationTypeName: String? = null,
        val Name: String? = null
    )

    data class LocationsCreateUpdateV1RequestItem(
        val Id: Int? = null,
        val LocationTypeName: String? = null,
        val Name: String? = null
    )

    data class LocationsUpdateV2RequestItem(
        val Id: Int? = null,
        val LocationTypeName: String? = null,
        val Name: String? = null
    )

    data class PatientsCreateV2RequestItem(
        val ActualDate: String? = null,
        val ConcentrateOuncesAllowed: Int? = null,
        val FlowerOuncesAllowed: Int? = null,
        val HasSalesLimitExemption: Boolean? = null,
        val InfusedOuncesAllowed: Int? = null,
        val LicenseEffectiveEndDate: String? = null,
        val LicenseEffectiveStartDate: String? = null,
        val LicenseNumber: String? = null,
        val MaxConcentrateThcPercentAllowed: Int? = null,
        val MaxFlowerThcPercentAllowed: Int? = null,
        val RecommendedPlants: Int? = null,
        val RecommendedSmokableQuantity: Int? = null,
        val ThcOuncesAllowed: Int? = null
    )

    data class PatientsCreateAddV1RequestItem(
        val ActualDate: String? = null,
        val ConcentrateOuncesAllowed: Int? = null,
        val FlowerOuncesAllowed: Int? = null,
        val HasSalesLimitExemption: Boolean? = null,
        val InfusedOuncesAllowed: Int? = null,
        val LicenseEffectiveEndDate: String? = null,
        val LicenseEffectiveStartDate: String? = null,
        val LicenseNumber: String? = null,
        val MaxConcentrateThcPercentAllowed: Int? = null,
        val MaxFlowerThcPercentAllowed: Int? = null,
        val RecommendedPlants: Int? = null,
        val RecommendedSmokableQuantity: Int? = null,
        val ThcOuncesAllowed: Int? = null
    )

    data class PatientsCreateUpdateV1RequestItem(
        val ActualDate: String? = null,
        val ConcentrateOuncesAllowed: Int? = null,
        val FlowerOuncesAllowed: Int? = null,
        val HasSalesLimitExemption: Boolean? = null,
        val InfusedOuncesAllowed: Int? = null,
        val LicenseEffectiveEndDate: String? = null,
        val LicenseEffectiveStartDate: String? = null,
        val LicenseNumber: String? = null,
        val MaxConcentrateThcPercentAllowed: Int? = null,
        val MaxFlowerThcPercentAllowed: Int? = null,
        val NewLicenseNumber: String? = null,
        val RecommendedPlants: Int? = null,
        val RecommendedSmokableQuantity: Int? = null,
        val ThcOuncesAllowed: Int? = null
    )

    data class PatientsUpdateV2RequestItem(
        val ActualDate: String? = null,
        val ConcentrateOuncesAllowed: Int? = null,
        val FlowerOuncesAllowed: Int? = null,
        val HasSalesLimitExemption: Boolean? = null,
        val InfusedOuncesAllowed: Int? = null,
        val LicenseEffectiveEndDate: String? = null,
        val LicenseEffectiveStartDate: String? = null,
        val LicenseNumber: String? = null,
        val MaxConcentrateThcPercentAllowed: Int? = null,
        val MaxFlowerThcPercentAllowed: Int? = null,
        val NewLicenseNumber: String? = null,
        val RecommendedPlants: Int? = null,
        val RecommendedSmokableQuantity: Int? = null,
        val ThcOuncesAllowed: Int? = null
    )

    data class RetailidCreateAssociateV2RequestItem(
        val PackageLabel: String? = null,
        val QrUrls: List<String>? = null
    )

    data class RetailidCreateGenerateV2Request(
        val PackageLabel: String? = null,
        val Quantity: Int? = null
    )

    data class RetailidCreateMergeV2Request(
        val packageLabels: List<String>? = null
    )

    data class RetailidCreatePackageInfoV2Request(
        val packageLabels: List<String>? = null
    )

    data class StrainsCreateV1RequestItem(
        val CbdLevel: Double? = null,
        val IndicaPercentage: Double? = null,
        val Name: String? = null,
        val SativaPercentage: Double? = null,
        val TestingStatus: String? = null,
        val ThcLevel: Double? = null
    )

    data class StrainsCreateV2RequestItem(
        val CbdLevel: Double? = null,
        val IndicaPercentage: Double? = null,
        val Name: String? = null,
        val SativaPercentage: Double? = null,
        val TestingStatus: String? = null,
        val ThcLevel: Double? = null
    )

    data class StrainsCreateUpdateV1RequestItem(
        val CbdLevel: Double? = null,
        val Id: Int? = null,
        val IndicaPercentage: Double? = null,
        val Name: String? = null,
        val SativaPercentage: Double? = null,
        val TestingStatus: String? = null,
        val ThcLevel: Double? = null
    )

    data class StrainsUpdateV2RequestItem(
        val CbdLevel: Double? = null,
        val Id: Int? = null,
        val IndicaPercentage: Double? = null,
        val Name: String? = null,
        val SativaPercentage: Double? = null,
        val TestingStatus: String? = null,
        val ThcLevel: Double? = null
    )

    data class PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantbatchesCreateAdditivesV1RequestItem(
        val ActiveIngredients: List<PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val PlantBatchName: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients(
        val Name: String? = null,
        val Percentage: Double? = null
    )

    data class PlantbatchesCreateAdditivesV2RequestItem(
        val ActiveIngredients: List<PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients>? = null,
        val ActualDate: String? = null,
        val AdditiveType: String? = null,
        val ApplicationDevice: String? = null,
        val EpaRegistrationNumber: String? = null,
        val PlantBatchName: String? = null,
        val ProductSupplier: String? = null,
        val ProductTradeName: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null
    )

    data class PlantbatchesCreateAdditivesUsingtemplateV2RequestItem(
        val ActualDate: String? = null,
        val AdditivesTemplateName: String? = null,
        val PlantBatchName: String? = null,
        val Rate: String? = null,
        val TotalAmountApplied: Int? = null,
        val TotalAmountUnitOfMeasure: String? = null,
        val Volume: String? = null
    )

    data class PlantbatchesCreateAdjustV1RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentReason: String? = null,
        val PlantBatchName: String? = null,
        val Quantity: Int? = null,
        val ReasonNote: String? = null
    )

    data class PlantbatchesCreateAdjustV2RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentReason: String? = null,
        val PlantBatchName: String? = null,
        val Quantity: Int? = null,
        val ReasonNote: String? = null
    )

    data class PlantbatchesCreateChangegrowthphaseV1RequestItem(
        val Count: Int? = null,
        val CountPerPlant: String? = null,
        val GrowthDate: String? = null,
        val GrowthPhase: String? = null,
        val Name: String? = null,
        val NewLocation: String? = null,
        val NewSublocation: String? = null,
        val PatientLicenseNumber: String? = null,
        val StartingTag: String? = null
    )

    data class PlantbatchesCreateGrowthphaseV2RequestItem(
        val Count: Int? = null,
        val CountPerPlant: String? = null,
        val GrowthDate: String? = null,
        val GrowthPhase: String? = null,
        val Name: String? = null,
        val NewLocation: String? = null,
        val NewSublocation: String? = null,
        val PatientLicenseNumber: String? = null,
        val StartingTag: String? = null
    )

    data class PlantbatchesCreatePackageV2RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val ExpirationDate: String? = null,
        val Id: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UseByDate: String? = null
    )

    data class PlantbatchesCreatePackageFrommotherplantV1RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val ExpirationDate: String? = null,
        val Id: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UseByDate: String? = null
    )

    data class PlantbatchesCreatePackageFrommotherplantV2RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val ExpirationDate: String? = null,
        val Id: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UseByDate: String? = null
    )

    data class PlantbatchesCreatePlantingsV2RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val Location: String? = null,
        val Name: String? = null,
        val PatientLicenseNumber: String? = null,
        val SourcePlantBatches: String? = null,
        val Strain: String? = null,
        val Sublocation: String? = null,
        val Type: String? = null
    )

    data class PlantbatchesCreateSplitV1RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val GroupName: String? = null,
        val Location: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val Strain: String? = null,
        val Sublocation: String? = null
    )

    data class PlantbatchesCreateSplitV2RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val GroupName: String? = null,
        val Location: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val Strain: String? = null,
        val Sublocation: String? = null
    )

    data class PlantbatchesCreateWasteV1RequestItem(
        val MixedMaterial: String? = null,
        val Note: String? = null,
        val PlantBatchName: String? = null,
        val ReasonName: String? = null,
        val UnitOfMeasureName: String? = null,
        val WasteDate: String? = null,
        val WasteMethodName: String? = null,
        val WasteWeight: Double? = null
    )

    data class PlantbatchesCreateWasteV2RequestItem(
        val MixedMaterial: String? = null,
        val Note: String? = null,
        val PlantBatchName: String? = null,
        val ReasonName: String? = null,
        val UnitOfMeasureName: String? = null,
        val WasteDate: String? = null,
        val WasteMethodName: String? = null,
        val WasteWeight: Double? = null
    )

    data class PlantbatchesCreatepackagesV1RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val ExpirationDate: String? = null,
        val Id: Int? = null,
        val IsDonation: Boolean? = null,
        val IsTradeSample: Boolean? = null,
        val Item: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PatientLicenseNumber: String? = null,
        val PlantBatch: String? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UseByDate: String? = null
    )

    data class PlantbatchesCreateplantingsV1RequestItem(
        val ActualDate: String? = null,
        val Count: Int? = null,
        val Location: String? = null,
        val Name: String? = null,
        val PatientLicenseNumber: String? = null,
        val SourcePlantBatches: String? = null,
        val Strain: String? = null,
        val Sublocation: String? = null,
        val Type: String? = null
    )

    data class PlantbatchesUpdateLocationV2RequestItem(
        val Location: String? = null,
        val MoveDate: String? = null,
        val Name: String? = null,
        val Sublocation: String? = null
    )

    data class PlantbatchesUpdateMoveplantbatchesV1RequestItem(
        val Location: String? = null,
        val MoveDate: String? = null,
        val Name: String? = null,
        val Sublocation: String? = null
    )

    data class PlantbatchesUpdateNameV2RequestItem(
        val Group: String? = null,
        val Id: Int? = null,
        val NewGroup: String? = null
    )

    data class PlantbatchesUpdateStrainV2RequestItem(
        val Id: Int? = null,
        val Name: String? = null,
        val StrainId: Int? = null,
        val StrainName: String? = null
    )

    data class PlantbatchesUpdateTagV2RequestItem(
        val Group: String? = null,
        val Id: Int? = null,
        val NewTag: String? = null,
        val ReplaceDate: String? = null,
        val TagId: Int? = null
    )

    data class ProcessingjobsCreateAdjustV1RequestItem_Packages(
        val Label: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreateAdjustV1RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentNote: String? = null,
        val AdjustmentReason: String? = null,
        val CountUnitOfMeasureName: String? = null,
        val Id: Int? = null,
        val Packages: List<ProcessingjobsCreateAdjustV1RequestItem_Packages>? = null,
        val VolumeUnitOfMeasureName: String? = null,
        val WeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsCreateAdjustV2RequestItem_Packages(
        val Label: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreateAdjustV2RequestItem(
        val AdjustmentDate: String? = null,
        val AdjustmentNote: String? = null,
        val AdjustmentReason: String? = null,
        val CountUnitOfMeasureName: String? = null,
        val Id: Int? = null,
        val Packages: List<ProcessingjobsCreateAdjustV2RequestItem_Packages>? = null,
        val VolumeUnitOfMeasureName: String? = null,
        val WeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsCreateJobtypesV1RequestItem(
        val Attributes: List<String>? = null,
        val Category: String? = null,
        val Description: String? = null,
        val Name: String? = null,
        val ProcessingSteps: String? = null
    )

    data class ProcessingjobsCreateJobtypesV2RequestItem(
        val Attributes: List<String>? = null,
        val Category: String? = null,
        val Description: String? = null,
        val Name: String? = null,
        val ProcessingSteps: String? = null
    )

    data class ProcessingjobsCreateStartV1RequestItem_Packages(
        val Label: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreateStartV1RequestItem(
        val CountUnitOfMeasure: String? = null,
        val JobName: String? = null,
        val JobType: String? = null,
        val Packages: List<ProcessingjobsCreateStartV1RequestItem_Packages>? = null,
        val StartDate: String? = null,
        val VolumeUnitOfMeasure: String? = null,
        val WeightUnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreateStartV2RequestItem_Packages(
        val Label: String? = null,
        val Quantity: Int? = null,
        val UnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreateStartV2RequestItem(
        val CountUnitOfMeasure: String? = null,
        val JobName: String? = null,
        val JobType: String? = null,
        val Packages: List<ProcessingjobsCreateStartV2RequestItem_Packages>? = null,
        val StartDate: String? = null,
        val VolumeUnitOfMeasure: String? = null,
        val WeightUnitOfMeasure: String? = null
    )

    data class ProcessingjobsCreatepackagesV1RequestItem(
        val ExpirationDate: String? = null,
        val FinishDate: String? = null,
        val FinishNote: String? = null,
        val FinishProcessingJob: Boolean? = null,
        val Item: String? = null,
        val JobName: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PackageDate: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val WasteCountQuantity: String? = null,
        val WasteCountUnitOfMeasureName: String? = null,
        val WasteVolumeQuantity: String? = null,
        val WasteVolumeUnitOfMeasureName: String? = null,
        val WasteWeightQuantity: String? = null,
        val WasteWeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsCreatepackagesV2RequestItem(
        val ExpirationDate: String? = null,
        val FinishDate: String? = null,
        val FinishNote: String? = null,
        val FinishProcessingJob: Boolean? = null,
        val Item: String? = null,
        val JobName: String? = null,
        val Location: String? = null,
        val Note: String? = null,
        val PackageDate: String? = null,
        val PatientLicenseNumber: String? = null,
        val ProductionBatchNumber: String? = null,
        val Quantity: Int? = null,
        val SellByDate: String? = null,
        val Sublocation: String? = null,
        val Tag: String? = null,
        val UnitOfMeasure: String? = null,
        val UseByDate: String? = null,
        val WasteCountQuantity: String? = null,
        val WasteCountUnitOfMeasureName: String? = null,
        val WasteVolumeQuantity: String? = null,
        val WasteVolumeUnitOfMeasureName: String? = null,
        val WasteWeightQuantity: String? = null,
        val WasteWeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsUpdateFinishV1RequestItem(
        val FinishDate: String? = null,
        val FinishNote: String? = null,
        val Id: Int? = null,
        val TotalCountWaste: String? = null,
        val TotalVolumeWaste: String? = null,
        val TotalWeightWaste: Int? = null,
        val WasteCountUnitOfMeasureName: String? = null,
        val WasteVolumeUnitOfMeasureName: String? = null,
        val WasteWeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsUpdateFinishV2RequestItem(
        val FinishDate: String? = null,
        val FinishNote: String? = null,
        val Id: Int? = null,
        val TotalCountWaste: String? = null,
        val TotalVolumeWaste: String? = null,
        val TotalWeightWaste: Int? = null,
        val WasteCountUnitOfMeasureName: String? = null,
        val WasteVolumeUnitOfMeasureName: String? = null,
        val WasteWeightUnitOfMeasureName: String? = null
    )

    data class ProcessingjobsUpdateJobtypesV1RequestItem(
        val Attributes: List<String>? = null,
        val CategoryName: String? = null,
        val Description: String? = null,
        val Id: Int? = null,
        val Name: String? = null,
        val ProcessingSteps: String? = null
    )

    data class ProcessingjobsUpdateJobtypesV2RequestItem(
        val Attributes: List<String>? = null,
        val CategoryName: String? = null,
        val Description: String? = null,
        val Id: Int? = null,
        val Name: String? = null,
        val ProcessingSteps: String? = null
    )

    data class ProcessingjobsUpdateUnfinishV1RequestItem(
        val Id: Int? = null
    )

    data class ProcessingjobsUpdateUnfinishV2RequestItem(
        val Id: Int? = null
    )
}
