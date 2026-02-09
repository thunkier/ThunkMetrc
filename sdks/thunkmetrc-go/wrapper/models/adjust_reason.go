package models

type AdjustReason struct {
    Name string `json:"Name,omitempty"`
    RequiresImmatureWasteWeight bool `json:"RequiresImmatureWasteWeight,omitempty"`
    RequiresMatureWasteWeight bool `json:"RequiresMatureWasteWeight,omitempty"`
    RequiresNote bool `json:"RequiresNote,omitempty"`
    RequiresWasteWeight bool `json:"RequiresWasteWeight,omitempty"`
}
