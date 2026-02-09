from typing import Any, Optional, Dict, List, Union
import urllib.parse

class WasteMethodsClient:
    def __init__(self, client):
        self.client = client
    def get_waste_methods_waste_methods(self) -> Any:
        """
        Retrieves all available waste methods.
        GET GetWasteMethodsWasteMethods
        Parameters:
        """
        query_params = {}
        path = f"/wastemethods/v2"
        return self.client._send("GET", path, body, query_params)

