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
/// Extension methods for SalesService.
/// </summary>
public static class SalesServiceExtensions
{
    /// <summary>
    /// Iterator for GetSalesActiveDeliveries.
    /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active deliveries.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<ActiveDelivery> IterateGetSalesActiveDeliveries(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesActiveDeliveries(
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
    /// Syncs GetSalesActiveDeliveries data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<ActiveDelivery>> SyncGetSalesActiveDeliveries(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<ActiveDelivery>();
        await foreach (var item in service.IterateGetSalesActiveDeliveries(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesActiveDeliveries across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<ActiveDelivery>>> SyncGetSalesActiveDeliveriesParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<ActiveDelivery>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<ActiveDelivery>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesActiveDeliveries(
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
    /// Iterator for GetSalesActiveDeliveriesRetailer.
    /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
    /// Permissions Required:
    /// - View Retailer Delivery
    /// - Manage Retailer Delivery
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return active retailer deliveries.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<ActiveDeliveriesRetailer> IterateGetSalesActiveDeliveriesRetailer(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesActiveDeliveriesRetailer(
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
    /// Syncs GetSalesActiveDeliveriesRetailer data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<ActiveDeliveriesRetailer>> SyncGetSalesActiveDeliveriesRetailer(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<ActiveDeliveriesRetailer>();
        await foreach (var item in service.IterateGetSalesActiveDeliveriesRetailer(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesActiveDeliveriesRetailer across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<ActiveDeliveriesRetailer>>> SyncGetSalesActiveDeliveriesRetailerParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<ActiveDeliveriesRetailer>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<ActiveDeliveriesRetailer>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesActiveDeliveriesRetailer(
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
    /// Iterator for GetSalesActiveReceipts.
    /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of active receipts.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<ActiveReceipt> IterateGetSalesActiveReceipts(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesActiveReceipts(
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
    /// Syncs GetSalesActiveReceipts data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<ActiveReceipt>> SyncGetSalesActiveReceipts(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<ActiveReceipt>();
        await foreach (var item in service.IterateGetSalesActiveReceipts(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesActiveReceipts across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<ActiveReceipt>>> SyncGetSalesActiveReceiptsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<ActiveReceipt>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<ActiveReceipt>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesActiveReceipts(
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
    /// Iterator for GetSalesDeliveriesReturnReasons.
    /// Returns a list of return reasons for sales deliveries based on the provided License Number.
    /// Permissions Required:
    /// - Sales Delivery
    /// </summary>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of delivery return reasons.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<DeliveriesReturnReason> IterateGetSalesDeliveriesReturnReasons(
        this SalesService service,
        string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesDeliveriesReturnReasons(
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
    /// Iterator for GetSalesInactiveDeliveries.
    /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive deliveries.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<InactiveDelivery> IterateGetSalesInactiveDeliveries(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesInactiveDeliveries(
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
    /// Syncs GetSalesInactiveDeliveries data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<InactiveDelivery>> SyncGetSalesInactiveDeliveries(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<InactiveDelivery>();
        await foreach (var item in service.IterateGetSalesInactiveDeliveries(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesInactiveDeliveries across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<InactiveDelivery>>> SyncGetSalesInactiveDeliveriesParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<InactiveDelivery>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<InactiveDelivery>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesInactiveDeliveries(
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
    /// Iterator for GetSalesInactiveDeliveriesRetailer.
    /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
    /// Permissions Required:
    /// - View Retailer Delivery
    /// - Manage Retailer Delivery
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return inactive retailer deliveries.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<InactiveDeliveriesRetailer> IterateGetSalesInactiveDeliveriesRetailer(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesInactiveDeliveriesRetailer(
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
    /// Syncs GetSalesInactiveDeliveriesRetailer data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<InactiveDeliveriesRetailer>> SyncGetSalesInactiveDeliveriesRetailer(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<InactiveDeliveriesRetailer>();
        await foreach (var item in service.IterateGetSalesInactiveDeliveriesRetailer(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesInactiveDeliveriesRetailer across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<InactiveDeliveriesRetailer>>> SyncGetSalesInactiveDeliveriesRetailerParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<InactiveDeliveriesRetailer>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<InactiveDeliveriesRetailer>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesInactiveDeliveriesRetailer(
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
    /// Iterator for GetSalesInactiveReceipts.
    /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// </summary>
    /// <param name="lastModifiedEnd">The last modified end timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="lastModifiedStart">The last modified start timestamp. If specified, also specifying any of the sales date parameters will result in an error.</param>
    /// <param name="licenseNumber">The License Number of the Facility for which to return the list of inactive receipts.</param>
    /// <param name="pageSize">The number of records to return per page. Pagination is currently disabled by default. You can enable pagination on this query by specifying a value that does not exceed 20.</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async IAsyncEnumerable<InactiveReceipt> IterateGetSalesInactiveReceipts(
        this SalesService service,
        string? lastModifiedEnd = null, string? lastModifiedStart = null, string? licenseNumber = null, string? pageSize = null, 
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        int page = 1;
        while (!cancellationToken.IsCancellationRequested)
        {
            var response = await service.GetSalesInactiveReceipts(
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
    /// Syncs GetSalesInactiveReceipts data for a specific time window.
    /// Automatically handles iteration and pagination.
    /// </summary>
    /// <param name="lastKnownSync">Last sync time. Defaults to 24 hours ago.</param>
    /// <param name="bufferMinutes">Buffer minutes (default 5).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<List<InactiveReceipt>> SyncGetSalesInactiveReceipts(
        this SalesService service,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5,
        string? licenseNumber = null,
    string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        var results = new List<InactiveReceipt>();
        await foreach (var item in service.IterateGetSalesInactiveReceipts(
             endStr, startStr, licenseNumber, pageSize,
             cancellationToken: cancellationToken))
        {
            results.Add(item);
        }
        return results;
    }

    /// <summary>
    /// Parallel Sync for GetSalesInactiveReceipts across multiple targets.
    /// </summary>
    /// <param name="targets">Targets to sync.</param>
    /// <param name="lastKnownSync">Last sync time.</param>
    /// <param name="bufferMinutes">Buffer minutes.</param>
    /// <param name="concurrencyLimit">Max concurrency (default 20).</param>
    /// <param name="cancellationToken">Cancellation token.</param>
    public static async Task<Dictionary<string, List<InactiveReceipt>>> SyncGetSalesInactiveReceiptsParallel(
        this IEnumerable<global::ThunkMetrc.Wrapper.SynchronizationTarget> targets,
        DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5, int concurrencyLimit = 20,
        string? pageSize = null,
    
        CancellationToken cancellationToken = default)
    {
        var (startStr, endStr) = MetrcExtensions.GetTimeWindow(lastKnownSync?.UtcDateTime, bufferMinutes);
        
        return await MetrcExtensions.ParallelSyncAsync<InactiveReceipt>(
            targets,
            concurrencyLimit,
            async (target, ct) => 
            {
               var list = new List<InactiveReceipt>();
               await foreach (var item in target.Wrapper.Sales.IterateGetSalesInactiveReceipts(
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
}
