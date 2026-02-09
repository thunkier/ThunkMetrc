from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PlantBatchesClient:
    def __init__(self, client):
        self.client = client
    def create_adjust_plant_batches(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST CreateAdjustPlantBatches
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/adjust"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_growth_phase(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates the growth phase of plants at a specified Facility based on tracking information.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        POST CreateGrowthPhase
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/growthphase"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_packages_from_mother_plant(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages from mother plants at the specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreatePackagesFromMotherPlant
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/packages/frommotherplant"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_additives(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records Additive usage details for plant batches at a specific Facility.
        Permissions Required:
        - Manage Plants Additives
        POST CreatePlantBatchesAdditives
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/additives"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_additives_using_template(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records Additive usage for plant batches at a Facility using predefined additive templates.
        Permissions Required:
        - Manage Plants Additives
        POST CreatePlantBatchesAdditivesUsingTemplate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/additives/usingtemplate"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST CreatePlantBatchesPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/packages"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_plantings(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new plantings for a Facility by generating plant batches based on provided planting details.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST CreatePlantBatchesPlantings
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/plantings"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_waste(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records waste information for plant batches based on the submitted data for the specified Facility.
        Permissions Required:
        - Manage Plants Waste
        POST CreatePlantBatchesWaste
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/waste"
        return self.client._send("POST", path, body, query_params)
    def create_plant_batches_split(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Splits an existing Plant Batch into multiple groups at the specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST CreateSplit
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/split"
        return self.client._send("POST", path, body, query_params)
    def delete_plant_batches(self, license_number: Optional[str] = None) -> Any:
        """
        Completes the destruction of plant batches based on the provided input data.
        Permissions Required:
        - View Immature Plants
        - Destroy Immature Plants
        DELETE DeletePlantBatches
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_plant_batches(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Immature Plants
        GET GetActivePlantBatches
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
        path = f"/plantbatches/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_plant_batches(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Immature Plants
        GET GetInactivePlantBatches
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
        path = f"/plantbatches/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_plant_batches_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Plant Batch by Id.
        Permissions Required:
        - View Immature Plants
        GET GetPlantBatchesById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_plant_batches_types(self, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of plant batch types.
        GET GetPlantBatchesTypes
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/plantbatches/v2/types"
        return self.client._send("GET", path, body, query_params)
    def get_plant_batches_waste(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves waste details associated with plant batches at a specified Facility.
        Permissions Required:
        - View Plants Waste
        GET GetPlantBatchesWaste
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
        path = f"/plantbatches/v2/waste"
        return self.client._send("GET", path, body, query_params)
    def get_plant_batches_waste_reasons(self, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        GET GetPlantBatchesWasteReasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/waste/reasons"
        return self.client._send("GET", path, body, query_params)
    def update_plant_batches_name(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Renames plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdateName
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/name"
        return self.client._send("PUT", path, body, query_params)
    def update_plant_batches_location(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Moves one or more plant batches to new locations with in a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants
        PUT UpdatePlantBatchesLocation
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/location"
        return self.client._send("PUT", path, body, query_params)
    def update_plant_batches_strain(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Changes the strain of plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdatePlantBatchesStrain
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/strain"
        return self.client._send("PUT", path, body, query_params)
    def update_plant_batches_tag(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Replaces tags for plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT UpdatePlantBatchesTag
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/plantbatches/v2/tag"
        return self.client._send("PUT", path, body, query_params)

