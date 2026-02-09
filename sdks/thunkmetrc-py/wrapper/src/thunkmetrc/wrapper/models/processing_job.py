from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ProcessingJob(TypedDict, total=False):
    CountUnitOfMeasureAbbreviation: str
    CountUnitOfMeasureId: int
    CountUnitOfMeasureName: str
    FinishNote: str
    FinishedDate: str
    Id: int
    IsFinished: bool
    JobTypeId: int
    JobTypeName: str
    Name: str
    Number: str
    Packages: List[Any]
    StartDate: str
    TotalCount: int
    TotalCountWaste: str
    TotalQuantity: float
    TotalUnitOfMeasureId: int
    TotalVolume: float
    TotalVolumeWaste: str
    TotalWeight: float
    TotalWeightWaste: str
    VolumeUnitOfMeasureAbbreviation: str
    VolumeUnitOfMeasureId: int
    VolumeUnitOfMeasureName: str
    WasteCountUnitOfMeasureAbbreviation: str
    WasteCountUnitOfMeasureId: int
    WasteCountUnitOfMeasureName: str
    WasteVolumeUnitOfMeasureAbbreviation: str
    WasteVolumeUnitOfMeasureId: int
    WasteVolumeUnitOfMeasureName: str
    WasteWeightUnitOfMeasureAbbreviation: str
    WasteWeightUnitOfMeasureId: int
    WasteWeightUnitOfMeasureName: str
    WeightUnitOfMeasureAbbreviation: str
    WeightUnitOfMeasureId: int
    WeightUnitOfMeasureName: str
