from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class DeliveryPackageWholesale(TypedDict, total=False):
    PackageId: int
    PackageLabel: str
    ReceiverWholesalePrice: str
    ShipperWholesalePrice: str
