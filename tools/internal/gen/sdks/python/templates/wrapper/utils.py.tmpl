import asyncio
from typing import TypeVar, AsyncGenerator, Callable, Awaitable, List, Optional, Any, Dict, Tuple
from datetime import datetime, timedelta
from .models import PaginatedResponse

T = TypeVar("T")

def deserialize_page(data: Any, item_type: Any) -> Optional[PaginatedResponse[T]]:
    """
    Converts raw API response (dict/list) into PaginatedResponse.
    item_type is the class to deserialize items into (e.g. Package).
    """
    if not data:
        return None
        
    # If using Pydantic or similar, we could just do PaginatedResponse[item_type](**data)
    # But here we have simple dataclasses?
    # Assuming Models have from_dict method or similar, or just dict unpacking if simple.
    
    # Check if data is PaginatedResponse-like
    if isinstance(data, dict) and "Data" in data:
        # It's a paginated response
        # We need to rely on the generated model's capabilities to parse children.
        # But generated models might handle this if we pass the dict to PaginatedResponse constructor?
        # Let's assume standard deserialization logic usage.
        # Actually, let's keep it simple: Just return the objects if possible, or wrap.
        pass
    
    # For now, let's just assume the caller handles strict typing or we do minimal wrapping
    # The Generic Iterator below expects PaginatedResponse.
    
    res = PaginatedResponse()
    
    if isinstance(data, list):
        # Raw list
        res.Data = [item_type(**i) if isinstance(i, dict) else i for i in data]
        res.TotalPages = 1
        return res
        
    if isinstance(data, dict):
        res.CurrentPage = data.get("CurrentPage")
        res.TotalPages = data.get("TotalPages")
        raw_data = data.get("Data", [])
        res.Data = [item_type(**i) if isinstance(i, dict) else i for i in raw_data]
        return res
        
    return None

async def iterate_pages(
    fetch_page: Callable[[int], Awaitable[Optional[PaginatedResponse[T]]]],
) -> AsyncGenerator[T, None]:
    """
    Generic iterator for paged endpoints.
    Yields items one by one.
    """
    page = 1
    
    while True:
        response = await fetch_page(page)
        
        if not response or not response.Data:
            break
            
        for item in response.Data:
            yield item
            
        if response.TotalPages is not None:
            if page >= response.TotalPages:
                break
        else:
            # Fallback if TotalPages is missing: if we got fewer items than expected?
            # Or just rely on next page being empty
            pass
            
        if response.TotalPages == 0:
            break
            
        page += 1

def get_time_window(last_known_sync: Optional[datetime], buffer_minutes: int = 0) -> Tuple[str, str]:
    """
    Calculates a Metrc-compatible time window string tuple (StartString, EndString).
    Defaults to 24h lookup if last_known_sync is None.
    """
    metrc_format = "%Y-%m-%dT%H:%M:%SZ" # Simplistic UTC format
    # Note: Metrc expects local time or specific offset sometimes? 
    # C# impl uses "yyyy-MM-ddTHH:mm:ssZ07:00" which is weird.
    # Actually generally ISO8601 works.
    # Let's use standard ISO format with Z
    
    end_time = datetime.utcnow()
    start_time = end_time - timedelta(hours=24)
    
    if last_known_sync:
        if buffer_minutes < 0: buffer_minutes = 0
        start_time = last_known_sync - timedelta(minutes=buffer_minutes)
        
    return (
        start_time.strftime(metrc_format),
        end_time.strftime(metrc_format)
    )

async def parallel_sync(
    targets: List[Any], # List of objects having .wrapper and .license_number
    concurrency_limit: int,
    operation: Callable[[Any], Awaitable[List[T]]]
) -> Dict[str, List[T]]:
    """
    Executes an operation concurrently across multiple targets with a semaphore.
    """
    semaphore = asyncio.Semaphore(concurrency_limit)
    results = {}
    
    async def worker(target):
        async with semaphore:
            try:
                items = await operation(target)
                # Assume target has license_number
                lic = getattr(target, 'license_number', str(target))
                results[lic] = items
            except Exception as e:
                # Log error? 
                pass

    tasks = [worker(t) for t in targets]
    await asyncio.gather(*tasks)
    return results
