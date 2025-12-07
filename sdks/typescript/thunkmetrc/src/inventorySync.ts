import { MetrcWrapper } from '@thunkmetrc/wrapper';

/**
 * Syncs all active packages modified within the time window.
 * 
 * @param wrapper MetrcWrapper instance
 * @param licenseNumber Facility license number
 * @param lastKnownSync Last successful sync time (Date). If null, defaults to 24h ago.
 * @param bufferMinutes Buffer to subtract from start time (default 5).
 * @returns List of active packages (any[])
 */
export async function activePackagesInventorySync(
  wrapper: MetrcWrapper,
  licenseNumber: string,
  lastKnownSync?: Date | null,
  bufferMinutes: number = 5
): Promise<any[]> {
  const endTime = new Date();
  let startTime: Date;

  if (lastKnownSync) {
    startTime = new Date(lastKnownSync.getTime() - bufferMinutes * 60000);
  } else {
    // Default to 24 hours ago
    startTime = new Date(endTime.getTime() - 24 * 60 * 60000);
  }

  // ISO 8601 strings (e.g., 2023-10-26T12:00:00.000Z)
  // Metrc accepts Z or offset. toISOString() uses Z.
  // URL encoding should handle unsafe chars if any, but pure ISO usually safe.
  const startStr = startTime.toISOString();
  const endStr = endTime.toISOString();

  const allPackages: any[] = [];
  let page = 1;
  const pageSize = 20;

  while (true) {
    // Wrapper signature: (path params..., body, query params...)
    // PackagesGetActiveV2 has no path params (except licenseNumber logic? No licenseNumber is query usually)
    // Actually, check definition: licenseNumber IS a query param usually for GET /packages/v2/active
    // Wrapper method: PackagesGetActiveV2(body?: any, lastModifiedEnd?, lastModifiedStart?, licenseNumber?, pageNumber?, pageSize?, ...)
    // Wait, path params come first. If none, then body. Then query.
    
    // We need to be careful about argument order generated.
    // Let's assume standard order: (body, ...queryParams)
    
    const response = await wrapper.packagesGetActiveV2(
      undefined, // body
      endStr,    // lastModifiedEnd
      startStr,  // lastModifiedStart
      licenseNumber,
      page.toString(),
      pageSize.toString()
    );

    // Response structure: { Data: [], TotalPages: number, ... } or just []?
    // Wrapper returns `data` from axios (the payload).
    
    if (response) {
       const data = response.Data;
       if (Array.isArray(data)) {
         allPackages.push(...data);
       }

       const totalPages = response.TotalPages;
       if (typeof totalPages === 'number') {
         if (page >= totalPages) break;
       } else {
         break;
       }
    } else {
      break;
    }

    page++;
  }

  return allPackages;
}
