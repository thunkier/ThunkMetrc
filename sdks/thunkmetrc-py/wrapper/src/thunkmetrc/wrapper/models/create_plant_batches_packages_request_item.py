from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantBatchesPackagesRequestItem(TypedDict, total=False):
    ActualDate: str
    Count: int
    ExpirationDate: str
    Id: int
    IsDonation: bool
    IsTradeSample: bool
    Item: str
    Location: str
    Note: str
    PatientLicenseNumber: str
    PlantBatch: str
    SellByDate: str
    Sublocation: str
    Tag: str
    UseByDate: str
