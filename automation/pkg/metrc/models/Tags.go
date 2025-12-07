package models

type Tag struct {
	FacilityId int `json:"FacilityId"`
	GroupTagTypeId *int `json:"GroupTagTypeId"`
	GroupTagTypeName *string `json:"GroupTagTypeName"`
	Id int `json:"Id"`
	Label string `json:"Label"`
	MaxGroupSize int `json:"MaxGroupSize"`
	TagInventoryTypeName string `json:"TagInventoryTypeName"`
	TagTypeId *int `json:"TagTypeId"`
	TagTypeName *string `json:"TagTypeName"`
	CommissionedDateTime *string `json:"CommissionedDateTime"`
	DetachedDateTime *string `json:"DetachedDateTime"`
	IsArchived bool `json:"IsArchived"`
	IsStaged bool `json:"IsStaged"`
	IsUsed bool `json:"IsUsed"`
	LastModified string `json:"LastModified"`
	ProductLabel *string `json:"ProductLabel"`
	QrCount *int `json:"QrCount"`
	StatusName *string `json:"StatusName"`
	UsedDateTime *string `json:"UsedDateTime"`
}

