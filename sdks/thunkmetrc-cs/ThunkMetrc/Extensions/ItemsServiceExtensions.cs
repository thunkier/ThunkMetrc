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
/// Extension methods for ItemsService.
/// </summary>
public static class ItemsServiceExtensions
{
    /// <summary>
    /// Iterator for GetActiveItems.
    /// Returns a list of active items for the specified Facility.
    /// Permissions Required:
    /// - Manage Items
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp</param>
    /// <param name="lastModifiedStart">The last modified start timestamp</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active items.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Item> IterateGetActiveItems(
        this ItemsService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetActiveItems(
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
    /// Syncs GetActiveItems data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<Item>> SyncGetActiveItems(
        this ItemsService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<Item>();
        await foreach (var item in service.IterateGetActiveItems(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetActiveItems across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<Item>>> SyncGetActiveItemsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<Item>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<Item>();
               await foreach (var item in target.Wrapper.Items.IterateGetActiveItems(
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
    /// Iterator for GetItemsBrands.
    /// Retrieves a list of active item brands for the specified Facility.
    /// Permissions Required:
    /// - Manage Items
    /// </summary>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active item brands.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Brand> IterateGetItemsBrands(
        this ItemsService service,
        string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetItemsBrands(
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
    /// <summary>
    /// Iterator for GetItemsCategories.
    /// Retrieves a list of item categories.
    /// </summary>
    /// <param name="licenseNumber">If specified, the categories will be retrieved for the specified License Number. If not specified, all item categories will be returned.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Category> IterateGetItemsCategories(
        this ItemsService service,
        string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetItemsCategories(
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
    /// <summary>
    /// Iterator for GetInactiveItems.
    /// Retrieves a list of inactive items for the specified Facility.
    /// Permissions Required:
    /// - Manage Items
    /// </summary>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive items.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<Item> IterateGetInactiveItems(
        this ItemsService service,
        string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetInactiveItems(
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
