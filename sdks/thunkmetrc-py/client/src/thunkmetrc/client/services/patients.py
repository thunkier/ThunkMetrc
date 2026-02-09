from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PatientsClient:
    def __init__(self, client):
        self.client = client
    def create_patients(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Adds new patients to a specified Facility.
        Permissions Required:
        - Manage Patients
        POST CreatePatients
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patients/v2"
        return self.client._send("POST", path, body, query_params)
    def delete_patients_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Removes a Patient, identified by an Id, from a specified Facility.
        Permissions Required:
        - Manage Patients
        DELETE DeletePatientsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patients/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_patients(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active patients for a specified Facility.
        Permissions Required:
        - Manage Patients
        GET GetActivePatients
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
        path = f"/patients/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_patients_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a Patient by Id.
        Permissions Required:
        - Manage Patients
        GET GetPatientsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patients/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def update_patients(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates Patient information for a specified Facility.
        Permissions Required:
        - Manage Patients
        PUT UpdatePatients
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patients/v2"
        return self.client._send("PUT", path, body, query_params)

