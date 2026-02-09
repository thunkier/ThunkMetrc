from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PackagesCreatePackagePlantingsRequest(TypedDict, total=False):
    LocationName: str
    PackageAdjustmentAmount: int
    PackageAdjustmentUnitOfMeasureName: str
    PackageLabel: str
    PatientLicenseNumber: str
    PlantBatchName: str
    PlantBatchType: str
    PlantCount: int
    PlantedDate: str
    StrainName: str
    SublocationName: str
    UnpackagedDate: str
