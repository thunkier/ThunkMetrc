package io.github.thunkier.thunkmetrc.wrapper;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.services.*;

/**
 * Metrc Wrapper with rate limiting.
 * Organized by services matching the API groups.
 */
public class MetrcWrapper {

    private final MetrcRateLimiter rateLimiter;

    public MetrcRateLimiter getRateLimiter() {
        return rateLimiter;
    }
    private final AdditivesTemplatesService additivesTemplates;
    private final CaregiversStatusService caregiversStatus;
    private final EmployeesService employees;
    private final FacilitiesService facilities;
    private final HarvestsService harvests;
    private final ItemsService items;
    private final LabTestsService labTests;
    private final LocationsService locations;
    private final PackagesService packages;
    private final PatientCheckInsService patientCheckIns;
    private final PatientsService patients;
    private final PatientsStatusService patientsStatus;
    private final PlantBatchesService plantBatches;
    private final PlantsService plants;
    private final ProcessingJobService processingJob;
    private final RetailIdService retailId;
    private final SalesService sales;
    private final SandboxService sandbox;
    private final StrainsService strains;
    private final SublocationsService sublocations;
    private final TagsService tags;
    private final TransfersService transfers;
    private final TransportersService transporters;
    private final UnitsOfMeasureService unitsOfMeasure;
    private final WasteMethodsService wasteMethods;
    private final WebhooksService webhooks;

    /**
     * Constructor with default config.
     * @param client MetrcClient instance
     */
    public MetrcWrapper(MetrcClient client) {
        this(client, new MetrcRateLimiter.Config());
    }

    /**
     * Convenience constructor that creates both client and wrapper.
     * @param baseUrl Metrc API base URL
     * @param vendorKey Vendor API key
     * @param userKey User API key
     */
    public MetrcWrapper(String baseUrl, String vendorKey, String userKey) {
        this(new MetrcClient(baseUrl, vendorKey, userKey));
    }

    /**
     * Constructor with custom config.
     * @param client MetrcClient instance
     * @param config RateLimiter configuration
     */
    public MetrcWrapper(MetrcClient client, MetrcRateLimiter.Config config) {
        this(client, new MetrcRateLimiter(config));
    }

    /**
     * Constructor with shared rate limiter.
     * @param client MetrcClient instance
     * @param rateLimiter Shared MetrcRateLimiter instance
     */
    public MetrcWrapper(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.rateLimiter = rateLimiter;
        this.additivesTemplates = new AdditivesTemplatesService(client, rateLimiter);
        this.caregiversStatus = new CaregiversStatusService(client, rateLimiter);
        this.employees = new EmployeesService(client, rateLimiter);
        this.facilities = new FacilitiesService(client, rateLimiter);
        this.harvests = new HarvestsService(client, rateLimiter);
        this.items = new ItemsService(client, rateLimiter);
        this.labTests = new LabTestsService(client, rateLimiter);
        this.locations = new LocationsService(client, rateLimiter);
        this.packages = new PackagesService(client, rateLimiter);
        this.patientCheckIns = new PatientCheckInsService(client, rateLimiter);
        this.patients = new PatientsService(client, rateLimiter);
        this.patientsStatus = new PatientsStatusService(client, rateLimiter);
        this.plantBatches = new PlantBatchesService(client, rateLimiter);
        this.plants = new PlantsService(client, rateLimiter);
        this.processingJob = new ProcessingJobService(client, rateLimiter);
        this.retailId = new RetailIdService(client, rateLimiter);
        this.sales = new SalesService(client, rateLimiter);
        this.sandbox = new SandboxService(client, rateLimiter);
        this.strains = new StrainsService(client, rateLimiter);
        this.sublocations = new SublocationsService(client, rateLimiter);
        this.tags = new TagsService(client, rateLimiter);
        this.transfers = new TransfersService(client, rateLimiter);
        this.transporters = new TransportersService(client, rateLimiter);
        this.unitsOfMeasure = new UnitsOfMeasureService(client, rateLimiter);
        this.wasteMethods = new WasteMethodsService(client, rateLimiter);
        this.webhooks = new WebhooksService(client, rateLimiter);
    }
    public AdditivesTemplatesService additivesTemplates() {
        return this.additivesTemplates;
    }
    public CaregiversStatusService caregiversStatus() {
        return this.caregiversStatus;
    }
    public EmployeesService employees() {
        return this.employees;
    }
    public FacilitiesService facilities() {
        return this.facilities;
    }
    public HarvestsService harvests() {
        return this.harvests;
    }
    public ItemsService items() {
        return this.items;
    }
    public LabTestsService labTests() {
        return this.labTests;
    }
    public LocationsService locations() {
        return this.locations;
    }
    public PackagesService packages() {
        return this.packages;
    }
    public PatientCheckInsService patientCheckIns() {
        return this.patientCheckIns;
    }
    public PatientsService patients() {
        return this.patients;
    }
    public PatientsStatusService patientsStatus() {
        return this.patientsStatus;
    }
    public PlantBatchesService plantBatches() {
        return this.plantBatches;
    }
    public PlantsService plants() {
        return this.plants;
    }
    public ProcessingJobService processingJob() {
        return this.processingJob;
    }
    public RetailIdService retailId() {
        return this.retailId;
    }
    public SalesService sales() {
        return this.sales;
    }
    public SandboxService sandbox() {
        return this.sandbox;
    }
    public StrainsService strains() {
        return this.strains;
    }
    public SublocationsService sublocations() {
        return this.sublocations;
    }
    public TagsService tags() {
        return this.tags;
    }
    public TransfersService transfers() {
        return this.transfers;
    }
    public TransportersService transporters() {
        return this.transporters;
    }
    public UnitsOfMeasureService unitsOfMeasure() {
        return this.unitsOfMeasure;
    }
    public WasteMethodsService wasteMethods() {
        return this.wasteMethods;
    }
    public WebhooksService webhooks() {
        return this.webhooks;
    }
}
