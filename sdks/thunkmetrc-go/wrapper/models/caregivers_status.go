package models

type CaregiversStatus struct {
    Active bool `json:"Active,omitempty"`
    CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
    Patients []string `json:"Patients,omitempty"`
}
