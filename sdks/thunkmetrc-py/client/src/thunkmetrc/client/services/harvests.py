from typing import Any, Optional, Dict, List, Union
import urllib.parse

class HarvestsClient:
    def __init__(self, client):
        self.client = client
    def create_harvests_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages from harvested products for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreateHarvestsPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/packages"
        return self.client._send("POST", path, body, query_params)
    def create_harvests_waste(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        POST CreateHarvestsWaste
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/waste"
        return self.client._send("POST", path, body, query_params)
    def create_harvests_packages_testing(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages for testing from harvested products for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreatePackagesTesting
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/packages/testing"
        return self.client._send("POST", path, body, query_params)
    def delete_harvests_waste_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Discontinues a specific harvest waste record by Id for the specified Facility.
        Permissions Required:
        - View Harvests
        - Discontinue Harvest Waste
        DELETE DeleteWasteById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/waste/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def finish_harvests_harvests(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Marks one or more harvests as finished for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT FinishHarvests
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/finish"
        return self.client._send("PUT", path, body, query_params)
    def get_active_harvests(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active harvests for a specified Facility.
        Permissions Required:
        - View Harvests
        GET GetActiveHarvests
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
        path = f"/harvests/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_harvests_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        Permissions Required:
        - View Harvests
        GET GetHarvestsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_harvests_waste(self, harvest_id: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
        Permissions Required:
        - View Harvests
        GET GetHarvestsWaste
        Parameters:
        - harvest_id (str, optional): Filter by harvestId
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if harvest_id is not None:
            query_params["harvestId"] = harvest_id
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/harvests/v2/waste"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_harvests(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive harvests for a specified Facility.
        Permissions Required:
        - View Harvests
        GET GetInactiveHarvests
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
        path = f"/harvests/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_on_hold_harvests(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of harvests on hold for a specified Facility.
        Permissions Required:
        - View Harvests
        GET GetOnHoldHarvests
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
        path = f"/harvests/v2/onhold"
        return self.client._send("GET", path, body, query_params)
    def get_harvests_waste_types(self, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of Waste types for harvests.
        GET GetWasteTypes
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/harvests/v2/waste/types"
        return self.client._send("GET", path, body, query_params)
    def unfinish_harvests_harvests(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Reopens one or more previously finished harvests for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT UnfinishHarvests
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/unfinish"
        return self.client._send("PUT", path, body, query_params)
    def update_harvests_location(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the Location of Harvest for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        PUT UpdateHarvestsLocation
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/location"
        return self.client._send("PUT", path, body, query_params)
    def update_harvests_rename(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Renames one or more harvests for the specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        PUT UpdateRename
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/rename"
        return self.client._send("PUT", path, body, query_params)
    def update_harvests_restore_harvested_plants(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Restores previously harvested plants to their original state for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT UpdateRestoreHarvestedPlants
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/harvests/v2/restore/harvestedplants"
        return self.client._send("PUT", path, body, query_params)

