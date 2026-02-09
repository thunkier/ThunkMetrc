using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Retrieves a list of employees for a specified Facility.
        /// Permissions Required:
        /// - Manage Employees
        /// - View Employees
        /// </summary>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the employees.</param>
        /// <param name="pageNumber">The number of the data page from which to return data.</param>
        /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
        public async Task<object> GetEmployees(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/employees/v2", null, queryParams);
        }
        /// <summary>
        /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        /// Permissions Required:
        /// - Manage Employees
        /// </summary>
        /// <param name="employeeLicenseNumber">The License Number of the Employee under which to get the Permission.</param>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the Employee.</param>
        public async Task<object> GetEmployeesPermissions(string? employeeLicenseNumber = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(employeeLicenseNumber)) queryParams["employeeLicenseNumber"] = employeeLicenseNumber;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/employees/v2/permissions", null, queryParams);
        }
    }
}

