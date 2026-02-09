package models

type Allotment struct {
	Allotment int `json:"Allotment"`
	FacilityId int `json:"FacilityId"`
	FacilityLicenseNumber string `json:"FacilityLicenseNumber"`
}

type AssociateRequest struct {
	PackageLabel string `json:"PackageLabel"`
	QrUrls []string `json:"QrUrls"`
}

type GenerateRequest struct {
	PackageLabel string `json:"PackageLabel"`
	Quantity float64 `json:"Quantity"`
}

type PackagesInfoRequest struct {
	PackageLabels []string `json:"packageLabels"`
}

type Receive struct {
	ChildTag string `json:"ChildTag"`
	Eaches []string `json:"Eaches"`
	LabelSource string `json:"LabelSource"`
	QrCount int `json:"QrCount"`
	Ranges [][]int `json:"Ranges"`
	RequiresVerification bool `json:"RequiresVerification"`
	SiblingTags []string `json:"SiblingTags"`
}

type ReceiveQrByShortCode struct {
	ChildTag string `json:"ChildTag"`
	Eaches []string `json:"Eaches"`
	LabelSource string `json:"LabelSource"`
	QrCount int `json:"QrCount"`
	Ranges [][]int `json:"Ranges"`
	RequiresVerification bool `json:"RequiresVerification"`
	SiblingTags []string `json:"SiblingTags"`
}

type RetailIdMergeRequest struct {
	PackageLabels []string `json:"packageLabels"`
}
