from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PlantsClient:
    def __init__(self, client):
        self.client = client
    def create_plants_additives_by_location(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records additive usage for plants based on their location within a specified Facility.
        Permissions Required:
        - Manage Plants
        - Manage Plants Additives
        POST CreateAdditivesByLocation
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/additives/bylocation"
        return self.client._send("POST", path, body, query_params)
    def create_plants_additives_by_location_using_template(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records additive usage for plants by location using a predefined additive template at a specified Facility.
        Permissions Required:
        - Manage Plants Additives
        POST CreateAdditivesByLocationUsingTemplate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/additives/bylocation/usingtemplate"
        return self.client._send("POST", path, body, query_params)
    def create_plants_manicure(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates harvest product records from plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        POST CreateManicure
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/manicure"
        return self.client._send("POST", path, body, query_params)
    def create_plants_plant_batch_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages from plant batches at a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreatePlantBatchPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/plantbatch/packages"
        return self.client._send("POST", path, body, query_params)
    def create_plants_additives(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records additive usage details applied to specific plants at a Facility.
        Permissions Required:
        - Manage Plants Additives
        POST CreatePlantsAdditives
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/additives"
        return self.client._send("POST", path, body, query_params)
    def create_plants_additives_using_template(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records additive usage for plants using predefined additive templates at a specified Facility.
        Permissions Required:
        - Manage Plants Additives
        POST CreatePlantsAdditivesUsingTemplate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/additives/usingtemplate"
        return self.client._send("POST", path, body, query_params)
    def create_plants_plantings(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new plant batches at a specified Facility from existing plant data.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        POST CreatePlantsPlantings
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/plantings"
        return self.client._send("POST", path, body, query_params)
    def create_plants_waste(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records waste events for plants at a Facility, including method, reason, and location details.
        Permissions Required:
        - Manage Plants Waste
        POST CreatePlantsWaste
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/waste"
        return self.client._send("POST", path, body, query_params)
    def delete_plants(self, license_number: Optional[str] = None) -> Any:
        """
        Removes plants from a Facility’s inventory while recording the reason for their disposal.
        Permissions Required:
        - View Veg/Flower Plants
        - Destroy Veg/Flower Plants
        DELETE DeletePlants
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2"
        return self.client._send("DELETE", path, body, query_params)
    def get_plants_additives(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves additive records applied to plants at a specified Facility.
        Permissions Required:
        - View/Manage Plants Additives
        GET GetAdditives
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
        path = f"/plants/v2/additives"
        return self.client._send("GET", path, body, query_params)
    def get_plants_additives_types(self) -> Any:
        """
        Retrieves a list of all plant additive types defined within a Facility.
        GET GetAdditivesTypes
        Parameters:
        """
        query_params = {}
        path = f"/plants/v2/additives/types"
        return self.client._send("GET", path, body, query_params)
    def get_flowering_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetFloweringPlants
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
        path = f"/plants/v2/flowering"
        return self.client._send("GET", path, body, query_params)
    def get_plants_growth_phases(self, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the list of growth phases supported by a specified Facility.
        GET GetGrowthPhases
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/growthphases"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves inactive plants at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetInactivePlants
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
        path = f"/plants/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_mother_inactive_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves inactive mother-phase plants at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET GetMotherInactivePlants
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
        path = f"/plants/v2/mother/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_mother_on_hold_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves mother-phase plants currently marked as on hold at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET GetMotherOnHoldPlants
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
        path = f"/plants/v2/mother/onhold"
        return self.client._send("GET", path, body, query_params)
    def get_mother_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves mother-phase plants at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET GetMotherPlants
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
        path = f"/plants/v2/mother"
        return self.client._send("GET", path, body, query_params)
    def get_on_hold_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves plants that are currently on hold at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetOnHoldPlants
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
        path = f"/plants/v2/onhold"
        return self.client._send("GET", path, body, query_params)
    def get_plants_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Plant by Id.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetPlantsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_plants_by_label(self, label: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Plant by label.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetPlantsByLabel
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/{urllib.parse.quote(label)}"
        return self.client._send("GET", path, body, query_params)
    def get_plants_waste(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of recorded plant waste events for a specific Facility.
        Permissions Required:
        - View Plants Waste
        GET GetPlantsWaste
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
        path = f"/plants/v2/waste"
        return self.client._send("GET", path, body, query_params)
    def get_plants_waste_methods(self, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of all available plant waste methods for use within a Facility.
        GET GetPlantsWasteMethods
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/plants/v2/waste/methods/all"
        return self.client._send("GET", path, body, query_params)
    def get_plants_waste_reasons(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retriveves available reasons for recording mature plant waste at a specified Facility.
        GET GetPlantsWasteReasons
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
        path = f"/plants/v2/waste/reasons"
        return self.client._send("GET", path, body, query_params)
    def get_vegetative_plants(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Veg/Flower Plants
        GET GetVegetativePlants
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
        path = f"/plants/v2/vegetative"
        return self.client._send("GET", path, body, query_params)
    def get_plants_waste_by_id(self, id: str, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
        Permissions Required:
        - View Plants Waste
        GET GetWasteById
        Parameters:
        - id (str): Path parameter id
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
        path = f"/plants/v2/waste/{urllib.parse.quote(id)}/plant"
        return self.client._send("GET", path, body, query_params)
    def get_plants_waste_package_by_id(self, id: str, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of package records linked to the specified plantWasteId for a given facility.
        Permissions Required:
        - View Plants Waste
        GET GetWastePackageById
        Parameters:
        - id (str): Path parameter id
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
        path = f"/plants/v2/waste/{urllib.parse.quote(id)}/package"
        return self.client._send("GET", path, body, query_params)
    def update_adjust_plants(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Adjusts the recorded count of plants at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdateAdjustPlants
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/adjust"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_growth_phase(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Changes the growth phases of plants within a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdateGrowthPhase
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/growthphase"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_harvest(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        PUT UpdateHarvest
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/harvest"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_merge(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Merges multiple plant groups into a single group within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        PUT UpdateMerge
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/merge"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_location(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Moves plant batches to new locations within a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdatePlantsLocation
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/location"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_strain(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the strain information for plants within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdatePlantsStrain
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/strain"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_tag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Replaces existing plant tags with new tags for plants within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdatePlantsTag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/tag"
        return self.client._send("PUT", path, body, query_params)
    def update_plants_split(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Splits an existing plant group into multiple groups within a Facility.
        Permissions Required:
        - View Plant
        PUT UpdateSplit
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plants/v2/split"
        return self.client._send("PUT", path, body, query_params)

