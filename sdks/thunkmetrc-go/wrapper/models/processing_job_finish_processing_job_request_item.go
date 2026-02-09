package models

type ProcessingJobFinishProcessingJobRequestItem struct {
    FinishDate string `json:"FinishDate,omitempty"`
    FinishNote string `json:"FinishNote,omitempty"`
    Id int `json:"Id,omitempty"`
    TotalCountWaste string `json:"TotalCountWaste,omitempty"`
    TotalVolumeWaste string `json:"TotalVolumeWaste,omitempty"`
    TotalWeightWaste int `json:"TotalWeightWaste,omitempty"`
    WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
    WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
    WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
}
