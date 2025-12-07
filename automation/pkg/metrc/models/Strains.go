package models

type Strain struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	CbdLevel *float64 `json:"CbdLevel"`
	Genetics string `json:"Genetics"`
	Id int `json:"Id"`
	IndicaPercentage float64 `json:"IndicaPercentage"`
	IsUsed bool `json:"IsUsed"`
	Name string `json:"Name"`
	SativaPercentage float64 `json:"SativaPercentage"`
	TestingStatus string `json:"TestingStatus"`
	ThcLevel *float64 `json:"ThcLevel"`
	CurrentPage int `json:"CurrentPage"`
	Data []StrainData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type StrainData struct {
	CbdLevel *float64 `json:"CbdLevel"`
	Genetics string `json:"Genetics"`
	Id int `json:"Id"`
	IndicaPercentage float64 `json:"IndicaPercentage"`
	IsUsed bool `json:"IsUsed"`
	Name string `json:"Name"`
	SativaPercentage float64 `json:"SativaPercentage"`
	TestingStatus string `json:"TestingStatus"`
	ThcLevel *float64 `json:"ThcLevel"`
}

