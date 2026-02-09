package io.github.thunkier.thunkmetrc.client

import okhttp3.*
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.RequestBody.Companion.toRequestBody
import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper
import org.slf4j.LoggerFactory

/**
 * Metrc API Client.
 * Provides access to all Metrc API endpoints via service-specific clients.
 */
class MetrcClient(
    private val baseUrl: String,
    vendorKey: String,
    userKey: String,
    client: OkHttpClient? = null
) {
    private val client = client ?: OkHttpClient.Builder()
        .readTimeout(100, java.util.concurrent.TimeUnit.SECONDS)
        .build()
    private val mapper = jacksonObjectMapper()
    private val auth = Credentials.basic(vendorKey, userKey)
    private val logger = LoggerFactory.getLogger(MetrcClient::class.java)
    val additivesTemplates = AdditivesTemplatesClient(this)
    val caregiversStatus = CaregiversStatusClient(this)
    val employees = EmployeesClient(this)
    val facilities = FacilitiesClient(this)
    val harvests = HarvestsClient(this)
    val items = ItemsClient(this)
    val labTests = LabTestsClient(this)
    val locations = LocationsClient(this)
    val packages = PackagesClient(this)
    val patientCheckIns = PatientCheckInsClient(this)
    val patients = PatientsClient(this)
    val patientsStatus = PatientsStatusClient(this)
    val plantBatches = PlantBatchesClient(this)
    val plants = PlantsClient(this)
    val processingJob = ProcessingJobClient(this)
    val retailId = RetailIdClient(this)
    val sales = SalesClient(this)
    val sandbox = SandboxClient(this)
    val strains = StrainsClient(this)
    val sublocations = SublocationsClient(this)
    val tags = TagsClient(this)
    val transfers = TransfersClient(this)
    val transporters = TransportersClient(this)
    val unitsOfMeasure = UnitsOfMeasureClient(this)
    val wasteMethods = WasteMethodsClient(this)
    val webhooks = WebhooksClient(this)

    internal inline fun <reified T> send(method: String, path: String, body: Any? = null): T? {
        val url = "${baseUrl.trimEnd('/')}$path"
        logger.debug("Request: {} {}", method, url)

        val builder = Request.Builder().url(url).header("Authorization", auth)

        when {
            body != null && (method == "POST" || method == "PUT" || method == "PATCH") -> {
                val json = mapper.writeValueAsString(body)
                builder.method(method, json.toRequestBody("application/json".toMediaType()))
            }
            method == "POST" || method == "PUT" || method == "PATCH" -> {
                builder.method(method, "".toRequestBody(null))
            }
            else -> {
                builder.method(method, null)
            }
        }

        client.newCall(builder.build()).execute().use { response ->
            if (!response.isSuccessful) {
                val code = response.code
                logger.warn("API Error: {}", code)

                var message = "Unexpected code $code"
                var validationErrors: List<String> = emptyList()

                val errorBody = response.body?.string().orEmpty()
                if (errorBody.isNotEmpty()) {
                    try {
                        val node = mapper.readTree(errorBody)
                        if (node.has("Message")) {
                            message = node.get("Message").asText()
                        }
                        if (node.has("ValidationErrors") && node.get("ValidationErrors").isArray) {
                            validationErrors = node.get("ValidationErrors").map { it.asText() }
                        }
                    } catch (_: Exception) {}
                }

                throw MetrcException(code, message, validationErrors)
            }

            if (response.code == 204) return null
            return mapper.readValue(response.body!!.string(), T::class.java)
        }
    }
}

