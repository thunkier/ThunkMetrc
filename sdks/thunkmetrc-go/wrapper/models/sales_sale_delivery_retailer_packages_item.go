package models

type SalesSaleDeliveryRetailerPackagesItem struct {
    ArchivedDate *string `json:"ArchivedDate,omitempty"`
    CompletedDateTime *string `json:"CompletedDateTime,omitempty"`
    IsOnInvestigationHold bool `json:"IsOnInvestigationHold,omitempty"`
    IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall,omitempty"`
    IsOnRecall *bool `json:"IsOnRecall,omitempty"`
    IsOnRecallCombined bool `json:"IsOnRecallCombined,omitempty"`
    ItemServingSize *string `json:"ItemServingSize,omitempty"`
    ItemStrainName *string `json:"ItemStrainName,omitempty"`
    ItemSupplyDurationDays *int64 `json:"ItemSupplyDurationDays,omitempty"`
    ItemUnitCbdContent *float64 `json:"ItemUnitCbdContent,omitempty"`
    ItemUnitCbdContentDose *float64 `json:"ItemUnitCbdContentDose,omitempty"`
    ItemUnitCbdContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdContentDoseUnitOfMeasureName,omitempty"`
    ItemUnitCbdContentUnitOfMeasureName *string `json:"ItemUnitCbdContentUnitOfMeasureName,omitempty"`
    ItemUnitCbdPercent *float64 `json:"ItemUnitCbdPercent,omitempty"`
    ItemUnitQuantity *float64 `json:"ItemUnitQuantity,omitempty"`
    ItemUnitQuantityUnitOfMeasureName *string `json:"ItemUnitQuantityUnitOfMeasureName,omitempty"`
    ItemUnitThcContent *float64 `json:"ItemUnitThcContent,omitempty"`
    ItemUnitThcContentDose *float64 `json:"ItemUnitThcContentDose,omitempty"`
    ItemUnitThcContentDoseUnitOfMeasureName *string `json:"ItemUnitThcContentDoseUnitOfMeasureName,omitempty"`
    ItemUnitThcContentUnitOfMeasureName *string `json:"ItemUnitThcContentUnitOfMeasureName,omitempty"`
    ItemUnitThcPercent *float64 `json:"ItemUnitThcPercent,omitempty"`
    ItemUnitVolume *float64 `json:"ItemUnitVolume,omitempty"`
    ItemUnitVolumeUnitOfMeasureName *string `json:"ItemUnitVolumeUnitOfMeasureName,omitempty"`
    ItemUnitWeight *float64 `json:"ItemUnitWeight,omitempty"`
    ItemUnitWeightUnitOfMeasureName *string `json:"ItemUnitWeightUnitOfMeasureName,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    ProductCategoryName *string `json:"ProductCategoryName,omitempty"`
    ProductName *string `json:"ProductName,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
    RecordedByUserName *string `json:"RecordedByUserName,omitempty"`
    RecordedDateTime string `json:"RecordedDateTime,omitempty"`
    RetailerDeliveryState *string `json:"RetailerDeliveryState,omitempty"`
    TotalPrice float64 `json:"TotalPrice,omitempty"`
    UnitOfMeasureAbbreviation *string `json:"UnitOfMeasureAbbreviation,omitempty"`
    UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
}
