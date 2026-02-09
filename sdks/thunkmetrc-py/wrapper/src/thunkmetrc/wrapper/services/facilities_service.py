from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class FacilitiesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_facilities(self, body: Any = None) -> List['Facility']:
        """
        This endpoint provides a list of facilities for which the authenticated user has access.
        GET /facilities/v2
        Parameters:
        """
        async def op():
            return await self.client.facilities.get_facilities(body=body,
            )
        return await self.limiter.execute(None, True, op)

