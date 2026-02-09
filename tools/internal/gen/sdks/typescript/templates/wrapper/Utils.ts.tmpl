import { PaginatedResponse } from './models';

export async function* iteratePages<T>(
    fetchPage: (page: number) => Promise<PaginatedResponse<T> | T[] | null>
): AsyncGenerator<T, void, unknown> {
    let page = 1;
    while (true) {
        const response = await fetchPage(page);

        if (!response) {
            break;
        }

        let items: T[] = [];
        let totalPages: number | undefined;

        if (Array.isArray(response)) {
            items = response;
            totalPages = 1;
        } else {
            // Check if it has Data property (PaginatedResponse)
            const paginated = response as any; // Cast to access optional props safely
            if (paginated.Data && Array.isArray(paginated.Data)) {
                items = paginated.Data;
                totalPages = paginated.TotalPages;
            } else if (Array.isArray(paginated)) {
                // Fallback if type overlap issues
                items = paginated;
            }
        }

        if (items.length === 0) {
            break;
        }

        for (const item of items) {
            yield item;
        }

        if (typeof totalPages === 'number' && page >= totalPages) {
            break;
        }
        
        // Safety break if we assume 1 page for raw arrays
        if (totalPages === 1) {
            break;
        }

        page++;
    }
}

export async function parallelSync<T>(
    targets: any[], 
    concurrencyLimit: number,
    operation: (target: any) => Promise<T[]>
): Promise<Record<string, T[]>> {
    const results: Record<string, T[]> = {};
    const queue = [...targets];
    
    // Simple semaphore-like worker pool
    const workers = Array(Math.min(concurrencyLimit, targets.length)).fill(null).map(async () => {
        while (queue.length > 0) {
            const target = queue.shift();
            if (!target) break;
            
            try {
                const items = await operation(target);
                const lic = target.licenseNumber || target.license || String(target);
                results[lic] = items;
            } catch (e) {
                console.error(`Error processing target ${target}:`, e);
            }
        }
    });

    await Promise.all(workers);
    return results;
}
