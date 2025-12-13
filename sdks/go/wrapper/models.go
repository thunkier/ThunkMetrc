package wrapper

type RetailIdCreateAssociateV2RequestItem struct {
	QrUrls []string `json:"QrUrls,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type RetailIdCreateGenerateV2Request struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type RetailIdCreateMergeV2Request struct {
	PackageLabels []string `json:"packageLabels,omitempty"`
}

type RetailIdCreatePackageInfoV2Request struct {
	PackageLabels []string `json:"packageLabels,omitempty"`
}

type StrainsCreateV1RequestItem struct {
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
}

type StrainsCreateV2RequestItem struct {
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
}

type StrainsCreateUpdateV1RequestItem struct {
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
}

type StrainsUpdateV2RequestItem struct {
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
}

type PlantBatchesCreateAdditivesV1RequestItem struct {
	AdditiveType string `json:"AdditiveType,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ActiveIngredients []PlantBatchesCreateAdditivesV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
}

type PlantBatchesCreateAdditivesV1RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantBatchesCreateAdditivesV2RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ActiveIngredients []PlantBatchesCreateAdditivesV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
}

type PlantBatchesCreateAdditivesV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantBatchesCreateAdditivesUsingtemplateV2RequestItem struct {
	Volume string `json:"Volume,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
}

type PlantBatchesCreateAdjustV1RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
}

type PlantBatchesCreateAdjustV2RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
}

type PlantBatchesCreateChangegrowthphaseV1RequestItem struct {
	StartingTag string `json:"StartingTag,omitempty"`
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	CountPerPlant string `json:"CountPerPlant,omitempty"`
	Count int `json:"Count,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesCreateGrowthphaseV2RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	CountPerPlant string `json:"CountPerPlant,omitempty"`
	StartingTag string `json:"StartingTag,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
	Name string `json:"Name,omitempty"`
	Count int `json:"Count,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
}

type PlantBatchesCreatePackageV2RequestItem struct {
	Location string `json:"Location,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Count int `json:"Count,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Item string `json:"Item,omitempty"`
	Tag string `json:"Tag,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Note string `json:"Note,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PlantBatchesCreatePackageFrommotherplantV1RequestItem struct {
	Sublocation string `json:"Sublocation,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Id int `json:"Id,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Note string `json:"Note,omitempty"`
	Tag string `json:"Tag,omitempty"`
}

type PlantBatchesCreatePackageFrommotherplantV2RequestItem struct {
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Tag string `json:"Tag,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
}

type PlantBatchesCreatePlantingsV2RequestItem struct {
	Type string `json:"Type,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	SourcePlantBatches string `json:"SourcePlantBatches,omitempty"`
	Count int `json:"Count,omitempty"`
	Strain string `json:"Strain,omitempty"`
	Location string `json:"Location,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesCreateSplitV1RequestItem struct {
	GroupName string `json:"GroupName,omitempty"`
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Strain string `json:"Strain,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
}

type PlantBatchesCreateSplitV2RequestItem struct {
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Strain string `json:"Strain,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	GroupName string `json:"GroupName,omitempty"`
}

type PlantBatchesCreateWasteV1RequestItem struct {
	Note string `json:"Note,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
}

type PlantBatchesCreateWasteV2RequestItem struct {
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
	Note string `json:"Note,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
}

type PlantBatchesCreatepackagesV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	Count int `json:"Count,omitempty"`
	Item string `json:"Item,omitempty"`
	Location string `json:"Location,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Note string `json:"Note,omitempty"`
	Id int `json:"Id,omitempty"`
}

type PlantBatchesCreateplantingsV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Strain string `json:"Strain,omitempty"`
	Location string `json:"Location,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	SourcePlantBatches string `json:"SourcePlantBatches,omitempty"`
	Type string `json:"Type,omitempty"`
	Count int `json:"Count,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesUpdateLocationV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	MoveDate string `json:"MoveDate,omitempty"`
}

type PlantBatchesUpdateMoveplantbatchesV1RequestItem struct {
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	MoveDate string `json:"MoveDate,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesUpdateNameV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Group string `json:"Group,omitempty"`
	NewGroup string `json:"NewGroup,omitempty"`
}

type PlantBatchesUpdateStrainV2RequestItem struct {
	StrainName string `json:"StrainName,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	StrainId int `json:"StrainId,omitempty"`
}

type PlantBatchesUpdateTagV2RequestItem struct {
	Group string `json:"Group,omitempty"`
	TagId int `json:"TagId,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	ReplaceDate string `json:"ReplaceDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type ProcessingJobsCreateAdjustV1RequestItem struct {
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	AdjustmentNote string `json:"AdjustmentNote,omitempty"`
	CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
	Packages []ProcessingJobsCreateAdjustV1RequestItem_Packages `json:"Packages,omitempty"`
	Id int `json:"Id,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
}

type ProcessingJobsCreateAdjustV1RequestItem_Packages struct {
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type ProcessingJobsCreateAdjustV2RequestItem struct {
	Packages []ProcessingJobsCreateAdjustV2RequestItem_Packages `json:"Packages,omitempty"`
	Id int `json:"Id,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	AdjustmentNote string `json:"AdjustmentNote,omitempty"`
	CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
}

type ProcessingJobsCreateAdjustV2RequestItem_Packages struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Label string `json:"Label,omitempty"`
}

type ProcessingJobsCreateJobtypesV1RequestItem struct {
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	Category string `json:"Category,omitempty"`
}

type ProcessingJobsCreateJobtypesV2RequestItem struct {
	Category string `json:"Category,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ProcessingJobsCreateStartV1RequestItem struct {
	Packages []ProcessingJobsCreateStartV1RequestItem_Packages `json:"Packages,omitempty"`
	StartDate string `json:"StartDate,omitempty"`
	JobName string `json:"JobName,omitempty"`
	JobType string `json:"JobType,omitempty"`
	CountUnitOfMeasure string `json:"CountUnitOfMeasure,omitempty"`
	VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure,omitempty"`
	WeightUnitOfMeasure string `json:"WeightUnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateStartV1RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateStartV2RequestItem struct {
	Packages []ProcessingJobsCreateStartV2RequestItem_Packages `json:"Packages,omitempty"`
	StartDate string `json:"StartDate,omitempty"`
	JobName string `json:"JobName,omitempty"`
	JobType string `json:"JobType,omitempty"`
	CountUnitOfMeasure string `json:"CountUnitOfMeasure,omitempty"`
	VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure,omitempty"`
	WeightUnitOfMeasure string `json:"WeightUnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateStartV2RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type ProcessingJobsCreatepackagesV1RequestItem struct {
	UseByDate string `json:"UseByDate,omitempty"`
	WasteVolumeQuantity string `json:"WasteVolumeQuantity,omitempty"`
	Location string `json:"Location,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	PackageDate string `json:"PackageDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	WasteCountQuantity string `json:"WasteCountQuantity,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Note string `json:"Note,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	FinishProcessingJob bool `json:"FinishProcessingJob,omitempty"`
	Item string `json:"Item,omitempty"`
	JobName string `json:"JobName,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	WasteWeightQuantity string `json:"WasteWeightQuantity,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
}

type ProcessingJobsCreatepackagesV2RequestItem struct {
	Location string `json:"Location,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	WasteCountQuantity string `json:"WasteCountQuantity,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	WasteWeightQuantity string `json:"WasteWeightQuantity,omitempty"`
	Item string `json:"Item,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	PackageDate string `json:"PackageDate,omitempty"`
	WasteVolumeQuantity string `json:"WasteVolumeQuantity,omitempty"`
	FinishProcessingJob bool `json:"FinishProcessingJob,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	JobName string `json:"JobName,omitempty"`
	Tag string `json:"Tag,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Note string `json:"Note,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
}

type ProcessingJobsUpdateFinishV1RequestItem struct {
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	Id int `json:"Id,omitempty"`
	TotalWeightWaste int `json:"TotalWeightWaste,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	TotalVolumeWaste string `json:"TotalVolumeWaste,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	TotalCountWaste string `json:"TotalCountWaste,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
}

type ProcessingJobsUpdateFinishV2RequestItem struct {
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	TotalVolumeWaste string `json:"TotalVolumeWaste,omitempty"`
	Id int `json:"Id,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	TotalCountWaste string `json:"TotalCountWaste,omitempty"`
	TotalWeightWaste int `json:"TotalWeightWaste,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
}

type ProcessingJobsUpdateJobtypesV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	CategoryName string `json:"CategoryName,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
}

type ProcessingJobsUpdateJobtypesV2RequestItem struct {
	Attributes []string `json:"Attributes,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	CategoryName string `json:"CategoryName,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
}

type ProcessingJobsUpdateUnfinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type ProcessingJobsUpdateUnfinishV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SublocationsCreateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
}

type SublocationsUpdateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	Id int `json:"Id,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem struct {
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Destinations []TransfersCreateExternalIncomingV1RequestItem_Destinations `json:"Destinations,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations struct {
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	Transporters []TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem struct {
	Destinations []TransfersCreateExternalIncomingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations struct {
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	Transporters []TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters struct {
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails []TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages struct {
	UseByDate string `json:"UseByDate,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem struct {
	Destinations []TransfersCreateTemplatesV1RequestItem_Destinations `json:"Destinations,omitempty"`
	Name string `json:"Name,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersCreateTemplatesV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateTemplatesV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations_Transporters struct {
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem struct {
	DriverName string `json:"DriverName,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Name string `json:"Name,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	Destinations []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations `json:"Destinations,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations struct {
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters struct {
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem struct {
	TransferId int `json:"TransferId,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	Destinations []TransfersUpdateExternalIncomingV1RequestItem_Destinations `json:"Destinations,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations struct {
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	Transporters []TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	Packages []TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages struct {
	ExternalId string `json:"ExternalId,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem struct {
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	Destinations []TransfersUpdateExternalIncomingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	TransferId int `json:"TransferId,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations struct {
	Packages []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters struct {
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	TransporterDetails []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages struct {
	UseByDate string `json:"UseByDate,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem struct {
	Name string `json:"Name,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	Destinations []TransfersUpdateTemplatesV1RequestItem_Destinations `json:"Destinations,omitempty"`
	TransferTemplateId int `json:"TransferTemplateId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations struct {
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	Packages []TransfersUpdateTemplatesV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations_Packages struct {
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem struct {
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	Destinations []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	TransferTemplateId int `json:"TransferTemplateId,omitempty"`
	Name string `json:"Name,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations struct {
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	Transporters []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterDetails []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
}

type PlantsCreateAdditivesV1RequestItem struct {
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	PlantLabels []string `json:"PlantLabels,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
}

type PlantsCreateAdditivesV1RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantsCreateAdditivesV2RequestItem struct {
	PlantLabels []string `json:"PlantLabels,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
}

type PlantsCreateAdditivesV2RequestItem_ActiveIngredients struct {
	Percentage float64 `json:"Percentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantsCreateAdditivesBylocationV1RequestItem struct {
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
}

type PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients struct {
	Percentage float64 `json:"Percentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantsCreateAdditivesBylocationV2RequestItem struct {
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
}

type PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem struct {
	SublocationName string `json:"SublocationName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
	Volume string `json:"Volume,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
}

type PlantsCreateAdditivesUsingtemplateV2RequestItem struct {
	Volume string `json:"Volume,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	PlantLabels []string `json:"PlantLabels,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
}

type PlantsCreateChangegrowthphasesV1RequestItem struct {
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
}

type PlantsCreateHarvestplantsV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Plant string `json:"Plant,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
}

type PlantsCreateManicureV2RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	Plant string `json:"Plant,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PlantsCreateManicureplantsV1RequestItem struct {
	PlantCount int `json:"PlantCount,omitempty"`
	Plant string `json:"Plant,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
}

type PlantsCreateMoveplantsV1RequestItem struct {
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PlantsCreatePlantbatchPackageV1RequestItem struct {
	PlantLabel string `json:"PlantLabel,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PackageTag string `json:"PackageTag,omitempty"`
	Item string `json:"Item,omitempty"`
	Location string `json:"Location,omitempty"`
	Note string `json:"Note,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	Count int `json:"Count,omitempty"`
}

type PlantsCreatePlantbatchPackageV2RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PackageTag string `json:"PackageTag,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	Location string `json:"Location,omitempty"`
	Note string `json:"Note,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Count int `json:"Count,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PlantsCreatePlantingsV1RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
}

type PlantsCreatePlantingsV2RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
}

type PlantsCreateWasteV1RequestItem struct {
	Note string `json:"Note,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	PlantLabels []interface{} `json:"PlantLabels,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
}

type PlantsCreateWasteV2RequestItem struct {
	ReasonName string `json:"ReasonName,omitempty"`
	PlantLabels []interface{} `json:"PlantLabels,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	Note string `json:"Note,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
}

type PlantsUpdateAdjustV2RequestItem struct {
	AdjustReason string `json:"AdjustReason,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	AdjustCount int `json:"AdjustCount,omitempty"`
}

type PlantsUpdateGrowthphaseV2RequestItem struct {
	NewTag string `json:"NewTag,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PlantsUpdateHarvestV2RequestItem struct {
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Plant string `json:"Plant,omitempty"`
}

type PlantsUpdateLocationV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type PlantsUpdateMergeV2RequestItem struct {
	TargetPlantGroupLabel string `json:"TargetPlantGroupLabel,omitempty"`
	SourcePlantGroupLabel string `json:"SourcePlantGroupLabel,omitempty"`
	MergeDate string `json:"MergeDate,omitempty"`
}

type PlantsUpdateSplitV2RequestItem struct {
	SplitDate string `json:"SplitDate,omitempty"`
	SourcePlantLabel string `json:"SourcePlantLabel,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	TagLabel string `json:"TagLabel,omitempty"`
	StrainLabel string `json:"StrainLabel,omitempty"`
}

type PlantsUpdateStrainV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	StrainId int `json:"StrainId,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	Id int `json:"Id,omitempty"`
}

type PlantsUpdateTagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	TagId int `json:"TagId,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	ReplaceDate string `json:"ReplaceDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type AdditivesTemplatesCreateV2RequestItem struct {
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ActiveIngredients []AdditivesTemplatesCreateV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	Name string `json:"Name,omitempty"`
	RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
}

type AdditivesTemplatesCreateV2RequestItem_ActiveIngredients struct {
	Percentage float64 `json:"Percentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type AdditivesTemplatesUpdateV2RequestItem struct {
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	Id int `json:"Id,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	Note string `json:"Note,omitempty"`
	RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
	ActiveIngredients []AdditivesTemplatesUpdateV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	Name string `json:"Name,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
}

type AdditivesTemplatesUpdateV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type HarvestsCreateFinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsCreatePackageV1RequestItem struct {
	Note string `json:"Note,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	Tag string `json:"Tag,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	Ingredients []HarvestsCreatePackageV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type HarvestsCreatePackageV1RequestItem_Ingredients struct {
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
}

type HarvestsCreatePackageV2RequestItem struct {
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	Location string `json:"Location,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	Ingredients []HarvestsCreatePackageV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Item string `json:"Item,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
}

type HarvestsCreatePackageV2RequestItem_Ingredients struct {
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
}

type HarvestsCreatePackageTestingV1RequestItem struct {
	Item string `json:"Item,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	Note string `json:"Note,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	Ingredients []HarvestsCreatePackageTestingV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Location string `json:"Location,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
}

type HarvestsCreatePackageTestingV1RequestItem_Ingredients struct {
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
}

type HarvestsCreatePackageTestingV2RequestItem struct {
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Note string `json:"Note,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	Ingredients []HarvestsCreatePackageTestingV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Item string `json:"Item,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Location string `json:"Location,omitempty"`
}

type HarvestsCreatePackageTestingV2RequestItem_Ingredients struct {
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
}

type HarvestsCreateRemoveWasteV1RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
	WasteType string `json:"WasteType,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
}

type HarvestsCreateUnfinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type HarvestsCreateWasteV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	WasteType string `json:"WasteType,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsUpdateFinishV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsUpdateLocationV2RequestItem struct {
	HarvestName string `json:"HarvestName,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type HarvestsUpdateMoveV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsUpdateRenameV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	OldName string `json:"OldName,omitempty"`
	NewName string `json:"NewName,omitempty"`
}

type HarvestsUpdateRenameV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	OldName string `json:"OldName,omitempty"`
	NewName string `json:"NewName,omitempty"`
}

type HarvestsUpdateRestoreHarvestedPlantsV2RequestItem struct {
	HarvestId int `json:"HarvestId,omitempty"`
	PlantTags []string `json:"PlantTags,omitempty"`
}

type HarvestsUpdateUnfinishV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type LabTestsCreateRecordV1RequestItem struct {
	Results []LabTestsCreateRecordV1RequestItem_Results `json:"Results,omitempty"`
	Label string `json:"Label,omitempty"`
	ResultDate string `json:"ResultDate,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsCreateRecordV1RequestItem_Results struct {
	Quantity float64 `json:"Quantity,omitempty"`
	Passed bool `json:"Passed,omitempty"`
	Notes string `json:"Notes,omitempty"`
	LabTestTypeName string `json:"LabTestTypeName,omitempty"`
}

type LabTestsCreateRecordV2RequestItem struct {
	Results []LabTestsCreateRecordV2RequestItem_Results `json:"Results,omitempty"`
	Label string `json:"Label,omitempty"`
	ResultDate string `json:"ResultDate,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsCreateRecordV2RequestItem_Results struct {
	LabTestTypeName string `json:"LabTestTypeName,omitempty"`
	Quantity float64 `json:"Quantity,omitempty"`
	Passed bool `json:"Passed,omitempty"`
	Notes string `json:"Notes,omitempty"`
}

type LabTestsUpdateLabtestdocumentV1RequestItem struct {
	LabTestResultId int `json:"LabTestResultId,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsUpdateLabtestdocumentV2RequestItem struct {
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
	LabTestResultId int `json:"LabTestResultId,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
}

type LabTestsUpdateResultReleaseV1RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type LabTestsUpdateResultReleaseV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesCreateDeliveryV1RequestItem struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transactions []SalesCreateDeliveryV1RequestItem_Transactions `json:"Transactions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
}

type SalesCreateDeliveryV1RequestItem_Transactions struct {
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryV2RequestItem struct {
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	Transactions []SalesCreateDeliveryV2RequestItem_Transactions `json:"Transactions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
}

type SalesCreateDeliveryV2RequestItem_Transactions struct {
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	Price string `json:"Price,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	Packages []SalesCreateDeliveryRetailerV1RequestItem_Packages `json:"Packages,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	Destinations []SalesCreateDeliveryRetailerV1RequestItem_Destinations `json:"Destinations,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Packages struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Destinations struct {
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transactions []SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions struct {
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem struct {
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Destinations []SalesCreateDeliveryRetailerV2RequestItem_Destinations `json:"Destinations,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Packages []SalesCreateDeliveryRetailerV2RequestItem_Packages `json:"Packages,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Destinations struct {
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transactions []SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions struct {
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	Price string `json:"Price,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Packages struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
}

type SalesCreateDeliveryRetailerDepartV1RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerDepartV2RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerEndV1RequestItem struct {
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
	Packages []SalesCreateDeliveryRetailerEndV1RequestItem_Packages `json:"Packages,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerEndV1RequestItem_Packages struct {
	EndQuantity int `json:"EndQuantity,omitempty"`
	EndUnitOfMeasure string `json:"EndUnitOfMeasure,omitempty"`
	Label string `json:"Label,omitempty"`
}

type SalesCreateDeliveryRetailerEndV2RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
	Packages []SalesCreateDeliveryRetailerEndV2RequestItem_Packages `json:"Packages,omitempty"`
}

type SalesCreateDeliveryRetailerEndV2RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	EndQuantity int `json:"EndQuantity,omitempty"`
	EndUnitOfMeasure string `json:"EndUnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV1RequestItem struct {
	Packages []SalesCreateDeliveryRetailerRestockV1RequestItem_Packages `json:"Packages,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Destinations string `json:"Destinations,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV1RequestItem_Packages struct {
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV2RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Destinations string `json:"Destinations,omitempty"`
	Packages []SalesCreateDeliveryRetailerRestockV2RequestItem_Packages `json:"Packages,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV2RequestItem_Packages struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV1RequestItem struct {
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	Transactions []SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions `json:"Transactions,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions struct {
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV2RequestItem struct {
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	Transactions []SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions `json:"Transactions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions struct {
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
}

type SalesCreateReceiptV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Transactions []SalesCreateReceiptV1RequestItem_Transactions `json:"Transactions,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
}

type SalesCreateReceiptV1RequestItem_Transactions struct {
	Price string `json:"Price,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
}

type SalesCreateReceiptV2RequestItem struct {
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Transactions []SalesCreateReceiptV2RequestItem_Transactions `json:"Transactions,omitempty"`
}

type SalesCreateReceiptV2RequestItem_Transactions struct {
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
}

type SalesCreateTransactionByDateV1RequestItem struct {
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
}

type SalesUpdateDeliveryV1RequestItem struct {
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	Id int `json:"Id,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	Transactions []SalesUpdateDeliveryV1RequestItem_Transactions `json:"Transactions,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type SalesUpdateDeliveryV1RequestItem_Transactions struct {
	Price string `json:"Price,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
}

type SalesUpdateDeliveryV2RequestItem struct {
	VehicleMake string `json:"VehicleMake,omitempty"`
	Transactions []SalesUpdateDeliveryV2RequestItem_Transactions `json:"Transactions,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	Id int `json:"Id,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
}

type SalesUpdateDeliveryV2RequestItem_Transactions struct {
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	Price string `json:"Price,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
}

type SalesUpdateDeliveryCompleteV1RequestItem struct {
	PaymentType string `json:"PaymentType,omitempty"`
	AcceptedPackages []string `json:"AcceptedPackages,omitempty"`
	ReturnedPackages []SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages `json:"ReturnedPackages,omitempty"`
	Id int `json:"Id,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
}

type SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages struct {
	ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure,omitempty"`
	ReturnReason string `json:"ReturnReason,omitempty"`
	ReturnReasonNote string `json:"ReturnReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
	ReturnQuantityVerified int `json:"ReturnQuantityVerified,omitempty"`
}

type SalesUpdateDeliveryCompleteV2RequestItem struct {
	PaymentType string `json:"PaymentType,omitempty"`
	AcceptedPackages []string `json:"AcceptedPackages,omitempty"`
	ReturnedPackages []SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages `json:"ReturnedPackages,omitempty"`
	Id int `json:"Id,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
}

type SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages struct {
	ReturnReasonNote string `json:"ReturnReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
	ReturnQuantityVerified int `json:"ReturnQuantityVerified,omitempty"`
	ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure,omitempty"`
	ReturnReason string `json:"ReturnReason,omitempty"`
}

type SalesUpdateDeliveryHubV1RequestItem struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	Id int `json:"Id,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterFacilityId string `json:"TransporterFacilityId,omitempty"`
}

type SalesUpdateDeliveryHubV2RequestItem struct {
	DriverName string `json:"DriverName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterFacilityId string `json:"TransporterFacilityId,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubAcceptV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubAcceptV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubDepartV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubDepartV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubVerifyIdV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	PaymentType string `json:"PaymentType,omitempty"`
}

type SalesUpdateDeliveryHubVerifyIdV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	PaymentType string `json:"PaymentType,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Destinations []SalesUpdateDeliveryRetailerV1RequestItem_Destinations `json:"Destinations,omitempty"`
	Packages []SalesUpdateDeliveryRetailerV1RequestItem_Packages `json:"Packages,omitempty"`
	Id int `json:"Id,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Destinations struct {
	DriverName string `json:"DriverName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	Transactions []SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	Id int `json:"Id,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	DriverEmployeeId int `json:"DriverEmployeeId,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Packages struct {
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	Destinations []SalesUpdateDeliveryRetailerV2RequestItem_Destinations `json:"Destinations,omitempty"`
	Packages []SalesUpdateDeliveryRetailerV2RequestItem_Packages `json:"Packages,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Id int `json:"Id,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Destinations struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	DriverEmployeeId int `json:"DriverEmployeeId,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	Transactions []SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	Id int `json:"Id,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions struct {
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Price string `json:"Price,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Packages struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesUpdateReceiptV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Id int `json:"Id,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	Transactions []SalesUpdateReceiptV1RequestItem_Transactions `json:"Transactions,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
}

type SalesUpdateReceiptV1RequestItem_Transactions struct {
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type SalesUpdateReceiptV2RequestItem struct {
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	Transactions []SalesUpdateReceiptV2RequestItem_Transactions `json:"Transactions,omitempty"`
	Id int `json:"Id,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type SalesUpdateReceiptV2RequestItem_Transactions struct {
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
}

type SalesUpdateReceiptFinalizeV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateReceiptUnfinalizeV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateTransactionByDateV1RequestItem struct {
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	Price string `json:"Price,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
}

type TransportersCreateDriverV2RequestItem struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	EmployeeId string `json:"EmployeeId,omitempty"`
	Name string `json:"Name,omitempty"`
}

type TransportersCreateVehicleV2RequestItem struct {
	Make string `json:"Make,omitempty"`
	Model string `json:"Model,omitempty"`
	LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
}

type TransportersUpdateDriverV2RequestItem struct {
	Id string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	EmployeeId string `json:"EmployeeId,omitempty"`
}

type TransportersUpdateVehicleV2RequestItem struct {
	LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
	Id string `json:"Id,omitempty"`
	Make string `json:"Make,omitempty"`
	Model string `json:"Model,omitempty"`
}

type PackagesCreateV1RequestItem struct {
	Tag string `json:"Tag,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Ingredients []PackagesCreateV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	Item string `json:"Item,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
}

type PackagesCreateV1RequestItem_Ingredients struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Package string `json:"Package,omitempty"`
}

type PackagesCreateV2RequestItem struct {
	Item string `json:"Item,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Tag string `json:"Tag,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Ingredients []PackagesCreateV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Location string `json:"Location,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PackagesCreateV2RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateAdjustV1RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesCreateAdjustV2RequestItem struct {
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type PackagesCreateChangeItemV1RequestItem struct {
	Label string `json:"Label,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PackagesCreateChangeLocationV1RequestItem struct {
	MoveDate string `json:"MoveDate,omitempty"`
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PackagesCreateFinishV1RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesCreatePlantingsV1RequestItem struct {
	LocationName string `json:"LocationName,omitempty"`
	PackageAdjustmentAmount int `json:"PackageAdjustmentAmount,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	PlantedDate string `json:"PlantedDate,omitempty"`
	UnpackagedDate string `json:"UnpackagedDate,omitempty"`
	PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
}

type PackagesCreatePlantingsV2RequestItem struct {
	LocationName string `json:"LocationName,omitempty"`
	UnpackagedDate string `json:"UnpackagedDate,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	PackageAdjustmentAmount int `json:"PackageAdjustmentAmount,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantedDate string `json:"PlantedDate,omitempty"`
}

type PackagesCreateRemediateV1RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	RemediationMethodName string `json:"RemediationMethodName,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
}

type PackagesCreateTestingV1RequestItem struct {
	Tag string `json:"Tag,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Ingredients []PackagesCreateTestingV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
}

type PackagesCreateTestingV1RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateTestingV2RequestItem struct {
	Note string `json:"Note,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	Ingredients []PackagesCreateTestingV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	Tag string `json:"Tag,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	Item string `json:"Item,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
}

type PackagesCreateTestingV2RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateUnfinishV1RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateAdjustV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
}

type PackagesUpdateChangeNoteV1RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Note string `json:"Note,omitempty"`
}

type PackagesUpdateDecontaminateV2RequestItem struct {
	DecontaminationMethodName string `json:"DecontaminationMethodName,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type PackagesUpdateDonationFlagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateDonationUnflagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateExternalidV2RequestItem struct {
	ExternalId string `json:"ExternalId,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type PackagesUpdateFinishV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PackagesUpdateItemV2RequestItem struct {
	Item string `json:"Item,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateLabTestRequiredV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
}

type PackagesUpdateLocationV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	MoveDate string `json:"MoveDate,omitempty"`
}

type PackagesUpdateNoteV2RequestItem struct {
	Note string `json:"Note,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type PackagesUpdateRemediateV2RequestItem struct {
	RemediationMethodName string `json:"RemediationMethodName,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type PackagesUpdateTradesampleFlagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateTradesampleUnflagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateUnfinishV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateUsebydateV2RequestItem struct {
	UseByDate string `json:"UseByDate,omitempty"`
	Label string `json:"Label,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
}

type PatientCheckInsCreateV1RequestItem struct {
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
}

type PatientCheckInsCreateV2RequestItem struct {
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PatientCheckInsUpdateV1RequestItem struct {
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	Id int `json:"Id,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
}

type PatientCheckInsUpdateV2RequestItem struct {
	CheckInDate string `json:"CheckInDate,omitempty"`
	Id int `json:"Id,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
}

type ItemsCreateV1RequestItem struct {
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	Strain string `json:"Strain,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	Description string `json:"Description,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	Name string `json:"Name,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
}

type ItemsCreateV2RequestItem struct {
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	Description string `json:"Description,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	Strain string `json:"Strain,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	Name string `json:"Name,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
}

type ItemsCreateBrandV2RequestItem struct {
	Name string `json:"Name,omitempty"`
}

type ItemsCreateFileV2RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreatePhotoV1RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreatePhotoV2RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreateUpdateV1RequestItem struct {
	Allergens string `json:"Allergens,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	Strain string `json:"Strain,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Description string `json:"Description,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	Id int `json:"Id,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	Name string `json:"Name,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
}

type ItemsUpdateV2RequestItem struct {
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	Description string `json:"Description,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	Strain string `json:"Strain,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	Id int `json:"Id,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	Name string `json:"Name,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
}

type ItemsUpdateBrandV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	Id int `json:"Id,omitempty"`
}

type LocationsCreateV1RequestItem struct {
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
}

type LocationsCreateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
}

type LocationsCreateUpdateV1RequestItem struct {
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
	Id int `json:"Id,omitempty"`
}

type LocationsUpdateV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
}

type PatientsCreateV2RequestItem struct {
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
}

type PatientsCreateAddV1RequestItem struct {
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
}

type PatientsCreateUpdateV1RequestItem struct {
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	NewLicenseNumber string `json:"NewLicenseNumber,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
}

type PatientsUpdateV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	NewLicenseNumber string `json:"NewLicenseNumber,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
}

