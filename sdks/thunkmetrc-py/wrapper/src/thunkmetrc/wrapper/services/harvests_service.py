from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class HarvestsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_harvests_packages(self, body: Optional[List['HarvestsCreateHarvestsPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages from harvested products for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /harvests/v2/packages
        Parameters:
        - body (List['HarvestsCreateHarvestsPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.create_harvests_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_harvests_waste(self, body: Optional[List['HarvestsCreateHarvestsWasteRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        POST /harvests/v2/waste
        Parameters:
        - body (List['HarvestsCreateHarvestsWasteRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.create_harvests_waste(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_harvests_packages_testing(self, body: Optional[List['HarvestsCreatePackagesTestingRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages for testing from harvested products for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /harvests/v2/packages/testing
        Parameters:
        - body (List['HarvestsCreatePackagesTestingRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.create_harvests_packages_testing(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_harvests_waste_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Discontinues a specific harvest waste record by Id for the specified Facility.
        Permissions Required:
        - View Harvests
        - Discontinue Harvest Waste
        DELETE /harvests/v2/waste/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.delete_harvests_waste_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def finish_harvests_harvests(self, body: Optional[List['HarvestsFinishHarvestsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Marks one or more harvests as finished for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT /harvests/v2/finish
        Parameters:
        - body (List['HarvestsFinishHarvestsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.finish_harvests_harvests(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_harvests(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Harvest']:
        """
        Retrieves a list of active harvests for a specified Facility.
        Permissions Required:
        - View Harvests
        GET /harvests/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.harvests.get_active_harvests(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_harvests_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Harvest':
        """
        Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
        Permissions Required:
        - View Harvests
        GET /harvests/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.get_harvests_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_harvests_waste(self, body: Any = None, harvest_id: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['HarvestsWaste']:
        """
        Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
        Permissions Required:
        - View Harvests
        GET /harvests/v2/waste
        Parameters:
        - harvest_id (str, optional): Filter by harvestId
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.harvests.get_harvests_waste(body=body,
                harvest_id=harvest_id,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_harvests(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Harvest']:
        """
        Retrieves a list of inactive harvests for a specified Facility.
        Permissions Required:
        - View Harvests
        GET /harvests/v2/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.harvests.get_inactive_harvests(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_on_hold_harvests(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Harvest']:
        """
        Retrieves a list of harvests on hold for a specified Facility.
        Permissions Required:
        - View Harvests
        GET /harvests/v2/onhold
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.harvests.get_on_hold_harvests(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_harvests_waste_types(self, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['WasteType']:
        """
        Retrieves a list of Waste types for harvests.
        GET /harvests/v2/waste/types
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.harvests.get_harvests_waste_types(body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def unfinish_harvests_harvests(self, body: Optional[List['HarvestsUnfinishHarvestsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Reopens one or more previously finished harvests for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT /harvests/v2/unfinish
        Parameters:
        - body (List['HarvestsUnfinishHarvestsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.unfinish_harvests_harvests(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_harvests_location(self, body: Optional[List['HarvestsUpdateHarvestsLocationRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the Location of Harvest for a specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        PUT /harvests/v2/location
        Parameters:
        - body (List['HarvestsUpdateHarvestsLocationRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.update_harvests_location(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_harvests_rename(self, body: Optional[List['HarvestsUpdateRenameRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Renames one or more harvests for the specified Facility.
        Permissions Required:
        - View Harvests
        - Manage Harvests
        PUT /harvests/v2/rename
        Parameters:
        - body (List['HarvestsUpdateRenameRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.update_harvests_rename(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_harvests_restore_harvested_plants(self, body: Optional[List['HarvestsUpdateRestoreHarvestedPlantsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Restores previously harvested plants to their original state for the specified Facility.
        Permissions Required:
        - View Harvests
        - Finish/Discontinue Harvests
        PUT /harvests/v2/restore/harvestedplants
        Parameters:
        - body (List['HarvestsUpdateRestoreHarvestedPlantsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.harvests.update_harvests_restore_harvested_plants(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

