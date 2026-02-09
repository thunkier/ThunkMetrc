from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PatientCheckInsCreatePatientCheckInsRequest(TypedDict, total=False):
    CheckInDate: str
    CheckInLocationId: int
    PatientLicenseNumber: str
    RegistrationExpiryDate: str
    RegistrationStartDate: str
