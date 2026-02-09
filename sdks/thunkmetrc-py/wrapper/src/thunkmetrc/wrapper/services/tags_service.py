from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class TagsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_tags_available_package(self, body: Any = None, license_number: Optional[str] = None) -> List['Tag']:
        """
        Returns a list of available package tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        GET /tags/v2/package/available
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.tags.get_tags_available_package(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_tags_available_plant(self, body: Any = None, license_number: Optional[str] = None) -> List['Tag']:
        """
        Returns a list of available plant tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        GET /tags/v2/plant/available
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.tags.get_tags_available_plant(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_staged_tags(self, body: Any = None, license_number: Optional[str] = None) -> List['Staged']:
        """
        Returns a list of staged tags. NOTE: This is a premium endpoint.
        Permissions Required:
        - Manage Tags
        - RetailId.AllowPackageStaging Key Value enabled
        GET /tags/v2/staged
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.tags.get_staged_tags(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

