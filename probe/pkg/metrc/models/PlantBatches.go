package models

type AdjustPlantBatchesRequest struct {
	AdjustmentDate string `json:"AdjustmentDate"`
	AdjustmentReason string `json:"AdjustmentReason"`
	PlantBatchName string `json:"PlantBatchName"`
	Quantity float64 `json:"Quantity"`
	ReasonNote *string `json:"ReasonNote"`
}

type NameRequest struct {
	Group string `json:"Group"`
	Id *int `json:"Id"`
	NewGroup string `json:"NewGroup"`
}

type PackagesFromMotherPlantRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	ExpirationDate string `json:"ExpirationDate"`
	Id *int `json:"Id"`
	IsDonation bool `json:"IsDonation"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PlantBatch string `json:"PlantBatch"`
	SellByDate string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UseByDate string `json:"UseByDate"`
}

type PlantBatchesAdditivesRequest struct {
	ActiveIngredients []PlantBatchesAdditivesRequestActiveIngredient `json:"ActiveIngredients"`
	ActualDate string `json:"ActualDate"`
	AdditiveType string `json:"AdditiveType"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber *string `json:"EpaRegistrationNumber"`
	PlantBatchName string `json:"PlantBatchName"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
}

type PlantBatchesAdditivesRequestActiveIngredient struct {
	Name string `json:"Name"`
	Percentage int `json:"Percentage"`
}

type PlantBatchesAdditivesUsingTemplateRequest struct {
	ActualDate string `json:"ActualDate"`
	AdditivesTemplateName string `json:"AdditivesTemplateName"`
	PlantBatchName string `json:"PlantBatchName"`
	Rate string `json:"Rate"`
	TotalAmountApplied int `json:"TotalAmountApplied"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure"`
	Volume string `json:"Volume"`
}

type PlantBatchesGrowthPhaseRequest struct {
	Count int `json:"Count"`
	GrowthDate string `json:"GrowthDate"`
	GrowthPhase string `json:"GrowthPhase"`
	Name string `json:"Name"`
	NewLocation string `json:"NewLocation"`
	NewSublocation string `json:"NewSublocation"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	StartingTag string `json:"StartingTag"`
}

type PlantBatchesLocationRequest struct {
	Location string `json:"Location"`
	MoveDate string `json:"MoveDate"`
	Name string `json:"Name"`
	Sublocation string `json:"Sublocation"`
}

type PlantBatchesPackagesRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	ExpirationDate string `json:"ExpirationDate"`
	Id *int `json:"Id"`
	IsDonation bool `json:"IsDonation"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PlantBatch string `json:"PlantBatch"`
	SellByDate string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UseByDate string `json:"UseByDate"`
}

type PlantBatchesPlantingsRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	Location *string `json:"Location"`
	Name string `json:"Name"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	SourcePlantBatches *string `json:"SourcePlantBatches"`
	Strain string `json:"Strain"`
	Sublocation *string `json:"Sublocation"`
	Type string `json:"Type"`
}

type PlantBatchesRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	PlantBatch string `json:"PlantBatch"`
	ReasonNote string `json:"ReasonNote"`
	WasteMaterialMixed string `json:"WasteMaterialMixed"`
	WasteMethodName string `json:"WasteMethodName"`
	WasteReasonName string `json:"WasteReasonName"`
	WasteUnitOfMeasure string `json:"WasteUnitOfMeasure"`
	WasteWeight float64 `json:"WasteWeight"`
}

type PlantBatchesSplitRequest struct {
	ActualDate string `json:"ActualDate"`
	Count int `json:"Count"`
	GroupName string `json:"GroupName"`
	Location string `json:"Location"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PlantBatch string `json:"PlantBatch"`
	Strain string `json:"Strain"`
	Sublocation string `json:"Sublocation"`
}

type PlantBatchesStrainRequest struct {
	Id int `json:"Id"`
	Name *string `json:"Name"`
	StrainId int `json:"StrainId"`
	StrainName *string `json:"StrainName"`
}

type PlantBatchesTagRequest struct {
	Group string `json:"Group"`
	Id *int `json:"Id"`
	NewTag string `json:"NewTag"`
	ReplaceDate string `json:"ReplaceDate"`
	TagId *int `json:"TagId"`
}

type PlantBatchesType struct {
	CanBeCloned bool `json:"CanBeCloned"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type PlantBatchesWaste struct {
	PlantBatchId int `json:"PlantBatchId"`
	PlantBatchName string `json:"PlantBatchName"`
	PlantCount int `json:"PlantCount"`
	PlantWasteNumber string `json:"PlantWasteNumber"`
	WasteDate string `json:"WasteDate"`
	WasteMethodName string `json:"WasteMethodName"`
	WasteReasonName string `json:"WasteReasonName"`
	WasteUnitOfMeasureName string `json:"WasteUnitOfMeasureName"`
	WasteWeight float64 `json:"WasteWeight"`
}

type PlantBatchesWasteRequest struct {
	MixedMaterial string `json:"MixedMaterial"`
	Note string `json:"Note"`
	PlantBatchName string `json:"PlantBatchName"`
	ReasonName string `json:"ReasonName"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
	WasteDate string `json:"WasteDate"`
	WasteMethodName string `json:"WasteMethodName"`
	WasteWeight float64 `json:"WasteWeight"`
}
