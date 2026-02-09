from typing import Any, Optional, Dict, List, Union
import urllib.parse

class CaregiversStatusClient:
    def __init__(self, client):
        self.client = client
    def get_caregivers_status_by_caregiver_license_number(self, caregiver_license_number: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        Permissions Required:
        - Lookup Caregivers
        GET GetCaregiversStatusByCaregiverLicenseNumber
        Parameters:
        - caregiver_license_number (str): Path parameter caregiverLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/caregivers/v2/status/{urllib.parse.quote(caregiver_license_number)}"
        return self.client._send("GET", path, body, query_params)

