from typing import Any, Optional, Dict, List, Union
import urllib.parse

class TransportersClient:
    def __init__(self, client):
        self.client = client
    def create_transporters_drivers(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new driver records for a Facility.
        Permissions Required:
        - Manage Transporters
        POST CreateDrivers
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/drivers"
        return self.client._send("POST", path, body, query_params)
    def create_transporters_vehicles(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new vehicle records for a Facility.
        Permissions Required:
        - Manage Transporters
        POST CreateVehicles
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/vehicles"
        return self.client._send("POST", path, body, query_params)
    def delete_transporters_driver_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        Permissions Required:
        - Manage Transporters
        DELETE DeleteDriverById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/drivers/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_transporters_vehicle_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        Permissions Required:
        - Manage Transporters
        DELETE DeleteVehicleById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/vehicles/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_transporters_driver_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        Permissions Required:
        - Transporters
        GET GetDriverById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/drivers/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_transporters_drivers(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of drivers for a Facility.
        Permissions Required:
        - Transporters
        GET GetDrivers
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transporters/v2/drivers"
        return self.client._send("GET", path, body, query_params)
    def get_transporters_vehicle_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        Permissions Required:
        - Transporters
        GET GetVehicleById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/vehicles/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_transporters_vehicles(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of vehicles for a Facility.
        Permissions Required:
        - Transporters
        GET GetVehicles
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/transporters/v2/vehicles"
        return self.client._send("GET", path, body, query_params)
    def update_transporters_drivers(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing driver records for a Facility.
        Permissions Required:
        - Manage Transporters
        PUT UpdateDrivers
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/drivers"
        return self.client._send("PUT", path, body, query_params)
    def update_transporters_vehicles(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing vehicle records for a facility.
        Permissions Required:
        - Manage Transporters
        PUT UpdateVehicles
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/transporters/v2/vehicles"
        return self.client._send("PUT", path, body, query_params)

