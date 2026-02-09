package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class SalesService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     * POST /sales/v2/deliveries
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createDeliveries(licenseNumber: String? = null, body: List<SalesCreateDeliveriesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesDeliveries(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST /sales/v2/deliveries/retailer/depart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createDeliveriesRetailerDepart(licenseNumber: String? = null, body: List<SalesCreateDeliveriesRetailerDepartRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesDeliveriesRetailerDepart(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST /sales/v2/deliveries/retailer/end
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createDeliveriesRetailerEnd(licenseNumber: String? = null, body: List<SalesCreateDeliveriesRetailerEndRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesDeliveriesRetailerEnd(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST /sales/v2/deliveries/retailer/restock
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createDeliveriesRetailerRestock(licenseNumber: String? = null, body: List<SalesCreateDeliveriesRetailerRestockRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesDeliveriesRetailerRestock(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales (Write)
     * - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
     * - Industry/Facility Type/Advanced Sales
     * - WebApi Sales Read Write State (All or WriteOnly)
     * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     * POST /sales/v2/receipts
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createReceipts(licenseNumber: String? = null, body: List<SalesCreateReceiptsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesReceipts(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     * - Manage Retailer Delivery
     * POST /sales/v2/deliveries/retailer
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createSalesDeliveriesRetailer(licenseNumber: String? = null, body: List<SalesCreateSalesDeliveriesRetailerRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.createSalesDeliveriesRetailer(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * Permissions Required:
     * - Manage Sales Delivery
     * DELETE /sales/v2/deliveries/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteDeliveryById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.deleteSalesDeliveryById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * DELETE /sales/v2/deliveries/retailer/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteDeliveryRetailerById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.deleteSalesDeliveryRetailerById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * Permissions Required:
     * - Manage Sales
     * DELETE /sales/v2/receipts/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteReceiptById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.deleteSalesReceiptById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET /sales/v2/deliveries/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ActiveDelivery>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesActiveDeliveries(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ActiveDelivery>
    }

    /**
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET /sales/v2/deliveries/retailer/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ActiveDeliveriesRetailer>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesActiveDeliveriesRetailer(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ActiveDeliveriesRetailer>
    }

    /**
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET /sales/v2/receipts/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getActiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<ActiveReceipt>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesActiveReceipts(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<ActiveReceipt>
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * GET /sales/v2/counties
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getCounties(): PaginatedResponse<County>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesCounties() 
        } as? PaginatedResponse<County>
    }

    /**
     * Returns a list of customer types.
     * GET /sales/v2/customertypes
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getCustomerTypes(): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesCustomerTypes() 
        } as? Any
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * Permissions Required:
     * - Sales Delivery
     * GET /sales/v2/deliveries/returnreasons
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveriesReturnReasons(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<DeliveriesReturnReason>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesDeliveriesReturnReasons(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<DeliveriesReturnReason>
    }

    /**
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET /sales/v2/deliveries/retailer/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDeliveryRetailerById(id: String, licenseNumber: String? = null): DeliveryRetailer? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesDeliveryRetailerById(id, licenseNumber, ) 
        } as? DeliveryRetailer
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET /sales/v2/deliveries/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<InactiveDelivery>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesInactiveDeliveries(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<InactiveDelivery>
    }

    /**
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET /sales/v2/deliveries/retailer/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<InactiveDeliveriesRetailer>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesInactiveDeliveriesRetailer(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<InactiveDeliveriesRetailer>
    }

    /**
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET /sales/v2/receipts/inactive
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getInactiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<InactiveReceipt>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesInactiveReceipts(lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<InactiveReceipt>
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * GET /sales/v2/patientregistration/locations
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPatientRegistrationLocations(): List<PatientRegistrationLocation>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesPatientRegistrationLocations() 
        } as? List<PatientRegistrationLocation>
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET /sales/v2/paymenttypes
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getPaymentTypes(licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesPaymentTypes(licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET /sales/v2/receipts/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getReceiptById(id: String, licenseNumber: String? = null): ActiveReceipt? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesReceiptById(id, licenseNumber, ) 
        } as? ActiveReceipt
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET /sales/v2/receipts/external/{externalNumber}
     * @param externalNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getReceiptsExternalByExternalNumber(externalNumber: String, licenseNumber: String? = null): List<ReceiptsExternalByExternalNumber>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesReceiptsExternalByExternalNumber(externalNumber, licenseNumber, ) 
        } as? List<ReceiptsExternalByExternalNumber>
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET /sales/v2/deliveries/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getSalesDeliveryById(id: String, licenseNumber: String? = null): SalesDelivery? {
        return rateLimiter.execute(null,true,
        ) { 
            client.sales.getSalesDeliveryById(id, licenseNumber, ) 
        } as? SalesDelivery
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT /sales/v2/deliveries
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveries(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveries(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT /sales/v2/deliveries/complete
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesComplete(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesCompleteRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesComplete(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery, Manage Sales Delivery Hub
     * PUT /sales/v2/deliveries/hub
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesHub(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesHubRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesHub(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT /sales/v2/deliveries/hub/accept
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesHubAccept(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesHubAcceptRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesHubAccept(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT /sales/v2/deliveries/hub/depart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesHubDepart(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesHubDepartRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesHubDepart(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT /sales/v2/deliveries/hub/verifyID
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesHubVerifyID(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesHubVerifyIDRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesHubVerifyID(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     * - Manage Retailer Delivery
     * PUT /sales/v2/deliveries/retailer
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDeliveriesRetailer(licenseNumber: String? = null, body: List<SalesUpdateDeliveriesRetailerRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesDeliveriesRetailer(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales
     * PUT /sales/v2/receipts
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateReceipts(licenseNumber: String? = null, body: List<SalesUpdateReceiptsRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesReceipts(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT /sales/v2/receipts/finalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateReceiptsFinalize(licenseNumber: String? = null, body: List<SalesUpdateReceiptsFinalizeRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesReceiptsFinalize(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT /sales/v2/receipts/unfinalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateReceiptsUnfinalize(licenseNumber: String? = null, body: List<SalesUpdateReceiptsUnfinalizeRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.sales.updateSalesReceiptsUnfinalize(licenseNumber, body) 
        } as? WriteResponse
    }
}

