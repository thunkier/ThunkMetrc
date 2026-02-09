package models

type HarvestsCreateHarvestPackagesRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
    DecontaminationDate string `json:"DecontaminationDate,omitempty"`
    DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    Ingredients []*HarvestsCreateHarvestPackagesRequestItemIngredient `json:"Ingredients,omitempty"`
    IsDonation bool `json:"IsDonation,omitempty"`
    IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
    IsTradeSample bool `json:"IsTradeSample,omitempty"`
    Item string `json:"Item,omitempty"`
    LabTestStageId int `json:"LabTestStageId,omitempty"`
    Location string `json:"Location,omitempty"`
    Note string `json:"Note,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
    ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
    ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
    ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
    RemediateProduct bool `json:"RemediateProduct,omitempty"`
    RemediationDate string `json:"RemediationDate,omitempty"`
    RemediationMethodId int `json:"RemediationMethodId,omitempty"`
    RemediationSteps string `json:"RemediationSteps,omitempty"`
    RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
    Tag string `json:"Tag,omitempty"`
    UnitOfWeight string `json:"UnitOfWeight,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
}
