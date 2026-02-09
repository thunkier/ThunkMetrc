from typing import Optional, Dict
import logging
from thunkmetrc.client import MetrcClient
from .wrapper import MetrcWrapper
from .ratelimiter import MetrcRateLimiter, RateLimiterConfig

class MetrcFactory:
    """
    Factory for creating MetrcWrappers that share a common RateLimiter (and optionally HttpClient).
    This is essential for multi-facility applications to respect the global integrator rate limit.
    """
    def __init__(self, 
                 logger: Optional[logging.Logger] = None, 
                 timeout: float = 100.0,
                 max_concurrent_requests: int = 150):
        self.logger = logger
        self.timeout = timeout
        
        # Shared Rate Limiter
        # We assume 150 calls/sec global limit for the integrator
        self.shared_limiter = MetrcRateLimiter(RateLimiterConfig(max_concurrent=max_concurrent_requests))
        
    def create(self, base_url: str, vendor_key: str, user_key: str) -> MetrcWrapper:
        """
        Creates a new Wrapper for a specific environment/user combination.
        It uses a fresh Client but shares the RateLimiter.
        """
        client = MetrcClient(
            base_url=base_url,
            vendor_key=vendor_key,
            user_key=user_key,
            timeout=self.timeout,
            logger=self.logger
        )
        
        # Create Wrapper injecting the shared limiter
        # Note: We need to modify Wrapper to accept an existing limiter instance
        wrapper = MetrcWrapper(client)
        wrapper.limiter = self.shared_limiter
        return wrapper
