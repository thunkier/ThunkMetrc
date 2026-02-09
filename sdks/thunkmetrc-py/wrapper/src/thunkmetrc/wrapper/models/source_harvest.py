from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class SourceHarvest(TypedDict, total=False):
    HarvestStartDate: str
    HarvestedByFacilityLicenseNumber: str
    HarvestedByFacilityName: str
    Name: str
