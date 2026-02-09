from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class UnitsOfMeasureService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_active_units_of_measure(self, body: Any = None) -> PaginatedResponse['UnitOfMeasure']:
        """
        Retrieves all active units of measure.
        GET /unitsofmeasure/v2/active
        Parameters:
        """
        async def op():
            return await self.client.units_of_measure.get_active_units_of_measure(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_units_of_measure(self, body: Any = None) -> PaginatedResponse['UnitOfMeasure']:
        """
        Retrieves all inactive units of measure.
        GET /unitsofmeasure/v2/inactive
        Parameters:
        """
        async def op():
            return await self.client.units_of_measure.get_inactive_units_of_measure(body=body,
            )
        return await self.limiter.execute(None, True, op)

