package models

type WriteResponse struct {
    Ids []int64 `json:"Ids,omitempty"`
    Warnings *string `json:"Warnings,omitempty"`
}
