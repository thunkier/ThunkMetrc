package models

type AdditivesTemplatesCreateAdditivesTemplatesRequestItem struct {
    ActiveIngredients []*AdditivesTemplatesCreateAdditivesTemplatesRequestItemActiveIngredient `json:"ActiveIngredients,omitempty"`
    AdditiveType string `json:"AdditiveType,omitempty"`
    ApplicationDevice string `json:"ApplicationDevice,omitempty"`
    EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
    Name string `json:"Name,omitempty"`
    Note string `json:"Note,omitempty"`
    ProductSupplier string `json:"ProductSupplier,omitempty"`
    ProductTradeName string `json:"ProductTradeName,omitempty"`
    RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
    RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
}
