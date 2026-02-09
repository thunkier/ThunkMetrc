package models

type PackagesInTransitProductLabel struct {
    IsActive bool `json:"IsActive,omitempty"`
    IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel,omitempty"`
    LabelSource *string `json:"LabelSource,omitempty"`
    LastLabelGenerationDate *string `json:"LastLabelGenerationDate,omitempty"`
    OriginalSourcePackageId *int64 `json:"OriginalSourcePackageId,omitempty"`
    OriginalSourcePackageLabel *string `json:"OriginalSourcePackageLabel,omitempty"`
    QrCodeRanges *string `json:"QrCodeRanges,omitempty"`
    QrCount int64 `json:"QrCount,omitempty"`
}
