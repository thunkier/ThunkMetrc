using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        /// <summary>
        /// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
        /// 
        ///   Permissions Required:
        ///   - None
        /// </summary>
        public async Task SandboxCreateIntegratorSetupV2(object? body = null, string? userKey = null)
        {
            var queryParams = new Dictionary<string, string>();
            if (!string.IsNullOrEmpty(userKey)) queryParams["userKey"] = userKey;
            await SendAsync(new HttpMethod("POST"), $"/sandbox/v2/integrator/setup", body, queryParams);
        }

    }
}
