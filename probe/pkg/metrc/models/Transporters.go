package models

type TransportersDriver struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EmployeeId string `json:"EmployeeId"`
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsArchived bool `json:"IsArchived"`
	LastModified string `json:"LastModified"`
	Name string `json:"Name"`
}

type TransportersDriversRequest struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EmployeeId string `json:"EmployeeId"`
	Name string `json:"Name"`
}

type TransportersVehicle struct {
	FacilityId int `json:"FacilityId"`
	Id int `json:"Id"`
	IsArchived bool `json:"IsArchived"`
	LastModified string `json:"LastModified"`
	LicensePlateNumber string `json:"LicensePlateNumber"`
	Make string `json:"Make"`
	Model string `json:"Model"`
}

type TransportersVehiclesRequest struct {
	Id string `json:"Id"`
	LicensePlateNumber string `json:"LicensePlateNumber"`
	Make string `json:"Make"`
	Model string `json:"Model"`
}
