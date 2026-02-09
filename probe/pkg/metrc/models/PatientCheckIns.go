package models

type PatientCheckIn struct {
	CheckInDate string `json:"CheckInDate"`
	CheckInLocationId int `json:"CheckInLocationId"`
	CheckInLocationName string `json:"CheckInLocationName"`
	Id int `json:"Id"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate"`
	RegistrationStartDate string `json:"RegistrationStartDate"`
}

type PatientCheckInsLocation struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

type PatientCheckInsRequest struct {
	CheckInDate string `json:"CheckInDate"`
	CheckInLocationId int `json:"CheckInLocationId"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate"`
	RegistrationStartDate string `json:"RegistrationStartDate"`
}
