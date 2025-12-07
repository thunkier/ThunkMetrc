package models

type RetailId struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	IssuanceId string `json:"IssuanceId"`
	Packages []RetailIdPackage `json:"Packages"`
	ChildTag string `json:"ChildTag"`
	Eaches []string `json:"Eaches"`
	LabelSource string `json:"LabelSource"`
	QrCount int `json:"QrCount"`
	Ranges string `json:"Ranges"`
	RequiresVerification bool `json:"RequiresVerification"`
	SiblingTags []string `json:"SiblingTags"`
}

type RetailIdPackage struct {
	EstimatedBalance int `json:"EstimatedBalance"`
	HasQrs bool `json:"HasQrs"`
	IssuanceId string `json:"IssuanceId"`
	Issuances []RetailIdPackageIssuance `json:"Issuances"`
	QrCount int `json:"QrCount"`
	RequiresVerification bool `json:"RequiresVerification"`
	SiblingCount int `json:"SiblingCount"`
	Tag string `json:"Tag"`
}

type RetailIdPackageIssuance struct {
	CreatedAt string `json:"CreatedAt"`
	LabelQuantity int `json:"LabelQuantity"`
	LabelSet int `json:"LabelSet"`
}

