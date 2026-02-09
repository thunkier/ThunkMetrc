using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// Retrieves all active units of measure.
        /// </summary>
        public async Task<object> GetActiveUnitsOfMeasure()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/unitsofmeasure/v2/active", null, queryParams);
        }
        /// <summary>
        /// Retrieves all inactive units of measure.
        /// </summary>
        public async Task<object> GetInactiveUnitsOfMeasure()
        {
            var queryParams = new Dictionary<string, string>();
            return await SendAsync<object>(new HttpMethod("GET"), $"/unitsofmeasure/v2/inactive", null, queryParams);
        }
    }
}

