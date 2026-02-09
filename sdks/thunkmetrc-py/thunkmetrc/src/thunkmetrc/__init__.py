# Enable PEP 420 implicit namespace package behavior
from pkgutil import extend_path
__path__ = extend_path(__path__, __name__)

from .inventory_sync import active_packages_inventory_sync, SynchronizationTarget

__all__ = ["active_packages_inventory_sync", "SynchronizationTarget"]
