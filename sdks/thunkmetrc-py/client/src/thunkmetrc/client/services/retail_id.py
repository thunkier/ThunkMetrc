from typing import Any, Optional, Dict, List, Union
import urllib.parse

class RetailIdClient:
    def __init__(self, client):
        self.client = client
    def create_retail_id_associate(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        POST CreateAssociate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/associate"
        return self.client._send("POST", path, body, query_params)
    def create_retail_id_generate(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        POST CreateGenerate
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/generate"
        return self.client._send("POST", path, body, query_params)
    def create_retail_id_merge(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Key Value Settings/Retail ID Merge Packages Enabled
        POST CreateMerge
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/merge"
        return self.client._send("POST", path, body, query_params)
    def create_retail_id_packages_info(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves Package information for given list of Package labels.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - Admin/Employees/Packages Page/Product Labels(Manage)
        POST CreatePackagesInfo
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/packages/info"
        return self.client._send("POST", path, body, query_params)
    def get_retail_id_allotment(self, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the available Retail Item ID quota for a facility.
        Permissions Required:
        - Download Product Labels
        - Manage Product Labels
        - Manage Tag Orders
        GET GetAllotment
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/allotment"
        return self.client._send("GET", path, body, query_params)
    def get_retail_id_receive_by_label(self, label: str, license_number: Optional[str] = None) -> Any:
        """
        Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Manage RetailId
        - WebApi Retail ID Read Write State (All or ReadOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        GET GetReceiveByLabel
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/receive/{urllib.parse.quote(label)}"
        return self.client._send("GET", path, body, query_params)
    def get_retail_id_receive_qr_by_short_code(self, short_code: str, license_number: Optional[str] = None) -> Any:
        """
        Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Manage RetailId
        - WebApi Retail ID Read Write State (All or ReadOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        GET GetReceiveQrByShortCode
        Parameters:
        - short_code (str): Path parameter shortCode
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/retailid/v2/receive/qr/{urllib.parse.quote(short_code)}"
        return self.client._send("GET", path, body, query_params)

