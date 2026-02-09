package models

type DeliveryPackage struct {
	ContainsRemediatedProduct bool `json:"ContainsRemediatedProduct"`
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId *int `json:"ExternalId"`
	GrossUnitOfWeightName *string `json:"GrossUnitOfWeightName"`
	IsDonation bool `json:"IsDonation"`
	IsTestingSample bool `json:"IsTestingSample"`
	IsTradeSample bool `json:"IsTradeSample"`
	IsTradeSamplePersistent bool `json:"IsTradeSamplePersistent"`
	ItemBrandName *string `json:"ItemBrandName"`
	ItemCategoryName string `json:"ItemCategoryName"`
	ItemId int `json:"ItemId"`
	ItemName string `json:"ItemName"`
	ItemServingSize *string `json:"ItemServingSize"`
	ItemStrainName *string `json:"ItemStrainName"`
	ItemSupplyDurationDays *int `json:"ItemSupplyDurationDays"`
	ItemUnitCbdAContent *float64 `json:"ItemUnitCbdAContent"`
	ItemUnitCbdAContentDose *float64 `json:"ItemUnitCbdAContentDose"`
	ItemUnitCbdAContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdAContentDoseUnitOfMeasureName"`
	ItemUnitCbdAContentUnitOfMeasureName *string `json:"ItemUnitCbdAContentUnitOfMeasureName"`
	ItemUnitCbdAPercent *float64 `json:"ItemUnitCbdAPercent"`
	ItemUnitCbdContent *float64 `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentDose *float64 `json:"ItemUnitCbdContentDose"`
	ItemUnitCbdContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdContentDoseUnitOfMeasureName"`
	ItemUnitCbdContentUnitOfMeasureName *string `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitCbdPercent *float64 `json:"ItemUnitCbdPercent"`
	ItemUnitQuantity *float64 `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName *string `json:"ItemUnitQuantityUnitOfMeasureName"`
	ItemUnitThcAContent *float64 `json:"ItemUnitThcAContent"`
	ItemUnitThcAContentDose *float64 `json:"ItemUnitThcAContentDose"`
	ItemUnitThcAContentDoseUnitOfMeasureName *string `json:"ItemUnitThcAContentDoseUnitOfMeasureName"`
	ItemUnitThcAContentUnitOfMeasureName *string `json:"ItemUnitThcAContentUnitOfMeasureName"`
	ItemUnitThcAPercent *float64 `json:"ItemUnitThcAPercent"`
	ItemUnitThcContent *float64 `json:"ItemUnitThcContent"`
	ItemUnitThcContentDose *float64 `json:"ItemUnitThcContentDose"`
	ItemUnitThcContentDoseUnitOfMeasureName *string `json:"ItemUnitThcContentDoseUnitOfMeasureName"`
	ItemUnitThcContentUnitOfMeasureName *string `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitThcPercent *float64 `json:"ItemUnitThcPercent"`
	ItemUnitVolume *float64 `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName *string `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight *float64 `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName *string `json:"ItemUnitWeightUnitOfMeasureName"`
	LabTestingState string `json:"LabTestingState"`
	PackageId int `json:"PackageId"`
	PackageLabel string `json:"PackageLabel"`
	PackageType string `json:"PackageType"`
	PackagedDate *string `json:"PackagedDate"`
	ProductLabel DeliveryPackageProductLabel `json:"ProductLabel"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	ReceivedQuantity float64 `json:"ReceivedQuantity"`
	ReceivedUnitOfMeasureName string `json:"ReceivedUnitOfMeasureName"`
	RemediationDate *string `json:"RemediationDate"`
	SellByDate *string `json:"SellByDate"`
	ShipmentPackageState string `json:"ShipmentPackageState"`
	ShippedQuantity float64 `json:"ShippedQuantity"`
	ShippedUnitOfMeasureName string `json:"ShippedUnitOfMeasureName"`
	SourceHarvestNames *string `json:"SourceHarvestNames"`
	SourcePackageIsDonation bool `json:"SourcePackageIsDonation"`
	SourcePackageIsTradeSample bool `json:"SourcePackageIsTradeSample"`
	SourcePackageLabels *string `json:"SourcePackageLabels"`
	SourceProductionBatchNumbers *string `json:"SourceProductionBatchNumbers"`
	UseByDate *string `json:"UseByDate"`
}

type DeliveryPackageProductLabel struct {
	IsActive bool `json:"IsActive"`
	IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel"`
	IsDiscontinued bool `json:"IsDiscontinued"`
	LastLabelGenerationDate *string `json:"LastLabelGenerationDate"`
	PackageId *int `json:"PackageId"`
	QrCount int `json:"QrCount"`
	TagId *int `json:"TagId"`
}

type DeliveryPackageRequiredlabtestbatch struct {
	LabTestBatchId int `json:"LabTestBatchId"`
	LabTestBatchName string `json:"LabTestBatchName"`
	PackageId int `json:"PackageId"`
}

type DeliveryPackageWholesale struct {
	PackageId int `json:"PackageId"`
	PackageLabel string `json:"PackageLabel"`
	ReceiverWholesalePrice *string `json:"ReceiverWholesalePrice"`
	ShipperWholesalePrice *string `json:"ShipperWholesalePrice"`
}

type DeliveryTransporter struct {
	TransporterDirection string `json:"TransporterDirection"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName string `json:"TransporterFacilityName"`
}

type DeliveryTransporterDetail struct {
	ActualDriverStartDateTime *string `json:"ActualDriverStartDateTime"`
	DriverLayoverLeg *string `json:"DriverLayoverLeg"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type Hub struct {
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	CreatedByUserName *string `json:"CreatedByUserName"`
	CreatedDateTime string `json:"CreatedDateTime"`
	DeliveryCount int `json:"DeliveryCount"`
	DeliveryId int `json:"DeliveryId"`
	DeliveryPackageCount int `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount int `json:"DeliveryReceivedPackageCount"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	Id int `json:"Id"`
	IsLayover bool `json:"IsLayover"`
	IsVoided bool `json:"IsVoided"`
	LastModified string `json:"LastModified"`
	ManifestNumber string `json:"ManifestNumber"`
	PackageCount int `json:"PackageCount"`
	ReceivedDateTime *string `json:"ReceivedDateTime"`
	ReceivedDeliveryCount int `json:"ReceivedDeliveryCount"`
	ReceivedPackageCount int `json:"ReceivedPackageCount"`
	RecipientFacilityLicenseNumber *string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName *string `json:"RecipientFacilityName"`
	RejectedPackagesReturned bool `json:"RejectedPackagesReturned"`
	ShipmentTransactionType int `json:"ShipmentTransactionType"`
	ShipmentTypeName *string `json:"ShipmentTypeName"`
	ShipperFacilityLicenseNumber string `json:"ShipperFacilityLicenseNumber"`
	ShipperFacilityName string `json:"ShipperFacilityName"`
	TransporterAcceptedDateTime *string `json:"TransporterAcceptedDateTime"`
	TransporterActualArrivalDateTime *string `json:"TransporterActualArrivalDateTime"`
	TransporterActualDepartureDateTime *string `json:"TransporterActualDepartureDateTime"`
	TransporterEstimatedArrivalDateTime *string `json:"TransporterEstimatedArrivalDateTime"`
	TransporterEstimatedDepartureDateTime *string `json:"TransporterEstimatedDepartureDateTime"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type HubArriveRequest struct {
	ShipmentDeliveryId int `json:"ShipmentDeliveryId"`
	TransporterDirection string `json:"TransporterDirection"`
}

type HubCheckinRequest struct {
	ShipmentDeliveryId int `json:"ShipmentDeliveryId"`
	TransporterDirection string `json:"TransporterDirection"`
}

type HubCheckoutRequest struct {
	ShipmentDeliveryId int `json:"ShipmentDeliveryId"`
	TransporterDirection string `json:"TransporterDirection"`
}

type HubDepartRequest struct {
	ShipmentDeliveryId int `json:"ShipmentDeliveryId"`
	TransporterDirection string `json:"TransporterDirection"`
}

type Template struct {
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	ContainsDonation bool `json:"ContainsDonation"`
	ContainsPlantPackage bool `json:"ContainsPlantPackage"`
	ContainsProductPackage bool `json:"ContainsProductPackage"`
	ContainsProductRequiresRemediation bool `json:"ContainsProductRequiresRemediation"`
	ContainsRemediatedProductPackage bool `json:"ContainsRemediatedProductPackage"`
	ContainsTestingSample bool `json:"ContainsTestingSample"`
	ContainsTradeSample bool `json:"ContainsTradeSample"`
	CreatedByUserName *string `json:"CreatedByUserName"`
	CreatedDateTime string `json:"CreatedDateTime"`
	DeliveryCount int `json:"DeliveryCount"`
	DeliveryId int `json:"DeliveryId"`
	DeliveryPackageCount int `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount int `json:"DeliveryReceivedPackageCount"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	Id int `json:"Id"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	IsVoided bool `json:"IsVoided"`
	LastModified string `json:"LastModified"`
	ManifestNumber *string `json:"ManifestNumber"`
	Name string `json:"Name"`
	OriginatingTemplateId *int `json:"OriginatingTemplateId"`
	PackageCount int `json:"PackageCount"`
	ReceivedDateTime *string `json:"ReceivedDateTime"`
	ReceivedDeliveryCount int `json:"ReceivedDeliveryCount"`
	ReceivedPackageCount int `json:"ReceivedPackageCount"`
	RecipientFacilityLicenseNumber *string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName *string `json:"RecipientFacilityName"`
	ShipmentLicenseType int `json:"ShipmentLicenseType"`
	ShipmentTransactionType *string `json:"ShipmentTransactionType"`
	ShipmentTypeName *string `json:"ShipmentTypeName"`
	ShipperFacilityLicenseNumber string `json:"ShipperFacilityLicenseNumber"`
	ShipperFacilityName string `json:"ShipperFacilityName"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TemplateDelivery struct {
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	DeliveryNumber int `json:"DeliveryNumber"`
	DeliveryPackageCount int `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount int `json:"DeliveryReceivedPackageCount"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	GrossUnitOfWeightId *int `json:"GrossUnitOfWeightId"`
	GrossUnitOfWeightName *string `json:"GrossUnitOfWeightName"`
	GrossWeight *float64 `json:"GrossWeight"`
	Id int `json:"Id"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	ManifestNumber *string `json:"ManifestNumber"`
	PDFDocumentFileSystemId *int `json:"PDFDocumentFileSystemId"`
	PlannedRoute string `json:"PlannedRoute"`
	ReceivedDateTime string `json:"ReceivedDateTime"`
	RecipientFacilityLicenseNumber string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName string `json:"RecipientFacilityName"`
	RejectedPackagesReturned bool `json:"RejectedPackagesReturned"`
	ShipmentTransactionType string `json:"ShipmentTransactionType"`
	ShipmentTypeName string `json:"ShipmentTypeName"`
}

type TemplateDeliveryPackage struct {
	ContainsRemediatedProduct bool `json:"ContainsRemediatedProduct"`
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId *int `json:"ExternalId"`
	GrossUnitOfWeightName *string `json:"GrossUnitOfWeightName"`
	IsDonation bool `json:"IsDonation"`
	IsTestingSample bool `json:"IsTestingSample"`
	IsTradeSample bool `json:"IsTradeSample"`
	IsTradeSamplePersistent bool `json:"IsTradeSamplePersistent"`
	ItemBrandName *string `json:"ItemBrandName"`
	ItemCategoryName string `json:"ItemCategoryName"`
	ItemId int `json:"ItemId"`
	ItemName string `json:"ItemName"`
	ItemServingSize *string `json:"ItemServingSize"`
	ItemStrainName *string `json:"ItemStrainName"`
	ItemSupplyDurationDays *int `json:"ItemSupplyDurationDays"`
	ItemUnitCbdAContent *float64 `json:"ItemUnitCbdAContent"`
	ItemUnitCbdAContentDose *float64 `json:"ItemUnitCbdAContentDose"`
	ItemUnitCbdAContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdAContentDoseUnitOfMeasureName"`
	ItemUnitCbdAContentUnitOfMeasureName *string `json:"ItemUnitCbdAContentUnitOfMeasureName"`
	ItemUnitCbdAPercent *float64 `json:"ItemUnitCbdAPercent"`
	ItemUnitCbdContent *float64 `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentDose *float64 `json:"ItemUnitCbdContentDose"`
	ItemUnitCbdContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdContentDoseUnitOfMeasureName"`
	ItemUnitCbdContentUnitOfMeasureName *string `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitCbdPercent *float64 `json:"ItemUnitCbdPercent"`
	ItemUnitQuantity *float64 `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName *string `json:"ItemUnitQuantityUnitOfMeasureName"`
	ItemUnitThcAContent *float64 `json:"ItemUnitThcAContent"`
	ItemUnitThcAContentDose *float64 `json:"ItemUnitThcAContentDose"`
	ItemUnitThcAContentDoseUnitOfMeasureName *string `json:"ItemUnitThcAContentDoseUnitOfMeasureName"`
	ItemUnitThcAContentUnitOfMeasureName *string `json:"ItemUnitThcAContentUnitOfMeasureName"`
	ItemUnitThcAPercent *float64 `json:"ItemUnitThcAPercent"`
	ItemUnitThcContent *float64 `json:"ItemUnitThcContent"`
	ItemUnitThcContentDose *float64 `json:"ItemUnitThcContentDose"`
	ItemUnitThcContentDoseUnitOfMeasureName *string `json:"ItemUnitThcContentDoseUnitOfMeasureName"`
	ItemUnitThcContentUnitOfMeasureName *string `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitThcPercent *float64 `json:"ItemUnitThcPercent"`
	ItemUnitVolume *float64 `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName *string `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight *float64 `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName *string `json:"ItemUnitWeightUnitOfMeasureName"`
	LabTestingState string `json:"LabTestingState"`
	PackageId int `json:"PackageId"`
	PackageLabel string `json:"PackageLabel"`
	PackageType string `json:"PackageType"`
	PackagedDate *string `json:"PackagedDate"`
	ProductLabel TemplateDeliveryPackageProductLabel `json:"ProductLabel"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation"`
	ProductionBatchNumber *string `json:"ProductionBatchNumber"`
	ReceivedQuantity float64 `json:"ReceivedQuantity"`
	ReceivedUnitOfMeasureName string `json:"ReceivedUnitOfMeasureName"`
	RemediationDate *string `json:"RemediationDate"`
	SellByDate *string `json:"SellByDate"`
	ShipmentPackageState string `json:"ShipmentPackageState"`
	ShippedQuantity float64 `json:"ShippedQuantity"`
	ShippedUnitOfMeasureName string `json:"ShippedUnitOfMeasureName"`
	SourceHarvestNames *string `json:"SourceHarvestNames"`
	SourcePackageIsDonation bool `json:"SourcePackageIsDonation"`
	SourcePackageIsTradeSample bool `json:"SourcePackageIsTradeSample"`
	SourcePackageLabels *string `json:"SourcePackageLabels"`
	SourceProductionBatchNumbers *string `json:"SourceProductionBatchNumbers"`
	UseByDate *string `json:"UseByDate"`
}

type TemplateDeliveryPackageProductLabel struct {
	IsActive bool `json:"IsActive"`
	IsChildFromParentWithLabel bool `json:"IsChildFromParentWithLabel"`
	IsDiscontinued bool `json:"IsDiscontinued"`
	LastLabelGenerationDate *string `json:"LastLabelGenerationDate"`
	PackageId *int `json:"PackageId"`
	QrCount int `json:"QrCount"`
	TagId *int `json:"TagId"`
}

type TemplateDeliveryTransporter struct {
	TransporterDirection string `json:"TransporterDirection"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName string `json:"TransporterFacilityName"`
}

type TemplateDeliveryTransporterDetail struct {
	ActualDriverStartDateTime *string `json:"ActualDriverStartDateTime"`
	DriverLayoverLeg *string `json:"DriverLayoverLeg"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber string `json:"DriverVehicleLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TransfersDelivery struct {
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	DeliveryNumber int `json:"DeliveryNumber"`
	DeliveryPackageCount int `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount int `json:"DeliveryReceivedPackageCount"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	GrossUnitOfWeightId *int `json:"GrossUnitOfWeightId"`
	GrossUnitOfWeightName *string `json:"GrossUnitOfWeightName"`
	GrossWeight *float64 `json:"GrossWeight"`
	Id int `json:"Id"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	ManifestNumber *string `json:"ManifestNumber"`
	PDFDocumentFileSystemId *int `json:"PDFDocumentFileSystemId"`
	PlannedRoute string `json:"PlannedRoute"`
	ReceivedDateTime string `json:"ReceivedDateTime"`
	RecipientFacilityLicenseNumber string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName string `json:"RecipientFacilityName"`
	RejectedPackagesReturned bool `json:"RejectedPackagesReturned"`
	ShipmentTransactionType string `json:"ShipmentTransactionType"`
	ShipmentTypeName string `json:"ShipmentTypeName"`
}

type TransfersExternalRequest struct {
	Destinations []TransfersExternalRequestDestination `json:"Destinations"`
	DriverLicenseNumber *string `json:"DriverLicenseNumber"`
	DriverName *string `json:"DriverName"`
	DriverOccupationalLicenseNumber *string `json:"DriverOccupationalLicenseNumber"`
	PhoneNumberForQuestions *string `json:"PhoneNumberForQuestions"`
	ShipperAddress1 string `json:"ShipperAddress1"`
	ShipperAddress2 *string `json:"ShipperAddress2"`
	ShipperAddressCity string `json:"ShipperAddressCity"`
	ShipperAddressPostalCode *string `json:"ShipperAddressPostalCode"`
	ShipperAddressState string `json:"ShipperAddressState"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber"`
	ShipperName string `json:"ShipperName"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	VehicleLicensePlateNumber *string `json:"VehicleLicensePlateNumber"`
	VehicleMake *string `json:"VehicleMake"`
	VehicleModel *string `json:"VehicleModel"`
}

type TransfersExternalRequestDestination struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	GrossUnitOfWeightId *int `json:"GrossUnitOfWeightId"`
	GrossWeight *float64 `json:"GrossWeight"`
	InvoiceNumber string `json:"InvoiceNumber"`
	Packages []TransfersExternalRequestDestinationPackage `json:"Packages"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber"`
	TransferTypeName string `json:"TransferTypeName"`
	Transporters []TransfersExternalRequestDestinationTransporter `json:"Transporters"`
}

type TransfersExternalRequestDestinationPackage struct {
	ExpirationDate *string `json:"ExpirationDate"`
	ExternalId string `json:"ExternalId"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName"`
	GrossWeight float64 `json:"GrossWeight"`
	ItemName string `json:"ItemName"`
	PackagedDate string `json:"PackagedDate"`
	Quantity float64 `json:"Quantity"`
	SellByDate *string `json:"SellByDate"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
	UseByDate *string `json:"UseByDate"`
	WholesalePrice *string `json:"WholesalePrice"`
}

type TransfersExternalRequestDestinationTransporter struct {
	DriverLayoverLeg string `json:"DriverLayoverLeg"`
	DriverLicenseNumber string `json:"DriverLicenseNumber"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	IsLayover bool `json:"IsLayover"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions"`
	TransporterDetails []TransfersExternalRequestDestinationTransporterTransporterDetail `json:"TransporterDetails"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TransfersExternalRequestDestinationTransporterTransporterDetail struct {
	DriverLayoverLeg string `json:"DriverLayoverLeg"`
	DriverLicenseNumber string `json:"DriverLicenseNumber"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TransfersTemplatesRequest struct {
	Destinations []TransfersTemplatesRequestDestination `json:"Destinations"`
	DriverLicenseNumber *string `json:"DriverLicenseNumber"`
	DriverName *string `json:"DriverName"`
	DriverOccupationalLicenseNumber *string `json:"DriverOccupationalLicenseNumber"`
	Name string `json:"Name"`
	PhoneNumberForQuestions *string `json:"PhoneNumberForQuestions"`
	TransferTemplateId int `json:"TransferTemplateId"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	VehicleLicensePlateNumber *string `json:"VehicleLicensePlateNumber"`
	VehicleMake *string `json:"VehicleMake"`
	VehicleModel *string `json:"VehicleModel"`
}

type TransfersTemplatesRequestDestination struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	InvoiceNumber string `json:"InvoiceNumber"`
	Packages []TransfersTemplatesRequestDestinationPackage `json:"Packages"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber"`
	TransferDestinationId int `json:"TransferDestinationId"`
	TransferTypeName string `json:"TransferTypeName"`
	Transporters []TransfersTemplatesRequestDestinationTransporter `json:"Transporters"`
}

type TransfersTemplatesRequestDestinationPackage struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName"`
	GrossWeight float64 `json:"GrossWeight"`
	PackageLabel string `json:"PackageLabel"`
	WholesalePrice *string `json:"WholesalePrice"`
}

type TransfersTemplatesRequestDestinationTransporter struct {
	DriverLayoverLeg string `json:"DriverLayoverLeg"`
	DriverLicenseNumber string `json:"DriverLicenseNumber"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	IsLayover bool `json:"IsLayover"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions"`
	TransporterDetails []TransfersTemplatesRequestDestinationTransporterTransporterDetail `json:"TransporterDetails"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TransfersTemplatesRequestDestinationTransporterTransporterDetail struct {
	DriverLayoverLeg string `json:"DriverLayoverLeg"`
	DriverLicenseNumber string `json:"DriverLicenseNumber"`
	DriverName string `json:"DriverName"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type TransfersType struct {
	BypassApproval bool `json:"BypassApproval"`
	ExternalIncomingCanRecordExternalIdentifier bool `json:"ExternalIncomingCanRecordExternalIdentifier"`
	ExternalIncomingExternalIdentifierRequired bool `json:"ExternalIncomingExternalIdentifierRequired"`
	ExternalOutgoingCanRecordExternalIdentifier bool `json:"ExternalOutgoingCanRecordExternalIdentifier"`
	ExternalOutgoingExternalIdentifierRequired bool `json:"ExternalOutgoingExternalIdentifierRequired"`
	ForExternalIncomingShipments bool `json:"ForExternalIncomingShipments"`
	ForExternalOutgoingShipments bool `json:"ForExternalOutgoingShipments"`
	ForLicensedShipments bool `json:"ForLicensedShipments"`
	Name string `json:"Name"`
	RequiresDestinationGrossWeight bool `json:"RequiresDestinationGrossWeight"`
	RequiresInvoiceNumber bool `json:"RequiresInvoiceNumber"`
	RequiresPDFDocument bool `json:"RequiresPDFDocument"`
	RequiresPackagesGrossWeight bool `json:"RequiresPackagesGrossWeight"`
	TransactionType string `json:"TransactionType"`
}
