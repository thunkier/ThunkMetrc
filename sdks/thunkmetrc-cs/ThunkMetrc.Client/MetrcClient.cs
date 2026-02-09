using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        private readonly HttpClient _http;
        private readonly string _baseUrl;
        private readonly JsonSerializerOptions _jsonOptions;
        private readonly ILogger? _logger;

        public MetrcClient(string baseUrl, string vendorKey, string userKey, HttpClient? http = null, ILogger? logger = null)
        {
            _baseUrl = baseUrl.TrimEnd('/');
            _http = http ?? new HttpClient();
            _logger = logger;
            var basicAuth = Convert.ToBase64String(Encoding.UTF8.GetBytes($"{vendorKey}:{userKey}"));
            _http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Basic", basicAuth);
            _jsonOptions = new JsonSerializerOptions
            {
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
                DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
                WriteIndented = true
            };
        }

        private async Task<T> SendAsync<T>(HttpMethod method, string path, object? body = null, Dictionary<string, string>? queryParams = null)
        {
            var url = BuildUrl(path, queryParams);
            var req = CreateRequest(method, url, body);

            _logger?.LogDebug("Sending {Method} request to {Url}", method, url);

            var resp = await _http.SendAsync(req);
            
            if (!resp.IsSuccessStatusCode)
            {
                await HandleErrorAsync(resp);
            }

            var content = await resp.Content.ReadAsStringAsync();
            if (string.IsNullOrWhiteSpace(content)) return default!;

            try 
            {
                return JsonSerializer.Deserialize<T>(content, _jsonOptions)!;
            }
            catch (JsonException ex)
            {
                _logger?.LogError(ex, "Failed to deserialize response from {Url}", url);
                throw;
            }
        }

        private async Task SendAsync(HttpMethod method, string path, object? body = null, Dictionary<string, string>? queryParams = null)
        {
            var url = BuildUrl(path, queryParams);
            var req = CreateRequest(method, url, body);

            _logger?.LogDebug("Sending {Method} request to {Url}", method, url);

            var resp = await _http.SendAsync(req);

            if (!resp.IsSuccessStatusCode)
            {
                await HandleErrorAsync(resp);
            }
        }

        private string BuildUrl(string path, Dictionary<string, string>? queryParams)
        {
            var url = $"{_baseUrl}{path}";
            if (queryParams != null && queryParams.Count > 0)
            {
                var q = string.Join("&", queryParams.Where(kv => !string.IsNullOrEmpty(kv.Value)).Select(kv => $"{kv.Key}={Uri.EscapeDataString(kv.Value)}"));
                if (!string.IsNullOrEmpty(q)) url += $"?{q}";
            }
            return url;
        }

        private HttpRequestMessage CreateRequest(HttpMethod method, string url, object? body)
        {
             var req = new HttpRequestMessage(method, url);
             if (body != null)
            {
                var json = JsonSerializer.Serialize(body, _jsonOptions);
                req.Content = new StringContent(json, Encoding.UTF8, "application/json");
            }
            return req;
        }

        private async Task HandleErrorAsync(HttpResponseMessage resp)
        {
            var errorContent = await resp.Content.ReadAsStringAsync();
            _logger?.LogError("Request failed with {StatusCode}. Body: {Body}", resp.StatusCode, errorContent);

            string? apiMessage = null;
            string[]? validationErrors = null;

            if (!string.IsNullOrWhiteSpace(errorContent))
            {
                try
                {
                    // Attempt to parse structured error
                    var errorObj = JsonSerializer.Deserialize<MetrcErrorDto>(errorContent, _jsonOptions);
                    apiMessage = errorObj?.Message;
                    validationErrors = errorObj?.ValidationErrors;
                }
                catch
                {
                    // Allow raw body fallback if parsing fails
                }
            }

            throw new MetrcException(resp.StatusCode, errorContent, resp.Headers, apiMessage, validationErrors);
        }

        // Internal DTO for parsing error responses
        private class MetrcErrorDto
        {
            public string? Message { get; set; }
            public string[]? ValidationErrors { get; set; }
        }
    }

    public class MetrcException : Exception
    {
        public HttpStatusCode StatusCode { get; }
        public string? ResponseBody { get; }
        public HttpResponseHeaders? Headers { get; }
        
        // Parsed fields
        public string? ApiMessage { get; }
        public string[]? ValidationErrors { get; }

        public MetrcException(HttpStatusCode statusCode, string? responseBody, HttpResponseHeaders? headers, string? apiMessage = null, string[]? validationErrors = null)
            : base(FormatMessage(statusCode, apiMessage, validationErrors))
        {
            StatusCode = statusCode;
            ResponseBody = responseBody;
            Headers = headers;
            ApiMessage = apiMessage;
            ValidationErrors = validationErrors;
        }

        private static string FormatMessage(HttpStatusCode statusCode, string? apiMessage, string[]? validationErrors)
        {
            var sb = new StringBuilder($"Metrc API Error: {statusCode}");
            if (!string.IsNullOrEmpty(apiMessage))
            {
                sb.Append($" - {apiMessage}");
            }
            if (validationErrors != null && validationErrors.Length > 0)
            {
                sb.Append(" | Validation Errors: ");
                sb.Append(string.Join("; ", validationErrors));
            }
            return sb.ToString();
        }
    }
}
