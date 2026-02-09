using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Threading;
using System.Threading.Tasks;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;

namespace ThunkMetrc
{
    public static class MetrcUtilities
    {
        /// <summary>
        /// Retrieves all active facilities and executes the specified operation in parallel across them.
        /// This eliminates the boilerplate of fetching facilities, deserializing them, and setting up SynchronizationTargets.
        /// </summary>
        /// <typeparam name="TResult">The type of the result returned by the operation.</typeparam>
        /// <param name="wrapper">The MetrcWrapper instance to use.</param>
        /// <param name="operation">The operation to perform for each facility. Returns a list of results for that facility.</param>
        /// <param name="concurrencyLimit">Max number of concurrent facility operations (default 20).</param>
        /// <param name="cancellationToken">Cancellation token.</param>
        /// <returns>A dictionary of LicenseNumber -> List of Results.</returns>
        public static async Task<Dictionary<string, List<TResult>>> ExecuteOnAllActiveFacilitiesAsync<TResult>(
            this MetrcWrapper wrapper,
            Func<SynchronizationTarget, CancellationToken, Task<List<TResult>>> operation,
            int concurrencyLimit = 20,
            CancellationToken cancellationToken = default)
        {
            // 1. Fetch Facilities
            var facilities = await wrapper.Facilities.GetFacilities(cancellationToken: cancellationToken);
            
            if (facilities == null) facilities = new List<Facility>();

            // 2. Filter for Active Facilities and Create Targets
            // Note: We use the *same* wrapper instance. 
            // If the user needs different creds per facility, they shouldn't use this helper.
            var targets = facilities
                .Where(f => f.License?.Number != null) // Ensure license exists. Add 'Active' check if property available?
                .Select(f => new SynchronizationTarget
                {
                    LicenseNumber = f.License.Number,
                    Wrapper = wrapper
                })
                .ToList();

            if (targets.Count == 0) return new Dictionary<string, List<TResult>>();

            // 3. Execute Parallel Sync
            return await MetrcExtensions.ParallelSyncAsync(targets, concurrencyLimit, operation, cancellationToken);
        }

        /// <summary>
        /// Chunks a collection into batches of a specified size. Useful for batching API requests (e.g. CreatePackages).
        /// </summary>
        public static IEnumerable<IEnumerable<T>> Batch<T>(this IEnumerable<T> source, int size)
        {
            if (size <= 0) throw new ArgumentException("Batch size must be positive", nameof(size));
            
            using var enumerator = source.GetEnumerator();
            while (enumerator.MoveNext())
            {
                yield return YieldBatchElements(enumerator, size);
            }
        }

        private static IEnumerable<T> YieldBatchElements<T>(IEnumerator<T> source, int size)
        {
            yield return source.Current;
            for (int i = 0; i < size - 1 && source.MoveNext(); i++)
            {
                yield return source.Current;
            }
        }
    }
}
