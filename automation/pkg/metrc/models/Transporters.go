package models

type Transporter struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EmployeeId string `json:"EmployeeId"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	CurrentPage int `json:"CurrentPage"`
	Data []TransporterData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
	LicensePlateNumber string `json:"LicensePlateNumber"`
	Make string `json:"Make"`
	Model string `json:"Model"`
}

type TransporterData struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EmployeeId string `json:"EmployeeId"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	Name string `json:"Name"`
	LicensePlateNumber string `json:"LicensePlateNumber"`
	Make string `json:"Make"`
	Model string `json:"Model"`
}

