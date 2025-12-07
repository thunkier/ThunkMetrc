package models

type AdditivesTemplate struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	ActiveIngredients []AdditivesTemplateActiveIngredient `json:"ActiveIngredients"`
	AdditiveType string `json:"AdditiveType"`
	AdditiveTypeName *string `json:"AdditiveTypeName"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsActive bool `json:"IsActive"`
	LastModified string `json:"LastModified"`
	Name *string `json:"Name"`
	Note *string `json:"Note"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription"`
	RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription"`
	CurrentPage int `json:"CurrentPage"`
	Data []AdditivesTemplateData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
}

type AdditivesTemplateActiveIngredient struct {
	Name string `json:"Name"`
	Percentage float64 `json:"Percentage"`
}

type AdditivesTemplateData struct {
	ActiveIngredients []AdditivesTemplateDataActiveIngredient `json:"ActiveIngredients"`
	AdditiveType string `json:"AdditiveType"`
	AdditiveTypeName *string `json:"AdditiveTypeName"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsActive bool `json:"IsActive"`
	LastModified string `json:"LastModified"`
	Name *string `json:"Name"`
	Note *string `json:"Note"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription"`
	RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription"`
}

type AdditivesTemplateDataActiveIngredient struct {
	Name string `json:"Name"`
	Percentage float64 `json:"Percentage"`
}

