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
/// Extension methods for PlantBatchesService.
/// </summary>
public static class PlantBatchesServiceExtensions
{
    /// <summary>
    /// Iterator for GetActivePlantBatches.
    /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
    /// Permissions Required:
    /// - View Immature Plants
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active plant batches.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<PlantBatch> IterateGetActivePlantBatches(
        this PlantBatchesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetActivePlantBatches(
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
    /// Syncs GetActivePlantBatches data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<PlantBatch>> SyncGetActivePlantBatches(
        this PlantBatchesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<PlantBatch>();
        await foreach (var item in service.IterateGetActivePlantBatches(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetActivePlantBatches across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<PlantBatch>>> SyncGetActivePlantBatchesParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<PlantBatch>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<PlantBatch>();
               await foreach (var item in target.Wrapper.PlantBatches.IterateGetActivePlantBatches(
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
    /// Iterator for GetInactivePlantBatches.
    /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
    /// Permissions Required:
    /// - View Immature Plants
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive plant batches.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<PlantBatch> IterateGetInactivePlantBatches(
        this PlantBatchesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetInactivePlantBatches(
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
    /// Syncs GetInactivePlantBatches data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<PlantBatch>> SyncGetInactivePlantBatches(
        this PlantBatchesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<PlantBatch>();
        await foreach (var item in service.IterateGetInactivePlantBatches(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetInactivePlantBatches across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<PlantBatch>>> SyncGetInactivePlantBatchesParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<PlantBatch>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<PlantBatch>();
               await foreach (var item in target.Wrapper.PlantBatches.IterateGetInactivePlantBatches(
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
    /// Iterator for GetPlantBatchesTypes.
    /// Retrieves a list of plant batch types.
    /// </summary>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<PlantBatchesType> IterateGetPlantBatchesTypes(
        this PlantBatchesService service,
        string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetPlantBatchesTypes(
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
    /// <summary>
    /// Iterator for GetPlantBatchesWaste.
    /// Retrieves waste details associated with plant batches at a specified Facility.
    /// Permissions Required:
    /// - View Plants Waste
    /// </summary>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of waste plants.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<PlantBatchesWaste> IterateGetPlantBatchesWaste(
        this PlantBatchesService service,
        string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetPlantBatchesWaste(
                licenseNumber, page.ToString(), pageSize,
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
