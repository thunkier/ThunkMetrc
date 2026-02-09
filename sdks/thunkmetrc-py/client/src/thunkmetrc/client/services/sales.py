from typing import Any, Optional, Dict, List, Union
import urllib.parse

class SalesClient:
    def __init__(self, client):
        self.client = client
    def create_sales_deliveries(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        POST CreateDeliveries
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries"
        return self.client._send("POST", path, body, query_params)
    def create_sales_deliveries_retailer_depart(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST CreateDeliveriesRetailerDepart
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer/depart"
        return self.client._send("POST", path, body, query_params)
    def create_sales_deliveries_retailer_end(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST CreateDeliveriesRetailerEnd
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer/end"
        return self.client._send("POST", path, body, query_params)
    def create_sales_deliveries_retailer_restock(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST CreateDeliveriesRetailerRestock
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer/restock"
        return self.client._send("POST", path, body, query_params)
    def create_sales_receipts(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales (Write)
        - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
        - Industry/Facility Type/Advanced Sales
        - WebApi Sales Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        POST CreateReceipts
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts"
        return self.client._send("POST", path, body, query_params)
    def create_sales_deliveries_retailer(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        - Manage Retailer Delivery
        POST CreateSalesDeliveriesRetailer
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer"
        return self.client._send("POST", path, body, query_params)
    def delete_sales_delivery_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        Permissions Required:
        - Manage Sales Delivery
        DELETE DeleteDeliveryById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_sales_delivery_retailer_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        DELETE DeleteDeliveryRetailerById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_sales_receipt_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        Permissions Required:
        - Manage Sales
        DELETE DeleteReceiptById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_sales_active_deliveries(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET GetActiveDeliveries
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/deliveries/active"
        return self.client._send("GET", path, body, query_params)
    def get_sales_active_deliveries_retailer(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET GetActiveDeliveriesRetailer
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/deliveries/retailer/active"
        return self.client._send("GET", path, body, query_params)
    def get_sales_active_receipts(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET GetActiveReceipts
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/receipts/active"
        return self.client._send("GET", path, body, query_params)
    def get_sales_counties(self) -> Any:
        """
        Returns a list of counties available for sales deliveries.
        GET GetCounties
        Parameters:
        """
        query_params = {}
        path = f"/sales/v2/counties"
        return self.client._send("GET", path, body, query_params)
    def get_sales_customer_types(self) -> Any:
        """
        Returns a list of customer types.
        GET GetCustomerTypes
        Parameters:
        """
        query_params = {}
        path = f"/sales/v2/customertypes"
        return self.client._send("GET", path, body, query_params)
    def get_sales_deliveries_return_reasons(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of return reasons for sales deliveries based on the provided License Number.
        Permissions Required:
        - Sales Delivery
        GET GetDeliveriesReturnReasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/deliveries/returnreasons"
        return self.client._send("GET", path, body, query_params)
    def get_sales_delivery_retailer_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a retailer delivery record by its ID, with an optional License Number.
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET GetDeliveryRetailerById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_sales_inactive_deliveries(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET GetInactiveDeliveries
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/deliveries/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_sales_inactive_deliveries_retailer(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET GetInactiveDeliveriesRetailer
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/deliveries/retailer/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_sales_inactive_receipts(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET GetInactiveReceipts
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/sales/v2/receipts/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_sales_patient_registration_locations(self) -> Any:
        """
        Returns a list of valid Patient registration locations for sales.
        GET GetPatientRegistrationLocations
        Parameters:
        """
        query_params = {}
        path = f"/sales/v2/patientregistration/locations"
        return self.client._send("GET", path, body, query_params)
    def get_sales_payment_types(self, license_number: Optional[str] = None) -> Any:
        """
        Returns a list of available payment types for the specified License Number.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET GetPaymentTypes
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/paymenttypes"
        return self.client._send("GET", path, body, query_params)
    def get_sales_receipt_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a sales receipt by its Id, with an optional License Number.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET GetReceiptById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_sales_receipts_external_by_external_number(self, external_number: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Sales Receipt by its external number, with an optional License Number.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET GetReceiptsExternalByExternalNumber
        Parameters:
        - external_number (str): Path parameter externalNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts/external/{urllib.parse.quote(external_number)}"
        return self.client._send("GET", path, body, query_params)
    def get_sales_delivery_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a sales delivery record by its Id, with an optional License Number.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET GetSalesDeliveryById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def update_sales_deliveries(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales Delivery
        PUT UpdateDeliveries
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_complete(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery
        PUT UpdateDeliveriesComplete
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/complete"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_hub(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales Delivery, Manage Sales Delivery Hub
        PUT UpdateDeliveriesHub
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/hub"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_hub_accept(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT UpdateDeliveriesHubAccept
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/hub/accept"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_hub_depart(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT UpdateDeliveriesHubDepart
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/hub/depart"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_hub_verify_id(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT UpdateDeliveriesHubVerifyID
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/hub/verifyID"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_deliveries_retailer(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        - Manage Retailer Delivery
        PUT UpdateDeliveriesRetailer
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/deliveries/retailer"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_receipts(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales
        PUT UpdateReceipts
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_receipts_finalize(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        Permissions Required:
        - Manage Sales
        PUT UpdateReceiptsFinalize
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts/finalize"
        return self.client._send("PUT", path, body, query_params)
    def update_sales_receipts_unfinalize(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        Permissions Required:
        - Manage Sales
        PUT UpdateReceiptsUnfinalize
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sales/v2/receipts/unfinalize"
        return self.client._send("PUT", path, body, query_params)

