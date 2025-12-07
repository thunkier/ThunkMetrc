package models

type PatientCheckIn struct {
	CheckInDate string `json:"CheckInDate"`
	CheckInLocationId int `json:"CheckInLocationId"`
	CheckInLocationName string `json:"CheckInLocationName"`
	Id int `json:"Id"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate"`
	RegistrationStartDate string `json:"RegistrationStartDate"`
	CurrentPage int `json:"CurrentPage"`
	Data []PatientCheckInData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
	Name string `json:"Name"`
}

type PatientCheckInData struct {
	CheckInDate string `json:"CheckInDate"`
	CheckInLocationId int `json:"CheckInLocationId"`
	CheckInLocationName string `json:"CheckInLocationName"`
	Id int `json:"Id"`
	PatientLicenseNumber string `json:"PatientLicenseNumber"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate"`
	RegistrationStartDate string `json:"RegistrationStartDate"`
}

