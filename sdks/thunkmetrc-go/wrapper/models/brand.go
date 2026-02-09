package models

type Brand struct {
    Id int64 `json:"Id,omitempty"`
    Name string `json:"Name,omitempty"`
    Status string `json:"Status,omitempty"`
}
