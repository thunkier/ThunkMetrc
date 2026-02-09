from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateDeliveriesRetailerEndRequest(TypedDict, total=False):
    ActualArrivalDateTime: str
    Packages: List[Any]
    RetailerDeliveryId: int
