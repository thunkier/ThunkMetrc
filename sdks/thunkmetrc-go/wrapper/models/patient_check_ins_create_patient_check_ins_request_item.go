package models

type PatientCheckInsCreatePatientCheckInsRequestItem struct {
    CheckInDate string `json:"CheckInDate,omitempty"`
    CheckInLocationId int `json:"CheckInLocationId,omitempty"`
    PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
    RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
    RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
}
