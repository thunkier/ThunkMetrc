package models

type PatientsStatu struct {
	Active bool `json:"Active"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed"`
	ConcentrateOuncesAvailable int `json:"ConcentrateOuncesAvailable"`
	ConcentrateOuncesPurchased int `json:"ConcentrateOuncesPurchased"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed"`
	FlowerOuncesAvailable int `json:"FlowerOuncesAvailable"`
	FlowerOuncesPurchased int `json:"FlowerOuncesPurchased"`
	IdentificationMethod *string `json:"IdentificationMethod"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed"`
	InfusedOuncesAvailable int `json:"InfusedOuncesAvailable"`
	InfusedOuncesPurchased int `json:"InfusedOuncesPurchased"`
	PatientLicenseExpirationDate *string `json:"PatientLicenseExpirationDate"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	PurchaseAmountDays int `json:"PurchaseAmountDays"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed"`
	ThcOuncesAvailable int `json:"ThcOuncesAvailable"`
	ThcOuncesPurchased int `json:"ThcOuncesPurchased"`
}

