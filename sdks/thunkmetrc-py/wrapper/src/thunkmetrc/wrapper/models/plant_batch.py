from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantBatch(TypedDict, total=False):
    DestroyedCount: int
    Id: int
    IsOnHold: bool
    IsOnInvestigation: bool
    IsOnInvestigationHold: bool
    IsOnInvestigationRecall: bool
    LastModified: str
    LocationId: int
    LocationName: str
    LocationTypeName: str
    MultiPlantBatch: bool
    Name: str
    PackagedCount: int
    PatientLicenseNumber: str
    PlantBatchTypeId: int
    PlantBatchTypeName: str
    PlantedDate: str
    SourcePackageId: int
    SourcePackageLabel: str
    SourcePlantBatchIds: List[Any]
    SourcePlantBatchNames: str
    SourcePlantId: int
    SourcePlantLabel: str
    StrainId: int
    StrainName: str
    SublocationId: int
    SublocationName: str
    TrackedCount: int
    UntrackedCount: int
