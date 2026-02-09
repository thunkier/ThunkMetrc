package models

type PatientCheckInsUpdatePatientCheckInsRequestItem struct {
    CheckInDate string `json:"CheckInDate,omitempty"`
    CheckInLocationId int `json:"CheckInLocationId,omitempty"`
    Id int `json:"Id,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
    RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
}
