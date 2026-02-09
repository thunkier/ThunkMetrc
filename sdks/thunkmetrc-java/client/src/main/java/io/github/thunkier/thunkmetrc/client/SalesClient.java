package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class SalesClient {
    private final MetrcClient client;

    public SalesClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
     * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
     * POST CreateDeliveries
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesDeliveries(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST CreateDeliveriesRetailerDepart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesDeliveriesRetailerDepart(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/depart");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST CreateDeliveriesRetailerEnd
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesDeliveriesRetailerEnd(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/end");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * POST CreateDeliveriesRetailerRestock
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesDeliveriesRetailerRestock(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/restock");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
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
     * POST CreateReceipts
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesReceipts(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
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
     * POST CreateSalesDeliveriesRetailer
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createSalesDeliveriesRetailer(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * Permissions Required:
     * - Manage Sales Delivery
     * DELETE DeleteDeliveryById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteSalesDeliveryById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
     * Permissions Required:
     * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
     * - Industry/Facility Type/Retailer Delivery
     * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
     * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
     * - Manage Retailer Delivery
     * DELETE DeleteDeliveryRetailerById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteSalesDeliveryRetailerById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * Permissions Required:
     * - Manage Sales
     * DELETE DeleteReceiptById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteSalesReceiptById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET GetActiveDeliveries
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesActiveDeliveries(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/active");
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
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET GetActiveDeliveriesRetailer
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesActiveDeliveriesRetailer(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/active");
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
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET GetActiveReceipts
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesActiveReceipts(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/active");
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
     * Returns a list of counties available for sales deliveries.
     * GET GetCounties
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesCounties(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/counties");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of customer types.
     * GET GetCustomerTypes
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesCustomerTypes(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/customertypes");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * Permissions Required:
     * - Sales Delivery
     * GET GetDeliveriesReturnReasons
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesDeliveriesReturnReasons(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/returnreasons");
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
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET GetDeliveryRetailerById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesDeliveryRetailerById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET GetInactiveDeliveries
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesInactiveDeliveries(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/inactive");
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
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * Permissions Required:
     * - View Retailer Delivery
     * - Manage Retailer Delivery
     * GET GetInactiveDeliveriesRetailer
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesInactiveDeliveriesRetailer(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer/inactive");
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
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET GetInactiveReceipts
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesInactiveReceipts(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/inactive");
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
     * Returns a list of valid Patient registration locations for sales.
     * GET GetPatientRegistrationLocations
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesPatientRegistrationLocations(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/patientregistration/locations");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET GetPaymentTypes
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesPaymentTypes(
        String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/paymenttypes");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET GetReceiptById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesReceiptById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * Permissions Required:
     * - View Sales
     * - Manage Sales
     * GET GetReceiptsExternalByExternalNumber
     * @param externalNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesReceiptsExternalByExternalNumber(
        String externalNumber, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/external/"+URLEncoder.encode(externalNumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * Permissions Required:
     * - View Sales Delivery
     * - Manage Sales Delivery
     * GET GetSalesDeliveryById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getSalesDeliveryById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT UpdateDeliveries
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveries(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT UpdateDeliveriesComplete
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesComplete(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/complete");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery, Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHub
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesHub(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubAccept
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesHubAccept(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/accept");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubDepart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesHubDepart(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/depart");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubVerifyID
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesHubVerifyID(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/hub/verifyID");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
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
     * PUT UpdateDeliveriesRetailer
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesDeliveriesRetailer(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/deliveries/retailer");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceipts
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesReceipts(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceiptsFinalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesReceiptsFinalize(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/finalize");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceiptsUnfinalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateSalesReceiptsUnfinalize(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/sales/v2/receipts/unfinalize");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

