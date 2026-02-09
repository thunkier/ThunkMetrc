from typing import Any, Optional, Dict, List, Union
import urllib.parse

class EmployeesClient:
    def __init__(self, client):
        self.client = client
    def get_employees(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of employees for a specified Facility.
        Permissions Required:
        - Manage Employees
        - View Employees
        GET GetEmployees
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
        path = f"/employees/v2"
        return self.client._send("GET", path, body, query_params)
    def get_employees_permissions(self, employee_license_number: Optional[str] = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        Permissions Required:
        - Manage Employees
        GET GetPermissions
        Parameters:
        - employee_license_number (str, optional): Filter by employeeLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if employee_license_number is not None:
            query_params["employeeLicenseNumber"] = employee_license_number
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/employees/v2/permissions"
        return self.client._send("GET", path, body, query_params)

