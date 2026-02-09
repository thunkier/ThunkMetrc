from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PlantBatchesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_adjust_plant_batches(self, body: Optional[List['PlantBatchesCreateAdjustPlantBatchesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST /plantbatches/v2/adjust
        Parameters:
        - body (List['PlantBatchesCreateAdjustPlantBatchesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_adjust_plant_batches(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_growth_phase(self, body: Optional[List['PlantBatchesCreateGrowthPhaseRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the growth phase of plants at a specified Facility based on tracking information.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        POST /plantbatches/v2/growthphase
        Parameters:
        - body (List['PlantBatchesCreateGrowthPhaseRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_growth_phase(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_packages_from_mother_plant(self, body: Optional[List['PlantBatchesCreatePackagesFromMotherPlantRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages from mother plants at the specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /plantbatches/v2/packages/frommotherplant
        Parameters:
        - body (List['PlantBatchesCreatePackagesFromMotherPlantRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_packages_from_mother_plant(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_additives(self, body: Optional[List['PlantBatchesCreatePlantBatchesAdditivesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records Additive usage details for plant batches at a specific Facility.
        Permissions Required:
        - Manage Plants Additives
        POST /plantbatches/v2/additives
        Parameters:
        - body (List['PlantBatchesCreatePlantBatchesAdditivesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_additives(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_additives_using_template(self, body: Optional[List['PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records Additive usage for plant batches at a Facility using predefined additive templates.
        Permissions Required:
        - Manage Plants Additives
        POST /plantbatches/v2/additives/usingtemplate
        Parameters:
        - body (List['PlantBatchesCreatePlantBatchesAdditivesUsingTemplateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_additives_using_template(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_packages(self, body: Optional[List['PlantBatchesCreatePlantBatchesPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /plantbatches/v2/packages
        Parameters:
        - body (List['PlantBatchesCreatePlantBatchesPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_plantings(self, body: Optional[List['PlantBatchesCreatePlantBatchesPlantingsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new plantings for a Facility by generating plant batches based on provided planting details.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST /plantbatches/v2/plantings
        Parameters:
        - body (List['PlantBatchesCreatePlantBatchesPlantingsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_plantings(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_waste(self, body: Optional[List['PlantBatchesCreatePlantBatchesWasteRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records waste information for plant batches based on the submitted data for the specified Facility.
        Permissions Required:
        - Manage Plants Waste
        POST /plantbatches/v2/waste
        Parameters:
        - body (List['PlantBatchesCreatePlantBatchesWasteRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_waste(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_plant_batches_split(self, body: Optional[List['PlantBatchesCreateSplitRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Splits an existing Plant Batch into multiple groups at the specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants Inventory
        POST /plantbatches/v2/split
        Parameters:
        - body (List['PlantBatchesCreateSplitRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.create_plant_batches_split(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_plant_batches(self, body: Any = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Completes the destruction of plant batches based on the provided input data.
        Permissions Required:
        - View Immature Plants
        - Destroy Immature Plants
        DELETE /plantbatches/v2
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.delete_plant_batches(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_plant_batches(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantBatch']:
        """
        Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Immature Plants
        GET /plantbatches/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plant_batches.get_active_plant_batches(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_plant_batches(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantBatch']:
        """
        Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
        Permissions Required:
        - View Immature Plants
        GET /plantbatches/v2/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plant_batches.get_inactive_plant_batches(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plant_batches_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'PlantBatch':
        """
        Retrieves a Plant Batch by Id.
        Permissions Required:
        - View Immature Plants
        GET /plantbatches/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.get_plant_batches_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plant_batches_types(self, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantBatchesType']:
        """
        Retrieves a list of plant batch types.
        GET /plantbatches/v2/types
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plant_batches.get_plant_batches_types(body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plant_batches_waste(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PlantBatchesWaste']:
        """
        Retrieves waste details associated with plant batches at a specified Facility.
        Permissions Required:
        - View Plants Waste
        GET /plantbatches/v2/waste
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.plant_batches.get_plant_batches_waste(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_plant_batches_waste_reasons(self, body: Any = None, license_number: Optional[str] = None) -> List['PlantBatchesWasteReason']:
        """
        Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
        GET /plantbatches/v2/waste/reasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.get_plant_batches_waste_reasons(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_plant_batches_name(self, body: Optional[List['PlantBatchesUpdateNameRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Renames plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plantbatches/v2/name
        Parameters:
        - body (List['PlantBatchesUpdateNameRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.update_plant_batches_name(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plant_batches_location(self, body: Optional[List['PlantBatchesUpdatePlantBatchesLocationRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Moves one or more plant batches to new locations with in a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants
        PUT /plantbatches/v2/location
        Parameters:
        - body (List['PlantBatchesUpdatePlantBatchesLocationRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.update_plant_batches_location(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plant_batches_strain(self, body: Optional[List['PlantBatchesUpdatePlantBatchesStrainRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Changes the strain of plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plantbatches/v2/strain
        Parameters:
        - body (List['PlantBatchesUpdatePlantBatchesStrainRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.update_plant_batches_strain(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_plant_batches_tag(self, body: Optional[List['PlantBatchesUpdatePlantBatchesTagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Replaces tags for plant batches at a specified Facility.
        Permissions Required:
        - View Veg/Flower Plants
        - Manage Veg/Flower Plants Inventory
        PUT /plantbatches/v2/tag
        Parameters:
        - body (List['PlantBatchesUpdatePlantBatchesTagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.plant_batches.update_plant_batches_tag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

