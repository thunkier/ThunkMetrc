package models

type AdditivesTemplate struct {
	ActiveIngredients []AdditivesTemplateActiveIngredient `json:"ActiveIngredients"`
	AdditiveType string `json:"AdditiveType"`
	AdditiveTypeName *string `json:"AdditiveTypeName"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsActive bool `json:"IsActive"`
	LastModified string `json:"LastModified"`
	Name *string `json:"Name"`
	Note *string `json:"Note"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription"`
	RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription"`
}

type AdditivesTemplateActiveIngredient struct {
	Name string `json:"Name"`
	Percentage float64 `json:"Percentage"`
}

type AdjustPackagesRequest struct {
	AdjustmentDate string `json:"AdjustmentDate"`
	AdjustmentReason string `json:"AdjustmentReason"`
	Label string `json:"Label"`
	Quantity float64 `json:"Quantity"`
	ReasonNote *string `json:"ReasonNote"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type Harvest struct {
	ArchivedDate *string `json:"ArchivedDate"`
	CurrentWeight float64 `json:"CurrentWeight"`
	DryingLocationId int `json:"DryingLocationId"`
	DryingLocationName string `json:"DryingLocationName"`
	DryingLocationTypeName *string `json:"DryingLocationTypeName"`
	DryingSublocationId int `json:"DryingSublocationId"`
	DryingSublocationName *string `json:"DryingSublocationName"`
	FinishedDate *string `json:"FinishedDate"`
	HarvestStartDate string `json:"HarvestStartDate"`
	HarvestType string `json:"HarvestType"`
	Id int `json:"Id"`
	IsOnHold bool `json:"IsOnHold"`
	IsOnInvestigation bool `json:"IsOnInvestigation"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	LabTestingState *string `json:"LabTestingState"`
	LabTestingStateDate *string `json:"LabTestingStateDate"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
	PackageCount int `json:"PackageCount"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlantCount int `json:"PlantCount"`
	SourceStrainCount int `json:"SourceStrainCount"`
	SourceStrainNames *string `json:"SourceStrainNames"`
	Strains []interface{} `json:"Strains"`
	TotalPackagedWeight float64 `json:"TotalPackagedWeight"`
	TotalRestoredWeight float64 `json:"TotalRestoredWeight"`
	TotalWasteWeight float64 `json:"TotalWasteWeight"`
	TotalWetWeight float64 `json:"TotalWetWeight"`
	UnitOfWeightName string `json:"UnitOfWeightName"`
}

type Item struct {
	AdministrationMethod *string `json:"AdministrationMethod"`
	Allergens *string `json:"Allergens"`
	ApprovalStatus string `json:"ApprovalStatus"`
	ApprovalStatusDateTime string `json:"ApprovalStatusDateTime"`
	DefaultLabTestingState string `json:"DefaultLabTestingState"`
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
	ProductCategoryType string `json:"ProductCategoryType"`
	ProductImages []interface{} `json:"ProductImages"`
	ProductPDFDocuments []interface{} `json:"ProductPDFDocuments"`
	ProductPhotoDescription *string `json:"ProductPhotoDescription"`
	PublicIngredients *string `json:"PublicIngredients"`
	QuantityType string `json:"QuantityType"`
	ServingSize *string `json:"ServingSize"`
	StrainId int `json:"StrainId"`
	StrainName string `json:"StrainName"`
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
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
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

type Mother struct {
	ClonedCount *int `json:"ClonedCount"`
	DescendedCount *int `json:"DescendedCount"`
	DestroyedByUserName *string `json:"DestroyedByUserName"`
	DestroyedDate *string `json:"DestroyedDate"`
	DestroyedNote *string `json:"DestroyedNote"`
	FloweringDate *string `json:"FloweringDate"`
	GroupTagTypeMax int `json:"GroupTagTypeMax"`
	GrowthPhase string `json:"GrowthPhase"`
	HarvestCount int `json:"HarvestCount"`
	HarvestId *int `json:"HarvestId"`
	HarvestedDate *string `json:"HarvestedDate"`
	HarvestedUnitOfWeightAbbreviation *string `json:"HarvestedUnitOfWeightAbbreviation"`
	HarvestedUnitOfWeightName *string `json:"HarvestedUnitOfWeightName"`
	HarvestedWetWeight *float64 `json:"HarvestedWetWeight"`
	Id int `json:"Id"`
	IsOnHold bool `json:"IsOnHold"`
	IsOnInvestigation bool `json:"IsOnInvestigation"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	Label string `json:"Label"`
	LastModified string `json:"LastModified"`
	LocationId int `json:"LocationId"`
	LocationName string `json:"LocationName"`
	LocationTypeName *string `json:"LocationTypeName"`
	MotherPlantDate string `json:"MotherPlantDate"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlantBatchId int `json:"PlantBatchId"`
	PlantBatchName string `json:"PlantBatchName"`
	PlantBatchTypeId int `json:"PlantBatchTypeId"`
	PlantBatchTypeName string `json:"PlantBatchTypeName"`
	PlantedDate string `json:"PlantedDate"`
	State string `json:"State"`
	StrainId int `json:"StrainId"`
	StrainName string `json:"StrainName"`
	SublocationId *int `json:"SublocationId"`
	SublocationName *string `json:"SublocationName"`
	SurvivedCount *int `json:"SurvivedCount"`
	TagTypeMax int `json:"TagTypeMax"`
	VegetativeDate *string `json:"VegetativeDate"`
}

type Patient struct {
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate"`
	LicenseNumber string `json:"LicenseNumber"`
	OtherFacilitiesCount int `json:"OtherFacilitiesCount"`
	PatientId int `json:"PatientId"`
	RecommendedPlants int `json:"RecommendedPlants"`
	RecommendedSmokableQuantity float64 `json:"RecommendedSmokableQuantity"`
	RegistrationDate string `json:"RegistrationDate"`
}

type Plant struct {
	ClonedCount *int `json:"ClonedCount"`
	DescendedCount *int `json:"DescendedCount"`
	DestroyedByUserName *string `json:"DestroyedByUserName"`
	DestroyedDate *string `json:"DestroyedDate"`
	DestroyedNote *string `json:"DestroyedNote"`
	FloweringDate *string `json:"FloweringDate"`
	GroupTagTypeMax int `json:"GroupTagTypeMax"`
	GrowthPhase string `json:"GrowthPhase"`
	HarvestCount int `json:"HarvestCount"`
	HarvestId *int `json:"HarvestId"`
	HarvestedDate *string `json:"HarvestedDate"`
	HarvestedUnitOfWeightAbbreviation *string `json:"HarvestedUnitOfWeightAbbreviation"`
	HarvestedUnitOfWeightName *string `json:"HarvestedUnitOfWeightName"`
	HarvestedWetWeight *float64 `json:"HarvestedWetWeight"`
	Id int `json:"Id"`
	IsOnHold bool `json:"IsOnHold"`
	IsOnInvestigation bool `json:"IsOnInvestigation"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	Label string `json:"Label"`
	LastModified string `json:"LastModified"`
	LocationId int `json:"LocationId"`
	LocationName string `json:"LocationName"`
	LocationTypeName *string `json:"LocationTypeName"`
	MotherPlantDate *string `json:"MotherPlantDate"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlantBatchId int `json:"PlantBatchId"`
	PlantBatchName string `json:"PlantBatchName"`
	PlantBatchTypeId int `json:"PlantBatchTypeId"`
	PlantBatchTypeName string `json:"PlantBatchTypeName"`
	PlantedDate string `json:"PlantedDate"`
	State string `json:"State"`
	StrainId int `json:"StrainId"`
	StrainName string `json:"StrainName"`
	SublocationId *int `json:"SublocationId"`
	SublocationName *string `json:"SublocationName"`
	SurvivedCount *int `json:"SurvivedCount"`
	TagTypeMax int `json:"TagTypeMax"`
	VegetativeDate string `json:"VegetativeDate"`
}

type PlantBatch struct {
	DestroyedCount int `json:"DestroyedCount"`
	Id int `json:"Id"`
	IsOnHold bool `json:"IsOnHold"`
	IsOnInvestigation bool `json:"IsOnInvestigation"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	LastModified string `json:"LastModified"`
	LocationId *int `json:"LocationId"`
	LocationName *string `json:"LocationName"`
	LocationTypeName *string `json:"LocationTypeName"`
	MultiPlantBatch bool `json:"MultiPlantBatch"`
	Name string `json:"Name"`
	PackagedCount int `json:"PackagedCount"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlantBatchTypeId int `json:"PlantBatchTypeId"`
	PlantBatchTypeName string `json:"PlantBatchTypeName"`
	PlantedDate string `json:"PlantedDate"`
	SourcePackageId *int `json:"SourcePackageId"`
	SourcePackageLabel *string `json:"SourcePackageLabel"`
	SourcePlantBatchIds []interface{} `json:"SourcePlantBatchIds"`
	SourcePlantBatchNames *string `json:"SourcePlantBatchNames"`
	SourcePlantId *int `json:"SourcePlantId"`
	SourcePlantLabel *string `json:"SourcePlantLabel"`
	StrainId int `json:"StrainId"`
	StrainName string `json:"StrainName"`
	SublocationId *int `json:"SublocationId"`
	SublocationName *string `json:"SublocationName"`
	TrackedCount int `json:"TrackedCount"`
	UntrackedCount int `json:"UntrackedCount"`
}

type ProcessingJob struct {
	CountUnitOfMeasureAbbreviation *string `json:"CountUnitOfMeasureAbbreviation"`
	CountUnitOfMeasureId int `json:"CountUnitOfMeasureId"`
	CountUnitOfMeasureName string `json:"CountUnitOfMeasureName"`
	FinishNote *string `json:"FinishNote"`
	FinishedDate *string `json:"FinishedDate"`
	Id int `json:"Id"`
	IsFinished bool `json:"IsFinished"`
	JobTypeId int `json:"JobTypeId"`
	JobTypeName *string `json:"JobTypeName"`
	Name string `json:"Name"`
	Number string `json:"Number"`
	Packages []interface{} `json:"Packages"`
	StartDate string `json:"StartDate"`
	TotalCount int `json:"TotalCount"`
	TotalCountWaste *string `json:"TotalCountWaste"`
	TotalVolume float64 `json:"TotalVolume"`
	TotalVolumeWaste *string `json:"TotalVolumeWaste"`
	TotalWeight float64 `json:"TotalWeight"`
	TotalWeightWaste *string `json:"TotalWeightWaste"`
	VolumeUnitOfMeasureAbbreviation *string `json:"VolumeUnitOfMeasureAbbreviation"`
	VolumeUnitOfMeasureId int `json:"VolumeUnitOfMeasureId"`
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName"`
	WasteCountUnitOfMeasureAbbreviation *string `json:"WasteCountUnitOfMeasureAbbreviation"`
	WasteCountUnitOfMeasureId *int `json:"WasteCountUnitOfMeasureId"`
	WasteCountUnitOfMeasureName *string `json:"WasteCountUnitOfMeasureName"`
	WasteVolumeUnitOfMeasureAbbreviation *string `json:"WasteVolumeUnitOfMeasureAbbreviation"`
	WasteVolumeUnitOfMeasureId *int `json:"WasteVolumeUnitOfMeasureId"`
	WasteVolumeUnitOfMeasureName *string `json:"WasteVolumeUnitOfMeasureName"`
	WasteWeightUnitOfMeasureAbbreviation *string `json:"WasteWeightUnitOfMeasureAbbreviation"`
	WasteWeightUnitOfMeasureId *int `json:"WasteWeightUnitOfMeasureId"`
	WasteWeightUnitOfMeasureName *string `json:"WasteWeightUnitOfMeasureName"`
	WeightUnitOfMeasureAbbreviation *string `json:"WeightUnitOfMeasureAbbreviation"`
	WeightUnitOfMeasureId int `json:"WeightUnitOfMeasureId"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName"`
}

type Strain struct {
	CbdLevel *string `json:"CbdLevel"`
	Genetics string `json:"Genetics"`
	Id int `json:"Id"`
	IndicaPercentage int `json:"IndicaPercentage"`
	IsUsed bool `json:"IsUsed"`
	Name string `json:"Name"`
	SativaPercentage int `json:"SativaPercentage"`
	TestingStatus string `json:"TestingStatus"`
	ThcLevel *string `json:"ThcLevel"`
}

type Sublocation struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

type Tag struct {
	FacilityId int `json:"FacilityId"`
	GroupTagTypeId *int `json:"GroupTagTypeId"`
	GroupTagTypeName *string `json:"GroupTagTypeName"`
	Id int `json:"Id"`
	Label string `json:"Label"`
	MaxGroupSize int `json:"MaxGroupSize"`
	TagInventoryTypeName string `json:"TagInventoryTypeName"`
	TagTypeId *int `json:"TagTypeId"`
	TagTypeName *string `json:"TagTypeName"`
}

type Transfer struct {
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	ContainsDonation bool `json:"ContainsDonation"`
	ContainsPlantPackage bool `json:"ContainsPlantPackage"`
	ContainsProductPackage bool `json:"ContainsProductPackage"`
	ContainsProductRequiresRemediation bool `json:"ContainsProductRequiresRemediation"`
	ContainsRemediatedProductPackage bool `json:"ContainsRemediatedProductPackage"`
	ContainsTestingSample bool `json:"ContainsTestingSample"`
	ContainsTradeSample bool `json:"ContainsTradeSample"`
	CreatedByUserName *string `json:"CreatedByUserName"`
	CreatedDateTime string `json:"CreatedDateTime"`
	DeliveryCount int `json:"DeliveryCount"`
	DeliveryId int `json:"DeliveryId"`
	DeliveryPackageCount int `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount int `json:"DeliveryReceivedPackageCount"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	Id int `json:"Id"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	IsVoided bool `json:"IsVoided"`
	LastModified string `json:"LastModified"`
	ManifestNumber string `json:"ManifestNumber"`
	Name *string `json:"Name"`
	OriginatingTemplateId *int `json:"OriginatingTemplateId"`
	PackageCount int `json:"PackageCount"`
	ReceivedDateTime *string `json:"ReceivedDateTime"`
	ReceivedDeliveryCount int `json:"ReceivedDeliveryCount"`
	ReceivedPackageCount int `json:"ReceivedPackageCount"`
	RecipientFacilityLicenseNumber *string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName *string `json:"RecipientFacilityName"`
	ShipmentLicenseType int `json:"ShipmentLicenseType"`
	ShipmentTransactionType *string `json:"ShipmentTransactionType"`
	ShipmentTypeName *string `json:"ShipmentTypeName"`
	ShipperFacilityLicenseNumber string `json:"ShipperFacilityLicenseNumber"`
	ShipperFacilityName string `json:"ShipperFacilityName"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type UnitOfMeasure struct {
	Abbreviation string `json:"Abbreviation"`
	Name string `json:"Name"`
	QuantityType string `json:"QuantityType"`
}

type WasteMethod struct {
	ForPlants bool `json:"ForPlants"`
	ForProductDestruction bool `json:"ForProductDestruction"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type WriteResponse struct {
	Ids []int `json:"Ids"`
	Warnings *string `json:"Warnings"`
}
