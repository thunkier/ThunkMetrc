from typing import Any, Optional, Dict, List, Union
import urllib.parse

class TransfersClient:
    def __init__(self, client):
        self.client = client
    def create_transfers_hub_arrive(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Arrive a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST CreateHubArrive
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/hub/arrive"
        return self.client._send("POST", path, body, query_params)
    def create_transfers_hub_checkin(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        CheckIn a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST CreateHubCheckin
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/hub/checkin"
        return self.client._send("POST", path, body, query_params)
    def create_transfers_hub_checkout(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        CheckOut a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST CreateHubCheckout
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/hub/checkout"
        return self.client._send("POST", path, body, query_params)
    def create_transfers_hub_depart(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Depart a transfer for a Facility.
        Permissions Required:
        - Manage Transfer Hub
        POST CreateHubDepart
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/hub/depart"
        return self.client._send("POST", path, body, query_params)
    def create_transfers_incoming_external(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates external incoming shipment plans for a Facility.
        Permissions Required:
        - Manage Transfers
        POST CreateIncomingExternal
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/external/incoming"
        return self.client._send("POST", path, body, query_params)
    def create_transfers_outgoing_templates(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new transfer templates for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        POST CreateOutgoingTemplates
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/templates/outgoing"
        return self.client._send("POST", path, body, query_params)
    def delete_transfers_incoming_external_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Voids an external incoming shipment plan for a Facility.
        Permissions Required:
        - Manage Transfers
        DELETE DeleteIncomingExternalById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/external/incoming/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_transfers_outgoing_template_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a transfer template for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        DELETE DeleteOutgoingTemplateById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/templates/outgoing/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_transfers_deliveries_packages_states(self) -> Any:
        """
        Returns a list of available shipment Package states.
        GET GetDeliveriesPackagesStates
        Parameters:
        """
        query_params = {}
        path = f"/transfers/v2/deliveries/packages/states"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_package_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetDeliveryPackageById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/deliveries/{urllib.parse.quote(id)}/packages"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_package_requiredlabtestbatch_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetDeliveryPackageRequiredlabtestbatchById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/deliveries/package/{urllib.parse.quote(id)}/requiredlabtestbatches"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_package_wholesale_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetDeliveryPackageWholesaleById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/deliveries/{urllib.parse.quote(id)}/packages/wholesale"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_transporter_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetDeliveryTransporterById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/deliveries/{urllib.parse.quote(id)}/transporters"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_transporter_detail_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetDeliveryTransporterDetailById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/deliveries/{urllib.parse.quote(id)}/transporters/details"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_hub(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetHub
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
        path = f"/transfers/v2/hub"
        return self.client._send("GET", path, body, query_params)
    def get_incoming_transfers(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetIncomingTransfers
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
        path = f"/transfers/v2/incoming"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_manifest_html_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetManifestHtmlById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/manifest/{urllib.parse.quote(id)}/html"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_manifest_pdf_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Get Transfer Manifest PDF for a given Transfer Id
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetManifestPdfById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/manifest/{urllib.parse.quote(id)}/pdf"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_outgoing_template_delivery_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of deliveries associated with a specific transfer template.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetOutgoingTemplateDeliveryById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/templates/outgoing/{urllib.parse.quote(id)}/deliveries"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_outgoing_template_delivery_package_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetOutgoingTemplateDeliveryPackageById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/templates/outgoing/deliveries/{urllib.parse.quote(id)}/packages"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_outgoing_template_delivery_transporter_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetOutgoingTemplateDeliveryTransporterById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/templates/outgoing/deliveries/{urllib.parse.quote(id)}/transporters"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_outgoing_template_delivery_transporter_detail_by_id(self, id: str) -> Any:
        """
        Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetOutgoingTemplateDeliveryTransporterDetailById
        Parameters:
        - id (str): Path parameter id
        """
        query_params = {}
        path = f"/transfers/v2/templates/outgoing/deliveries/{urllib.parse.quote(id)}/transporters/details"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_outgoing_templates(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfer Templates
        - View Transfer Templates
        GET GetOutgoingTemplates
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
        path = f"/transfers/v2/templates/outgoing"
        return self.client._send("GET", path, body, query_params)
    def get_outgoing_transfers(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetOutgoingTransfers
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
        path = f"/transfers/v2/outgoing"
        return self.client._send("GET", path, body, query_params)
    def get_rejected_transfers(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of shipments with rejected packages for a Facility.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetRejectedTransfers
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
        path = f"/transfers/v2/rejected"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_delivery_by_id(self, id: str, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
        Permissions Required:
        - Manage Transfers
        - View Transfers
        GET GetTransfersDeliveryById
        Parameters:
        - id (str): Path parameter id
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transfers/v2/{urllib.parse.quote(id)}/deliveries"
        return self.client._send("GET", path, body, query_params)
    def get_transfers_types(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of available transfer types for a Facility based on its license number.
        GET GetTransfersTypes
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
        path = f"/transfers/v2/types"
        return self.client._send("GET", path, body, query_params)
    def update_transfers_incoming_external(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates external incoming shipment plans for a Facility.
        Permissions Required:
        - Manage Transfers
        PUT UpdateIncomingExternal
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/external/incoming"
        return self.client._send("PUT", path, body, query_params)
    def update_transfers_outgoing_templates(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing transfer templates for a Facility.
        Permissions Required:
        - Manage Transfer Templates
        PUT UpdateOutgoingTemplates
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transfers/v2/templates/outgoing"
        return self.client._send("PUT", path, body, query_params)

