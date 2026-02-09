package models

type PackageProductLabel struct {
    IsActive bool `json:"IsActive,omitempty"`
    IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel,omitempty"`
    IsDiscontinued bool `json:"IsDiscontinued,omitempty"`
    LastLabelGenerationDate *string `json:"LastLabelGenerationDate,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
    QrCount int64 `json:"QrCount,omitempty"`
    TagId *int64 `json:"TagId,omitempty"`
}
