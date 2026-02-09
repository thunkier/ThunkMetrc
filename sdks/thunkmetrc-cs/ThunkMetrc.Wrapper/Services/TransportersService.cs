using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using ThunkMetrc.Client;
using Microsoft.Extensions.Logging;
using System.Text.Json;
using System.Threading;
using System.Runtime.CompilerServices;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;
using File = ThunkMetrc.Wrapper.Models.File;

namespace ThunkMetrc.Wrapper.Services
{
    public class TransportersService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Creates new driver records for a Facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a Driver.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateTransportersDrivers(List<TransportersCreateDriversRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateTransportersDrivers(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Creates new vehicle records for a Facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to create a Vehicle.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task CreateTransportersVehicles(List<TransportersCreateVehiclesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.CreateTransportersVehicles(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to archive a Driver.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteTransportersDriverById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteTransportersDriverById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number for which to archive a Vehicle</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task DeleteTransportersVehicleById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.DeleteTransportersVehicleById(id, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a Driver.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<TransportersDriver> GetTransportersDriverById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTransportersDriverById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<TransportersDriver>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of drivers for a Facility.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of drivers.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<TransportersDriver>> GetTransportersDrivers(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTransportersDrivers(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<TransportersDriver>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="id">id</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to return a Vehicle.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<TransportersVehicle> GetTransportersVehicleById(string id, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTransportersVehicleById(id, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<TransportersVehicle>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves a list of vehicles for a Facility.
        /// Permissions Required:
        /// - Transporters
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility for which to return the list of vehicles.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<TransportersVehicle>> GetTransportersVehicles(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetTransportersVehicles(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<TransportersVehicle>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Updates existing driver records for a Facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of drivers.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateTransportersDrivers(List<TransportersUpdateDriversRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateTransportersDrivers(body, licenseNumber); return null; }, cancellationToken);
        }
        /// <summary>
        /// Updates existing vehicle records for a facility.
        /// Permissions Required:
        /// - Manage Transporters
        /// </summary>
        /// <param name="body">Request body.</param>
        /// <param name="licenseNumber">The License Number of the Facility for which to update the list of vehicles.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task UpdateTransportersVehicles(List<TransportersUpdateVehiclesRequest>? body = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            await _rateLimiter.ExecuteAsync<object?>(null, false, async () => { await _client.UpdateTransportersVehicles(body, licenseNumber); return null; }, cancellationToken);
        }
    }
}

