package models

type ItemCategory struct {
    CanBeDecontaminated bool `json:"CanBeDecontaminated,omitempty"`
    CanBeDestroyed bool `json:"CanBeDestroyed,omitempty"`
    CanBeRemediated bool `json:"CanBeRemediated,omitempty"`
    CanContainSeeds bool `json:"CanContainSeeds,omitempty"`
    LabTestBatchNames []interface{} `json:"LabTestBatchNames,omitempty"`
    Name string `json:"Name,omitempty"`
    ProductCategoryType string `json:"ProductCategoryType,omitempty"`
    QuantityType string `json:"QuantityType,omitempty"`
    RequiresAdministrationMethod bool `json:"RequiresAdministrationMethod,omitempty"`
    RequiresAllergens bool `json:"RequiresAllergens,omitempty"`
    RequiresDescription bool `json:"RequiresDescription,omitempty"`
    RequiresItemBrand bool `json:"RequiresItemBrand,omitempty"`
    RequiresLabelPhotoDescription bool `json:"RequiresLabelPhotoDescription,omitempty"`
    RequiresLabelPhotos int64 `json:"RequiresLabelPhotos,omitempty"`
    RequiresNumberOfDoses bool `json:"RequiresNumberOfDoses,omitempty"`
    RequiresPackagingPhotoDescription bool `json:"RequiresPackagingPhotoDescription,omitempty"`
    RequiresPackagingPhotos int64 `json:"RequiresPackagingPhotos,omitempty"`
    RequiresProductPDFDocuments int64 `json:"RequiresProductPDFDocuments,omitempty"`
    RequiresProductPhotoDescription bool `json:"RequiresProductPhotoDescription,omitempty"`
    RequiresProductPhotos int64 `json:"RequiresProductPhotos,omitempty"`
    RequiresPublicIngredients bool `json:"RequiresPublicIngredients,omitempty"`
    RequiresServingSize bool `json:"RequiresServingSize,omitempty"`
    RequiresStrain bool `json:"RequiresStrain,omitempty"`
    RequiresSupplyDurationDays bool `json:"RequiresSupplyDurationDays,omitempty"`
    RequiresUnitCbdAContent bool `json:"RequiresUnitCbdAContent,omitempty"`
    RequiresUnitCbdAContentDose bool `json:"RequiresUnitCbdAContentDose,omitempty"`
    RequiresUnitCbdAPercent bool `json:"RequiresUnitCbdAPercent,omitempty"`
    RequiresUnitCbdContent bool `json:"RequiresUnitCbdContent,omitempty"`
    RequiresUnitCbdContentDose bool `json:"RequiresUnitCbdContentDose,omitempty"`
    RequiresUnitCbdPercent bool `json:"RequiresUnitCbdPercent,omitempty"`
    RequiresUnitThcAContent bool `json:"RequiresUnitThcAContent,omitempty"`
    RequiresUnitThcAContentDose bool `json:"RequiresUnitThcAContentDose,omitempty"`
    RequiresUnitThcAPercent bool `json:"RequiresUnitThcAPercent,omitempty"`
    RequiresUnitThcContent bool `json:"RequiresUnitThcContent,omitempty"`
    RequiresUnitThcContentDose bool `json:"RequiresUnitThcContentDose,omitempty"`
    RequiresUnitThcPercent bool `json:"RequiresUnitThcPercent,omitempty"`
    RequiresUnitVolume bool `json:"RequiresUnitVolume,omitempty"`
    RequiresUnitWeight bool `json:"RequiresUnitWeight,omitempty"`
}
