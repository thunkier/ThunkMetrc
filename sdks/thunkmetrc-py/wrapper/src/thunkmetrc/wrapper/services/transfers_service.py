from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class TransfersService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_transfers_hub_arrive(self, body: Optional[List['TransfersCreateHubArriveRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Arrive a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST /transfers/v2/hub/arrive
        Parameters:
        - body (List['TransfersCreateHubArriveRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_hub_arrive(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transfers_hub_checkin(self, body: Optional[List['TransfersCreateHubCheckinRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        CheckIn a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST /transfers/v2/hub/checkin
        Parameters:
        - body (List['TransfersCreateHubCheckinRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_hub_checkin(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transfers_hub_checkout(self, body: Optional[List['TransfersCreateHubCheckoutRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        CheckOut a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST /transfers/v2/hub/checkout
        Parameters:
        - body (List['TransfersCreateHubCheckoutRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_hub_checkout(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transfers_hub_depart(self, body: Optional[List['TransfersCreateHubDepartRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Depart a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST /transfers/v2/hub/depart
        Parameters:
        - body (List['TransfersCreateHubDepartRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_hub_depart(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transfers_incoming_external(self, body: Optional[List['TransfersCreateIncomingExternalRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates external incoming shipment plans for a Facility.
        Permissions Required:
        - Manage Transfers
        POST /transfers/v2/external/incoming
        Parameters:
        - body (List['TransfersCreateIncomingExternalRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_incoming_external(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transfers_outgoing_templates(self, body: Optional[List['TransfersCreateOutgoingTemplatesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new transfer templates for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        POST /transfers/v2/templates/outgoing
        Parameters:
        - body (List['TransfersCreateOutgoingTemplatesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.create_transfers_outgoing_templates(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_transfers_incoming_external_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Voids an external incoming shipment plan for a Facility.
        Permissions Required:
        - Manage Transfers
        DELETE /transfers/v2/external/incoming/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.delete_transfers_incoming_external_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_transfers_outgoing_template_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a transfer template for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        DELETE /transfers/v2/templates/outgoing/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.delete_transfers_outgoing_template_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_transfers_deliveries_packages_states(self, body: Any = None) -> Any:
        """
        Returns a list of available shipment Package states.
        GET /transfers/v2/deliveries/packages/states
        Parameters:
        """
        async def op():
            return await self.client.transfers.get_transfers_deliveries_packages_states(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_package_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveryPackage']:
        """
        Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/deliveries/{id}/packages
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_package_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_package_requiredlabtestbatch_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveryPackageRequiredlabtestbatch']:
        """
        Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/deliveries/package/{id}/requiredlabtestbatches
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_package_requiredlabtestbatch_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_package_wholesale_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveryPackageWholesale']:
        """
        Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/deliveries/{id}/packages/wholesale
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_package_wholesale_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_transporter_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveryTransporter']:
        """
        Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/deliveries/{id}/transporters
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_transporter_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_transporter_detail_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['DeliveryTransporterDetail']:
        """
        Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/deliveries/{id}/transporters/details
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_transporter_detail_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_hub(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Hub']:
        """
        Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/hub
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_hub(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_incoming_transfers(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Transfer']:
        """
        Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/incoming
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_incoming_transfers(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_manifest_html_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/manifest/{id}/html
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.get_transfers_manifest_html_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_manifest_pdf_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'ManifestPdf':
        """
        Get Transfer Manifest PDF for a given Transfer Id
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/manifest/{id}/pdf
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.get_transfers_manifest_pdf_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_outgoing_template_delivery_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TemplateDelivery']:
        """
        Retrieves a list of deliveries associated with a specific transfer template.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/templates/outgoing/{id}/deliveries
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_outgoing_template_delivery_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_outgoing_template_delivery_package_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TemplateDeliveryPackage']:
        """
        Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/templates/outgoing/deliveries/{id}/packages
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_outgoing_template_delivery_package_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_outgoing_template_delivery_transporter_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TemplateDeliveryTransporter']:
        """
        Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_outgoing_template_delivery_transporter_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_outgoing_template_delivery_transporter_detail_by_id(self, id: str, body: Any = None) -> 'TemplateDeliveryTransporterDetail':
        """
        Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
        Parameters:
        - id (str): Path parameter id
        """
        async def op():
            return await self.client.transfers.get_transfers_outgoing_template_delivery_transporter_detail_by_id(id, body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_outgoing_templates(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Template']:
        """
        Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET /transfers/v2/templates/outgoing
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_outgoing_templates(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_outgoing_transfers(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Transfer']:
        """
        Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/outgoing
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_outgoing_transfers(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_rejected_transfers(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Transfer']:
        """
        Retrieves a list of shipments with rejected packages for a Facility.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/rejected
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_rejected_transfers(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_delivery_by_id(self, id: str, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TransfersDelivery']:
        """
        Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET /transfers/v2/{id}/deliveries
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_delivery_by_id(id, body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transfers_types(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TransfersType']:
        """
        Retrieves a list of available transfer types for a Facility based on its license number.
        GET /transfers/v2/types
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transfers.get_transfers_types(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def update_transfers_incoming_external(self, body: Optional[List['TransfersUpdateIncomingExternalRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates external incoming shipment plans for a Facility.
        Permissions Required:
        - Manage Transfers
        PUT /transfers/v2/external/incoming
        Parameters:
        - body (List['TransfersUpdateIncomingExternalRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.update_transfers_incoming_external(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_transfers_outgoing_templates(self, body: Optional[List['TransfersUpdateOutgoingTemplatesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing transfer templates for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        PUT /transfers/v2/templates/outgoing
        Parameters:
        - body (List['TransfersUpdateOutgoingTemplatesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transfers.update_transfers_outgoing_templates(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

