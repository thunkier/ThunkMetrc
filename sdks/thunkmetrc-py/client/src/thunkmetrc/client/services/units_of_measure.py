from typing import Any, Optional, Dict, List, Union
import urllib.parse

class UnitsOfMeasureClient:
    def __init__(self, client):
        self.client = client
    def get_active_units_of_measure(self) -> Any:
        """
        Retrieves all active units of measure.
        GET GetActiveUnitsOfMeasure
        Parameters:
        """
        query_params = {}
        path = f"/unitsofmeasure/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_units_of_measure(self) -> Any:
        """
        Retrieves all inactive units of measure.
        GET GetInactiveUnitsOfMeasure
        Parameters:
        """
        query_params = {}
        path = f"/unitsofmeasure/v2/inactive"
        return self.client._send("GET", path, body, query_params)

