package models

type ProcessingJob struct {
    CountUnitOfMeasureAbbreviation *string `json:"CountUnitOfMeasureAbbreviation,omitempty"`
    CountUnitOfMeasureId int64 `json:"CountUnitOfMeasureId,omitempty"`
    CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
    FinishNote *string `json:"FinishNote,omitempty"`
    FinishedDate *string `json:"FinishedDate,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsFinished bool `json:"IsFinished,omitempty"`
    JobTypeId int64 `json:"JobTypeId,omitempty"`
    JobTypeName *string `json:"JobTypeName,omitempty"`
    Name string `json:"Name,omitempty"`
    Number string `json:"Number,omitempty"`
    Packages []interface{} `json:"Packages,omitempty"`
    StartDate string `json:"StartDate,omitempty"`
    TotalCount int64 `json:"TotalCount,omitempty"`
    TotalCountWaste *string `json:"TotalCountWaste,omitempty"`
    TotalQuantity float64 `json:"TotalQuantity,omitempty"`
    TotalUnitOfMeasureId int64 `json:"TotalUnitOfMeasureId,omitempty"`
    TotalVolume float64 `json:"TotalVolume,omitempty"`
    TotalVolumeWaste *string `json:"TotalVolumeWaste,omitempty"`
    TotalWeight float64 `json:"TotalWeight,omitempty"`
    TotalWeightWaste *string `json:"TotalWeightWaste,omitempty"`
    VolumeUnitOfMeasureAbbreviation *string `json:"VolumeUnitOfMeasureAbbreviation,omitempty"`
    VolumeUnitOfMeasureId int64 `json:"VolumeUnitOfMeasureId,omitempty"`
    VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
    WasteCountUnitOfMeasureAbbreviation *string `json:"WasteCountUnitOfMeasureAbbreviation,omitempty"`
    WasteCountUnitOfMeasureId *int64 `json:"WasteCountUnitOfMeasureId,omitempty"`
    WasteCountUnitOfMeasureName *string `json:"WasteCountUnitOfMeasureName,omitempty"`
    WasteVolumeUnitOfMeasureAbbreviation *string `json:"WasteVolumeUnitOfMeasureAbbreviation,omitempty"`
    WasteVolumeUnitOfMeasureId *int64 `json:"WasteVolumeUnitOfMeasureId,omitempty"`
    WasteVolumeUnitOfMeasureName *string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
    WasteWeightUnitOfMeasureAbbreviation *string `json:"WasteWeightUnitOfMeasureAbbreviation,omitempty"`
    WasteWeightUnitOfMeasureId *int64 `json:"WasteWeightUnitOfMeasureId,omitempty"`
    WasteWeightUnitOfMeasureName *string `json:"WasteWeightUnitOfMeasureName,omitempty"`
    WeightUnitOfMeasureAbbreviation *string `json:"WeightUnitOfMeasureAbbreviation,omitempty"`
    WeightUnitOfMeasureId int64 `json:"WeightUnitOfMeasureId,omitempty"`
    WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
}
