from typing import Any, Optional, Dict, List, Union
import urllib.parse

class StrainsClient:
    def __init__(self, client):
        self.client = client
    def create_strains(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new strain records for a specified Facility.
        Permissions Required:
        - Manage Strains
        POST CreateStrains
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/strains/v2"
        return self.client._send("POST", path, body, query_params)
    def delete_strains_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives an existing strain record for a Facility
        Permissions Required:
        - Manage Strains
        DELETE DeleteStrainsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/strains/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_strains(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Strains
        GET GetActiveStrains
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
        path = f"/strains/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_strains(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Strains
        GET GetInactiveStrains
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
        path = f"/strains/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_strains_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Strain record by its Id, with an optional license number.
        Permissions Required:
        - Manage Strains
        GET GetStrainsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/strains/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def update_strains(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing strain records for a specified Facility.
        Permissions Required:
        - Manage Strains
        PUT UpdateStrains
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/strains/v2"
        return self.client._send("PUT", path, body, query_params)

