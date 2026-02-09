package models

type SalesUpdateDeliveriesCompleteRequestItem struct {
    AcceptedPackages []string `json:"AcceptedPackages,omitempty"`
    ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
    Id int `json:"Id,omitempty"`
    PaymentType string `json:"PaymentType,omitempty"`
    ReturnedPackages []*SalesUpdateDeliveriesCompleteRequestItemReturnedPackage `json:"ReturnedPackages,omitempty"`
}
