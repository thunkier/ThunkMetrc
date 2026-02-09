from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PackagesClient:
    def __init__(self, client):
        self.client = client
    def create_adjust_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records a list of adjustments for packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        POST CreateAdjustPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/adjust"
        return self.client._send("POST", path, body, query_params)
    def create_packages_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new packages for a specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreatePackagesPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2"
        return self.client._send("POST", path, body, query_params)
    def create_packages_plantings(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new plantings from packages for a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants
        - View Packages
        - Manage Packages Inventory
        POST CreatePackagesPlantings
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/plantings"
        return self.client._send("POST", path, body, query_params)
    def create_packages_testing(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new packages for testing for a specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreateTesting
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/testing"
        return self.client._send("POST", path, body, query_params)
    def delete_packages_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Discontinues a Package at a specific Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        DELETE DeletePackagesById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def finish_packages_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates a list of packages as finished for a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT FinishPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/finish"
        return self.client._send("PUT", path, body, query_params)
    def finishedgood_flag_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Flags one or more Packages at the specified Facility as Finished Goods.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT FinishedgoodFlag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/finishedgood/flag"
        return self.client._send("PUT", path, body, query_params)
    def finishedgood_unflag_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Removes the Finished Good flag one or more Packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT FinishedgoodUnflag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/finishedgood/unflag"
        return self.client._send("PUT", path, body, query_params)
    def get_active_packages(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active packages for a specified Facility.
        Permissions Required:
        - View Packages
        GET GetActivePackages
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
        path = f"/packages/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_packages_adjust_reasons(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of adjustment reasons for packages at a specified Facility.
        GET GetAdjustReasons
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
        path = f"/packages/v2/adjust/reasons"
        return self.client._send("GET", path, body, query_params)
    def get_packages_adjustments(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves the Package Adjustments for a Facility
        Permissions Required:
        - View Packages
        GET GetAdjustments
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
        path = f"/packages/v2/adjustments"
        return self.client._send("GET", path, body, query_params)
    def get_in_transit_packages(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of packages in transit for a specified Facility.
        Permissions Required:
        - View Packages
        GET GetInTransitPackages
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
        path = f"/packages/v2/intransit"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_packages(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive packages for a specified Facility.
        Permissions Required:
        - View Packages
        GET GetInactivePackages
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
        path = f"/packages/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_packages_lab_samples(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        Permissions Required:
        - View Packages
        GET GetLabSamples
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
        path = f"/packages/v2/labsamples"
        return self.client._send("GET", path, body, query_params)
    def get_on_hold_packages(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of packages on hold for a specified Facility.
        Permissions Required:
        - View Packages
        GET GetOnHoldPackages
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
        path = f"/packages/v2/onhold"
        return self.client._send("GET", path, body, query_params)
    def get_packages_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Package by its Id.
        Permissions Required:
        - View Packages
        GET GetPackagesById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_packages_by_label(self, label: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Package by its label.
        Permissions Required:
        - View Packages
        GET GetPackagesByLabel
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/{urllib.parse.quote(label)}"
        return self.client._send("GET", path, body, query_params)
    def get_packages_types(self) -> Any:
        """
        Retrieves a list of available Package types.
        GET GetPackagesTypes
        Parameters:
        """
        query_params = {}
        path = f"/packages/v2/types"
        return self.client._send("GET", path, body, query_params)
    def get_packages_source_harvest_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the source harvests for a Package by its Id.
        Permissions Required:
        - View Package Source Harvests
        GET GetSourceHarvestById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/{urllib.parse.quote(id)}/source/harvests"
        return self.client._send("GET", path, body, query_params)
    def get_packages_transferred(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of transferred packages for a specific Facility.
        Permissions Required:
        - View Packages
        GET GetTransferred
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
        path = f"/packages/v2/transferred"
        return self.client._send("GET", path, body, query_params)
    def unfinish_packages_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates a list of packages as unfinished for a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UnfinishPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/unfinish"
        return self.client._send("PUT", path, body, query_params)
    def update_adjust_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Set the final quantity for a Package.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateAdjustPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/adjust"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_decontaminate(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the Product decontaminate information for a list of packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateDecontaminate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/decontaminate"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_donation_flag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Flags one or more packages for donation at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateDonationFlag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/donation/flag"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_donation_unflag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Removes the donation flag from one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateDonationUnflag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/donation/unflag"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_externalid(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the external identifiers for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Package Inventory
        - External Id Enabled
        PUT UpdateExternalid
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/externalid"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_item(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the associated Item for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT UpdateItem
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/item"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_labtests_required(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the list of required lab test batches for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT UpdateLabtestsRequired
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/labtests/required"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_note(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates notes associated with one or more packages for the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        - Manage Package Notes
        PUT UpdateNote
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/note"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_location(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the Location and Sublocation for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT UpdatePackagesLocation
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/location"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_remediate(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates a list of Product remediations for packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateRemediate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/remediate"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_trade_sample_flag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Flags or unflags one or more packages at the specified Facility as trade samples.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateTradeSampleFlag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/tradesample/flag"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_trade_sample_unflag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Removes the trade sample flag from one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateTradeSampleUnflag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/tradesample/unflag"
        return self.client._send("PUT", path, body, query_params)
    def update_packages_use_by_date(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the use-by date for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT UpdateUseByDate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/packages/v2/usebydate"
        return self.client._send("PUT", path, body, query_params)

