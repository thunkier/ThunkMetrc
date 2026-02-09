from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PatientCheckInsUpdatePatientCheckInsRequest(TypedDict, total=False):
    CheckInDate: str
    CheckInLocationId: int
    Id: int
    PatientLicenseNumber: str
    RegistrationExpiryDate: str
    RegistrationStartDate: str
