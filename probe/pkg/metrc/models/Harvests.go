package models

type FinishHarvestsRequest struct {
	ActualDate string `json:"ActualDate"`
	Id int `json:"Id"`
}

type HarvestsLocationRequest struct {
	ActualDate string `json:"ActualDate"`
	DryingLocation string `json:"DryingLocation"`
	DryingSublocation string `json:"DryingSublocation"`
	HarvestName *string `json:"HarvestName"`
	Id int `json:"Id"`
}

type HarvestsPackagesRequest struct {
	ActualDate string `json:"ActualDate"`
	DecontaminateProduct bool `json:"DecontaminateProduct"`
	DecontaminationDate *string `json:"DecontaminationDate"`
	DecontaminationSteps *string `json:"DecontaminationSteps"`
	ExpirationDate *string `json:"ExpirationDate"`
	Ingredients []HarvestsPackagesRequestIngredient `json:"Ingredients"`
	IsDonation bool `json:"IsDonation"`
	IsProductionBatch bool `json:"IsProductionBatch"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	LabTestStageId int `json:"LabTestStageId"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	RemediateProduct bool `json:"RemediateProduct"`
	RemediationDate *string `json:"RemediationDate"`
	RemediationMethodId *int `json:"RemediationMethodId"`
	RemediationSteps *string `json:"RemediationSteps"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches"`
	SellByDate *string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UnitOfWeight string `json:"UnitOfWeight"`
	UseByDate *string `json:"UseByDate"`
}

type HarvestsPackagesRequestIngredient struct {
	HarvestId int `json:"HarvestId"`
	HarvestName *string `json:"HarvestName"`
	UnitOfWeight string `json:"UnitOfWeight"`
	Weight float64 `json:"Weight"`
}

type HarvestsWaste struct {
	ActualDate string `json:"ActualDate"`
	Id int `json:"Id"`
	UnitOfWeightName string `json:"UnitOfWeightName"`
	WasteTypeName string `json:"WasteTypeName"`
	WasteWeight float64 `json:"WasteWeight"`
}

type HarvestsWasteRequest struct {
	ActualDate string `json:"ActualDate"`
	Id int `json:"Id"`
	UnitOfWeight string `json:"UnitOfWeight"`
	WasteType string `json:"WasteType"`
	WasteWeight float64 `json:"WasteWeight"`
}

type PackagesTestingRequest struct {
	ActualDate string `json:"ActualDate"`
	DecontaminateProduct bool `json:"DecontaminateProduct"`
	DecontaminationDate *string `json:"DecontaminationDate"`
	DecontaminationSteps *string `json:"DecontaminationSteps"`
	ExpirationDate *string `json:"ExpirationDate"`
	Ingredients []PackagesTestingRequestIngredient `json:"Ingredients"`
	IsDonation bool `json:"IsDonation"`
	IsProductionBatch bool `json:"IsProductionBatch"`
	IsTradeSample bool `json:"IsTradeSample"`
	Item string `json:"Item"`
	LabTestStageId int `json:"LabTestStageId"`
	Location *string `json:"Location"`
	Note string `json:"Note"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	RemediateProduct bool `json:"RemediateProduct"`
	RemediationDate *string `json:"RemediationDate"`
	RemediationMethodId *int `json:"RemediationMethodId"`
	RemediationSteps *string `json:"RemediationSteps"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches"`
	SellByDate *string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UnitOfWeight string `json:"UnitOfWeight"`
	UseByDate *string `json:"UseByDate"`
}

type PackagesTestingRequestIngredient struct {
	HarvestId int `json:"HarvestId"`
	HarvestName *string `json:"HarvestName"`
	UnitOfWeight string `json:"UnitOfWeight"`
	Weight float64 `json:"Weight"`
}

type RenameRequest struct {
	Id int `json:"Id"`
	NewName string `json:"NewName"`
	OldName *string `json:"OldName"`
}

type RestoreHarvestedPlantsRequest struct {
	HarvestId int `json:"HarvestId"`
	PlantTags []string `json:"PlantTags"`
}

type UnfinishHarvestsRequest struct {
	Id int `json:"Id"`
}

type WasteType struct {
	Name string `json:"Name"`
}
