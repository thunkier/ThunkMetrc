package models

type ProcessingJob struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
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
	CurrentPage int `json:"CurrentPage"`
	Data []ProcessingJobData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type ProcessingJobData struct {
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
	Attributes []ProcessingJobDataAttribute `json:"Attributes"`
	CategoryName string `json:"CategoryName"`
	Description string `json:"Description"`
	ForItems bool `json:"ForItems"`
	ForProcessingJobs bool `json:"ForProcessingJobs"`
	ProcessingSteps string `json:"ProcessingSteps"`
	RequiresProcessingJobAttributes bool `json:"RequiresProcessingJobAttributes"`
	LastModified string `json:"LastModified"`
}

type ProcessingJobDataAttribute struct {
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

