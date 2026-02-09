package models

type PatientCheckIn struct {
    CheckInDate string `json:"CheckInDate,omitempty"`
    CheckInLocationId int64 `json:"CheckInLocationId,omitempty"`
    CheckInLocationName string `json:"CheckInLocationName,omitempty"`
    Id int64 `json:"Id,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
    RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
}
