import { MetrcWrapper } from "@thunkier/thunkmetrc-wrapper";
import pLimit from "p-limit";

export interface SynchronizationTarget {
  licenseNumber: string;
  wrapper: MetrcWrapper;
}

export async function activePackagesInventorySync(
  targets: SynchronizationTarget[],
  lastKnownSync?: Date,
  bufferMinutes: number = 20
): Promise<Map<string, any[]>> {
  
  const endTime = new Date();
  const startTime = lastKnownSync 
    ? new Date(lastKnownSync.getTime() - bufferMinutes * 60000)
    : new Date(endTime.getTime() - 24 * 60 * 60 * 1000);

  const startStr = startTime.toISOString();
  const endStr = endTime.toISOString();

  const results = new Map<string, any[]>();
  const limit = pLimit(20);

  const tasks = targets.map(target => limit(async () => {
    const packages: any[] = [];
    
    const iter = target.wrapper.packages.iterateGetActivePackages(
        undefined, // body
        endStr,    // lastModifiedEnd
        startStr,  // lastModifiedStart
        target.licenseNumber,
        undefined, // pageNumber (ignored by iterator logic which uses local counter)
        ""  // pageSize
    );

    for await (const page of iter) {
        if (page && page.Data && Array.isArray(page.Data)) {
            packages.push(...page.Data);
        } else if (Array.isArray(page)) {
             packages.push(...page);
        }
    }
    results.set(target.licenseNumber, packages);
  }));

  await Promise.all(tasks);
  return results;
}
