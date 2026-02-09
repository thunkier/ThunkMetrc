from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class SandboxService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_sandbox_integrator_setup(self, body: Any = None, user_key: Optional[str] = None) -> Any:
        """
        This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
        POST /sandbox/v2/integrator/setup
        Parameters:
        - body (Any): Request body
        - user_key (str, optional): Filter by userKey
        """
        async def op():
            return await self.client.sandbox.create_sandbox_integrator_setup(body=body,
                user_key=user_key,
            )
        return await self.limiter.execute(None, False, op)

