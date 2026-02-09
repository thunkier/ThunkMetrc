from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Mother(TypedDict, total=False):
    ClonedCount: int
    DescendedCount: int
    DestroyedByUserName: str
    DestroyedDate: str
    DestroyedNote: str
    FloweringDate: str
    GroupTagTypeMax: int
    GrowthPhase: str
    HarvestCount: int
    HarvestId: int
    HarvestedDate: str
    HarvestedUnitOfWeightAbbreviation: str
    HarvestedUnitOfWeightName: str
    HarvestedWetWeight: float
    Id: int
    IsOnHold: bool
    IsOnInvestigation: bool
    IsOnInvestigationHold: bool
    IsOnInvestigationRecall: bool
    Label: str
    LastModified: str
    LocationId: int
    LocationName: str
    LocationTypeName: str
    MotherPlantDate: str
    PatientLicenseNumber: str
    PlantBatchId: int
    PlantBatchName: str
    PlantBatchTypeId: int
    PlantBatchTypeName: str
    PlantedDate: str
    State: str
    StrainId: int
    StrainName: str
    SublocationId: int
    SublocationName: str
    SurvivedCount: int
    TagTypeMax: int
    VegetativeDate: str
