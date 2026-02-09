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
        /// </summary>
        public async Task<object> GetFacilities()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/facilities/v2", null, queryParams);
        }
    }
}

