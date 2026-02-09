from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PackagesPackage(TypedDict, total=False):
    ActualDepartureDateTime: str
    ExternalId: int
    GrossUnitOfWeightAbbreviation: str
    GrossWeight: float
    Id: int
    ItemStrainName: str
    LabTestingStateName: str
    ManifestNumber: str
    PackageId: int
    PackageLabel: str
    ProcessingJobTypeName: str
    ProductCategoryName: str
    ProductName: str
    ReceivedDateTime: str
    ReceivedQuantity: float
    ReceivedUnitOfMeasureAbbreviation: str
    ReceiverWholesalePrice: float
    RecipientFacilityLicenseNumber: str
    RecipientFacilityName: str
    ShipmentPackageStateName: str
    ShippedQuantity: float
    ShippedUnitOfMeasureAbbreviation: str
    ShipperWholesalePrice: float
    SourceHarvestNames: str
    SourcePackageLabels: str
