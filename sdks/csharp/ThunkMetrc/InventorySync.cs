using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Threading.Tasks;
using ThunkMetrc.Wrapper;

namespace ThunkMetrc;

public class ThunkMetrc
{
    private readonly MetrcWrapper _wrapper;

    public ThunkMetrc(MetrcWrapper wrapper)
    {
        _wrapper = wrapper;
    }

    /// <summary>
    /// Syncs all active packages modified within the time window.
    /// Handles pagination automatically.
    /// </summary>
    /// <param name="licenseNumber">Facility license number.</param>
    /// <param name="lastKnownSync">Last successful sync time. Default: 24h ago.</param>
    /// <param name="bufferMinutes">Buffer to subtract from start time. Default: 5.</param>
    /// <returns>List of active packages (as dynamic objects).</returns>
    public async Task<List<object>> ActivePackagesInventorySyncAsync(string licenseNumber, DateTimeOffset? lastKnownSync = null, int bufferMinutes = 5)
    {
        var endTime = DateTimeOffset.UtcNow;
        var startTime = lastKnownSync?.AddMinutes(-bufferMinutes) ?? endTime.AddHours(-24);

        var startStr = startTime.ToString("yyyy-MM-ddTHH:mm:ssK");
        var endStr = endTime.ToString("yyyy-MM-ddTHH:mm:ssK");

        var allPackages = new List<object>();
        int page = 1;
        const int pageSize = 20;

        while (true)
        {
            var result = await _wrapper.PackagesGetActiveV2(
                lastModifiedEnd: endStr,
                lastModifiedStart: startStr,
                licenseNumber: licenseNumber,
                pageNumber: page.ToString(),
                pageSize: pageSize.ToString()
            );

            if (result is JsonElement root)
            {
                if (root.TryGetProperty("Data", out var dataElement) && dataElement.ValueKind == JsonValueKind.Array)
                {
                    foreach (var item in dataElement.EnumerateArray())
                    {
                        allPackages.Add(item.Clone()); 
                    }
                }

                if (root.TryGetProperty("TotalPages", out var totalPagesElement))
                {
                   if (totalPagesElement.TryGetInt32(out var totalPages))
                   {
                       if (page >= totalPages) break;
                   }
                   else
                   {
                       break; 
                   }
                }
                else
                {
                    break; 
                }
            }
            else
            {
                break;
            }

            page++;
        }

        return allPackages;
    }
}
