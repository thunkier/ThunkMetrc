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
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task<object> WasteMethodsGetAllV2(string? No = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(No)) queryParams["No"] = No;
            return await SendAsync<object>(new HttpMethod("GET"), $"/wastemethods/v2", null, queryParams);
        }

    }
}
