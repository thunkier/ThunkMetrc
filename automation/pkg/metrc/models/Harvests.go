package models

type Harvest struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
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
	CurrentPage int `json:"CurrentPage"`
	Data []HarvestData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type HarvestData struct {
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
	ActualDate string `json:"ActualDate"`
	WasteTypeName string `json:"WasteTypeName"`
	WasteWeight float64 `json:"WasteWeight"`
}

