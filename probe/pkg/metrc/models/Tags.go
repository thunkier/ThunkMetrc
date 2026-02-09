package models

type Staged struct {
	CommissionedDateTime *string `json:"CommissionedDateTime"`
	DetachedDateTime *string `json:"DetachedDateTime"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsArchived bool `json:"IsArchived"`
	IsStaged bool `json:"IsStaged"`
	IsUsed bool `json:"IsUsed"`
	Label string `json:"Label"`
	LastModified string `json:"LastModified"`
	MaxGroupSize int `json:"MaxGroupSize"`
	ProductLabel *string `json:"ProductLabel"`
	QrCount *int `json:"QrCount"`
	StatusName *string `json:"StatusName"`
	TagInventoryTypeName string `json:"TagInventoryTypeName"`
	TagTypeId *int `json:"TagTypeId"`
	TagTypeName *string `json:"TagTypeName"`
	UsedDateTime *string `json:"UsedDateTime"`
}
