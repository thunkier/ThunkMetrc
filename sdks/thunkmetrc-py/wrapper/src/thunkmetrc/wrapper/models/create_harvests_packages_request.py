from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateHarvestsPackagesRequest(TypedDict, total=False):
    ActualDate: str
    DecontaminateProduct: bool
    DecontaminationDate: str
    DecontaminationSteps: str
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
    ProductRequiresDecontamination: bool
    ProductRequiresRemediation: bool
    ProductionBatchNumber: str
    RemediateProduct: bool
    RemediationDate: str
    RemediationMethodId: int
    RemediationSteps: str
    RequiredLabTestBatches: List[Any]
    SellByDate: str
    Sublocation: str
    Tag: str
    UnitOfWeight: str
    UseByDate: str
