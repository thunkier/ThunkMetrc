package models

type PlantBatch struct {
    DestroyedCount int64 `json:"DestroyedCount,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsOnHold bool `json:"IsOnHold,omitempty"`
    IsOnInvestigation bool `json:"IsOnInvestigation,omitempty"`
    IsOnInvestigationHold bool `json:"IsOnInvestigationHold,omitempty"`
    IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    LocationId *int64 `json:"LocationId,omitempty"`
    LocationName *string `json:"LocationName,omitempty"`
    LocationTypeName *string `json:"LocationTypeName,omitempty"`
    MultiPlantBatch bool `json:"MultiPlantBatch,omitempty"`
    Name string `json:"Name,omitempty"`
    PackagedCount int64 `json:"PackagedCount,omitempty"`
    PatientLicenseNumber *string `json:"PatientLicenseNumber,omitempty"`
    PlantBatchTypeId int64 `json:"PlantBatchTypeId,omitempty"`
    PlantBatchTypeName string `json:"PlantBatchTypeName,omitempty"`
    PlantedDate string `json:"PlantedDate,omitempty"`
    SourcePackageId *int64 `json:"SourcePackageId,omitempty"`
    SourcePackageLabel *string `json:"SourcePackageLabel,omitempty"`
    SourcePlantBatchIds []interface{} `json:"SourcePlantBatchIds,omitempty"`
    SourcePlantBatchNames *string `json:"SourcePlantBatchNames,omitempty"`
    SourcePlantId *int64 `json:"SourcePlantId,omitempty"`
    SourcePlantLabel *string `json:"SourcePlantLabel,omitempty"`
    StrainId int64 `json:"StrainId,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
    SublocationId *int64 `json:"SublocationId,omitempty"`
    SublocationName *string `json:"SublocationName,omitempty"`
    TrackedCount int64 `json:"TrackedCount,omitempty"`
    UntrackedCount int64 `json:"UntrackedCount,omitempty"`
}
