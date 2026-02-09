from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransferTemplateDelivery(TypedDict, total=False):
    ActualArrivalDateTime: str
    ActualDepartureDateTime: str
    ActualReturnArrivalDateTime: str
    ActualReturnDepartureDateTime: str
    DeliveryNumber: int
    DeliveryPackageCount: int
    DeliveryReceivedPackageCount: int
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
    EstimatedReturnArrivalDateTime: str
    EstimatedReturnDepartureDateTime: str
    GrossUnitOfWeightId: int
    GrossUnitOfWeightName: str
    GrossWeight: float
    Id: int
    InvoiceNumber: str
    ManifestNumber: str
    PDFDocumentFileSystemId: int
    PlannedRoute: str
    ReceivedDateTime: str
    RecipientFacilityLicenseNumber: str
    RecipientFacilityName: str
    RejectedPackagesReturned: bool
    ShipmentTransactionType: str
    ShipmentTypeName: str
