package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class UnitsOfMeasureClient(private val client: MetrcClient) {
    /**
     * Retrieves all active units of measure.
     * GET GetActiveUnitsOfMeasure
     * @throws IOException If request fails
     * @return Response object
     */
    fun getActiveUnitsOfMeasure(): Any? {
        val path = StringBuilder("/unitsofmeasure/v2/active")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves all inactive units of measure.
     * GET GetInactiveUnitsOfMeasure
     * @throws IOException If request fails
     * @return Response object
     */
    fun getInactiveUnitsOfMeasure(): Any? {
        val path = StringBuilder("/unitsofmeasure/v2/inactive")
        val query = ArrayList<String>()
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

}

