from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class StrainsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_strains(self, body: Optional[List['StrainsCreateStrainsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new strain records for a specified Facility.
        Permissions Required:
        - Manage Strains
        POST /strains/v2
        Parameters:
        - body (List['StrainsCreateStrainsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.strains.create_strains(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_strains_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives an existing strain record for a Facility
        Permissions Required:
        - Manage Strains
        DELETE /strains/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.strains.delete_strains_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_strains(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Strain']:
        """
        Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Strains
        GET /strains/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.strains.get_active_strains(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_strains(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Strain']:
        """
        Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
        Permissions Required:
        - Manage Strains
        GET /strains/v2/inactive
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.strains.get_inactive_strains(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_strains_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Strain':
        """
        Retrieves a Strain record by its Id, with an optional license number.
        Permissions Required:
        - Manage Strains
        GET /strains/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.strains.get_strains_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_strains(self, body: Optional[List['StrainsUpdateStrainsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing strain records for a specified Facility.
        Permissions Required:
        - Manage Strains
        PUT /strains/v2
        Parameters:
        - body (List['StrainsUpdateStrainsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.strains.update_strains(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

