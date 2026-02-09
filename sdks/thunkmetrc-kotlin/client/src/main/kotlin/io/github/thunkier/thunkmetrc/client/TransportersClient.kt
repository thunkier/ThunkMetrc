package io.github.thunkier.thunkmetrc.client

import java.net.URLEncoder
import java.nio.charset.StandardCharsets

class TransportersClient(private val client: MetrcClient) {
    /**
     * Creates new driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST CreateDrivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createTransportersDrivers(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
    }

    /**
     * Creates new vehicle records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST CreateVehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun createTransportersVehicles(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("POST", path.toString(), body)
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE DeleteDriverById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteTransportersDriverById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/drivers/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE DeleteVehicleById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun deleteTransportersVehicleById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/vehicles/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("DELETE", path.toString(), null)
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Transporters
     * GET GetDriverById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTransportersDriverById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/drivers/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of drivers for a Facility.
     * Permissions Required:
     * - Transporters
     * GET GetDrivers
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTransportersDrivers(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Transporters
     * GET GetVehicleById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTransportersVehicleById(id: String, licenseNumber: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/vehicles/${URLEncoder.encode(id, StandardCharsets.UTF_8)}")
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
     * Retrieves a list of vehicles for a Facility.
     * Permissions Required:
     * - Transporters
     * GET GetVehicles
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    fun getTransportersVehicles(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null, ): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (pageNumber != null) {
            query.add("pageNumber=${URLEncoder.encode(pageNumber, StandardCharsets.UTF_8)}")
        }
        if (pageSize != null) {
            query.add("pageSize=${URLEncoder.encode(pageSize, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("GET", path.toString(), null)
    }

    /**
     * Updates existing driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT UpdateDrivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateTransportersDrivers(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/drivers")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

    /**
     * Updates existing vehicle records for a facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT UpdateVehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    fun updateTransportersVehicles(licenseNumber: String? = null, body: Any? = null): Any? {
        val path = StringBuilder("/transporters/v2/vehicles")
        val query = ArrayList<String>()
        if (licenseNumber != null) {
            query.add("licenseNumber=${URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8)}")
        }
        if (query.isNotEmpty()) {
            path.append("?").append(query.joinToString("&"))
        }

        return client.send<Any>("PUT", path.toString(), body)
    }

}

