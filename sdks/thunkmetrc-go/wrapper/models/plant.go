package models

type Plant struct {
    ClonedCount *int64 `json:"ClonedCount,omitempty"`
    DescendedCount *int64 `json:"DescendedCount,omitempty"`
    DestroyedByUserName *string `json:"DestroyedByUserName,omitempty"`
    DestroyedDate *string `json:"DestroyedDate,omitempty"`
    DestroyedNote *string `json:"DestroyedNote,omitempty"`
    FloweringDate string `json:"FloweringDate,omitempty"`
    GroupTagTypeMax int64 `json:"GroupTagTypeMax,omitempty"`
    GrowthPhase string `json:"GrowthPhase,omitempty"`
    HarvestCount int64 `json:"HarvestCount,omitempty"`
    HarvestId *int64 `json:"HarvestId,omitempty"`
    HarvestedDate *string `json:"HarvestedDate,omitempty"`
    HarvestedUnitOfWeightAbbreviation *string `json:"HarvestedUnitOfWeightAbbreviation,omitempty"`
    HarvestedUnitOfWeightName *string `json:"HarvestedUnitOfWeightName,omitempty"`
    HarvestedWetWeight *float64 `json:"HarvestedWetWeight,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsOnHold bool `json:"IsOnHold,omitempty"`
    IsOnInvestigation bool `json:"IsOnInvestigation,omitempty"`
    IsOnInvestigationHold bool `json:"IsOnInvestigationHold,omitempty"`
    IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall,omitempty"`
    Label string `json:"Label,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    LocationId int64 `json:"LocationId,omitempty"`
    LocationName string `json:"LocationName,omitempty"`
    LocationTypeName *string `json:"LocationTypeName,omitempty"`
    MotherPlantDate *string `json:"MotherPlantDate,omitempty"`
    PatientLicenseNumber *string `json:"PatientLicenseNumber,omitempty"`
    PlantBatchId int64 `json:"PlantBatchId,omitempty"`
    PlantBatchName string `json:"PlantBatchName,omitempty"`
    PlantBatchTypeId int64 `json:"PlantBatchTypeId,omitempty"`
    PlantBatchTypeName string `json:"PlantBatchTypeName,omitempty"`
    PlantedDate string `json:"PlantedDate,omitempty"`
    State string `json:"State,omitempty"`
    StrainId int64 `json:"StrainId,omitempty"`
    StrainName string `json:"StrainName,omitempty"`
    SublocationId *int64 `json:"SublocationId,omitempty"`
    SublocationName *string `json:"SublocationName,omitempty"`
    SurvivedCount *int64 `json:"SurvivedCount,omitempty"`
    TagTypeMax int64 `json:"TagTypeMax,omitempty"`
    VegetativeDate string `json:"VegetativeDate,omitempty"`
}
