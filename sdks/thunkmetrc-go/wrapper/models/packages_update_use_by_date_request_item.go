package models

type PackagesUpdateUseByDateRequestItem struct {
    ExpirationDate string `json:"ExpirationDate,omitempty"`
    Label string `json:"Label,omitempty"`
    SellByDate string `json:"SellByDate,omitempty"`
    UseByDate string `json:"UseByDate,omitempty"`
}
