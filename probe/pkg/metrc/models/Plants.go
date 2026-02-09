package models

type Additive struct {
	AdditiveTypeName *string `json:"AdditiveTypeName"`
	AmountUnitOfMeasure string `json:"AmountUnitOfMeasure"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber *string `json:"EpaRegistrationNumber"`
	Note *string `json:"Note"`
	PlantBatchId *int `json:"PlantBatchId"`
	PlantBatchName *string `json:"PlantBatchName"`
	PlantCount int `json:"PlantCount"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	Rate *string `json:"Rate"`
	RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription"`
	RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	Volume *float64 `json:"Volume"`
}

type AdditivesByLocationRequest struct {
	ActiveIngredients []AdditivesByLocationRequestActiveIngredient `json:"ActiveIngredients"`
	ActualDate string `json:"ActualDate"`
	AdditiveType string `json:"AdditiveType"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber *string `json:"EpaRegistrationNumber"`
	LocationName string `json:"LocationName"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	SublocationName *string `json:"SublocationName"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
}

type AdditivesByLocationRequestActiveIngredient struct {
	Name string `json:"Name"`
	Percentage int `json:"Percentage"`
}

type AdditivesByLocationUsingTemplateRequest struct {
	ActualDate string `json:"ActualDate"`
	AdditivesTemplateName string `json:"AdditivesTemplateName"`
	LocationName string `json:"LocationName"`
	Rate string `json:"Rate"`
	SublocationName *string `json:"SublocationName"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
	Volume string `json:"Volume"`
}

type AdditivesRequest struct {
	ActiveIngredients []AdditivesRequestActiveIngredient `json:"ActiveIngredients"`
	ActualDate string `json:"ActualDate"`
	AdditiveType string `json:"AdditiveType"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber *string `json:"EpaRegistrationNumber"`
	PlantLabels []string `json:"PlantLabels"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
}

type AdditivesRequestActiveIngredient struct {
	Name string `json:"Name"`
	Percentage int `json:"Percentage"`
}

type AdditivesUsingTemplateRequest struct {
	ActualDate string `json:"ActualDate"`
	AdditivesTemplateName string `json:"AdditivesTemplateName"`
	PlantLabels []string `json:"PlantLabels"`
	Rate string `json:"Rate"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
	Volume string `json:"Volume"`
}

type AdjustPlantsRequest struct {
	AdjustCount int `json:"AdjustCount"`
	AdjustReason string `json:"AdjustReason"`
	AdjustmentDate string `json:"AdjustmentDate"`
	Id int `json:"Id"`
	Label *string `json:"Label"`
	ReasonNote string `json:"ReasonNote"`
}

type HarvestRequest struct {
	ActualDate string `json:"ActualDate"`
	DryingLocation string `json:"DryingLocation"`
	DryingSublocation string `json:"DryingSublocation"`
	HarvestName string `json:"HarvestName"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	Plant string `json:"Plant"`
	UnitOfWeight string `json:"UnitOfWeight"`
	Weight float64 `json:"Weight"`
}

type ManicureRequest struct {
	ActualDate string `json:"ActualDate"`
	DryingLocation string `json:"DryingLocation"`
	DryingSublocation string `json:"DryingSublocation"`
	HarvestName *string `json:"HarvestName"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	Plant string `json:"Plant"`
	PlantCount int `json:"PlantCount"`
	UnitOfWeight string `json:"UnitOfWeight"`
	Weight float64 `json:"Weight"`
}

type PlantBatchPackagesRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	IsDonation bool `json:"IsDonation"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	Location string `json:"Location"`
	Note *string `json:"Note"`
	PackageTag string `json:"PackageTag"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlantBatchType string `json:"PlantBatchType"`
	PlantLabel string `json:"PlantLabel"`
	Sublocation *string `json:"Sublocation"`
}

type PlantsGrowthPhaseRequest struct {
	GrowthDate string `json:"GrowthDate"`
	GrowthPhase string `json:"GrowthPhase"`
	Id *int `json:"Id"`
	Label string `json:"Label"`
	NewLocation string `json:"NewLocation"`
	NewSublocation *string `json:"NewSublocation"`
	NewTag string `json:"NewTag"`
}

type PlantsLocationRequest struct {
	ActualDate string `json:"ActualDate"`
	Id *int `json:"Id"`
	Label string `json:"Label"`
	Location string `json:"Location"`
	Sublocation string `json:"Sublocation"`
}

type PlantsMergeRequest struct {
	MergeDate string `json:"MergeDate"`
	SourcePlantGroupLabel string `json:"SourcePlantGroupLabel"`
	TargetPlantGroupLabel string `json:"TargetPlantGroupLabel"`
}

type PlantsPlantingsRequest struct {
	ActualDate string `json:"ActualDate"`
	LocationName string `json:"LocationName"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PlantBatchName string `json:"PlantBatchName"`
	PlantBatchType string `json:"PlantBatchType"`
	PlantCount int `json:"PlantCount"`
	PlantLabel string `json:"PlantLabel"`
	StrainName string `json:"StrainName"`
	SublocationName string `json:"SublocationName"`
}

type PlantsRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	Id *int `json:"Id"`
	Label string `json:"Label"`
	ReasonNote string `json:"ReasonNote"`
	WasteMaterialMixed string `json:"WasteMaterialMixed"`
	WasteMethodName string `json:"WasteMethodName"`
	WasteReasonName string `json:"WasteReasonName"`
	WasteUnitOfMeasureName string `json:"WasteUnitOfMeasureName"`
	WasteWeight float64 `json:"WasteWeight"`
}

type PlantsSplitRequest struct {
	PlantCount int `json:"PlantCount"`
	SourcePlantLabel string `json:"SourcePlantLabel"`
	SplitDate string `json:"SplitDate"`
	StrainLabel string `json:"StrainLabel"`
	TagLabel string `json:"TagLabel"`
}

type PlantsWaste struct {
	GrowthPhase int `json:"GrowthPhase"`
	Label string `json:"Label"`
	LocationId int `json:"LocationId"`
	LocationName string `json:"LocationName"`
	PlantId int `json:"PlantId"`
	PlantWasteId int `json:"PlantWasteId"`
	StateName string `json:"StateName"`
	SublocationId int `json:"SublocationId"`
	SublocationName *string `json:"SublocationName"`
	WasteAmount int `json:"WasteAmount"`
	WasteUnitOfMeasureAbbreviation string `json:"WasteUnitOfMeasureAbbreviation"`
}

type PlantsWasteRequest struct {
	LocationName string `json:"LocationName"`
	MixedMaterial string `json:"MixedMaterial"`
	Note string `json:"Note"`
	PlantLabels []interface{} `json:"PlantLabels"`
	ReasonName string `json:"ReasonName"`
	SublocationName string `json:"SublocationName"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
	WasteDate string `json:"WasteDate"`
	WasteMethodName string `json:"WasteMethodName"`
	WasteWeight float64 `json:"WasteWeight"`
}

type StrainRequest struct {
	Id int `json:"Id"`
	Label *string `json:"Label"`
	StrainId int `json:"StrainId"`
	StrainName *string `json:"StrainName"`
}

type TagRequest struct {
	Id *int `json:"Id"`
	Label string `json:"Label"`
	NewTag string `json:"NewTag"`
	ReplaceDate string `json:"ReplaceDate"`
	TagId *int `json:"TagId"`
}

type WastePackage struct {
	ArchivedDate *string `json:"ArchivedDate"`
	ContainsDecontaminatedProduct bool `json:"ContainsDecontaminatedProduct"`
	ContainsRemediatedProduct bool `json:"ContainsRemediatedProduct"`
	DecontaminationDate *string `json:"DecontaminationDate"`
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId *int `json:"ExternalId"`
	FinishedDate *string `json:"FinishedDate"`
	Id int `json:"Id"`
	InitialLabTestingState string `json:"InitialLabTestingState"`
	IsDonation bool `json:"IsDonation"`
	IsDonationPersistent bool `json:"IsDonationPersistent"`
	IsFinished bool `json:"IsFinished"`
	IsOnHold bool `json:"IsOnHold"`
	IsOnHoldCombined bool `json:"IsOnHoldCombined"`
	IsOnInvestigation bool `json:"IsOnInvestigation"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	IsOnRecall *bool `json:"IsOnRecall"`
	IsOnRecallCombined bool `json:"IsOnRecallCombined"`
	IsOnRetailerDelivery bool `json:"IsOnRetailerDelivery"`
	IsProcessValidationTestingSample bool `json:"IsProcessValidationTestingSample"`
	IsProductionBatch bool `json:"IsProductionBatch"`
	IsTestingSample bool `json:"IsTestingSample"`
	IsTradeSample bool `json:"IsTradeSample"`
	IsTradeSamplePersistent bool `json:"IsTradeSamplePersistent"`
	Item WastePackageItem `json:"Item"`
	ItemFromFacilityLicenseNumber *string `json:"ItemFromFacilityLicenseNumber"`
	ItemFromFacilityName *string `json:"ItemFromFacilityName"`
	LabTestResultExpirationDateTime *string `json:"LabTestResultExpirationDateTime"`
	LabTestStage *string `json:"LabTestStage"`
	LabTestStageId *int `json:"LabTestStageId"`
	LabTestingPerformedDate *string `json:"LabTestingPerformedDate"`
	LabTestingRecordedDate *string `json:"LabTestingRecordedDate"`
	LabTestingState string `json:"LabTestingState"`
	LabTestingStateDate string `json:"LabTestingStateDate"`
	Label string `json:"Label"`
	LabelsLastGeneratedDateTime *string `json:"LabelsLastGeneratedDateTime"`
	LastModified string `json:"LastModified"`
	LocationId *int `json:"LocationId"`
	LocationName *string `json:"LocationName"`
	LocationTypeName *string `json:"LocationTypeName"`
	Note *string `json:"Note"`
	OriginalPackageQuantity float64 `json:"OriginalPackageQuantity"`
	PackageForProductDestruction *string `json:"PackageForProductDestruction"`
	PackageType string `json:"PackageType"`
	PackagedDate string `json:"PackagedDate"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	ProductLabel WastePackageProductLabel `json:"ProductLabel"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	Quantity float64 `json:"Quantity"`
	ReceivedDateTime *string `json:"ReceivedDateTime"`
	ReceivedFromFacilityLicenseNumber *string `json:"ReceivedFromFacilityLicenseNumber"`
	ReceivedFromFacilityName *string `json:"ReceivedFromFacilityName"`
	ReceivedFromManifestNumber *string `json:"ReceivedFromManifestNumber"`
	RemediationDate *string `json:"RemediationDate"`
	SellByDate *string `json:"SellByDate"`
	SourceHarvestCount int `json:"SourceHarvestCount"`
	SourceHarvestNames *string `json:"SourceHarvestNames"`
	SourcePackageCount int `json:"SourcePackageCount"`
	SourcePackageIsDonation bool `json:"SourcePackageIsDonation"`
	SourcePackageIsTradeSample bool `json:"SourcePackageIsTradeSample"`
	SourcePackageLabels *string `json:"SourcePackageLabels"`
	SourceProcessingJobCount int `json:"SourceProcessingJobCount"`
	SourceProductionBatchNumbers *string `json:"SourceProductionBatchNumbers"`
	SublocationId *int `json:"SublocationId"`
	SublocationName *string `json:"SublocationName"`
	UnitOfMeasureAbbreviation string `json:"UnitOfMeasureAbbreviation"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
	UseByDate *string `json:"UseByDate"`
}

type WastePackageItem struct {
	AdministrationMethod *string `json:"AdministrationMethod"`
	Allergens *string `json:"Allergens"`
	ApprovalStatus int `json:"ApprovalStatus"`
	ApprovalStatusDateTime string `json:"ApprovalStatusDateTime"`
	DefaultLabTestingState int `json:"DefaultLabTestingState"`
	Description *string `json:"Description"`
	GlobalProductName *string `json:"GlobalProductName"`
	GlobalProductNumber *string `json:"GlobalProductNumber"`
	HasExpirationDate bool `json:"HasExpirationDate"`
	HasSellByDate bool `json:"HasSellByDate"`
	HasUseByDate bool `json:"HasUseByDate"`
	Id int `json:"Id"`
	IsExpirationDateRequired bool `json:"IsExpirationDateRequired"`
	IsSellByDateRequired bool `json:"IsSellByDateRequired"`
	IsUseByDateRequired bool `json:"IsUseByDateRequired"`
	IsUsed bool `json:"IsUsed"`
	ItemBrandId int `json:"ItemBrandId"`
	ItemBrandName *string `json:"ItemBrandName"`
	LabTestBatchNames []interface{} `json:"LabTestBatchNames"`
	LabelImages []interface{} `json:"LabelImages"`
	LabelPhotoDescription *string `json:"LabelPhotoDescription"`
	Name string `json:"Name"`
	NumberOfDoses *string `json:"NumberOfDoses"`
	PackagingImages []interface{} `json:"PackagingImages"`
	PackagingPhotoDescription *string `json:"PackagingPhotoDescription"`
	ProcessingJobCategoryName *string `json:"ProcessingJobCategoryName"`
	ProcessingJobTypeName *string `json:"ProcessingJobTypeName"`
	ProductBrandName *string `json:"ProductBrandName"`
	ProductCategoryName string `json:"ProductCategoryName"`
	ProductCategoryType int `json:"ProductCategoryType"`
	ProductImages []interface{} `json:"ProductImages"`
	ProductPDFDocuments []interface{} `json:"ProductPDFDocuments"`
	ProductPhotoDescription *string `json:"ProductPhotoDescription"`
	PublicIngredients *string `json:"PublicIngredients"`
	QuantityType int `json:"QuantityType"`
	ServingSize *string `json:"ServingSize"`
	StrainId *int `json:"StrainId"`
	StrainName *string `json:"StrainName"`
	UnitCbdAContent *float64 `json:"UnitCbdAContent"`
	UnitCbdAContentDose *float64 `json:"UnitCbdAContentDose"`
	UnitCbdAContentDoseUnitOfMeasureName *string `json:"UnitCbdAContentDoseUnitOfMeasureName"`
	UnitCbdAContentUnitOfMeasureName *string `json:"UnitCbdAContentUnitOfMeasureName"`
	UnitCbdAPercent *float64 `json:"UnitCbdAPercent"`
	UnitCbdContent *float64 `json:"UnitCbdContent"`
	UnitCbdContentDose *float64 `json:"UnitCbdContentDose"`
	UnitCbdContentDoseUnitOfMeasureName *string `json:"UnitCbdContentDoseUnitOfMeasureName"`
	UnitCbdContentUnitOfMeasureName *string `json:"UnitCbdContentUnitOfMeasureName"`
	UnitCbdPercent *float64 `json:"UnitCbdPercent"`
	UnitOfMeasureName *string `json:"UnitOfMeasureName"`
	UnitQuantity *float64 `json:"UnitQuantity"`
	UnitQuantityUnitOfMeasureName *string `json:"UnitQuantityUnitOfMeasureName"`
	UnitThcAContent *float64 `json:"UnitThcAContent"`
	UnitThcAContentDose *float64 `json:"UnitThcAContentDose"`
	UnitThcAContentDoseUnitOfMeasureId *int `json:"UnitThcAContentDoseUnitOfMeasureId"`
	UnitThcAContentDoseUnitOfMeasureName *string `json:"UnitThcAContentDoseUnitOfMeasureName"`
	UnitThcAContentUnitOfMeasureName *string `json:"UnitThcAContentUnitOfMeasureName"`
	UnitThcAPercent *float64 `json:"UnitThcAPercent"`
	UnitThcContent *float64 `json:"UnitThcContent"`
	UnitThcContentDose *float64 `json:"UnitThcContentDose"`
	UnitThcContentDoseUnitOfMeasureId *int `json:"UnitThcContentDoseUnitOfMeasureId"`
	UnitThcContentDoseUnitOfMeasureName *string `json:"UnitThcContentDoseUnitOfMeasureName"`
	UnitThcContentUnitOfMeasureName *string `json:"UnitThcContentUnitOfMeasureName"`
	UnitThcPercent *float64 `json:"UnitThcPercent"`
	UnitVolume *float64 `json:"UnitVolume"`
	UnitVolumeUnitOfMeasureName *string `json:"UnitVolumeUnitOfMeasureName"`
	UnitWeight *float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasureName *string `json:"UnitWeightUnitOfMeasureName"`
}

type WastePackageProductLabel struct {
	IsActive bool `json:"IsActive"`
	IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel"`
	IsDiscontinued bool `json:"IsDiscontinued"`
	LastLabelGenerationDate *string `json:"LastLabelGenerationDate"`
	PackageId int `json:"PackageId"`
	QrCount int `json:"QrCount"`
	TagId *int `json:"TagId"`
}

type WasteReason struct {
	Name string `json:"Name"`
	RequiresImmatureWasteWeight bool `json:"RequiresImmatureWasteWeight"`
	RequiresMatureWasteWeight bool `json:"RequiresMatureWasteWeight"`
	RequiresNote bool `json:"RequiresNote"`
	RequiresWasteWeight bool `json:"RequiresWasteWeight"`
}
