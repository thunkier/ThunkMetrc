from typing import Any, Optional, Dict, List, Union
import urllib.parse

class TagsClient:
    def __init__(self, client):
        self.client = client
    def get_tags_available_package(self, license_number: Optional[str] = None) -> Any:
        """
        Returns a list of available package tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        GET GetAvailablePackage
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/tags/v2/package/available"
        return self.client._send("GET", path, body, query_params)
    def get_tags_available_plant(self, license_number: Optional[str] = None) -> Any:
        """
        Returns a list of available plant tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        GET GetAvailablePlant
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/tags/v2/plant/available"
        return self.client._send("GET", path, body, query_params)
    def get_staged_tags(self, license_number: Optional[str] = None) -> Any:
        """
        Returns a list of staged tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        - RetailId.AllowPackageStaging Key Value enabled
        GET GetStagedTags
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/tags/v2/staged"
        return self.client._send("GET", path, body, query_params)

