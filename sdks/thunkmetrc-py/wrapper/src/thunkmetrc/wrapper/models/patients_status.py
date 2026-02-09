from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PatientsStatus(TypedDict, total=False):
    Active: bool
    ConcentrateOuncesAllowed: int
    ConcentrateOuncesAvailable: int
    ConcentrateOuncesPurchased: int
    FlowerOuncesAllowed: int
    FlowerOuncesAvailable: int
    FlowerOuncesPurchased: int
    IdentificationMethod: str
    InfusedOuncesAllowed: int
    InfusedOuncesAvailable: int
    InfusedOuncesPurchased: int
    PatientLicenseExpirationDate: str
    PatientLicenseNumber: str
    PatientRegistrationStartDate: str
    PurchaseAmountDays: int
    RegistrationNumber: str
    ThcOuncesAllowed: int
    ThcOuncesAvailable: int
    ThcOuncesPurchased: int
