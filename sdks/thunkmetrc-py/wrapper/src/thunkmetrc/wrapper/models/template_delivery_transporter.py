from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TemplateDeliveryTransporter(TypedDict, total=False):
    ShipmentDeliveryId: int
    TransporterDirection: str
    TransporterFacilityLicenseNumber: str
    TransporterFacilityName: str
