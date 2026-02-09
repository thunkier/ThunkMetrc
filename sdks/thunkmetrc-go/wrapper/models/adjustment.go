package models

type Adjustment struct {
    AdjustedByUserName string `json:"AdjustedByUserName,omitempty"`
    AdjustmentDate string `json:"AdjustmentDate,omitempty"`
    AdjustmentNote string `json:"AdjustmentNote,omitempty"`
    AdjustmentQuantity float64 `json:"AdjustmentQuantity,omitempty"`
    AdjustmentReasonName string `json:"AdjustmentReasonName,omitempty"`
    AdjustmentUnitOfMeasureAbbreviation string `json:"AdjustmentUnitOfMeasureAbbreviation,omitempty"`
    AdjustmentUnitOfMeasureName string `json:"AdjustmentUnitOfMeasureName,omitempty"`
    ItemCategoryName string `json:"ItemCategoryName,omitempty"`
    ItemName string `json:"ItemName,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
    PackageLabTestResultExpirationDateTime *string `json:"PackageLabTestResultExpirationDateTime,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    PackagedDate string `json:"PackagedDate,omitempty"`
}
