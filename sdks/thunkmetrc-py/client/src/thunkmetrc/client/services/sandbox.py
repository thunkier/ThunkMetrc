from typing import Any, Optional, Dict, List, Union
import urllib.parse

class SandboxClient:
    def __init__(self, client):
        self.client = client
    def create_sandbox_integrator_setup(self, body: Any = None, user_key: Optional[str] = None) -> Any:
        """
        This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
        POST CreateIntegratorSetup
        Parameters:
        - body (Any): Request body
        - user_key (str, optional): Filter by userKey
        """
        query_params = {}
        if user_key is not None:
            query_params["userKey"] = user_key
        path = f"/sandbox/v2/integrator/setup"
        return self.client._send("POST", path, body, query_params)

