from typing import Optional
from thunkmetrc.client import MetrcClient
from .ratelimiter import MetrcRateLimiter, RateLimiterConfig
from .services import *

class MetrcWrapper:
    def __init__(self, client: MetrcClient, config: Optional[RateLimiterConfig] = None):
        self.client = client
        self.limiter = MetrcRateLimiter(config)
        self.additives_templates = AdditivesTemplatesService(self.client, self.limiter)
        self.caregivers_status = CaregiversStatusService(self.client, self.limiter)
        self.employees = EmployeesService(self.client, self.limiter)
        self.facilities = FacilitiesService(self.client, self.limiter)
        self.harvests = HarvestsService(self.client, self.limiter)
        self.items = ItemsService(self.client, self.limiter)
        self.lab_tests = LabTestsService(self.client, self.limiter)
        self.locations = LocationsService(self.client, self.limiter)
        self.packages = PackagesService(self.client, self.limiter)
        self.patient_check_ins = PatientCheckInsService(self.client, self.limiter)
        self.patients = PatientsService(self.client, self.limiter)
        self.patients_status = PatientsStatusService(self.client, self.limiter)
        self.plant_batches = PlantBatchesService(self.client, self.limiter)
        self.plants = PlantsService(self.client, self.limiter)
        self.processing_job = ProcessingJobService(self.client, self.limiter)
        self.retail_id = RetailIdService(self.client, self.limiter)
        self.sales = SalesService(self.client, self.limiter)
        self.sandbox = SandboxService(self.client, self.limiter)
        self.strains = StrainsService(self.client, self.limiter)
        self.sublocations = SublocationsService(self.client, self.limiter)
        self.tags = TagsService(self.client, self.limiter)
        self.transfers = TransfersService(self.client, self.limiter)
        self.transporters = TransportersService(self.client, self.limiter)
        self.units_of_measure = UnitsOfMeasureService(self.client, self.limiter)
        self.waste_methods = WasteMethodsService(self.client, self.limiter)
        self.webhooks = WebhooksService(self.client, self.limiter)
    @property
    def AdditivesTemplates(self) -> AdditivesTemplatesService:
        return self.additives_templates
    @property
    def CaregiversStatus(self) -> CaregiversStatusService:
        return self.caregivers_status
    @property
    def Employees(self) -> EmployeesService:
        return self.employees
    @property
    def Facilities(self) -> FacilitiesService:
        return self.facilities
    @property
    def Harvests(self) -> HarvestsService:
        return self.harvests
    @property
    def Items(self) -> ItemsService:
        return self.items
    @property
    def LabTests(self) -> LabTestsService:
        return self.lab_tests
    @property
    def Locations(self) -> LocationsService:
        return self.locations
    @property
    def Packages(self) -> PackagesService:
        return self.packages
    @property
    def PatientCheckIns(self) -> PatientCheckInsService:
        return self.patient_check_ins
    @property
    def Patients(self) -> PatientsService:
        return self.patients
    @property
    def PatientsStatus(self) -> PatientsStatusService:
        return self.patients_status
    @property
    def PlantBatches(self) -> PlantBatchesService:
        return self.plant_batches
    @property
    def Plants(self) -> PlantsService:
        return self.plants
    @property
    def ProcessingJob(self) -> ProcessingJobService:
        return self.processing_job
    @property
    def RetailId(self) -> RetailIdService:
        return self.retail_id
    @property
    def Sales(self) -> SalesService:
        return self.sales
    @property
    def Sandbox(self) -> SandboxService:
        return self.sandbox
    @property
    def Strains(self) -> StrainsService:
        return self.strains
    @property
    def Sublocations(self) -> SublocationsService:
        return self.sublocations
    @property
    def Tags(self) -> TagsService:
        return self.tags
    @property
    def Transfers(self) -> TransfersService:
        return self.transfers
    @property
    def Transporters(self) -> TransportersService:
        return self.transporters
    @property
    def UnitsOfMeasure(self) -> UnitsOfMeasureService:
        return self.units_of_measure
    @property
    def WasteMethods(self) -> WasteMethodsService:
        return self.waste_methods
    @property
    def Webhooks(self) -> WebhooksService:
        return self.webhooks
