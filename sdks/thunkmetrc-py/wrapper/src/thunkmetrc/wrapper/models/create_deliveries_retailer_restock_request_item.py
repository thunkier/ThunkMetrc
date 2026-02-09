from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateDeliveriesRetailerRestockRequestItem(TypedDict, total=False):
    DateTime: str
    Destinations: str
    EstimatedDepartureDateTime: str
    Packages: List[Any]
    RetailerDeliveryId: int
