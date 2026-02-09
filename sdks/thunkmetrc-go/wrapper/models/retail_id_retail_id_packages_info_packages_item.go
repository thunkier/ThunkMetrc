package models

type RetailIdRetailIdPackagesInfoPackagesItem struct {
    EstimatedBalance int64 `json:"EstimatedBalance,omitempty"`
    HasQrs bool `json:"HasQrs,omitempty"`
    IssuanceId string `json:"IssuanceId,omitempty"`
    Issuances []RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem `json:"Issuances,omitempty"`
    QrCount int64 `json:"QrCount,omitempty"`
    RequiresVerification bool `json:"RequiresVerification,omitempty"`
    SiblingCount int64 `json:"SiblingCount,omitempty"`
    Tag string `json:"Tag,omitempty"`
}
