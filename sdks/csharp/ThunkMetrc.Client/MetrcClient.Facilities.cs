using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// This endpoint provides a list of facilities for which the authenticated user has access.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> FacilitiesGetAllV1(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/facilities/v1", null, queryParams);
        }

        /// <summary>
        /// This endpoint provides a list of facilities for which the authenticated user has access.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> FacilitiesGetAllV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/facilities/v2", null, queryParams);
        }

    }
}
