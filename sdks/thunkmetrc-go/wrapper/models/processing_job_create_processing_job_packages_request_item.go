package models

type ProcessingJobCreateProcessingJobPackagesRequestItem struct {
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    FinishDate string `json:"FinishDate,omitempty"`
    FinishNote string `json:"FinishNote,omitempty"`
    FinishProcessingJob bool `json:"FinishProcessingJob,omitempty"`
    IsFinishedGood bool `json:"IsFinishedGood,omitempty"`
    Item string `json:"Item,omitempty"`
    JobName string `json:"JobName,omitempty"`
    Location string `json:"Location,omitempty"`
    Note string `json:"Note,omitempty"`
    PackageDate string `json:"PackageDate,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
    Tag string `json:"Tag,omitempty"`
    UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
    WasteCountQuantity float64 `json:"WasteCountQuantity,omitempty"`
    WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
    WasteVolumeQuantity float64 `json:"WasteVolumeQuantity,omitempty"`
    WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
    WasteWeightQuantity float64 `json:"WasteWeightQuantity,omitempty"`
    WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
}
