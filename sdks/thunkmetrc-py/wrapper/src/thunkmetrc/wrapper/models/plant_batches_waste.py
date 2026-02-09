from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantBatchesWaste(TypedDict, total=False):
    PlantBatchId: int
    PlantBatchName: str
    PlantCount: int
    PlantWasteNumber: str
    WasteDate: str
    WasteMethodName: str
    WasteReasonName: str
    WasteUnitOfMeasureName: str
    WasteWeight: float
