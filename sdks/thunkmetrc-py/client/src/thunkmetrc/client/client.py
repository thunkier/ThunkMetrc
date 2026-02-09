import httpx
import logging
import base64
import urllib.parse
from typing import Any, Optional, Dict, List, Union
from typing_extensions import TypedDict 
from typing import TypedDict
from .services.additives_templates import AdditivesTemplatesClient
from .services.caregivers_status import CaregiversStatusClient
from .services.employees import EmployeesClient
from .services.facilities import FacilitiesClient
from .services.harvests import HarvestsClient
from .services.items import ItemsClient
from .services.lab_tests import LabTestsClient
from .services.locations import LocationsClient
from .services.packages import PackagesClient
from .services.patient_check_ins import PatientCheckInsClient
from .services.patients import PatientsClient
from .services.patients_status import PatientsStatusClient
from .services.plant_batches import PlantBatchesClient
from .services.plants import PlantsClient
from .services.processing_job import ProcessingJobClient
from .services.retail_id import RetailIdClient
from .services.sales import SalesClient
from .services.sandbox import SandboxClient
from .services.strains import StrainsClient
from .services.sublocations import SublocationsClient
from .services.tags import TagsClient
from .services.transfers import TransfersClient
from .services.transporters import TransportersClient
from .services.units_of_measure import UnitsOfMeasureClient
from .services.waste_methods import WasteMethodsClient
from .services.webhooks import WebhooksClient

class MetrcError(Exception):
    def __init__(self, status_code: int, message: str, validation_errors: List[str] = None):
        self.status_code = status_code
        self.message = message
        self.validation_errors = validation_errors or []
        super().__init__(f"API Error {status_code}: {message}")

class MetrcClient:
    def __init__(self, base_url: str, vendor_key: str, user_key: str, 
                 client: Optional[httpx.Client] = None, 
                 timeout: float = 100.0,
                 user_agent: str = "ThunkMetrc/0.1 Python",
                 logger: Optional[logging.Logger] = None):
        self.base_url = base_url.rstrip('/')
        self.vendor_key = vendor_key
        self.user_key = user_key
        self.logger = logger
        
        if client:
            self.client = client
        else:
            self.client = httpx.Client(timeout=timeout)
            
        self.client.auth = (vendor_key, user_key)
        self.client.headers.update({"User-Agent": user_agent})

        # Initialize Service Clients
        self.additives_templates = AdditivesTemplatesClient(self)
        self.caregivers_status = CaregiversStatusClient(self)
        self.employees = EmployeesClient(self)
        self.facilities = FacilitiesClient(self)
        self.harvests = HarvestsClient(self)
        self.items = ItemsClient(self)
        self.lab_tests = LabTestsClient(self)
        self.locations = LocationsClient(self)
        self.packages = PackagesClient(self)
        self.patient_check_ins = PatientCheckInsClient(self)
        self.patients = PatientsClient(self)
        self.patients_status = PatientsStatusClient(self)
        self.plant_batches = PlantBatchesClient(self)
        self.plants = PlantsClient(self)
        self.processing_job = ProcessingJobClient(self)
        self.retail_id = RetailIdClient(self)
        self.sales = SalesClient(self)
        self.sandbox = SandboxClient(self)
        self.strains = StrainsClient(self)
        self.sublocations = SublocationsClient(self)
        self.tags = TagsClient(self)
        self.transfers = TransfersClient(self)
        self.transporters = TransportersClient(self)
        self.units_of_measure = UnitsOfMeasureClient(self)
        self.waste_methods = WasteMethodsClient(self)
        self.webhooks = WebhooksClient(self)

    def _send(self, method: str, path: str, body: Any = None, query_params: Dict[str, Any] = None) -> Any:
        url = f"{self.base_url}{path}"
        
        if self.logger:
            self.logger.debug(f"Sending Request: {method} {url}")
            
        try:
            response = self.client.request(method, url, json=body, params=query_params)
            
            if response.status_code >= 400:
                if self.logger:
                    self.logger.warning(f"API Error Response: {response.status_code}")
                
                # Try parsing structured error
                try:
                    error_data = response.json()
                    # Check if it matches Metrc Error format
                    if isinstance(error_data, dict) and ('Message' in error_data or 'ValidationErrors' in error_data):
                        raise MetrcError(
                            status_code=response.status_code,
                            message=error_data.get('Message', 'Unknown Error'),
                            validation_errors=error_data.get('ValidationErrors', [])
                        )
                except ValueError:
                    pass # Not JSON
                except MetrcError:
                    raise # Rethrow valid MetrcError

                # Fallback
                response.raise_for_status()

            if response.status_code == 204:
                return None
            return response.json()
            
        except httpx.RequestError as e:
            if self.logger:
                self.logger.error(f"Request Failed: {e}")
            raise
