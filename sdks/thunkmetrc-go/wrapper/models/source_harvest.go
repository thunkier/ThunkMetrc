package models

type SourceHarvest struct {
    HarvestStartDate string `json:"HarvestStartDate,omitempty"`
    HarvestedByFacilityLicenseNumber string `json:"HarvestedByFacilityLicenseNumber,omitempty"`
    HarvestedByFacilityName string `json:"HarvestedByFacilityName,omitempty"`
    Name string `json:"Name,omitempty"`
}
