package models

type HarvestsUpdateHarvestRenameRequestItem struct {
    Id int `json:"Id,omitempty"`
    NewName string `json:"NewName,omitempty"`
    OldName string `json:"OldName,omitempty"`
}
