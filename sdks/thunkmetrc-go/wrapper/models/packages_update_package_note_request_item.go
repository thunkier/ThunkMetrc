package models

type PackagesUpdatePackageNoteRequestItem struct {
    Note string `json:"Note,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
}
