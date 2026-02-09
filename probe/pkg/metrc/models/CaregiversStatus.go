package models

type CaregiversStatus struct {
	Active bool `json:"Active"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber"`
	Patients []string `json:"Patients"`
}
