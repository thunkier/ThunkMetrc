package models

type WebhooksUpdateWebhooksRequest struct {
    ErrorResponseJsonTemplate string `json:"errorResponseJsonTemplate,omitempty"`
    FacilityLicenseNumbers string `json:"facilityLicenseNumbers,omitempty"`
    ObjectType string `json:"objectType,omitempty"`
    ServerPublicKeyFingerprint string `json:"serverPublicKeyFingerprint,omitempty"`
    Status string `json:"status,omitempty"`
    Template string `json:"template,omitempty"`
    TpiApiKey string `json:"tpiApiKey,omitempty"`
    Url string `json:"url,omitempty"`
    UserApiKey string `json:"userApiKey,omitempty"`
    Verb string `json:"verb,omitempty"`
}
