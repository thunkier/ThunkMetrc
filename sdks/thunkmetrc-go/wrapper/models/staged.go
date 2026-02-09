package models

type Staged struct {
    CommissionedDateTime *string `json:"CommissionedDateTime,omitempty"`
    DetachedDateTime *string `json:"DetachedDateTime,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsArchived bool `json:"IsArchived,omitempty"`
    IsStaged bool `json:"IsStaged,omitempty"`
    IsUsed bool `json:"IsUsed,omitempty"`
    Label string `json:"Label,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    MaxGroupSize int64 `json:"MaxGroupSize,omitempty"`
    ProductLabel *string `json:"ProductLabel,omitempty"`
    QrCount *int64 `json:"QrCount,omitempty"`
    StatusName *string `json:"StatusName,omitempty"`
    TagInventoryTypeName string `json:"TagInventoryTypeName,omitempty"`
    TagTypeId *int64 `json:"TagTypeId,omitempty"`
    TagTypeName *string `json:"TagTypeName,omitempty"`
    UsedDateTime *string `json:"UsedDateTime,omitempty"`
}
