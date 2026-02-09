package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class TagsClient(private val client: MetrcClient) {
    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET GetAvailablePackage
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTagsAvailablePackage(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/tags/v2/package/available")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * GET GetAvailablePlant
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTagsAvailablePlant(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/tags/v2/plant/available")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * Permissions Required:
     * - Manage Tags
     * - RetailId.AllowPackageStaging Key Value enabled
     * GET GetStagedTags
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getStagedTags(licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/tags/v2/staged")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

