from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class SalesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_sales_deliveries(self, body: Optional[List['SalesCreateDeliveriesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        POST /sales/v2/deliveries
        Parameters:
        - body (List['SalesCreateDeliveriesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_deliveries(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_sales_deliveries_retailer_depart(self, body: Optional[List['SalesCreateDeliveriesRetailerDepartRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST /sales/v2/deliveries/retailer/depart
        Parameters:
        - body (List['SalesCreateDeliveriesRetailerDepartRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_deliveries_retailer_depart(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_sales_deliveries_retailer_end(self, body: Optional[List['SalesCreateDeliveriesRetailerEndRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST /sales/v2/deliveries/retailer/end
        Parameters:
        - body (List['SalesCreateDeliveriesRetailerEndRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_deliveries_retailer_end(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_sales_deliveries_retailer_restock(self, body: Optional[List['SalesCreateDeliveriesRetailerRestockRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        POST /sales/v2/deliveries/retailer/restock
        Parameters:
        - body (List['SalesCreateDeliveriesRetailerRestockRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_deliveries_retailer_restock(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_sales_receipts(self, body: Optional[List['SalesCreateReceiptsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales (Write)
        - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
        - Industry/Facility Type/Advanced Sales
        - WebApi Sales Read Write State (All or WriteOnly)
        - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
        POST /sales/v2/receipts
        Parameters:
        - body (List['SalesCreateReceiptsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_receipts(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_sales_deliveries_retailer(self, body: Optional[List['SalesCreateSalesDeliveriesRetailerRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
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
        POST /sales/v2/deliveries/retailer
        Parameters:
        - body (List['SalesCreateSalesDeliveriesRetailerRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.create_sales_deliveries_retailer(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_sales_delivery_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Voids a sales delivery for a Facility using the provided License Number and delivery Id.
        Permissions Required:
        - Manage Sales Delivery
        DELETE /sales/v2/deliveries/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.delete_sales_delivery_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_sales_delivery_retailer_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
        - Industry/Facility Type/Retailer Delivery
        - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
        - WebApi Sales Deliveries Read Write State (All or WriteOnly)
        - Manage Retailer Delivery
        DELETE /sales/v2/deliveries/retailer/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.delete_sales_delivery_retailer_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_sales_receipt_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a sales receipt for a Facility using the provided License Number and receipt Id.
        Permissions Required:
        - Manage Sales
        DELETE /sales/v2/receipts/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.delete_sales_receipt_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_sales_active_deliveries(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ActiveDelivery']:
        """
        Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET /sales/v2/deliveries/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_active_deliveries(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_active_deliveries_retailer(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ActiveDeliveriesRetailer']:
        """
        Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET /sales/v2/deliveries/retailer/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_active_deliveries_retailer(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_active_receipts(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ActiveReceipt']:
        """
        Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET /sales/v2/receipts/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_active_receipts(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_counties(self, body: Any = None) -> PaginatedResponse['County']:
        """
        Returns a list of counties available for sales deliveries.
        GET /sales/v2/counties
        Parameters:
        """
        async def op():
            return await self.client.sales.get_sales_counties(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_customer_types(self, body: Any = None) -> Any:
        """
        Returns a list of customer types.
        GET /sales/v2/customertypes
        Parameters:
        """
        async def op():
            return await self.client.sales.get_sales_customer_types(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_deliveries_return_reasons(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveriesReturnReason']:
        """
        Returns a list of return reasons for sales deliveries based on the provided License Number.
        Permissions Required:
        - Sales Delivery
        GET /sales/v2/deliveries/returnreasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_deliveries_return_reasons(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_delivery_retailer_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'DeliveryRetailer':
        """
        Retrieves a retailer delivery record by its ID, with an optional License Number.
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET /sales/v2/deliveries/retailer/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.get_sales_delivery_retailer_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_inactive_deliveries(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['InactiveDelivery']:
        """
        Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET /sales/v2/deliveries/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_inactive_deliveries(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_inactive_deliveries_retailer(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['InactiveDeliveriesRetailer']:
        """
        Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
        Permissions Required:
        - View Retailer Delivery
        - Manage Retailer Delivery
        GET /sales/v2/deliveries/retailer/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_inactive_deliveries_retailer(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_inactive_receipts(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['InactiveReceipt']:
        """
        Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET /sales/v2/receipts/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sales.get_sales_inactive_receipts(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_patient_registration_locations(self, body: Any = None) -> List['PatientRegistrationLocation']:
        """
        Returns a list of valid Patient registration locations for sales.
        GET /sales/v2/patientregistration/locations
        Parameters:
        """
        async def op():
            return await self.client.sales.get_sales_patient_registration_locations(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_payment_types(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Returns a list of available payment types for the specified License Number.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET /sales/v2/paymenttypes
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.get_sales_payment_types(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_receipt_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'ActiveReceipt':
        """
        Retrieves a sales receipt by its Id, with an optional License Number.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET /sales/v2/receipts/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.get_sales_receipt_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_receipts_external_by_external_number(self, external_number: str, body: Any = None, license_number: Optional[str] = None) -> List['ReceiptsExternalByExternalNumber']:
        """
        Retrieves a Sales Receipt by its external number, with an optional License Number.
        Permissions Required:
        - View Sales
        - Manage Sales
        GET /sales/v2/receipts/external/{externalNumber}
        Parameters:
        - external_number (str): Path parameter externalNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.get_sales_receipts_external_by_external_number(external_number, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sales_delivery_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'SalesDelivery':
        """
        Retrieves a sales delivery record by its Id, with an optional License Number.
        Permissions Required:
        - View Sales Delivery
        - Manage Sales Delivery
        GET /sales/v2/deliveries/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.get_sales_delivery_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_sales_deliveries(self, body: Optional[List['SalesUpdateDeliveriesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales Delivery
        PUT /sales/v2/deliveries
        Parameters:
        - body (List['SalesUpdateDeliveriesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_complete(self, body: Optional[List['SalesUpdateDeliveriesCompleteRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery
        PUT /sales/v2/deliveries/complete
        Parameters:
        - body (List['SalesUpdateDeliveriesCompleteRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_complete(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_hub(self, body: Optional[List['SalesUpdateDeliveriesHubRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales Delivery, Manage Sales Delivery Hub
        PUT /sales/v2/deliveries/hub
        Parameters:
        - body (List['SalesUpdateDeliveriesHubRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_hub(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_hub_accept(self, body: Optional[List['SalesUpdateDeliveriesHubAcceptRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT /sales/v2/deliveries/hub/accept
        Parameters:
        - body (List['SalesUpdateDeliveriesHubAcceptRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_hub_accept(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_hub_depart(self, body: Optional[List['SalesUpdateDeliveriesHubDepartRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT /sales/v2/deliveries/hub/depart
        Parameters:
        - body (List['SalesUpdateDeliveriesHubDepartRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_hub_depart(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_hub_verify_id(self, body: Optional[List['SalesUpdateDeliveriesHubVerifyIDRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
        Permissions Required:
        - Manage Sales Delivery Hub
        PUT /sales/v2/deliveries/hub/verifyID
        Parameters:
        - body (List['SalesUpdateDeliveriesHubVerifyIDRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_hub_verify_id(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_deliveries_retailer(self, body: Optional[List['SalesUpdateDeliveriesRetailerRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
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
        PUT /sales/v2/deliveries/retailer
        Parameters:
        - body (List['SalesUpdateDeliveriesRetailerRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_deliveries_retailer(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_receipts(self, body: Optional[List['SalesUpdateReceiptsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
        Permissions Required:
        - Manage Sales
        PUT /sales/v2/receipts
        Parameters:
        - body (List['SalesUpdateReceiptsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_receipts(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_receipts_finalize(self, body: Optional[List['SalesUpdateReceiptsFinalizeRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        Permissions Required:
        - Manage Sales
        PUT /sales/v2/receipts/finalize
        Parameters:
        - body (List['SalesUpdateReceiptsFinalizeRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_receipts_finalize(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_sales_receipts_unfinalize(self, body: Optional[List['SalesUpdateReceiptsUnfinalizeRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
        Permissions Required:
        - Manage Sales
        PUT /sales/v2/receipts/unfinalize
        Parameters:
        - body (List['SalesUpdateReceiptsUnfinalizeRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sales.update_sales_receipts_unfinalize(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

