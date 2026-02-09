using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using ThunkMetrc.Wrapper.Models;

namespace ThunkMetrc.Wrapper
{
    public static class MetrcExtensions
    {
        private const string MetrcTimeFormat = "yyyy-MM-ddTHH:mm:ssZ07:00";

        /// <summary>
        /// Calculates a Metrc-compatible time window string tuple (StartString, EndString) based on the last known sync time.
        /// Handles the buffer logic and RFC3339 formatting.
        /// </summary>
        /// <param name="lastKnownSync">The last time the sync ran. If null, defaults to 24 hours ago.</param>
        /// <param name="bufferMinutes">Minutes to subtract from start time for safety (default 0).</param>
        /// <returns>A tuple of (StartString, EndString) formatted for Metrc API.</returns>
        public static (string StartString, string EndString) GetTimeWindow(DateTime? lastKnownSync, int bufferMinutes = 0)
        {
            var endTime = DateTime.UtcNow;
            var startTime = endTime.AddHours(-24);

            if (lastKnownSync.HasValue)
            {
                if (bufferMinutes < 0) bufferMinutes = 0;
                startTime = lastKnownSync.Value.AddMinutes(-bufferMinutes);
            }

            return (
                startTime.ToString(MetrcTimeFormat),
                endTime.ToString(MetrcTimeFormat)
            );
        }

        /// <summary>
        /// Executes an operation concurrently across multiple targets with a specified concurrency limit.
        /// </summary>
        /// <typeparam name="TResult">The type of the result item (e.g. Package).</typeparam>
        /// <param name="targets">List of targets (license + wrapper).</param>
        /// <param name="concurrencyLimit">Maximum number of concurrent facility requests.</param>
        /// <param name="operation">The operation to perform for each target.</param>
        /// <param name="cancellationToken">Cancellation token.</param>
        /// <returns>A dictionary of LicenseNumber -> List of Results.</returns>
        public static async Task<Dictionary<string, List<TResult>>> ParallelSyncAsync<TResult>(
            IEnumerable<SynchronizationTarget> targets,
            int concurrencyLimit,
            Func<SynchronizationTarget, CancellationToken, Task<List<TResult>>> operation,
            CancellationToken cancellationToken = default)
        {
            if (concurrencyLimit <= 0) concurrencyLimit = 20;

            var results = new ConcurrentDictionary<string, List<TResult>>();
            var semaphore = new SemaphoreSlim(concurrencyLimit);
            var tasks = new List<Task>();

            foreach (var target in targets)
            {
                tasks.Add(Task.Run(async () =>
                {
                    await semaphore.WaitAsync(cancellationToken);
                    try
                    {
                        var items = await operation(target, cancellationToken);
                        results.TryAdd(target.LicenseNumber, items);
                    }
                    finally
                    {
                        semaphore.Release();
                    }
                }, cancellationToken));
            }

            await Task.WhenAll(tasks);
            return results.ToDictionary(k => k.Key, v => v.Value);
        }

        /// <summary>
        /// Generic iterator for paged endpoints.
        /// Handles getting pages, yielding items, and checking TotalPages.
        /// </summary>
        public static async IAsyncEnumerable<T> IteratePagesAsync<T>(
            Func<int, CancellationToken, Task<PaginatedResponse<T>?>> fetchPage, 
            [System.Runtime.CompilerServices.EnumeratorCancellation] CancellationToken cancellationToken = default)
        {
            int page = 1;
            while (!cancellationToken.IsCancellationRequested)
            {
                var response = await fetchPage(page, cancellationToken);
                
                // Fallback for empty/null responses
                if (response == null || response.Data == null || response.Data.Count == 0) yield break;

                foreach (var item in response.Data)
                {
                    yield return item;
                }

                if (response.TotalPages > 0 && page >= response.TotalPages) yield break;
                
                // Safety break for single page responses that don't report TotalPages
                if (response.TotalPages == 0) yield break;

                page++;
            }
        }

        /// <summary>
        /// Helper to deserialize a JSON result into a PaginatedResponse, handling both wrapped and raw list formats.
        /// </summary>
        public static PaginatedResponse<T>? DeserializePage<T>(object? result, int page)
        {
            if (result is not JsonElement root || root.ValueKind == JsonValueKind.Null) return null;

            try
            {
                // 1. Try deserializing as PaginatedResponse
                var response = JsonSerializer.Deserialize<PaginatedResponse<T>>(root);
                
                // 2. Fallback: If Data is empty but it looks like a raw array, try deserializing as List<T>
                if ((response?.Data == null || response.Data.Count == 0) && root.ValueKind == JsonValueKind.Array)
                {
                    var items = JsonSerializer.Deserialize<List<T>>(root);
                    if (items != null)
                    {
                        return new PaginatedResponse<T>
                        {
                            Data = items,
                            CurrentPage = page,
                            TotalPages = 1 // Assume single page for raw lists
                        };
                    }
                }

                return response;
            }
            catch
            {
                // Deserialization failure
                return null;
            }
        }
    }

    public class SynchronizationTarget
    {
        public string LicenseNumber { get; set; } = string.Empty;
        public MetrcWrapper Wrapper { get; set; } = default!;
    }
}
