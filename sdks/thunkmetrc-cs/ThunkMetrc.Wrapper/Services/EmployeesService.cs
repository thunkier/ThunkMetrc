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
    public class EmployeesService(MetrcClient client, IMetrcRateLimiter rateLimiter)
    {
        private readonly MetrcClient _client = client;
        private readonly IMetrcRateLimiter _rateLimiter = rateLimiter;
        /// <summary>
        /// Retrieves a list of employees for a specified Facility.
        /// Permissions Required:
        /// - Manage Employees
        /// - View Employees
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the employees.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<PaginatedResponse<Employee>> GetEmployees(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetEmployees(licenseNumber, pageNumber, pageSize), cancellationToken);
            if (result is JsonElement json)
            {
                return json.Deserialize<PaginatedResponse<Employee>>() ?? throw new InvalidOperationException("Failed to deserialize response");
            }
            throw new InvalidOperationException("Unexpected response type");
        }
        /// <summary>
        /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        /// Permissions Required:
        /// - Manage Employees
        /// </summary>
        /// <param name="employeeLicenseNumber">The License Number of the Employee under which to get the Permission.</param>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the Employee.</param>
        /// <param name="cancellationToken">The cancellation token to cancel the operation.</param>
        public async Task<object> GetEmployeesPermissions(string? employeeLicenseNumber = null, string? licenseNumber = null, CancellationToken cancellationToken = default)
        {
            var result = await _rateLimiter.ExecuteAsync<object>(null, true, () => _client.GetEmployeesPermissions(employeeLicenseNumber, licenseNumber), cancellationToken);
            if (result is JsonElement json)
            {
                return json;
            }
            throw new InvalidOperationException("Unexpected response type");
        }
    }
}

