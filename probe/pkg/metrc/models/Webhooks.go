package models

type WebhooksRequest struct {
	ErrorResponseJsonTemplate string `json:"errorResponseJsonTemplate"`
	FacilityLicenseNumbers string `json:"facilityLicenseNumbers"`
	ObjectType string `json:"objectType"`
	ServerPublicKeyFingerprint string `json:"serverPublicKeyFingerprint"`
	Status string `json:"status"`
	Template string `json:"template"`
	TpiApiKey string `json:"tpiApiKey"`
	Url string `json:"url"`
	UserApiKey string `json:"userApiKey"`
	Verb string `json:"verb"`
}
