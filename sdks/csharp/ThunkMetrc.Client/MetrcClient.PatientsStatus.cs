using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Patients
        /// </summary>
        public async Task<object> PatientsStatusGetStatusesByPatientLicenseNumberV1(string patientLicenseNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v1/statuses/{Uri.EscapeDataString(patientLicenseNumber)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Patients
        /// </summary>
        public async Task<object> PatientsStatusGetStatusesByPatientLicenseNumberV2(string patientLicenseNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/statuses/{Uri.EscapeDataString(patientLicenseNumber)}", null, queryParams);
        }

    }
}
