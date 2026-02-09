from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class SublocationsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_sublocations(self, body: Optional[List['SublocationsCreateSublocationsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new sublocation records for a Facility.
        Permissions Required:
        - Manage Locations
        POST /sublocations/v2
        Parameters:
        - body (List['SublocationsCreateSublocationsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sublocations.create_sublocations(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_sublocations_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives an existing Sublocation record for a Facility.
        Permissions Required:
        - Manage Locations
        DELETE /sublocations/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sublocations.delete_sublocations_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_sublocations(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Sublocation']:
        """
        Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Locations
        GET /sublocations/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sublocations.get_active_sublocations(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_sublocations(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Sublocation']:
        """
        Retrieves a list of inactive sublocations for the specified Facility.
        Permissions Required:
        - Manage Locations
        GET /sublocations/v2/inactive
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.sublocations.get_inactive_sublocations(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_sublocations_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Sublocation':
        """
        Retrieves a Sublocation by its Id, with an optional license number.
        Permissions Required:
        - Manage Locations
        GET /sublocations/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sublocations.get_sublocations_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_sublocations(self, body: Optional[List['SublocationsUpdateSublocationsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing sublocation records for a specified Facility.
        Permissions Required:
        - Manage Locations
        PUT /sublocations/v2
        Parameters:
        - body (List['SublocationsUpdateSublocationsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.sublocations.update_sublocations(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

