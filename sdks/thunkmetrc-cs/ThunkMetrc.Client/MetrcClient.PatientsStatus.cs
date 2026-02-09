using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// Permissions Required:
        /// - Lookup Patients
        /// </summary>
        /// <param name="patientLicenseNumber">patientLicenseNumber</param>
        /// <param name="licenseNumber">The License Number of the Facility under which to get the Patient Status.</param>
        public async Task<object> GetPatientsStatusesByPatientLicenseNumber(string patientLicenseNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/patients/v2/statuses/{Uri.EscapeDataString(patientLicenseNumber)}", null, queryParams);
        }
    }
}

