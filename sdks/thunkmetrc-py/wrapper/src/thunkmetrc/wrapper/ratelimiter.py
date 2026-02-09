import asyncio
import time
import random
import math
import re
from typing import Optional, Dict, Any, Callable

class RateLimiterConfig:
    def __init__(self, 
                 enabled: bool = False,
                 max_get_per_second_per_facility: float = 50,
                 max_get_per_second_integrator: float = 150,
                 max_concurrent_get_per_facility: int = 10,
                 max_concurrent_get_integrator: int = 30,
                 max_retries: int = 5,
                 initial_backoff: float = 0.5,
                 max_backoff: float = 10.0,
                 backoff_exponent: float = 1.5):
        self.enabled = enabled
        self.max_get_per_second_per_facility = max_get_per_second_per_facility
        self.max_get_per_second_integrator = max_get_per_second_integrator
        self.max_concurrent_get_per_facility = max_concurrent_get_per_facility
        self.max_concurrent_get_integrator = max_concurrent_get_integrator
        self.max_retries = max_retries
        self.initial_backoff = initial_backoff
        self.max_backoff = max_backoff
        self.backoff_exponent = backoff_exponent

class TokenBucket:
    def __init__(self, rate: float, capacity: float):
        self.rate = rate
        self.capacity = capacity
        self.tokens = capacity
        self.last_refill = time.monotonic()
        self.lock = asyncio.Lock()

    async def wait(self):
        async with self.lock:
            now = time.monotonic()
            elapsed = now - self.last_refill
            self.tokens = min(self.capacity, self.tokens + elapsed * self.rate)
            self.last_refill = now

            if self.tokens >= 1:
                self.tokens -= 1
                return
            
            # Wait time
            missing = 1.0 - self.tokens
            wait_time = missing / self.rate
        
        await asyncio.sleep(wait_time)
        await self.wait() # Recurse to re-check

class MetrcRateLimiter:
    def __init__(self, config: Optional[RateLimiterConfig] = None):
        if config is None:
            config = RateLimiterConfig()
        self.config = config
        
        self.integrator_rate = TokenBucket(config.max_get_per_second_integrator, config.max_get_per_second_integrator)
        self.integrator_sem = asyncio.Semaphore(config.max_concurrent_get_integrator)
        
        self.facility_rates: Dict[str, TokenBucket] = {}
        self.facility_sems: Dict[str, asyncio.Semaphore] = {}
        self.lock = asyncio.Lock()

    async def execute(self, facility: Optional[str], is_get: bool, op: Callable[[], Any]):
        if not self.config.enabled or not is_get:
            return await op()

        async with self.integrator_sem:
            async with self.get_facility_sem(facility):
                await self.integrator_rate.wait()
                if facility:
                    await self.get_facility_rate(facility).wait()
                
                # Robust Retry Loop
                retries = 0
                while True:
                    try:
                        return await op()
                    except Exception as e:
                        msg = str(e).lower()
                        # Check for 429
                        if "429" in msg or "too many requests" in msg:
                            # Try to find Retry-After in message (hacky if no header access)
                            # Assuming e might be a specialized error or just a string
                            retry_after = 0.0
                            match = re.search(r"retry-after[:\s]+(\d+)", msg)
                            if match:
                                retry_after = float(match.group(1))
                            
                            if retries >= self.config.max_retries:
                                raise e
                            
                            wait = retry_after
                            if wait <= 0:
                                # Exponential backoff + Jitter
                                backoff = self.config.initial_backoff * (self.config.backoff_exponent ** retries)
                                backoff = min(backoff, self.config.max_backoff)
                                jitter = random.uniform(0, backoff * 0.1)
                                wait = backoff + jitter
                            
                            await asyncio.sleep(wait)
                            retries += 1
                            continue
                        
                        # 5xx Errors logic could go here
                        
                        raise e

    def get_facility_rate(self, facility: str) -> TokenBucket:
        if facility not in self.facility_rates:
             self.facility_rates[facility] = TokenBucket(self.config.max_get_per_second_per_facility, self.config.max_get_per_second_per_facility)
        return self.facility_rates[facility]

    def get_facility_sem(self, facility: Optional[str]) -> asyncio.Semaphore:
        if not facility:
             return asyncio.Semaphore(9999) 
        if facility not in self.facility_sems:
             self.facility_sems[facility] = asyncio.Semaphore(self.config.max_concurrent_get_per_facility)
        return self.facility_sems[facility]
