from typing import Any, Optional, Dict, List, Union
import urllib.parse

class LocationsClient:
    def __init__(self, client):
        self.client = client
    def create_locations(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        POST CreateLocations
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/locations/v2"
        return self.client._send("POST", path, body, query_params)
    def delete_locations_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a specified Location, identified by its Id, for a Facility.
        Permissions Required:
        - Manage Locations
        DELETE DeleteLocationsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/locations/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_locations(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET GetActiveLocations
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
        path = f"/locations/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_locations(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET GetInactiveLocations
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
        path = f"/locations/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_locations_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Location by its Id.
        Permissions Required:
        - Manage Locations
        GET GetLocationsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/locations/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_locations_types(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active location types for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET GetLocationsTypes
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
        path = f"/locations/v2/types"
        return self.client._send("GET", path, body, query_params)
    def update_locations(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        PUT UpdateLocations
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/locations/v2"
        return self.client._send("PUT", path, body, query_params)

