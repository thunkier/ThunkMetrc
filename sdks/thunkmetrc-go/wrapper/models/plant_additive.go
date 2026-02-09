package models

type PlantAdditive struct {
    AdditiveTypeName *string `json:"AdditiveTypeName,omitempty"`
    AmountUnitOfMeasure string `json:"AmountUnitOfMeasure,omitempty"`
    ApplicationDevice string `json:"ApplicationDevice,omitempty"`
    EpaRegistrationNumber *string `json:"EpaRegistrationNumber,omitempty"`
    Note *string `json:"Note,omitempty"`
    PlantBatchId *int64 `json:"PlantBatchId,omitempty"`
    PlantBatchName *string `json:"PlantBatchName,omitempty"`
    PlantCount int64 `json:"PlantCount,omitempty"`
    ProductSupplier string `json:"ProductSupplier,omitempty"`
    ProductTradeName string `json:"ProductTradeName,omitempty"`
    Rate *string `json:"Rate,omitempty"`
    RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
    RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
    TotalAmountApplied int64 `json:"TotalAmountApplied,omitempty"`
    Volume *float64 `json:"Volume,omitempty"`
}
