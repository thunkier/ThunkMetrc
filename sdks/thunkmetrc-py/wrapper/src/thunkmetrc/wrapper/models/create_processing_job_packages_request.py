from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateProcessingJobPackagesRequest(TypedDict, total=False):
    ExpirationDate: str
    FinishDate: str
    FinishNote: str
    FinishProcessingJob: bool
    Item: str
    JobName: str
    Location: str
    Note: str
    PackageDate: str
    PatientLicenseNumber: str
    ProductionBatchNumber: str
    Quantity: int
    SellByDate: str
    Sublocation: str
    Tag: str
    UnitOfMeasure: str
    UseByDate: str
    WasteCountQuantity: str
    WasteCountUnitOfMeasureName: str
    WasteVolumeQuantity: str
    WasteVolumeUnitOfMeasureName: str
    WasteWeightQuantity: str
    WasteWeightUnitOfMeasureName: str
