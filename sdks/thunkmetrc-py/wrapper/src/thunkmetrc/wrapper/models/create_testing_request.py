from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateTestingRequest(TypedDict, total=False):
    ActualDate: str
    ExpirationDate: str
    Ingredients: List[Any]
    IsDonation: bool
    IsProductionBatch: bool
    IsTradeSample: bool
    Item: str
    LabTestStageId: int
    Location: str
    Note: str
    PatientLicenseNumber: str
    ProcessingJobTypeId: int
    ProductRequiresRemediation: bool
    ProductionBatchNumber: str
    Quantity: int
    RequiredLabTestBatches: List[Any]
    SellByDate: str
    Sublocation: str
    Tag: str
    UnitOfMeasure: str
    UseByDate: str
    UseSameItem: bool
