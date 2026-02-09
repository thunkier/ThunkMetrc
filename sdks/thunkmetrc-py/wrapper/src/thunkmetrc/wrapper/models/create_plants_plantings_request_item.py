from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantsPlantingsRequestItem(TypedDict, total=False):
    ActualDate: str
    LocationName: str
    PatientLicenseNumber: str
    PlantBatchName: str
    PlantBatchType: str
    PlantCount: int
    PlantLabel: str
    StrainName: str
    SublocationName: str
