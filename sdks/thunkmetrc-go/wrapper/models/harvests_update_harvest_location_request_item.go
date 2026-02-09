package models

type HarvestsUpdateHarvestLocationRequestItem struct {
    ActualDate string `json:"ActualDate,omitempty"`
    DryingLocation string `json:"DryingLocation,omitempty"`
    DryingSublocation string `json:"DryingSublocation,omitempty"`
    HarvestName string `json:"HarvestName,omitempty"`
    Id int `json:"Id,omitempty"`
}
