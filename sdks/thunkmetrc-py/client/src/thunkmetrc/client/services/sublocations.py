from typing import Any, Optional, Dict, List, Union
import urllib.parse

class SublocationsClient:
    def __init__(self, client):
        self.client = client
    def create_sublocations(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new sublocation records for a Facility.
        Permissions Required:
        - Manage Locations
        POST CreateSublocations
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sublocations/v2"
        return self.client._send("POST", path, body, query_params)
    def delete_sublocations_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives an existing Sublocation record for a Facility.
        Permissions Required:
        - Manage Locations
        DELETE DeleteSublocationsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sublocations/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_sublocations(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Locations
        GET GetActiveSublocations
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
        path = f"/sublocations/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_sublocations(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive sublocations for the specified Facility.
        Permissions Required:
        - Manage Locations
        GET GetInactiveSublocations
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
        path = f"/sublocations/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_sublocations_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Sublocation by its Id, with an optional license number.
        Permissions Required:
        - Manage Locations
        GET GetSublocationsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sublocations/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def update_sublocations(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing sublocation records for a specified Facility.
        Permissions Required:
        - Manage Locations
        PUT UpdateSublocations
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/sublocations/v2"
        return self.client._send("PUT", path, body, query_params)

