package models

type AdjustReason struct {
	Name string `json:"Name"`
	RequiresImmatureWasteWeight bool `json:"RequiresImmatureWasteWeight"`
	RequiresMatureWasteWeight bool `json:"RequiresMatureWasteWeight"`
	RequiresNote bool `json:"RequiresNote"`
	RequiresWasteWeight bool `json:"RequiresWasteWeight"`
}

type DecontaminateRequest struct {
	DecontaminationDate string `json:"DecontaminationDate"`
	DecontaminationMethodName string `json:"DecontaminationMethodName"`
	DecontaminationSteps string `json:"DecontaminationSteps"`
	PackageLabel string `json:"PackageLabel"`
}

type DonationFlagRequest struct {
	Label string `json:"Label"`
}

type DonationUnflagRequest struct {
	Label string `json:"Label"`
}

type ExternalidRequest struct {
	ExternalId string `json:"ExternalId"`
	PackageLabel string `json:"PackageLabel"`
}

type FinishPackagesRequest struct {
	ActualDate string `json:"ActualDate"`
	Label string `json:"Label"`
}

type FinishedgoodFlagRequest struct {
	Label string `json:"Label"`
}

type FinishedgoodUnflagRequest struct {
	Label string `json:"Label"`
}

type InTransit struct {
	ArchivedDate *string `json:"ArchivedDate"`
	ContainsDecontaminatedProduct bool `json:"ContainsDecontaminatedProduct"`
	ContainsRemediatedProduct bool `json:"ContainsRemediatedProduct"`
	DecontaminationDate *string `json:"DecontaminationDate"`
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId *int `json:"ExternalId"`
	FinishedDate string `json:"FinishedDate"`
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
	Item InTransitItem `json:"Item"`
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
	ProductLabel InTransitProductLabel `json:"ProductLabel"`
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

type InTransitItem struct {
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

type InTransitProductLabel struct {
	IsActive bool `json:"IsActive"`
	IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel"`
	IsDiscontinued bool `json:"IsDiscontinued"`
	LastLabelGenerationDate *string `json:"LastLabelGenerationDate"`
	PackageId int `json:"PackageId"`
	QrCount int `json:"QrCount"`
	TagId *int `json:"TagId"`
}

type ItemRequest struct {
	Item string `json:"Item"`
	Label string `json:"Label"`
}

type LabtestsRequiredRequest struct {
	Label string `json:"Label"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches"`
}

type NoteRequest struct {
	Note string `json:"Note"`
	PackageLabel string `json:"PackageLabel"`
}

type PackagesLocationRequest struct {
	Label string `json:"Label"`
	Location string `json:"Location"`
	MoveDate string `json:"MoveDate"`
	Sublocation string `json:"Sublocation"`
}

type PackagesPackage struct {
	ArchivedDate *string `json:"ArchivedDate"`
	ContainsDecontaminatedProduct bool `json:"ContainsDecontaminatedProduct"`
	ContainsRemediatedProduct bool `json:"ContainsRemediatedProduct"`
	DecontaminationDate *string `json:"DecontaminationDate"`
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId *int `json:"ExternalId"`
	FinishedDate string `json:"FinishedDate"`
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
	Item PackagesPackageItem `json:"Item"`
	ItemFromFacilityLicenseNumber *string `json:"ItemFromFacilityLicenseNumber"`
	ItemFromFacilityName *string `json:"ItemFromFacilityName"`
	LabFacilityLicenseNumber *string `json:"LabFacilityLicenseNumber"`
	LabFacilityName *string `json:"LabFacilityName"`
	LabTestResultExpirationDateTime *string `json:"LabTestResultExpirationDateTime"`
	LabTestStage string `json:"LabTestStage"`
	LabTestStageId int `json:"LabTestStageId"`
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
	OverallPassed bool `json:"OverallPassed"`
	PackageForProductDestruction *string `json:"PackageForProductDestruction"`
	PackageType string `json:"PackageType"`
	PackagedDate string `json:"PackagedDate"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	ProductLabel PackagesPackageProductLabel `json:"ProductLabel"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	Quantity float64 `json:"Quantity"`
	ReceivedDateTime *string `json:"ReceivedDateTime"`
	ReceivedFromFacilityLicenseNumber *string `json:"ReceivedFromFacilityLicenseNumber"`
	ReceivedFromFacilityName *string `json:"ReceivedFromFacilityName"`
	ReceivedFromManifestNumber *string `json:"ReceivedFromManifestNumber"`
	RemediationDate *string `json:"RemediationDate"`
	ResultReleaseDateTime string `json:"ResultReleaseDateTime"`
	ResultReleased bool `json:"ResultReleased"`
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
	TestComment *string `json:"TestComment"`
	TestInformationalOnly bool `json:"TestInformationalOnly"`
	TestPassed bool `json:"TestPassed"`
	TestPerformedDate string `json:"TestPerformedDate"`
	TestResultLevel int `json:"TestResultLevel"`
	TestTypeName string `json:"TestTypeName"`
	UnitOfMeasureAbbreviation string `json:"UnitOfMeasureAbbreviation"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
	UseByDate *string `json:"UseByDate"`
}

type PackagesPackageItem struct {
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

type PackagesPackageProductLabel struct {
	IsActive bool `json:"IsActive"`
	IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel"`
	IsDiscontinued bool `json:"IsDiscontinued"`
	LastLabelGenerationDate *string `json:"LastLabelGenerationDate"`
	PackageId int `json:"PackageId"`
	QrCount int `json:"QrCount"`
	TagId *int `json:"TagId"`
}

type PackagesPackagesRequest struct {
	ActualDate string `json:"ActualDate"`
	ExpirationDate *string `json:"ExpirationDate"`
	Ingredients []PackagesPackagesRequestIngredient `json:"Ingredients"`
	IsDonation bool `json:"IsDonation"`
	IsProductionBatch bool `json:"IsProductionBatch"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	LabTestStageId *int `json:"LabTestStageId"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	ProcessingJobTypeId *int `json:"ProcessingJobTypeId"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	Quantity float64 `json:"Quantity"`
	RequiredLabTestBatches *bool `json:"RequiredLabTestBatches"`
	SellByDate *string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UseByDate *string `json:"UseByDate"`
	UseSameItem bool `json:"UseSameItem"`
}

type PackagesPackagesRequestIngredient struct {
	Package string `json:"Package"`
	Quantity float64 `json:"Quantity"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type PackagesPlantingsRequest struct {
	LocationName string `json:"LocationName"`
	PackageAdjustmentAmount int `json:"PackageAdjustmentAmount"`
	PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName"`
	PackageLabel string `json:"PackageLabel"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PlantBatchName string `json:"PlantBatchName"`
	PlantBatchType string `json:"PlantBatchType"`
	PlantCount int `json:"PlantCount"`
	PlantedDate string `json:"PlantedDate"`
	StrainName string `json:"StrainName"`
	SublocationName string `json:"SublocationName"`
	UnpackagedDate string `json:"UnpackagedDate"`
}

type RemediateRequest struct {
	PackageLabel string `json:"PackageLabel"`
	RemediationDate string `json:"RemediationDate"`
	RemediationMethodName string `json:"RemediationMethodName"`
	RemediationSteps string `json:"RemediationSteps"`
}

type SourceHarvest struct {
	HarvestStartDate string `json:"HarvestStartDate"`
	HarvestedByFacilityLicenseNumber string `json:"HarvestedByFacilityLicenseNumber"`
	HarvestedByFacilityName string `json:"HarvestedByFacilityName"`
	Name string `json:"Name"`
}

type TestingRequest struct {
	ActualDate string `json:"ActualDate"`
	ExpirationDate *string `json:"ExpirationDate"`
	Ingredients []TestingRequestIngredient `json:"Ingredients"`
	IsDonation bool `json:"IsDonation"`
	IsProductionBatch bool `json:"IsProductionBatch"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	LabTestStageId *int `json:"LabTestStageId"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	ProcessingJobTypeId *int `json:"ProcessingJobTypeId"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	Quantity float64 `json:"Quantity"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches"`
	SellByDate *string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UseByDate *string `json:"UseByDate"`
	UseSameItem bool `json:"UseSameItem"`
}

type TestingRequestIngredient struct {
	Package string `json:"Package"`
	Quantity float64 `json:"Quantity"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type TradeSampleFlagRequest struct {
	Label string `json:"Label"`
}

type TradeSampleUnflagRequest struct {
	Label string `json:"Label"`
}

type UnfinishPackagesRequest struct {
	Label string `json:"Label"`
}

type UseByDateRequest struct {
	ExpirationDate string `json:"ExpirationDate"`
	Label string `json:"Label"`
	SellByDate string `json:"SellByDate"`
	UseByDate string `json:"UseByDate"`
}
