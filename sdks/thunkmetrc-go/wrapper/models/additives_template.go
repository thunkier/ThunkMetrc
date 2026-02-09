package models

type AdditivesTemplate struct {
    ActiveIngredients []AdditivesTemplateActiveIngredientsItem `json:"ActiveIngredients,omitempty"`
    AdditiveType string `json:"AdditiveType,omitempty"`
    AdditiveTypeName *string `json:"AdditiveTypeName,omitempty"`
    ApplicationDevice string `json:"ApplicationDevice,omitempty"`
    EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
    FacilityId int64 `json:"FacilityId,omitempty"`
    Id int64 `json:"Id,omitempty"`
    IsActive bool `json:"IsActive,omitempty"`
    LastModified string `json:"LastModified,omitempty"`
    Name *string `json:"Name,omitempty"`
    Note *string `json:"Note,omitempty"`
    ProductSupplier string `json:"ProductSupplier,omitempty"`
    ProductTradeName string `json:"ProductTradeName,omitempty"`
    RestrictiveEntryIntervalQuantityDescription *string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
    RestrictiveEntryIntervalTimeDescription *string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
}
