package models

type PlantBatche struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
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
	CurrentPage int `json:"CurrentPage"`
	Data []PlantBatcheData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type PlantBatcheData struct {
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
	CanBeCloned bool `json:"CanBeCloned"`
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

