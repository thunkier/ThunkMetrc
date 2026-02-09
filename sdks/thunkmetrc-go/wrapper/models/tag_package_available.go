package models

type TagPackageAvailable struct {
    FacilityId int64 `json:"FacilityId,omitempty"`
    GroupTagTypeId string `json:"GroupTagTypeId,omitempty"`
    GroupTagTypeName string `json:"GroupTagTypeName,omitempty"`
    Id int64 `json:"Id,omitempty"`
    Label string `json:"Label,omitempty"`
    MaxGroupSize int64 `json:"MaxGroupSize,omitempty"`
    TagInventoryTypeName string `json:"TagInventoryTypeName,omitempty"`
    TagTypeId string `json:"TagTypeId,omitempty"`
    TagTypeName string `json:"TagTypeName,omitempty"`
}
