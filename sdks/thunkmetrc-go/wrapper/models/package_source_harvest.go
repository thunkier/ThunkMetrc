package models

type PackageSourceHarvest struct {
    HarvestStartDate string `json:"HarvestStartDate,omitempty"`
    HarvestedByFacilityLicenseNumber string `json:"HarvestedByFacilityLicenseNumber,omitempty"`
    HarvestedByFacilityName string `json:"HarvestedByFacilityName,omitempty"`
    Name string `json:"Name,omitempty"`
}
