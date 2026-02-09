from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PlantsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_plants_additives_by_location(self, body: Optional[List['PlantsCreateAdditivesByLocationRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records additive usage for plants based on their location within a specified Facility.
        Permissions Required:
        - Manage Plants
        - Manage Plants Additives
        POST /plants/v2/additives/bylocation
        Parameters:
        - body (List['PlantsCreateAdditivesByLocationRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_additives_by_location(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_additives_by_location_using_template(self, body: Optional[List['PlantsCreateAdditivesByLocationUsingTemplateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records additive usage for plants by location using a predefined additive template at a specified Facility.
        Permissions Required:
        - Manage Plants Additives
        POST /plants/v2/additives/bylocation/usingtemplate
        Parameters:
        - body (List['PlantsCreateAdditivesByLocationUsingTemplateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_additives_by_location_using_template(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_manicure(self, body: Optional[List['PlantsCreateManicureRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates harvest product records from plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        POST /plants/v2/manicure
        Parameters:
        - body (List['PlantsCreateManicureRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_manicure(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_plant_batch_packages(self, body: Optional[List['PlantsCreatePlantBatchPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages from plant batches at a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /plants/v2/plantbatch/packages
        Parameters:
        - body (List['PlantsCreatePlantBatchPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_plant_batch_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_additives(self, body: Optional[List['PlantsCreatePlantsAdditivesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records additive usage details applied to specific plants at a Facility.
        Permissions Required:
        - Manage Plants Additives
        POST /plants/v2/additives
        Parameters:
        - body (List['PlantsCreatePlantsAdditivesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_additives(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_additives_using_template(self, body: Optional[List['PlantsCreatePlantsAdditivesUsingTemplateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records additive usage for plants using predefined additive templates at a specified Facility.
        Permissions Required:
        - Manage Plants Additives
        POST /plants/v2/additives/usingtemplate
        Parameters:
        - body (List['PlantsCreatePlantsAdditivesUsingTemplateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_additives_using_template(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_plantings(self, body: Optional[List['PlantsCreatePlantsPlantingsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new plant batches at a specified Facility from existing plant data.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        POST /plants/v2/plantings
        Parameters:
        - body (List['PlantsCreatePlantsPlantingsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_plantings(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plants_waste(self, body: Optional[List['PlantsCreatePlantsWasteRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records waste events for plants at a Facility, including method, reason, and location details.
        Permissions Required:
        - Manage Plants Waste
        POST /plants/v2/waste
        Parameters:
        - body (List['PlantsCreatePlantsWasteRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.create_plants_waste(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_plants(self, body: Any = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Removes plants from a Facility’s inventory while recording the reason for their disposal.
        Permissions Required:
        - View Veg/Flower Plants
        - Destroy Veg/Flower Plants
        DELETE /plants/v2
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.delete_plants(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_plants_additives(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Additive']:
        """
        Retrieves additive records applied to plants at a specified Facility.
        Permissions Required:
        - View/Manage Plants Additives
        GET /plants/v2/additives
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_additives(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_additives_types(self, body: Any = None) -> Any:
        """
        Retrieves a list of all plant additive types defined within a Facility.
        GET /plants/v2/additives/types
        Parameters:
        """
        async def op():
            return await self.client.plants.get_plants_additives_types(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_flowering_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Plant']:
        """
        Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/flowering
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_flowering_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_growth_phases(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the list of growth phases supported by a specified Facility.
        GET /plants/v2/growthphases
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.get_plants_growth_phases(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Plant']:
        """
        Retrieves inactive plants at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_inactive_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_mother_inactive_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Mother']:
        """
        Retrieves inactive mother-phase plants at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET /plants/v2/mother/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_mother_inactive_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_mother_on_hold_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Mother']:
        """
        Retrieves mother-phase plants currently marked as on hold at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET /plants/v2/mother/onhold
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_mother_on_hold_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_mother_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Mother']:
        """
        Retrieves mother-phase plants at a specified Facility.
        Permissions Required:
        - View Mother Plants
        GET /plants/v2/mother
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_mother_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_on_hold_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Plant']:
        """
        Retrieves plants that are currently on hold at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/onhold
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_on_hold_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Plant':
        """
        Retrieves a Plant by Id.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.get_plants_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_by_label(self, label: str, body: Any = None, license_number: Optional[str] = None) -> List['Plant']:
        """
        Retrieves a Plant by label.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/{label}
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.get_plants_by_label(label, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_waste(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantsWaste']:
        """
        Retrieves a list of recorded plant waste events for a specific Facility.
        Permissions Required:
        - View Plants Waste
        GET /plants/v2/waste
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_waste(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_waste_methods(self, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['WasteMethod']:
        """
        Retrieves a list of all available plant waste methods for use within a Facility.
        GET /plants/v2/waste/methods/all
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_waste_methods(body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_waste_reasons(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['WasteReason']:
        """
        Retriveves available reasons for recording mature plant waste at a specified Facility.
        GET /plants/v2/waste/reasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_waste_reasons(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_vegetative_plants(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Plant']:
        """
        Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Veg/Flower Plants
        GET /plants/v2/vegetative
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_vegetative_plants(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_waste_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantsWaste']:
        """
        Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
        Permissions Required:
        - View Plants Waste
        GET /plants/v2/waste/{id}/plant
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_waste_by_id(id, body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plants_waste_package_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['WastePackage']:
        """
        Retrieves a list of package records linked to the specified plantWasteId for a given facility.
        Permissions Required:
        - View Plants Waste
        GET /plants/v2/waste/{id}/package
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plants.get_plants_waste_package_by_id(id, body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def update_adjust_plants(self, body: Optional[List['PlantsUpdateAdjustPlantsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Adjusts the recorded count of plants at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plants/v2/adjust
        Parameters:
        - body (List['PlantsUpdateAdjustPlantsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_adjust_plants(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_growth_phase(self, body: Optional[List['PlantsUpdateGrowthPhaseRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Changes the growth phases of plants within a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plants/v2/growthphase
        Parameters:
        - body (List['PlantsUpdateGrowthPhaseRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_growth_phase(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_harvest(self, body: Optional[List['PlantsUpdateHarvestRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        PUT /plants/v2/harvest
        Parameters:
        - body (List['PlantsUpdateHarvestRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_harvest(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_merge(self, body: Optional[List['PlantsUpdateMergeRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Merges multiple plant groups into a single group within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manicure/Harvest Veg/Flower Plants
        PUT /plants/v2/merge
        Parameters:
        - body (List['PlantsUpdateMergeRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_merge(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_location(self, body: Optional[List['PlantsUpdatePlantsLocationRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Moves plant batches to new locations within a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plants/v2/location
        Parameters:
        - body (List['PlantsUpdatePlantsLocationRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_location(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_strain(self, body: Optional[List['PlantsUpdatePlantsStrainRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the strain information for plants within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plants/v2/strain
        Parameters:
        - body (List['PlantsUpdatePlantsStrainRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_strain(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_tag(self, body: Optional[List['PlantsUpdatePlantsTagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Replaces existing plant tags with new tags for plants within a Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plants/v2/tag
        Parameters:
        - body (List['PlantsUpdatePlantsTagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_tag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plants_split(self, body: Optional[List['PlantsUpdateSplitRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Splits an existing plant group into multiple groups within a Facility.
        Permissions Required:
        - View Plant
        PUT /plants/v2/split
        Parameters:
        - body (List['PlantsUpdateSplitRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plants.update_plants_split(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

