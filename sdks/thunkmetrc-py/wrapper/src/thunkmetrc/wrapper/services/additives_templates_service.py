from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class AdditivesTemplatesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_additives_templates(self, body: Optional[List['AdditivesTemplatesCreateAdditivesTemplatesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        POST /additivestemplates/v2
        Parameters:
        - body (List['AdditivesTemplatesCreateAdditivesTemplatesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.additives_templates.create_additives_templates(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_additives_templates(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['AdditivesTemplate']:
        """
        Retrieves a list of active additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        GET /additivestemplates/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.additives_templates.get_active_additives_templates(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_additives_templates_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'AdditivesTemplate':
        """
        Retrieves an Additive Template by its Id.
        Permissions Required:
        - Manage Additives
        GET /additivestemplates/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.additives_templates.get_additives_templates_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_additives_templates(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['AdditivesTemplate']:
        """
        Retrieves a list of inactive additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        GET /additivestemplates/v2/inactive
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.additives_templates.get_inactive_additives_templates(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def update_additives_templates(self, body: Optional[List['AdditivesTemplatesUpdateAdditivesTemplatesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing additive templates for a specified Facility.
        Permissions Required:
        - Manage Additives
        PUT /additivestemplates/v2
        Parameters:
        - body (List['AdditivesTemplatesUpdateAdditivesTemplatesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.additives_templates.update_additives_templates(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

