package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class SalesClient(private val client: MetrcClient) {
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
    fun createSalesDeliveries(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries")
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
    fun createSalesDeliveriesRetailerDepart(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/depart")
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
    fun createSalesDeliveriesRetailerEnd(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/end")
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
    fun createSalesDeliveriesRetailerRestock(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/restock")
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
    fun createSalesReceipts(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts")
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
    fun createSalesDeliveriesRetailer(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer")
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
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * Permissions Required:
     * - Manage Sales Delivery
     * DELETE DeleteDeliveryById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteSalesDeliveryById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
    fun deleteSalesDeliveryRetailerById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * Permissions Required:
     * - Manage Sales
     * DELETE DeleteReceiptById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteSalesReceiptById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
    fun getSalesActiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/active")
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
    fun getSalesActiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/active")
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
    fun getSalesActiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/receipts/active")
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
     * Returns a list of counties available for sales deliveries.
     * GET GetCounties
     * @throws IOException If request fails
     * @return Response object
     */
    fun getSalesCounties(): Any? {
        val path = StringBuilder("/sales/v2/counties")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of customer types.
     * GET GetCustomerTypes
     * @throws IOException If request fails
     * @return Response object
     */
    fun getSalesCustomerTypes(): Any? {
        val path = StringBuilder("/sales/v2/customertypes")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getSalesDeliveriesReturnReasons(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/returnreasons")
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
    fun getSalesDeliveryRetailerById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
    fun getSalesInactiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/inactive")
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
    fun getSalesInactiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer/inactive")
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
    fun getSalesInactiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/receipts/inactive")
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
     * Returns a list of valid Patient registration locations for sales.
     * GET GetPatientRegistrationLocations
     * @throws IOException If request fails
     * @return Response object
     */
    fun getSalesPatientRegistrationLocations(): Any? {
        val path = StringBuilder("/sales/v2/patientregistration/locations")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
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
    fun getSalesPaymentTypes(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/paymenttypes")
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
    fun getSalesReceiptById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/receipts/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
    fun getSalesReceiptsExternalByExternalNumber(externalNumber: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/receipts/external/${URLEncoder.encode(externalNumber, StandardCharsets.UTF_8)}")
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
    fun getSalesDeliveryById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/sales/v2/deliveries/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT UpdateDeliveries
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveries(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries")
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
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery
     * PUT UpdateDeliveriesComplete
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveriesComplete(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/complete")
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
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales Delivery, Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHub
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveriesHub(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub")
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
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubAccept
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveriesHubAccept(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/accept")
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
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubDepart
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveriesHubDepart(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/depart")
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
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * Permissions Required:
     * - Manage Sales Delivery Hub
     * PUT UpdateDeliveriesHubVerifyID
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesDeliveriesHubVerifyID(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/hub/verifyID")
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
    fun updateSalesDeliveriesRetailer(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/deliveries/retailer")
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
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceipts
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesReceipts(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts")
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
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceiptsFinalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesReceiptsFinalize(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/finalize")
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
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * Permissions Required:
     * - Manage Sales
     * PUT UpdateReceiptsUnfinalize
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateSalesReceiptsUnfinalize(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/sales/v2/receipts/unfinalize")
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

