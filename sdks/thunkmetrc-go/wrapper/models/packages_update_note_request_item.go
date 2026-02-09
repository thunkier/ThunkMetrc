package models

type PackagesUpdateNoteRequestItem struct {
    Note string `json:"Note,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
}
