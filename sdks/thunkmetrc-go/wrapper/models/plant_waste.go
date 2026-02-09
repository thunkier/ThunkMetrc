package models

type PlantWaste struct {
    GrowthPhase int64 `json:"GrowthPhase,omitempty"`
    Label string `json:"Label,omitempty"`
    LocationId int64 `json:"LocationId,omitempty"`
    LocationName string `json:"LocationName,omitempty"`
    PlantId int64 `json:"PlantId,omitempty"`
    PlantWasteId int64 `json:"PlantWasteId,omitempty"`
    StateName string `json:"StateName,omitempty"`
    SublocationId int64 `json:"SublocationId,omitempty"`
    SublocationName *string `json:"SublocationName,omitempty"`
    WasteAmount int64 `json:"WasteAmount,omitempty"`
    WasteUnitOfMeasureAbbreviation string `json:"WasteUnitOfMeasureAbbreviation,omitempty"`
}
