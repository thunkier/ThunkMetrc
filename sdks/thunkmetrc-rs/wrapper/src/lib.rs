pub use thunkmetrc_client::MetrcClient;
use std::sync::Arc;
mod ratelimiter;
pub use ratelimiter::RateLimiterConfig;
pub use ratelimiter::MetrcRateLimiter;

pub mod models;
pub mod utils;
pub mod services;
use services::additives_templates_service::AdditivesTemplatesService;
use services::caregivers_status_service::CaregiversStatusService;
use services::employees_service::EmployeesService;
use services::facilities_service::FacilitiesService;
use services::harvests_service::HarvestsService;
use services::items_service::ItemsService;
use services::lab_tests_service::LabTestsService;
use services::locations_service::LocationsService;
use services::packages_service::PackagesService;
use services::patient_check_ins_service::PatientCheckInsService;
use services::patients_service::PatientsService;
use services::patients_status_service::PatientsStatusService;
use services::plant_batches_service::PlantBatchesService;
use services::plants_service::PlantsService;
use services::processing_job_service::ProcessingJobService;
use services::retail_id_service::RetailIdService;
use services::sales_service::SalesService;
use services::sandbox_service::SandboxService;
use services::strains_service::StrainsService;
use services::sublocations_service::SublocationsService;
use services::tags_service::TagsService;
use services::transfers_service::TransfersService;
use services::transporters_service::TransportersService;
use services::units_of_measure_service::UnitsOfMeasureService;
use services::waste_methods_service::WasteMethodsService;
use services::webhooks_service::WebhooksService;

pub struct MetrcWrapper {
    pub additives_templates: AdditivesTemplatesService,
    pub caregivers_status: CaregiversStatusService,
    pub employees: EmployeesService,
    pub facilities: FacilitiesService,
    pub harvests: HarvestsService,
    pub items: ItemsService,
    pub lab_tests: LabTestsService,
    pub locations: LocationsService,
    pub packages: PackagesService,
    pub patient_check_ins: PatientCheckInsService,
    pub patients: PatientsService,
    pub patients_status: PatientsStatusService,
    pub plant_batches: PlantBatchesService,
    pub plants: PlantsService,
    pub processing_job: ProcessingJobService,
    pub retail_id: RetailIdService,
    pub sales: SalesService,
    pub sandbox: SandboxService,
    pub strains: StrainsService,
    pub sublocations: SublocationsService,
    pub tags: TagsService,
    pub transfers: TransfersService,
    pub transporters: TransportersService,
    pub units_of_measure: UnitsOfMeasureService,
    pub waste_methods: WasteMethodsService,
    pub webhooks: WebhooksService,
}

impl MetrcWrapper {
    pub fn new(client: MetrcClient) -> Self {
        let rate_limiter = Arc::new(MetrcRateLimiter::new(Some(RateLimiterConfig::default())));
        MetrcWrapper {
            additives_templates: AdditivesTemplatesService::new(client.clone(), rate_limiter.clone()),
            caregivers_status: CaregiversStatusService::new(client.clone(), rate_limiter.clone()),
            employees: EmployeesService::new(client.clone(), rate_limiter.clone()),
            facilities: FacilitiesService::new(client.clone(), rate_limiter.clone()),
            harvests: HarvestsService::new(client.clone(), rate_limiter.clone()),
            items: ItemsService::new(client.clone(), rate_limiter.clone()),
            lab_tests: LabTestsService::new(client.clone(), rate_limiter.clone()),
            locations: LocationsService::new(client.clone(), rate_limiter.clone()),
            packages: PackagesService::new(client.clone(), rate_limiter.clone()),
            patient_check_ins: PatientCheckInsService::new(client.clone(), rate_limiter.clone()),
            patients: PatientsService::new(client.clone(), rate_limiter.clone()),
            patients_status: PatientsStatusService::new(client.clone(), rate_limiter.clone()),
            plant_batches: PlantBatchesService::new(client.clone(), rate_limiter.clone()),
            plants: PlantsService::new(client.clone(), rate_limiter.clone()),
            processing_job: ProcessingJobService::new(client.clone(), rate_limiter.clone()),
            retail_id: RetailIdService::new(client.clone(), rate_limiter.clone()),
            sales: SalesService::new(client.clone(), rate_limiter.clone()),
            sandbox: SandboxService::new(client.clone(), rate_limiter.clone()),
            strains: StrainsService::new(client.clone(), rate_limiter.clone()),
            sublocations: SublocationsService::new(client.clone(), rate_limiter.clone()),
            tags: TagsService::new(client.clone(), rate_limiter.clone()),
            transfers: TransfersService::new(client.clone(), rate_limiter.clone()),
            transporters: TransportersService::new(client.clone(), rate_limiter.clone()),
            units_of_measure: UnitsOfMeasureService::new(client.clone(), rate_limiter.clone()),
            waste_methods: WasteMethodsService::new(client.clone(), rate_limiter.clone()),
            webhooks: WebhooksService::new(client.clone(), rate_limiter.clone()),
        }
    }

    pub fn new_with_config(client: MetrcClient, config: RateLimiterConfig) -> Self {
        let rate_limiter = Arc::new(MetrcRateLimiter::new(Some(config)));
        MetrcWrapper {
            additives_templates: AdditivesTemplatesService::new(client.clone(), rate_limiter.clone()),
            caregivers_status: CaregiversStatusService::new(client.clone(), rate_limiter.clone()),
            employees: EmployeesService::new(client.clone(), rate_limiter.clone()),
            facilities: FacilitiesService::new(client.clone(), rate_limiter.clone()),
            harvests: HarvestsService::new(client.clone(), rate_limiter.clone()),
            items: ItemsService::new(client.clone(), rate_limiter.clone()),
            lab_tests: LabTestsService::new(client.clone(), rate_limiter.clone()),
            locations: LocationsService::new(client.clone(), rate_limiter.clone()),
            packages: PackagesService::new(client.clone(), rate_limiter.clone()),
            patient_check_ins: PatientCheckInsService::new(client.clone(), rate_limiter.clone()),
            patients: PatientsService::new(client.clone(), rate_limiter.clone()),
            patients_status: PatientsStatusService::new(client.clone(), rate_limiter.clone()),
            plant_batches: PlantBatchesService::new(client.clone(), rate_limiter.clone()),
            plants: PlantsService::new(client.clone(), rate_limiter.clone()),
            processing_job: ProcessingJobService::new(client.clone(), rate_limiter.clone()),
            retail_id: RetailIdService::new(client.clone(), rate_limiter.clone()),
            sales: SalesService::new(client.clone(), rate_limiter.clone()),
            sandbox: SandboxService::new(client.clone(), rate_limiter.clone()),
            strains: StrainsService::new(client.clone(), rate_limiter.clone()),
            sublocations: SublocationsService::new(client.clone(), rate_limiter.clone()),
            tags: TagsService::new(client.clone(), rate_limiter.clone()),
            transfers: TransfersService::new(client.clone(), rate_limiter.clone()),
            transporters: TransportersService::new(client.clone(), rate_limiter.clone()),
            units_of_measure: UnitsOfMeasureService::new(client.clone(), rate_limiter.clone()),
            waste_methods: WasteMethodsService::new(client.clone(), rate_limiter.clone()),
            webhooks: WebhooksService::new(client.clone(), rate_limiter.clone()),
        }
    }
    pub fn new_with_limiter(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        MetrcWrapper {
            additives_templates: AdditivesTemplatesService::new(client.clone(), rate_limiter.clone()),
            caregivers_status: CaregiversStatusService::new(client.clone(), rate_limiter.clone()),
            employees: EmployeesService::new(client.clone(), rate_limiter.clone()),
            facilities: FacilitiesService::new(client.clone(), rate_limiter.clone()),
            harvests: HarvestsService::new(client.clone(), rate_limiter.clone()),
            items: ItemsService::new(client.clone(), rate_limiter.clone()),
            lab_tests: LabTestsService::new(client.clone(), rate_limiter.clone()),
            locations: LocationsService::new(client.clone(), rate_limiter.clone()),
            packages: PackagesService::new(client.clone(), rate_limiter.clone()),
            patient_check_ins: PatientCheckInsService::new(client.clone(), rate_limiter.clone()),
            patients: PatientsService::new(client.clone(), rate_limiter.clone()),
            patients_status: PatientsStatusService::new(client.clone(), rate_limiter.clone()),
            plant_batches: PlantBatchesService::new(client.clone(), rate_limiter.clone()),
            plants: PlantsService::new(client.clone(), rate_limiter.clone()),
            processing_job: ProcessingJobService::new(client.clone(), rate_limiter.clone()),
            retail_id: RetailIdService::new(client.clone(), rate_limiter.clone()),
            sales: SalesService::new(client.clone(), rate_limiter.clone()),
            sandbox: SandboxService::new(client.clone(), rate_limiter.clone()),
            strains: StrainsService::new(client.clone(), rate_limiter.clone()),
            sublocations: SublocationsService::new(client.clone(), rate_limiter.clone()),
            tags: TagsService::new(client.clone(), rate_limiter.clone()),
            transfers: TransfersService::new(client.clone(), rate_limiter.clone()),
            transporters: TransportersService::new(client.clone(), rate_limiter.clone()),
            units_of_measure: UnitsOfMeasureService::new(client.clone(), rate_limiter.clone()),
            waste_methods: WasteMethodsService::new(client.clone(), rate_limiter.clone()),
            webhooks: WebhooksService::new(client.clone(), rate_limiter.clone()),
        }
    }
}

