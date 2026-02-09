package models

type AdditivesTemplatesRequest struct {
	ActiveIngredients []AdditivesTemplatesRequestActiveIngredient `json:"ActiveIngredients"`
	AdditiveType string `json:"AdditiveType"`
	ApplicationDevice string `json:"ApplicationDevice"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	Note string `json:"Note"`
	ProductSupplier string `json:"ProductSupplier"`
	ProductTradeName string `json:"ProductTradeName"`
	RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription"`
	RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription"`
}

type AdditivesTemplatesRequestActiveIngredient struct {
	Name string `json:"Name"`
	Percentage float64 `json:"Percentage"`
}
