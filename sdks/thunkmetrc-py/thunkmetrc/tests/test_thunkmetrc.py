"""
ThunkMetrc Python SDK Integration Tests

Uses PEP 420 namespace packages - works with editable installs.
"""
import os
import pytest
from pathlib import Path


def get_env(key: str, fallback: str = "") -> str:
    """Get environment variable with fallback."""
    return os.environ.get(key, fallback)


def load_dotenv():
    """Load .env file from repo root."""
    try:
        from dotenv import load_dotenv as _load
        env_path = Path(__file__).parent.parent.parent.parent.parent / ".env"
        if env_path.exists():
            _load(env_path)
    except ImportError:
        pass  # dotenv not installed, rely on env vars


@pytest.fixture
def env_config():
    """Load environment configuration."""
    load_dotenv()
    
    base_url = get_env("METRC_BASE_URL", "https://sandbox-api-or.metrc.com")
    vendor_key = get_env("METRC_VENDOR_API_KEY")
    user_key = get_env("METRC_USER_API_KEY")
    license_number = get_env("METRC_FACILITY_LICENSE_NUMBER")
    
    return {
        "base_url": base_url,
        "vendor_key": vendor_key,
        "user_key": user_key,
        "license_number": license_number
    }


def test_thunkmetrc_inventory_module():
    """Verify inventory module is importable."""
    from thunkmetrc.inventory import active_packages_inventory_sync
    assert active_packages_inventory_sync is not None
    assert callable(active_packages_inventory_sync)


def test_thunkmetrc_wrapper_module():
    """Verify wrapper module is importable."""
    from thunkmetrc.wrapper import MetrcWrapper
    assert MetrcWrapper is not None


def test_thunkmetrc_client_module():
    """Verify client module is importable."""
    from thunkmetrc.client import MetrcClient
    assert MetrcClient is not None


@pytest.mark.asyncio
async def test_active_packages_inventory_sync(env_config):
    """Integration test for active packages inventory sync."""
    if not env_config["vendor_key"] or not env_config["user_key"] or not env_config["license_number"]:
        pytest.skip("Missing METRC_VENDOR_API_KEY, METRC_USER_API_KEY, or METRC_FACILITY_LICENSE_NUMBER")
    
    print(f"\nConfig: {env_config['base_url']}, License: {env_config['license_number']}")
    
    from thunkmetrc.client import MetrcClient
    from thunkmetrc.wrapper import MetrcWrapper
    from thunkmetrc.inventory import active_packages_inventory_sync
    
    # Create the full stack
    client = MetrcClient(env_config["base_url"], env_config["vendor_key"], env_config["user_key"])
    wrapper = MetrcWrapper(client)
    
    print(f"Starting sync for license: {env_config['license_number']}...")
    
    import time
    start = time.time()
    
    try:
        packages = await active_packages_inventory_sync(
            wrapper,
            env_config["license_number"],
            None,
            5
        )
        
        duration = time.time() - start
        package_count = len(packages) if packages else 0
        print(f"Sync completed in {duration:.2f}s. Found {package_count} packages.")
        
        # Just assert we didn't crash. Count might be 0 if sandbox is empty.
        if packages is not None and len(packages) > 0:
            print(f"First package sample: {packages[0]}")
    except Exception as e:
        if "429" in str(e):
            print("WARNING: Test hit rate limit (429). Passing test as environment is restricted.")
        else:
            raise
