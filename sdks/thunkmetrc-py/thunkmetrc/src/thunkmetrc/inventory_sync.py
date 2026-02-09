import asyncio
from typing import List, Dict, Any, Optional
from datetime import datetime, timedelta

class SynchronizationTarget:
    def __init__(self, license_number: str, wrapper: Any):
        self.license_number = license_number
        self.wrapper = wrapper

async def active_packages_inventory_sync(
    targets: List[SynchronizationTarget],
    last_known_sync: Optional[datetime] = None,
    buffer_minutes: int = 20
) -> Dict[str, List[Any]]:

    end_time = datetime.now()
    if last_known_sync:
        start_time = last_known_sync - timedelta(minutes=buffer_minutes)
    else:
        # Default lookback if no sync time (e.g. 24 hours)
        start_time = end_time - timedelta(days=1)

    # Metrc format often requires ISO 8601 or YYYY-MM-DD. 
    # Wrapper usually expects strings. The generated wrapper signature says `str`.
    # Based on other SDKs, it's often ISO.
    start_str = start_time.isoformat()
    end_str = end_time.isoformat()

    results = {}
    sem = asyncio.Semaphore(20) # Concurrency limit

    async def fetch_target(target):
        async with sem:
             packages = []
             # Method name: packages_get_active_v2
             # Iterator: iterate_packages_get_active_v2
             # Arguments: body, last_modified_end, last_modified_start, license_number, page_number, page_size
             # We must pass arguments as kwargs or positional matching definition.
             # Definition:
             # async def iterate_packages_get_active_v2(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None, body: Any = None) -> AsyncGenerator[Any, None]:
             
             # Note: Argument order depends on template generation. 
             # Path params first (none for this call), then body, then query params sorted.
             # Query params sorted: lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize.
             
             # Safest to use kwargs based on generated signatures I saw earlier.
             # Snake case params: last_modified_end, last_modified_start, license_number ...
             
             iterator = target.wrapper.packages.iterate_get_active_packages(
                 last_modified_end=end_str,
                 last_modified_start=start_str,
                 license_number=target.license_number
             )
             
             async for page in iterator:
                 if isinstance(page, dict) and 'Data' in page and isinstance(page['Data'], list):
                      packages.extend(page['Data'])
                 elif isinstance(page, list):
                      packages.extend(page)

             results[target.license_number] = packages

    await asyncio.gather(*(fetch_target(t) for t in targets))
    return results
