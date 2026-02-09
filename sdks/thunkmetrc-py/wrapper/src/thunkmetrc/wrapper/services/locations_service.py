from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class LocationsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_locations(self, body: Optional[List['LocationsCreateLocationsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        POST /locations/v2
        Parameters:
        - body (List['LocationsCreateLocationsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.locations.create_locations(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_locations_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a specified Location, identified by its Id, for a Facility.
        Permissions Required:
        - Manage Locations
        DELETE /locations/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.locations.delete_locations_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_locations(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['LocationsLocation']:
        """
        Retrieves a list of active locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET /locations/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.locations.get_active_locations(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_locations(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['LocationsLocation']:
        """
        Retrieves a list of inactive locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET /locations/v2/inactive
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.locations.get_inactive_locations(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_locations_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'LocationsLocation':
        """
        Retrieves a Location by its Id.
        Permissions Required:
        - Manage Locations
        GET /locations/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.locations.get_locations_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_locations_types(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['LocationsType']:
        """
        Retrieves a list of active location types for a specified Facility.
        Permissions Required:
        - Manage Locations
        GET /locations/v2/types
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.locations.get_locations_types(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def update_locations(self, body: Optional[List['LocationsUpdateLocationsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing locations for a specified Facility.
        Permissions Required:
        - Manage Locations
        PUT /locations/v2
        Parameters:
        - body (List['LocationsUpdateLocationsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.locations.update_locations(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

