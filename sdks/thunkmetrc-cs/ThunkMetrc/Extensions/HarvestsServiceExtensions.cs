using System;
using System.Collections.Generic;
using System.Runtime.CompilerServices;
using System.Threading;
using System.Threading.Tasks;
using ThunkMetrc.Wrapper;
using ThunkMetrc.Wrapper.Models;
using ThunkMetrc.Wrapper.Services;

namespace ThunkMetrc.Extensions;

/// <summary>
/// Extension methods for HarvestsService.
/// </summary>
public static class HarvestsServiceExtensions
{
    /// <summary>
    /// Iterator for GetActiveHarvests.
    /// Retrieves a list of active harvests for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active harvests.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Harvest> IterateGetActiveHarvests(
        this HarvestsService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetActiveHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, page.ToString(), pageSize,
                cancellationToken: cancellationToken
            );

            if (response?.Data == null || response.Data.Count == 0)
            {
                yield break;
            }

            foreach (var item in response.Data)
            {
                yield return item;
            }

            if (response.Data.Count < 20)
            {
                // Assuming default page size of 20 implies end of data
            }
            
            page++;
        }
    }
    /// <summary>
    /// Syncs GetActiveHarvests data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<Harvest>> SyncGetActiveHarvests(
        this HarvestsService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<Harvest>();
        await foreach (var item in service.IterateGetActiveHarvests(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetActiveHarvests across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<Harvest>>> SyncGetActiveHarvestsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<Harvest>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<Harvest>();
               await foreach (var item in target.Wrapper.Harvests.IterateGetActiveHarvests(
                   endStr, startStr, target.LicenseNumber, pageSize,
                   cancellationToken: ct
               ))
               {
                   list.Add(item);
               }
               return list;
            },
            cancellationToken
        );
    }
    /// <summary>
    /// Iterator for GetHarvestsWaste.
    /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
    /// Permissions Required:
    /// - View Harvests
    /// </summary>
    /// <param name="harvestId">The harvestId is the unique identifier for each harvest.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste harvests.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<HarvestsWaste> IterateGetHarvestsWaste(
        this HarvestsService service,
        string? harvestId = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetHarvestsWaste(
                harvestId, licenseNumber, page.ToString(), pageSize,
                cancellationToken: cancellationToken
            );

            if (response?.Data == null || response.Data.Count == 0)
            {
                yield break;
            }

            foreach (var item in response.Data)
            {
                yield return item;
            }

            if (response.Data.Count < 20)
            {
                // Assuming default page size of 20 implies end of data
            }
            
            page++;
        }
    }
    /// <summary>
    /// Iterator for GetInactiveHarvests.
    /// Retrieves a list of inactive harvests for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive harvests.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Harvest> IterateGetInactiveHarvests(
        this HarvestsService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetInactiveHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, page.ToString(), pageSize,
                cancellationToken: cancellationToken
            );

            if (response?.Data == null || response.Data.Count == 0)
            {
                yield break;
            }

            foreach (var item in response.Data)
            {
                yield return item;
            }

            if (response.Data.Count < 20)
            {
                // Assuming default page size of 20 implies end of data
            }
            
            page++;
        }
    }
    /// <summary>
    /// Syncs GetInactiveHarvests data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<Harvest>> SyncGetInactiveHarvests(
        this HarvestsService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<Harvest>();
        await foreach (var item in service.IterateGetInactiveHarvests(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetInactiveHarvests across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<Harvest>>> SyncGetInactiveHarvestsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<Harvest>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<Harvest>();
               await foreach (var item in target.Wrapper.Harvests.IterateGetInactiveHarvests(
                   endStr, startStr, target.LicenseNumber, pageSize,
                   cancellationToken: ct
               ))
               {
                   list.Add(item);
               }
               return list;
            },
            cancellationToken
        );
    }
    /// <summary>
    /// Iterator for GetOnHoldHarvests.
    /// Retrieves a list of harvests on hold for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of harvests on hold.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Harvest> IterateGetOnHoldHarvests(
        this HarvestsService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetOnHoldHarvests(
                lastModifiedEnd, lastModifiedStart, licenseNumber, page.ToString(), pageSize,
                cancellationToken: cancellationToken
            );

            if (response?.Data == null || response.Data.Count == 0)
            {
                yield break;
            }

            foreach (var item in response.Data)
            {
                yield return item;
            }

            if (response.Data.Count < 20)
            {
                // Assuming default page size of 20 implies end of data
            }
            
            page++;
        }
    }
    /// <summary>
    /// Syncs GetOnHoldHarvests data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<Harvest>> SyncGetOnHoldHarvests(
        this HarvestsService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<Harvest>();
        await foreach (var item in service.IterateGetOnHoldHarvests(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetOnHoldHarvests across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<Harvest>>> SyncGetOnHoldHarvestsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<Harvest>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<Harvest>();
               await foreach (var item in target.Wrapper.Harvests.IterateGetOnHoldHarvests(
                   endStr, startStr, target.LicenseNumber, pageSize,
                   cancellationToken: ct
               ))
               {
                   list.Add(item);
               }
               return list;
            },
            cancellationToken
        );
    }
    /// <summary>
    /// Iterator for GetHarvestsWasteTypes.
    /// Retrieves a list of Waste types for harvests.
    /// </summary>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<WasteType> IterateGetHarvestsWasteTypes(
        this HarvestsService service,
        string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetHarvestsWasteTypes(
                page.ToString(), pageSize,
                cancellationToken: cancellationToken
            );

            if (response?.Data == null || response.Data.Count == 0)
            {
                yield break;
            }

            foreach (var item in response.Data)
            {
                yield return item;
            }

            if (response.Data.Count < 20)
            {
                // Assuming default page size of 20 implies end of data
            }
            
            page++;
        }
    }
}
