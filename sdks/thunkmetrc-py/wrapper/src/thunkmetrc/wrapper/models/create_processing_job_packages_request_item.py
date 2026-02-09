from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateProcessingJobPackagesRequestItem(TypedDict, total=False):
    ExpirationDate: str
    FinishDate: str
    FinishNote: str
    FinishProcessingJob: bool
    IsFinishedGood: bool
    Item: str
    JobName: str
    Location: str
    Note: str
    PackageDate: str
    PatientLicenseNumber: str
    ProductionBatchNumber: str
    Quantity: float
    SellByDate: str
    Sublocation: str
    Tag: str
    UnitOfMeasure: str
    UseByDate: str
    WasteCountQuantity: float
    WasteCountUnitOfMeasureName: str
    WasteVolumeQuantity: float
    WasteVolumeUnitOfMeasureName: str
    WasteWeightQuantity: float
    WasteWeightUnitOfMeasureName: str
