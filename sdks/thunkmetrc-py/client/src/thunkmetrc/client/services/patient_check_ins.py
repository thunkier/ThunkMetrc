from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PatientCheckInsClient:
    def __init__(self, client):
        self.client = client
    def create_patient_check_ins(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Records patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        POST CreatePatientCheckIns
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patient-checkins/v2"
        return self.client._send("POST", path, body, query_params)
    def delete_patient_check_ins_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a Patient Check-In, identified by its Id, for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        DELETE DeletePatientCheckInsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patient-checkins/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_patient_check_ins_locations(self) -> Any:
        """
        Retrieves a list of Patient Check-In locations.
        GET GetLocations
        Parameters:
        """
        query_params = {}
        path = f"/patient-checkins/v2/locations"
        return self.client._send("GET", path, body, query_params)
    def get_patient_check_ins(self, checkin_date_end: Optional[str] = None, checkin_date_start: Optional[str] = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a list of patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        GET GetPatientCheckIns
        Parameters:
        - checkin_date_end (str, optional): Filter by checkinDateEnd
        - checkin_date_start (str, optional): Filter by checkinDateStart
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if checkin_date_end is not None:
            query_params["checkinDateEnd"] = checkin_date_end
        if checkin_date_start is not None:
            query_params["checkinDateStart"] = checkin_date_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patient-checkins/v2"
        return self.client._send("GET", path, body, query_params)
    def update_patient_check_ins(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        PUT UpdatePatientCheckIns
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patient-checkins/v2"
        return self.client._send("PUT", path, body, query_params)

