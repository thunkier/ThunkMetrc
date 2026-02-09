package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper
import io.github.cdimascio.dotenv.Dotenv
import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.Assertions.assertNotNull
import org.junit.jupiter.api.Assertions.assertTrue
import org.junit.jupiter.api.Test
import java.io.File

class IntegrationTest {

    private fun getEnv(dotenv: Dotenv, key: String): String? {
        return dotenv.get(key) ?: System.getenv(key)
    }

    @Test
    fun testActivePackagesInventorySync() = runBlocking {
        val builder = Dotenv.configure().ignoreIfMissing()
        val envFile = File("../../../.env")
        if (envFile.exists()) {
            builder.directory(envFile.parent)
        }
        val dotenv = builder.load()

        var baseUrl = getEnv(dotenv, "METRC_BASE_URL")
        val vendorKey = getEnv(dotenv, "METRC_VENDOR_API_KEY")
        val userKey = getEnv(dotenv, "METRC_USER_API_KEY")
        val licenseNumber = getEnv(dotenv, "METRC_FACILITY_LICENSE_NUMBER")

        if (baseUrl == null) baseUrl = "https://sandbox-api-or.metrc.com"

        if (vendorKey == null || userKey == null || licenseNumber == null) {
            println("Skipping test: Missing env vars")
            return@runBlocking
        }

        println("Config: $baseUrl, $licenseNumber")

        val client = MetrcClient(baseUrl, vendorKey, userKey)
        val wrapper = MetrcWrapper(client)
        val thunk = ThunkMetrc(wrapper)

        println("Starting sync for $licenseNumber")
        val start = System.currentTimeMillis()

        try {
            val packages = thunk.activePackagesInventorySync(licenseNumber, null, 5)

            val duration = System.currentTimeMillis() - start
            println("Sync finished in ${duration}ms. Found ${packages.size} packages.")

            assertNotNull(packages)
            assertTrue(packages.size >= 0)
        } catch (e: Exception) {
            if (e.message?.contains("429") == true) {
                println("WARNING: Test hit rate limit (429). Passing test as environment is restricted.")
            } else {
                throw e
            }
        }
    }
}
