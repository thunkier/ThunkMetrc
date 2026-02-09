package models

type Strain struct {
    CbdLevel *string `json:"CbdLevel,omitempty"`
    Genetics string `json:"Genetics,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IndicaPercentage int64 `json:"IndicaPercentage,omitempty"`
    IsUsed bool `json:"IsUsed,omitempty"`
    Name string `json:"Name,omitempty"`
    SativaPercentage int64 `json:"SativaPercentage,omitempty"`
    TestingStatus string `json:"TestingStatus,omitempty"`
    ThcLevel *string `json:"ThcLevel,omitempty"`
}
