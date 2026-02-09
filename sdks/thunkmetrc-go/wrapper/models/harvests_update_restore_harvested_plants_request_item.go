package models

type HarvestsUpdateRestoreHarvestedPlantsRequestItem struct {
    HarvestId int `json:"HarvestId,omitempty"`
    PlantTags []string `json:"PlantTags,omitempty"`
}
