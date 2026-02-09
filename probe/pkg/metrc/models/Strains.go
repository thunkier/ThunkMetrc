package models

type StrainsRequest struct {
	CbdLevel float64 `json:"CbdLevel"`
	Id int `json:"Id"`
	IndicaPercentage int `json:"IndicaPercentage"`
	Name string `json:"Name"`
	SativaPercentage int `json:"SativaPercentage"`
	TestingStatus string `json:"TestingStatus"`
	ThcLevel float64 `json:"ThcLevel"`
}
