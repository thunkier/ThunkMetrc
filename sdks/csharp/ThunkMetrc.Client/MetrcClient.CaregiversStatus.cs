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
        ///   - Lookup Caregivers
        /// </summary>
        public async Task<object> CaregiversStatusGetByCaregiverLicenseNumberV1(string caregiverLicenseNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/caregivers/v1/status/{Uri.EscapeDataString(caregiverLicenseNumber)}", null, queryParams);
        }

        /// <summary>
        /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        /// 
        ///   Permissions Required:
        ///   - Lookup Caregivers
        /// </summary>
        public async Task<object> CaregiversStatusGetByCaregiverLicenseNumberV2(string caregiverLicenseNumber, string? licenseNumber = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(licenseNumber)) queryParams["licenseNumber"] = licenseNumber;
            return await SendAsync<object>(new HttpMethod("GET"), $"/caregivers/v2/status/{Uri.EscapeDataString(caregiverLicenseNumber)}", null, queryParams);
        }

    }
}
