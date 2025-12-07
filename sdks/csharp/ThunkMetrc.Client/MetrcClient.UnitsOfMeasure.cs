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
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetActiveV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/unitsofmeasure/v1/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves all active units of measure.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetActiveV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/unitsofmeasure/v2/active", null, queryParams);
        }

        /// <summary>
        /// Retrieves all inactive units of measure.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> UnitsOfMeasureGetInactiveV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/unitsofmeasure/v2/inactive", null, queryParams);
        }

    }
}
