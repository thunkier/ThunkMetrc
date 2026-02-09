from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateDeliveriesCompleteRequestItem(TypedDict, total=False):
    AcceptedPackages: List[Any]
    ActualArrivalDateTime: str
    Id: int
    PaymentType: str
    ReturnedPackages: List[Any]
