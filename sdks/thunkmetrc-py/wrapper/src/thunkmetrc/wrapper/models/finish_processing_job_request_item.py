from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class FinishProcessingJobRequestItem(TypedDict, total=False):
    FinishDate: str
    FinishNote: str
    Id: int
    TotalCountWaste: str
    TotalVolumeWaste: str
    TotalWeightWaste: int
    WasteCountUnitOfMeasureName: str
    WasteVolumeUnitOfMeasureName: str
    WasteWeightUnitOfMeasureName: str
