using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Retrieves all available waste methods.
        /// </summary>
        public async Task<object> GetWasteMethodsWasteMethods()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/wastemethods/v2", null, queryParams);
        }
    }
}

