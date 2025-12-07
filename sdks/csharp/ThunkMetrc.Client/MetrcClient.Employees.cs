using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Permissions Required:
        ///   - Manage Employees
        /// </summary>
        public async Task<object> EmployeesGetAllV1(string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/employees/v1", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of employees for a specified Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Employees
        ///   - View Employees
        /// </summary>
        public async Task<object> EmployeesGetAllV2(string? licenseNumber = null, string? pageNumber = null, string? pageSize = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            if (!string.IsNullOrEmpty(pageNumber)) queryParams["pageNumber"] = pageNumber;
            if (!string.IsNullOrEmpty(pageSize)) queryParams["pageSize"] = pageSize;
            return await SendAsync<object>(new HttpMethod("GET"), $"/employees/v2", null, queryParams);
        }

        /// <summary>
        /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        /// 
        ///   Permissions Required:
        ///   - Manage Employees
        /// </summary>
        public async Task<object> EmployeesGetPermissionsV2(string? employeeLicenseNumber = null, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(employeeLicenseNumber)) queryParams["employeeLicenseNumber"] = employeeLicenseNumber;
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/employees/v2/permissions", null, queryParams);
        }

    }
}
