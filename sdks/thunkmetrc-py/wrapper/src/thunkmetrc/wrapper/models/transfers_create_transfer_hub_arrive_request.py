from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransfersCreateTransferHubArriveRequest(TypedDict, total=False):
    ShipmentDeliveryId: int
    TransporterDirection: str
