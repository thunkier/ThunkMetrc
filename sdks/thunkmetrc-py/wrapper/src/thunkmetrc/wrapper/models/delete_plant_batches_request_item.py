from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class DeletePlantBatchesRequestItem(TypedDict, total=False):
    ActualDate: str
    Count: int
    PlantBatch: str
    ReasonNote: str
    WasteMaterialMixed: str
    WasteMethodName: str
    WasteReasonName: str
    WasteUnitOfMeasure: str
    WasteWeight: float
