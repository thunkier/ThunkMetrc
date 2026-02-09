from typing import Any, Optional, Dict, List, Union
import urllib.parse

class PatientsStatusClient:
    def __init__(self, client):
        self.client = client
    def get_patients_statuses_by_patient_license_number(self, patient_license_number: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        Permissions Required:
        - Lookup Patients
        GET GetPatientsStatusesByPatientLicenseNumber
        Parameters:
        - patient_license_number (str): Path parameter patientLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/patients/v2/statuses/{urllib.parse.quote(patient_license_number)}"
        return self.client._send("GET", path, body, query_params)

