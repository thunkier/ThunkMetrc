from typing import Any, Optional, Dict, List, Union
import urllib.parse

class LabTestsClient:
    def __init__(self, client):
        self.client = client
    def create_lab_tests_record(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        POST CreateRecord
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/labtests/v2/record"
        return self.client._send("POST", path, body, query_params)
    def get_lab_tests_batches(self, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of Lab Test batches.
        GET GetBatches
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/labtests/v2/batches"
        return self.client._send("GET", path, body, query_params)
    def get_lab_tests_lab_test_document_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a specific Lab Test result document by its Id for a given Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        GET GetLabTestDocumentById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/labtests/v2/labtestdocument/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_lab_tests_types(self, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of Lab Test types.
        GET GetLabTestsTypes
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/labtests/v2/types"
        return self.client._send("GET", path, body, query_params)
    def get_lab_tests_results(self, license_number: Optional[str] = None, package_id: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves Lab Test results for a specified Package.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        GET GetResults
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - package_id (str, optional): Filter by packageId
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if package_id is not None:
            query_params["packageId"] = package_id
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/labtests/v2/results"
        return self.client._send("GET", path, body, query_params)
    def get_lab_tests_states(self) -> Any:
        """
        Returns a list of all lab testing states.
        GET GetStates
        Parameters:
        """
        query_params = {}
        path = f"/labtests/v2/states"
        return self.client._send("GET", path, body, query_params)
    def update_lab_tests_lab_test_document(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates one or more documents for previously submitted lab tests.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateLabTestDocument
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/labtests/v2/labtestdocument"
        return self.client._send("PUT", path, body, query_params)
    def update_lab_tests_results_release(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Releases Lab Test results for one or more packages.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT UpdateResultsRelease
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/labtests/v2/results/release"
        return self.client._send("PUT", path, body, query_params)

