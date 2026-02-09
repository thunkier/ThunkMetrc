from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class WasteMethodsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_waste_methods_waste_methods(self, body: Any = None) -> List['WasteMethod']:
        """
        Retrieves all available waste methods.
        GET /wastemethods/v2
        Parameters:
        """
        async def op():
            return await self.client.waste_methods.get_waste_methods_waste_methods(body=body,
            )
        return await self.limiter.execute(None, True, op)

