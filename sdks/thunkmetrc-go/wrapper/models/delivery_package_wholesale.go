package models

type DeliveryPackageWholesale struct {
    PackageId int64 `json:"PackageId,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
    ReceiverWholesalePrice *string `json:"ReceiverWholesalePrice,omitempty"`
    ShipperWholesalePrice *string `json:"ShipperWholesalePrice,omitempty"`
}
