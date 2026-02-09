package models

type StrainsCreateStrainsRequestItem struct {
    CbdLevel float64 `json:"CbdLevel,omitempty"`
    IndicaPercentage int `json:"IndicaPercentage,omitempty"`
    Name string `json:"Name,omitempty"`
    SativaPercentage int `json:"SativaPercentage,omitempty"`
    TestingStatus string `json:"TestingStatus,omitempty"`
    ThcLevel float64 `json:"ThcLevel,omitempty"`
}
