package io.github.thunkier.thunkmetrc.wrapper

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.services.*

/**
 * Metrc Wrapper with rate limiting.
 * Organized by services matching the API groups.
 */
class MetrcWrapper(
    private val client: MetrcClient,
    val rateLimiter: MetrcRateLimiter
) {
    /**
     * Constructor with default config.
     * @param client MetrcClient instance
     */
    constructor(client: MetrcClient) : this(client, MetrcRateLimiter())

    /**
     * Constructor with custom config.
     * @param client MetrcClient instance
     * @param config RateLimiter configuration
     */
    constructor(client: MetrcClient, config: RateLimiterConfig) : this(client, MetrcRateLimiter(config))

    /**
     * Convenience constructor that creates both client and wrapper.
     * @param baseUrl Metrc API base URL
     * @param vendorKey Vendor API key
     * @param userKey User API key
     */
    constructor(baseUrl: String, vendorKey: String, userKey: String) : this(MetrcClient(baseUrl, vendorKey, userKey))
    val additivesTemplates = AdditivesTemplatesService(client, rateLimiter)
    val caregiversStatus = CaregiversStatusService(client, rateLimiter)
    val employees = EmployeesService(client, rateLimiter)
    val facilities = FacilitiesService(client, rateLimiter)
    val harvests = HarvestsService(client, rateLimiter)
    val items = ItemsService(client, rateLimiter)
    val labTests = LabTestsService(client, rateLimiter)
    val locations = LocationsService(client, rateLimiter)
    val packages = PackagesService(client, rateLimiter)
    val patientCheckIns = PatientCheckInsService(client, rateLimiter)
    val patients = PatientsService(client, rateLimiter)
    val patientsStatus = PatientsStatusService(client, rateLimiter)
    val plantBatches = PlantBatchesService(client, rateLimiter)
    val plants = PlantsService(client, rateLimiter)
    val processingJob = ProcessingJobService(client, rateLimiter)
    val retailId = RetailIdService(client, rateLimiter)
    val sales = SalesService(client, rateLimiter)
    val sandbox = SandboxService(client, rateLimiter)
    val strains = StrainsService(client, rateLimiter)
    val sublocations = SublocationsService(client, rateLimiter)
    val tags = TagsService(client, rateLimiter)
    val transfers = TransfersService(client, rateLimiter)
    val transporters = TransportersService(client, rateLimiter)
    val unitsOfMeasure = UnitsOfMeasureService(client, rateLimiter)
    val wasteMethods = WasteMethodsService(client, rateLimiter)
    val webhooks = WebhooksService(client, rateLimiter)
}
