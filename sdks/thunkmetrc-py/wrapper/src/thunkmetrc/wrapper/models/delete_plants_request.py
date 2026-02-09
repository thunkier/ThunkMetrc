from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class DeletePlantsRequest(TypedDict, total=False):
    ActualDate: str
    Count: int
    Id: int
    Label: str
    ReasonNote: str
    WasteMaterialMixed: str
    WasteMethodName: str
    WasteReasonName: str
    WasteUnitOfMeasureName: str
    WasteWeight: float
