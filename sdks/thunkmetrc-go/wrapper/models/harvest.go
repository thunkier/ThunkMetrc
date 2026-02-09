package models

type Harvest struct {
    ArchivedDate *string `json:"ArchivedDate,omitempty"`
    CurrentWeight float64 `json:"CurrentWeight,omitempty"`
    DryingLocationId int64 `json:"DryingLocationId,omitempty"`
    DryingLocationName string `json:"DryingLocationName,omitempty"`
    DryingLocationTypeName *string `json:"DryingLocationTypeName,omitempty"`
    DryingSublocationId int64 `json:"DryingSublocationId,omitempty"`
    DryingSublocationName *string `json:"DryingSublocationName,omitempty"`
    FinishedDate *string `json:"FinishedDate,omitempty"`
    HarvestStartDate string `json:"HarvestStartDate,omitempty"`
    HarvestType string `json:"HarvestType,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsOnHold bool `json:"IsOnHold,omitempty"`
    IsOnInvestigation bool `json:"IsOnInvestigation,omitempty"`
    IsOnInvestigationHold bool `json:"IsOnInvestigationHold,omitempty"`
    IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall,omitempty"`
    LabTestingState *string `json:"LabTestingState,omitempty"`
    LabTestingStateDate *string `json:"LabTestingStateDate,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Name string `json:"Name,omitempty"`
    PackageCount int64 `json:"PackageCount,omitempty"`
    PatientLicenseNumber *string `json:"PatientLicenseNumber,omitempty"`
    PlantCount int64 `json:"PlantCount,omitempty"`
    SourceStrainCount int64 `json:"SourceStrainCount,omitempty"`
    SourceStrainNames *string `json:"SourceStrainNames,omitempty"`
    Strains []interface{} `json:"Strains,omitempty"`
    TotalPackagedWeight float64 `json:"TotalPackagedWeight,omitempty"`
    TotalRestoredWeight float64 `json:"TotalRestoredWeight,omitempty"`
    TotalWasteWeight float64 `json:"TotalWasteWeight,omitempty"`
    TotalWetWeight float64 `json:"TotalWetWeight,omitempty"`
    UnitOfWeightName string `json:"UnitOfWeightName,omitempty"`
}
