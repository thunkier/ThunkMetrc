from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateItemsRequest(TypedDict, total=False):
    AdministrationMethod: str
    Allergens: str
    Description: str
    GlobalProductName: str
    Id: int
    ItemBrand: str
    ItemCategory: str
    ItemIngredients: str
    LabelImageFileSystemIds: str
    LabelPhotoDescription: str
    Name: str
    NumberOfDoses: str
    PackagingImageFileSystemIds: str
    PackagingPhotoDescription: str
    ProcessingJobCategoryName: str
    ProcessingJobTypeName: str
    ProductImageFileSystemIds: str
    ProductPDFFileSystemIds: str
    ProductPhotoDescription: str
    PublicIngredients: str
    ServingSize: str
    Strain: str
    SupplyDurationDays: int
    UnitCbdAContent: float
    UnitCbdAContentDose: float
    UnitCbdAContentDoseUnitOfMeasure: str
    UnitCbdAContentUnitOfMeasure: str
    UnitCbdAPercent: float
    UnitCbdContent: float
    UnitCbdContentDose: float
    UnitCbdContentDoseUnitOfMeasure: str
    UnitCbdContentUnitOfMeasure: str
    UnitCbdPercent: float
    UnitOfMeasure: str
    UnitThcAContent: float
    UnitThcAContentDose: float
    UnitThcAContentDoseUnitOfMeasure: str
    UnitThcAContentUnitOfMeasure: str
    UnitThcAPercent: float
    UnitThcContent: float
    UnitThcContentDose: float
    UnitThcContentDoseUnitOfMeasure: str
    UnitThcContentUnitOfMeasure: str
    UnitThcPercent: float
    UnitVolume: float
    UnitVolumeUnitOfMeasure: str
    UnitWeight: float
    UnitWeightUnitOfMeasure: str
