package models

type Brand struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
	Status string `json:"Status"`
}

type Category struct {
	CanBeDecontaminated bool `json:"CanBeDecontaminated"`
	CanBeDestroyed bool `json:"CanBeDestroyed"`
	CanBeRemediated bool `json:"CanBeRemediated"`
	CanContainSeeds bool `json:"CanContainSeeds"`
	LabTestBatchNames []string `json:"LabTestBatchNames"`
	Name string `json:"Name"`
	ProductCategoryType string `json:"ProductCategoryType"`
	QuantityType string `json:"QuantityType"`
	RequiresAdministrationMethod bool `json:"RequiresAdministrationMethod"`
	RequiresAllergens bool `json:"RequiresAllergens"`
	RequiresDescription bool `json:"RequiresDescription"`
	RequiresItemBrand bool `json:"RequiresItemBrand"`
	RequiresLabelPhotoDescription bool `json:"RequiresLabelPhotoDescription"`
	RequiresLabelPhotos int `json:"RequiresLabelPhotos"`
	RequiresNumberOfDoses bool `json:"RequiresNumberOfDoses"`
	RequiresPackagingPhotoDescription bool `json:"RequiresPackagingPhotoDescription"`
	RequiresPackagingPhotos int `json:"RequiresPackagingPhotos"`
	RequiresProductPDFDocuments int `json:"RequiresProductPDFDocuments"`
	RequiresProductPhotoDescription bool `json:"RequiresProductPhotoDescription"`
	RequiresProductPhotos int `json:"RequiresProductPhotos"`
	RequiresPublicIngredients bool `json:"RequiresPublicIngredients"`
	RequiresServingSize bool `json:"RequiresServingSize"`
	RequiresStrain bool `json:"RequiresStrain"`
	RequiresSupplyDurationDays bool `json:"RequiresSupplyDurationDays"`
	RequiresUnitCbdAContent bool `json:"RequiresUnitCbdAContent"`
	RequiresUnitCbdAContentDose bool `json:"RequiresUnitCbdAContentDose"`
	RequiresUnitCbdAPercent bool `json:"RequiresUnitCbdAPercent"`
	RequiresUnitCbdContent bool `json:"RequiresUnitCbdContent"`
	RequiresUnitCbdContentDose bool `json:"RequiresUnitCbdContentDose"`
	RequiresUnitCbdPercent bool `json:"RequiresUnitCbdPercent"`
	RequiresUnitThcAContent bool `json:"RequiresUnitThcAContent"`
	RequiresUnitThcAContentDose bool `json:"RequiresUnitThcAContentDose"`
	RequiresUnitThcAPercent bool `json:"RequiresUnitThcAPercent"`
	RequiresUnitThcContent bool `json:"RequiresUnitThcContent"`
	RequiresUnitThcContentDose bool `json:"RequiresUnitThcContentDose"`
	RequiresUnitThcPercent bool `json:"RequiresUnitThcPercent"`
	RequiresUnitVolume bool `json:"RequiresUnitVolume"`
	RequiresUnitWeight bool `json:"RequiresUnitWeight"`
}

type File struct {
	ContentType string `json:"ContentType"`
	FileContents string `json:"FileContents"`
	FileDownloadName string `json:"FileDownloadName"`
}

type FileRequest struct {
	EncodedImageBase64 string `json:"EncodedImageBase64"`
	FileName string `json:"FileName"`
}

type ItemsBrandRequest struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

type ItemsRequest struct {
	AdministrationMethod *string `json:"AdministrationMethod"`
	Allergens *string `json:"Allergens"`
	Description *string `json:"Description"`
	GlobalProductName *string `json:"GlobalProductName"`
	Id int `json:"Id"`
	ItemBrand *string `json:"ItemBrand"`
	ItemCategory string `json:"ItemCategory"`
	ItemIngredients *string `json:"ItemIngredients"`
	LabelImageFileSystemIds *string `json:"LabelImageFileSystemIds"`
	LabelPhotoDescription *string `json:"LabelPhotoDescription"`
	Name string `json:"Name"`
	NumberOfDoses *string `json:"NumberOfDoses"`
	PackagingImageFileSystemIds *string `json:"PackagingImageFileSystemIds"`
	PackagingPhotoDescription *string `json:"PackagingPhotoDescription"`
	ProcessingJobCategoryName *string `json:"ProcessingJobCategoryName"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName"`
	ProductImageFileSystemIds *string `json:"ProductImageFileSystemIds"`
	ProductPDFFileSystemIds *string `json:"ProductPDFFileSystemIds"`
	ProductPhotoDescription *string `json:"ProductPhotoDescription"`
	PublicIngredients *string `json:"PublicIngredients"`
	ServingSize *string `json:"ServingSize"`
	Strain string `json:"Strain"`
	SupplyDurationDays *int `json:"SupplyDurationDays"`
	UnitCbdAContent *float64 `json:"UnitCbdAContent"`
	UnitCbdAContentDose *float64 `json:"UnitCbdAContentDose"`
	UnitCbdAContentDoseUnitOfMeasure *string `json:"UnitCbdAContentDoseUnitOfMeasure"`
	UnitCbdAContentUnitOfMeasure *string `json:"UnitCbdAContentUnitOfMeasure"`
	UnitCbdAPercent *float64 `json:"UnitCbdAPercent"`
	UnitCbdContent *float64 `json:"UnitCbdContent"`
	UnitCbdContentDose *float64 `json:"UnitCbdContentDose"`
	UnitCbdContentDoseUnitOfMeasure *string `json:"UnitCbdContentDoseUnitOfMeasure"`
	UnitCbdContentUnitOfMeasure *string `json:"UnitCbdContentUnitOfMeasure"`
	UnitCbdPercent *float64 `json:"UnitCbdPercent"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UnitThcAContent *float64 `json:"UnitThcAContent"`
	UnitThcAContentDose *float64 `json:"UnitThcAContentDose"`
	UnitThcAContentDoseUnitOfMeasure *string `json:"UnitThcAContentDoseUnitOfMeasure"`
	UnitThcAContentUnitOfMeasure *string `json:"UnitThcAContentUnitOfMeasure"`
	UnitThcAPercent *float64 `json:"UnitThcAPercent"`
	UnitThcContent *float64 `json:"UnitThcContent"`
	UnitThcContentDose *float64 `json:"UnitThcContentDose"`
	UnitThcContentDoseUnitOfMeasure *string `json:"UnitThcContentDoseUnitOfMeasure"`
	UnitThcContentUnitOfMeasure *string `json:"UnitThcContentUnitOfMeasure"`
	UnitThcPercent *float64 `json:"UnitThcPercent"`
	UnitVolume *float64 `json:"UnitVolume"`
	UnitVolumeUnitOfMeasure *string `json:"UnitVolumeUnitOfMeasure"`
	UnitWeight *float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasure *string `json:"UnitWeightUnitOfMeasure"`
}

type Photo struct {
	ContentType string `json:"ContentType"`
	FileContents string `json:"FileContents"`
	FileDownloadName string `json:"FileDownloadName"`
}

type PhotoRequest struct {
	EncodedImageBase64 string `json:"EncodedImageBase64"`
	FileName string `json:"FileName"`
}
