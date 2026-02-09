package models

type RetailIdCreateAssociateRequestItem struct {
    PackageLabel string `json:"PackageLabel,omitempty"`
    QrUrls []string `json:"QrUrls,omitempty"`
}
