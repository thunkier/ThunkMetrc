from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Harvest(TypedDict, total=False):
    ArchivedDate: str
    CurrentWeight: float
    DryingLocationId: int
    DryingLocationName: str
    DryingLocationTypeName: str
    DryingSublocationId: int
    DryingSublocationName: str
    FinishedDate: str
    HarvestStartDate: str
    HarvestType: str
    Id: int
    IsOnHold: bool
    IsOnInvestigation: bool
    IsOnInvestigationHold: bool
    IsOnInvestigationRecall: bool
    LabTestingState: str
    LabTestingStateDate: str
    LastModified: str
    Name: str
    PackageCount: int
    PatientLicenseNumber: str
    PlantCount: int
    SourceStrainCount: int
    SourceStrainNames: str
    Strains: List[Any]
    TotalPackagedWeight: float
    TotalRestoredWeight: float
    TotalWasteWeight: float
    TotalWetWeight: float
    UnitOfWeightName: str
