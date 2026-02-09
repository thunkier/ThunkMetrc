from typing import Any, Optional, Dict, List, Union
import urllib.parse

class AdditivesTemplatesClient:
    def __init__(self, client):
        self.client = client
    def create_additives_templates(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        POST CreateAdditivesTemplates
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/additivestemplates/v2"
        return self.client._send("POST", path, body, query_params)
    def get_active_additives_templates(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        GET GetActiveAdditivesTemplates
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
        path = f"/additivestemplates/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_additives_templates_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves an Additive Template by its Id.
        Permissions Required:
        - Manage Additives
        GET GetAdditivesTemplatesById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/additivestemplates/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_additives_templates(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        GET GetInactiveAdditivesTemplates
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
        path = f"/additivestemplates/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def update_additives_templates(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        PUT UpdateAdditivesTemplates
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/additivestemplates/v2"
        return self.client._send("PUT", path, body, query_params)

