package models

type PackagesUpdatePackageLocationRequestItem struct {
    Label string `json:"Label,omitempty"`
    Location string `json:"Location,omitempty"`
    MoveDate string `json:"MoveDate,omitempty"`
    Sublocation string `json:"Sublocation,omitempty"`
}
