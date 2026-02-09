from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantBatchPackagesRequestItem(TypedDict, total=False):
    ActualDate: str
    Count: int
    IsDonation: bool
    IsTradeSample: bool
    Item: str
    Location: str
    Note: str
    PackageTag: str
    PatientLicenseNumber: str
    PlantBatchType: str
    PlantLabel: str
    Sublocation: str
