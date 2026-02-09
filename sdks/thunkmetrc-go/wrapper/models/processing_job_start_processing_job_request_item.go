package models

type ProcessingJobStartProcessingJobRequestItem struct {
    CountUnitOfMeasure string `json:"CountUnitOfMeasure,omitempty"`
    JobName string `json:"JobName,omitempty"`
    JobType string `json:"JobType,omitempty"`
    Packages []*ProcessingJobStartProcessingJobRequestItemPackage `json:"Packages,omitempty"`
    StartDate string `json:"StartDate,omitempty"`
    VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure,omitempty"`
    WeightUnitOfMeasure string `json:"WeightUnitOfMeasure,omitempty"`
}
