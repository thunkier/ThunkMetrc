package models

type PackagesCreatePackageTestingRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    Ingredients []*PackagesCreatePackageTestingRequestItemIngredient `json:"Ingredients,omitempty"`
    IsDonation bool `json:"IsDonation,omitempty"`
    IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
    IsTradeSample bool `json:"IsTradeSample,omitempty"`
    Item string `json:"Item,omitempty"`
    LabTestStageId int `json:"LabTestStageId,omitempty"`
    Location string `json:"Location,omitempty"`
    Note string `json:"Note,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
    ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
    ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
    Quantity int `json:"Quantity,omitempty"`
    RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
    Tag string `json:"Tag,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
    UseSameItem bool `json:"UseSameItem,omitempty"`
}
