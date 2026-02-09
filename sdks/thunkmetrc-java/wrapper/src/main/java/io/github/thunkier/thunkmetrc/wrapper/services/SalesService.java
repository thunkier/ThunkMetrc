package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class SalesService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public SalesService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

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
    public WriteResponse createSalesDeliveries(
        String licenseNumber, List<SalesCreateDeliveriesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesDeliveries(
                licenseNumber, body
            )
        );
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
    public WriteResponse createSalesDeliveriesRetailerDepart(
        String licenseNumber, List<SalesCreateDeliveriesRetailerDepartRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesDeliveriesRetailerDepart(
                licenseNumber, body
            )
        );
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
    public WriteResponse createSalesDeliveriesRetailerEnd(
        String licenseNumber, List<SalesCreateDeliveriesRetailerEndRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesDeliveriesRetailerEnd(
                licenseNumber, body
            )
        );
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
    public WriteResponse createSalesDeliveriesRetailerRestock(
        String licenseNumber, List<SalesCreateDeliveriesRetailerRestockRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesDeliveriesRetailerRestock(
                licenseNumber, body
            )
        );
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
    public WriteResponse createSalesReceipts(
        String licenseNumber, List<SalesCreateReceiptsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesReceipts(
                licenseNumber, body
            )
        );
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
    public WriteResponse createSalesDeliveriesRetailer(
        String licenseNumber, List<SalesCreateSalesDeliveriesRetailerRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.createSalesDeliveriesRetailer(
                licenseNumber, body
            )
        );
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
    public Object deleteSalesDeliveryById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.sales.deleteSalesDeliveryById(
                id, licenseNumber
            )
        );
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
    public Object deleteSalesDeliveryRetailerById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.sales.deleteSalesDeliveryRetailerById(
                id, licenseNumber
            )
        );
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
    public Object deleteSalesReceiptById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.sales.deleteSalesReceiptById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ActiveDelivery> getSalesActiveDeliveries(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ActiveDelivery>) client.sales.getSalesActiveDeliveries(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ActiveDeliveriesRetailer> getSalesActiveDeliveriesRetailer(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ActiveDeliveriesRetailer>) client.sales.getSalesActiveDeliveriesRetailer(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<ActiveReceipt> getSalesActiveReceipts(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<ActiveReceipt>) client.sales.getSalesActiveReceipts(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * GET /sales/v2/counties
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<County> getSalesCounties(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<County>) client.sales.getSalesCounties(
                
            )
        );
    }

    /**
     * Returns a list of customer types.
     * GET /sales/v2/customertypes
     * @throws Exception If request fails
     * @return Response object
     */
    public Object getSalesCustomerTypes(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.sales.getSalesCustomerTypes(
                
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<DeliveriesReturnReason> getSalesDeliveriesReturnReasons(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<DeliveriesReturnReason>) client.sales.getSalesDeliveriesReturnReasons(
                licenseNumber, pageNumber, pageSize
            )
        );
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
    public DeliveryRetailer getSalesDeliveryRetailerById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (DeliveryRetailer) client.sales.getSalesDeliveryRetailerById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<InactiveDelivery> getSalesInactiveDeliveries(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<InactiveDelivery>) client.sales.getSalesInactiveDeliveries(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<InactiveDeliveriesRetailer> getSalesInactiveDeliveriesRetailer(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<InactiveDeliveriesRetailer>) client.sales.getSalesInactiveDeliveriesRetailer(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
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
    @SuppressWarnings("unchecked")
    public PaginatedResponse<InactiveReceipt> getSalesInactiveReceipts(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<InactiveReceipt>) client.sales.getSalesInactiveReceipts(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * GET /sales/v2/patientregistration/locations
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public List<PatientRegistrationLocation> getSalesPatientRegistrationLocations(
        
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<PatientRegistrationLocation>) client.sales.getSalesPatientRegistrationLocations(
                
            )
        );
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
    public Object getSalesPaymentTypes(
        String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (Object) client.sales.getSalesPaymentTypes(
                licenseNumber
            )
        );
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
    public ActiveReceipt getSalesReceiptById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (ActiveReceipt) client.sales.getSalesReceiptById(
                id, licenseNumber
            )
        );
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
    @SuppressWarnings("unchecked")
    public List<ReceiptsExternalByExternalNumber> getSalesReceiptsExternalByExternalNumber(
        String externalNumber, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (List<ReceiptsExternalByExternalNumber>) client.sales.getSalesReceiptsExternalByExternalNumber(
                externalNumber, licenseNumber
            )
        );
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
    public SalesDelivery getSalesDeliveryById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (SalesDelivery) client.sales.getSalesDeliveryById(
                id, licenseNumber
            )
        );
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
    public WriteResponse updateSalesDeliveries(
        String licenseNumber, List<SalesUpdateDeliveriesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveries(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesComplete(
        String licenseNumber, List<SalesUpdateDeliveriesCompleteRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesComplete(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesHub(
        String licenseNumber, List<SalesUpdateDeliveriesHubRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesHub(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesHubAccept(
        String licenseNumber, List<SalesUpdateDeliveriesHubAcceptRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesHubAccept(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesHubDepart(
        String licenseNumber, List<SalesUpdateDeliveriesHubDepartRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesHubDepart(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesHubVerifyID(
        String licenseNumber, List<SalesUpdateDeliveriesHubVerifyIDRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesHubVerifyID(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesDeliveriesRetailer(
        String licenseNumber, List<SalesUpdateDeliveriesRetailerRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesDeliveriesRetailer(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesReceipts(
        String licenseNumber, List<SalesUpdateReceiptsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesReceipts(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesReceiptsFinalize(
        String licenseNumber, List<SalesUpdateReceiptsFinalizeRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesReceiptsFinalize(
                licenseNumber, body
            )
        );
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
    public WriteResponse updateSalesReceiptsUnfinalize(
        String licenseNumber, List<SalesUpdateReceiptsUnfinalizeRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.sales.updateSalesReceiptsUnfinalize(
                licenseNumber, body
            )
        );
    }

}

