package models

type ProcessingJobCreateAdjustProcessingJobRequestItem struct {
    AdjustmentDate string `json:"AdjustmentDate,omitempty"`
    AdjustmentNote string `json:"AdjustmentNote,omitempty"`
    AdjustmentReason string `json:"AdjustmentReason,omitempty"`
    CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
    Id int `json:"Id,omitempty"`
    Packages []*ProcessingJobCreateAdjustProcessingJobRequestItemPackage `json:"Packages,omitempty"`
    VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
    WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
}
