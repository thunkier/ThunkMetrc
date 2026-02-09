package io.github.thunkier.thunkmetrc.wrapper.services

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import kotlinx.coroutines.flow.Flow

class TransportersService(private val client: MetrcClient, private val rateLimiter: MetrcRateLimiter) {

    /**
     * Creates new driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createDrivers(licenseNumber: String? = null, body: List<TransportersCreateDriversRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.createTransportersDrivers(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Creates new vehicle records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun createVehicles(licenseNumber: String? = null, body: List<TransportersCreateVehiclesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.createTransportersVehicles(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE /transporters/v2/drivers/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteDriverById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.deleteTransportersDriverById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE /transporters/v2/vehicles/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun deleteVehicleById(id: String, licenseNumber: String? = null): Any? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.deleteTransportersVehicleById(id, licenseNumber, ) 
        } as? Any
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/drivers/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDriverById(id: String, licenseNumber: String? = null): TransportersDriver? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transporters.getTransportersDriverById(id, licenseNumber, ) 
        } as? TransportersDriver
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getDrivers(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TransportersDriver>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transporters.getTransportersDrivers(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TransportersDriver>
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/vehicles/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getVehicleById(id: String, licenseNumber: String? = null): TransportersVehicle? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transporters.getTransportersVehicleById(id, licenseNumber, ) 
        } as? TransportersVehicle
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun getVehicles(licenseNumber: String? = null, pageNumber: String? = null, pageSize: String? = null): PaginatedResponse<TransportersVehicle>? {
        return rateLimiter.execute(null,true,
        ) { 
            client.transporters.getTransportersVehicles(licenseNumber, pageNumber, pageSize, ) 
        } as? PaginatedResponse<TransportersVehicle>
    }

    /**
     * Updates existing driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateDrivers(licenseNumber: String? = null, body: List<TransportersUpdateDriversRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.updateTransportersDrivers(licenseNumber, body) 
        } as? WriteResponse
    }

    /**
     * Updates existing vehicle records for a facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    suspend fun updateVehicles(licenseNumber: String? = null, body: List<TransportersUpdateVehiclesRequestItem>): WriteResponse? {
        return rateLimiter.execute(null,false,
        ) { 
            client.transporters.updateTransportersVehicles(licenseNumber, body) 
        } as? WriteResponse
    }
}

