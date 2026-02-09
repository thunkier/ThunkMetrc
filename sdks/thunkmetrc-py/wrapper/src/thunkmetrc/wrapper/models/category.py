from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Category(TypedDict, total=False):
    CanBeDecontaminated: bool
    CanBeDestroyed: bool
    CanBeRemediated: bool
    CanContainSeeds: bool
    LabTestBatchNames: List[Any]
    Name: str
    ProductCategoryType: str
    QuantityType: str
    RequiresAdministrationMethod: bool
    RequiresAllergens: bool
    RequiresDescription: bool
    RequiresItemBrand: bool
    RequiresLabelPhotoDescription: bool
    RequiresLabelPhotos: int
    RequiresNumberOfDoses: bool
    RequiresPackagingPhotoDescription: bool
    RequiresPackagingPhotos: int
    RequiresProductPDFDocuments: int
    RequiresProductPhotoDescription: bool
    RequiresProductPhotos: int
    RequiresPublicIngredients: bool
    RequiresServingSize: bool
    RequiresStrain: bool
    RequiresSupplyDurationDays: bool
    RequiresUnitCbdAContent: bool
    RequiresUnitCbdAContentDose: bool
    RequiresUnitCbdAPercent: bool
    RequiresUnitCbdContent: bool
    RequiresUnitCbdContentDose: bool
    RequiresUnitCbdPercent: bool
    RequiresUnitThcAContent: bool
    RequiresUnitThcAContentDose: bool
    RequiresUnitThcAPercent: bool
    RequiresUnitThcContent: bool
    RequiresUnitThcContentDose: bool
    RequiresUnitThcPercent: bool
    RequiresUnitVolume: bool
    RequiresUnitWeight: bool
