package models

type CaregiversStatu struct {
	Active bool `json:"Active"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber"`
	Patients []string `json:"Patients"`
}

