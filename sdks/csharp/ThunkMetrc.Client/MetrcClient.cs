using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace ThunkMetrc.Client
{
    public partial class MetrcClient
    {
        private readonly HttpClient _http;
        private readonly string _baseUrl;
        private readonly JsonSerializerOptions _jsonOptions;

        public MetrcClient(string baseUrl, string vendorKey, string userKey, HttpClient? http = null)
        {
            _baseUrl = baseUrl.TrimEnd('/');
            _http = http ?? new HttpClient();
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
            var url = $"{_baseUrl}{path}";
            if (queryParams != null && queryParams.Count > 0)
            {
                var q = string.Join("&", queryParams.Where(kv => !string.IsNullOrEmpty(kv.Value)).Select(kv => $"{kv.Key}={Uri.EscapeDataString(kv.Value)}"));
                if (!string.IsNullOrEmpty(q)) url += $"?{q}";
            }

            var req = new HttpRequestMessage(method, url);

            if (body != null)
            {
                var json = JsonSerializer.Serialize(body, _jsonOptions);
                req.Content = new StringContent(json, Encoding.UTF8, "application/json");
            }

            var resp = await _http.SendAsync(req);
            resp.EnsureSuccessStatusCode();

            var content = await resp.Content.ReadAsStringAsync();
            if (string.IsNullOrWhiteSpace(content)) return default!;

            return JsonSerializer.Deserialize<T>(content, _jsonOptions)!;
        }

        private async Task SendAsync(HttpMethod method, string path, object? body = null, Dictionary<string, string>? queryParams = null)
        {
             var url = $"{_baseUrl}{path}";
            if (queryParams != null && queryParams.Count > 0)
            {
                var q = string.Join("&", queryParams.Where(kv => !string.IsNullOrEmpty(kv.Value)).Select(kv => $"{kv.Key}={Uri.EscapeDataString(kv.Value)}"));
                if (!string.IsNullOrEmpty(q)) url += $"?{q}";
            }

            var req = new HttpRequestMessage(method, url);

            if (body != null)
            {
                var json = JsonSerializer.Serialize(body, _jsonOptions);
                req.Content = new StringContent(json, Encoding.UTF8, "application/json");
            }

            var resp = await _http.SendAsync(req);
            resp.EnsureSuccessStatusCode();
        }
    }
}
