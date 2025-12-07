package models

type Sublocation struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	CurrentPage int `json:"CurrentPage"`
	Data []SublocationData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type SublocationData struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

