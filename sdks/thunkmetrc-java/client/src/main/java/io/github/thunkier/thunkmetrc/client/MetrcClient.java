package io.github.thunkier.thunkmetrc.client;

import okhttp3.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * Metrc API Client.
 * Provides access to all Metrc API endpoints via service-specific clients.
 */
public class MetrcClient {
    private static final Logger logger = LoggerFactory.getLogger(MetrcClient.class);
    private final OkHttpClient client;
    private final ObjectMapper mapper;
    private final String baseUrl;
    private final String apiKey;
    public final AdditivesTemplatesClient additivesTemplates;
    public final CaregiversStatusClient caregiversStatus;
    public final EmployeesClient employees;
    public final FacilitiesClient facilities;
    public final HarvestsClient harvests;
    public final ItemsClient items;
    public final LabTestsClient labTests;
    public final LocationsClient locations;
    public final PackagesClient packages;
    public final PatientCheckInsClient patientCheckIns;
    public final PatientsClient patients;
    public final PatientsStatusClient patientsStatus;
    public final PlantBatchesClient plantBatches;
    public final PlantsClient plants;
    public final ProcessingJobClient processingJob;
    public final RetailIdClient retailId;
    public final SalesClient sales;
    public final SandboxClient sandbox;
    public final StrainsClient strains;
    public final SublocationsClient sublocations;
    public final TagsClient tags;
    public final TransfersClient transfers;
    public final TransportersClient transporters;
    public final UnitsOfMeasureClient unitsOfMeasure;
    public final WasteMethodsClient wasteMethods;
    public final WebhooksClient webhooks;

    public MetrcClient(String baseUrl, String vendorKey, String userKey) {
        this(baseUrl, vendorKey, userKey, new OkHttpClient.Builder()
            .readTimeout(100, java.util.concurrent.TimeUnit.SECONDS)
            .build());
    }

    public MetrcClient(String baseUrl, String vendorKey, String userKey, OkHttpClient client) {
        this.baseUrl = baseUrl.replaceAll("/$", "");
        this.apiKey = Credentials.basic(vendorKey, userKey);
        this.client = client;
        this.mapper = new ObjectMapper();
        this.additivesTemplates = new AdditivesTemplatesClient(this);
        this.caregiversStatus = new CaregiversStatusClient(this);
        this.employees = new EmployeesClient(this);
        this.facilities = new FacilitiesClient(this);
        this.harvests = new HarvestsClient(this);
        this.items = new ItemsClient(this);
        this.labTests = new LabTestsClient(this);
        this.locations = new LocationsClient(this);
        this.packages = new PackagesClient(this);
        this.patientCheckIns = new PatientCheckInsClient(this);
        this.patients = new PatientsClient(this);
        this.patientsStatus = new PatientsStatusClient(this);
        this.plantBatches = new PlantBatchesClient(this);
        this.plants = new PlantsClient(this);
        this.processingJob = new ProcessingJobClient(this);
        this.retailId = new RetailIdClient(this);
        this.sales = new SalesClient(this);
        this.sandbox = new SandboxClient(this);
        this.strains = new StrainsClient(this);
        this.sublocations = new SublocationsClient(this);
        this.tags = new TagsClient(this);
        this.transfers = new TransfersClient(this);
        this.transporters = new TransportersClient(this);
        this.unitsOfMeasure = new UnitsOfMeasureClient(this);
        this.wasteMethods = new WasteMethodsClient(this);
        this.webhooks = new WebhooksClient(this);
    }

    public OkHttpClient getHttpClient() {
        return client;
    }

    public ObjectMapper getMapper() {
        return mapper;
    }

    protected Object send(String method, String path, Object body) throws IOException {
        String url = this.baseUrl + path;
        logger.debug("Request: {} {}", method, url);

        Request.Builder builder = new Request.Builder()
            .url(url)
            .header("Authorization", this.apiKey);

        if (body != null && (method.equals("POST") || method.equals("PUT") || method.equals("PATCH"))) {
            String json = this.mapper.writeValueAsString(body);
            builder.method(method, RequestBody.create(json, MediaType.parse("application/json")));
        } else if (method.equals("POST") || method.equals("PUT") || method.equals("PATCH")) {
            builder.method(method, RequestBody.create("", null));
        } else {
            builder.method(method, null);
        }

        try (Response response = client.newCall(builder.build()).execute()) {
            if (!response.isSuccessful()) {
                int code = response.code();
                logger.warn("API Error: {}", code);

                String message = "Unexpected code " + code;
                List<String> validationErrors = null;

                String errorBody = response.body() != null ? response.body().string() : "";
                if (!errorBody.isEmpty()) {
                    try {
                        JsonNode node = mapper.readTree(errorBody);
                        if (node.has("Message")) {
                            message = node.get("Message").asText();
                        }
                        if (node.has("ValidationErrors") && node.get("ValidationErrors").isArray()) {
                            validationErrors = new ArrayList<>();
                            for (JsonNode err : node.get("ValidationErrors")) {
                                validationErrors.add(err.asText());
                            }
                        }
                    } catch (Exception ignored) {}
                }

                throw new MetrcException(code, message, validationErrors);
            }

            if (response.code() == 204) {
                return null;
            }

            return this.mapper.readValue(response.body().string(), Object.class);
        }
    }
}

