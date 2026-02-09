from typing import Any, Optional, Dict, List, Union
import urllib.parse

class FacilitiesClient:
    def __init__(self, client):
        self.client = client
    def get_facilities(self) -> Any:
        """
        This endpoint provides a list of facilities for which the authenticated user has access.
        GET GetFacilities
        Parameters:
        """
        query_params = {}
        path = f"/facilities/v2"
        return self.client._send("GET", path, body, query_params)

