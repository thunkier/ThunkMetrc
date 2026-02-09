package models

type ActiveJobType struct {
	Attributes []ActiveJobTypeAttribute `json:"Attributes"`
	CategoryName string `json:"CategoryName"`
	Description string `json:"Description"`
	ForItems bool `json:"ForItems"`
	ForProcessingJobs bool `json:"ForProcessingJobs"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	ProcessingSteps string `json:"ProcessingSteps"`
	RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes"`
}

type ActiveJobTypeAttribute struct {
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type AdjustProcessingJobRequest struct {
	AdjustmentDate string `json:"AdjustmentDate"`
	AdjustmentNote *string `json:"AdjustmentNote"`
	AdjustmentReason string `json:"AdjustmentReason"`
	CountUnitOfMeasureName *string `json:"CountUnitOfMeasureName"`
	Id int `json:"Id"`
	Packages []AdjustProcessingJobRequestPackage `json:"Packages"`
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName"`
}

type AdjustProcessingJobRequestPackage struct {
	Label string `json:"Label"`
	Quantity float64 `json:"Quantity"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type FinishProcessingJobRequest struct {
	FinishDate string `json:"FinishDate"`
	FinishNote string `json:"FinishNote"`
	Id int `json:"Id"`
	TotalCountWaste *string `json:"TotalCountWaste"`
	TotalVolumeWaste *string `json:"TotalVolumeWaste"`
	TotalWeightWaste int `json:"TotalWeightWaste"`
	WasteCountUnitOfMeasureName *string `json:"WasteCountUnitOfMeasureName"`
	WasteVolumeUnitOfMeasureName *string `json:"WasteVolumeUnitOfMeasureName"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName"`
}

type InactiveJobType struct {
	Attributes []InactiveJobTypeAttribute `json:"Attributes"`
	CategoryName string `json:"CategoryName"`
	Description string `json:"Description"`
	ForItems bool `json:"ForItems"`
	ForProcessingJobs bool `json:"ForProcessingJobs"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	ProcessingSteps string `json:"ProcessingSteps"`
	RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes"`
}

type InactiveJobTypeAttribute struct {
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type JobTypesAttribute struct {
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type JobTypesCategory struct {
	ForItems bool `json:"ForItems"`
	ForProcessingJobs bool `json:"ForProcessingJobs"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes"`
}

type ProcessingJobJobTypesRequest struct {
	Attributes []string `json:"Attributes"`
	CategoryName string `json:"CategoryName"`
	Description string `json:"Description"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	ProcessingSteps string `json:"ProcessingSteps"`
}

type ProcessingJobPackagesRequest struct {
	ExpirationDate *string `json:"ExpirationDate"`
	FinishDate *string `json:"FinishDate"`
	FinishNote *string `json:"FinishNote"`
	FinishProcessingJob bool `json:"FinishProcessingJob"`
	Item *string `json:"Item"`
	JobName string `json:"JobName"`
	Location *string `json:"Location"`
	Note *string `json:"Note"`
	PackageDate *string `json:"PackageDate"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	Quantity float64 `json:"Quantity"`
	SellByDate *string `json:"SellByDate"`
	Sublocation *string `json:"Sublocation"`
	Tag string `json:"Tag"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UseByDate *string `json:"UseByDate"`
	WasteCountQuantity *float64 `json:"WasteCountQuantity"`
	WasteCountUnitOfMeasureName *string `json:"WasteCountUnitOfMeasureName"`
	WasteVolumeQuantity *float64 `json:"WasteVolumeQuantity"`
	WasteVolumeUnitOfMeasureName *string `json:"WasteVolumeUnitOfMeasureName"`
	WasteWeightQuantity *float64 `json:"WasteWeightQuantity"`
	WasteWeightUnitOfMeasureName *string `json:"WasteWeightUnitOfMeasureName"`
}

type StartProcessingJobRequest struct {
	CountUnitOfMeasure string `json:"CountUnitOfMeasure"`
	JobName string `json:"JobName"`
	JobType string `json:"JobType"`
	Packages []StartProcessingJobRequestPackage `json:"Packages"`
	StartDate string `json:"StartDate"`
	VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure"`
	WeightUnitOfMeasure string `json:"WeightUnitOfMeasure"`
}

type StartProcessingJobRequestPackage struct {
	Label string `json:"Label"`
	Quantity float64 `json:"Quantity"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type UnfinishProcessingJobRequest struct {
	Id int `json:"Id"`
}
