import asyncio
from typing import TypeVar, List, Callable, Awaitable, Tuple, Dict
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

T = TypeVar("T")
C = TypeVar("C")

async def sync_parallel(
    targets: List[Tuple[C, str]],
    concurrency_limit: int,
    operation: Callable[[C, MetrcRateLimiter, str], Awaitable[List[T]]]
) -> Dict[str, List[T]]:
    """
    Executes a synchronization operation in parallel across multiple targets.
    
    Args:
        targets: List of (Client, LicenseNumber) tuples.
        concurrency_limit: Maximum number of concurrent operations.
        operation: Async function taking (Client, RateLimiter, LicenseNumber) and returning List[T].
        
    Returns:
        Dictionary mapping license number to list of results.
    """
    sem = asyncio.Semaphore(concurrency_limit)
    results: Dict[str, List[T]] = {}
    
    async def worker(target: Tuple[C, str]):
        client, license_number = target
        # Assuming client has a rate_limiter property or similar, or we might need to pass it?
        # In Python Wrapper, MetrcWrapper exposes .rate_limiter.
        # But here 'C' is likely MetrcWrapper.
        # Let's assume C has a .rate_limiter property.
        limiter = getattr(client, "rate_limiter", None)
        if limiter is None:
             raise ValueError("Client must have a rate_limiter property")

        async with sem:
            res = await operation(client, limiter, license_number)
            results[license_number] = res

    tasks = [worker(t) for t in targets]
    await asyncio.gather(*tasks)
    
    return results
