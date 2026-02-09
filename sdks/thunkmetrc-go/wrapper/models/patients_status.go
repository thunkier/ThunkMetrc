package models

type PatientsStatus struct {
    Active bool `json:"Active,omitempty"`
    ConcentrateOuncesAllowed int64 `json:"ConcentrateOuncesAllowed,omitempty"`
    ConcentrateOuncesAvailable int64 `json:"ConcentrateOuncesAvailable,omitempty"`
    ConcentrateOuncesPurchased int64 `json:"ConcentrateOuncesPurchased,omitempty"`
    FlowerOuncesAllowed int64 `json:"FlowerOuncesAllowed,omitempty"`
    FlowerOuncesAvailable int64 `json:"FlowerOuncesAvailable,omitempty"`
    FlowerOuncesPurchased int64 `json:"FlowerOuncesPurchased,omitempty"`
    IdentificationMethod *string `json:"IdentificationMethod,omitempty"`
    InfusedOuncesAllowed int64 `json:"InfusedOuncesAllowed,omitempty"`
    InfusedOuncesAvailable int64 `json:"InfusedOuncesAvailable,omitempty"`
    InfusedOuncesPurchased int64 `json:"InfusedOuncesPurchased,omitempty"`
    PatientLicenseExpirationDate string `json:"PatientLicenseExpirationDate,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    PatientRegistrationStartDate string `json:"PatientRegistrationStartDate,omitempty"`
    PurchaseAmountDays int64 `json:"PurchaseAmountDays,omitempty"`
    RegistrationNumber string `json:"RegistrationNumber,omitempty"`
    ThcOuncesAllowed int64 `json:"ThcOuncesAllowed,omitempty"`
    ThcOuncesAvailable int64 `json:"ThcOuncesAvailable,omitempty"`
    ThcOuncesPurchased int64 `json:"ThcOuncesPurchased,omitempty"`
}
