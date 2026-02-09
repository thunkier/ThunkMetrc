package models

type ActiveDeliveriesRetailer struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	AllowFullEdit bool `json:"AllowFullEdit"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	DateTime string `json:"DateTime"`
	Destinations []interface{} `json:"Destinations"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	FacilityLicenseNumber string `json:"FacilityLicenseNumber"`
	FacilityName string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Leg int `json:"Leg"`
	Packages []interface{} `json:"Packages"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	RestockDateTime *string `json:"RestockDateTime"`
	RetailerDeliveryNumber string `json:"RetailerDeliveryNumber"`
	RetailerDeliveryState string `json:"RetailerDeliveryState"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	TotalPriceSold int `json:"TotalPriceSold"`
	VehicleInfo string `json:"VehicleInfo"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type ActiveDelivery struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	ConsumerId string `json:"ConsumerId"`
	DeliveryNumber *string `json:"DeliveryNumber"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	FacilityLicenseNumber *string `json:"FacilityLicenseNumber"`
	FacilityName *string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientDeliveryZoneId *int `json:"RecipientDeliveryZoneId"`
	RecipientDeliveryZoneName *string `json:"RecipientDeliveryZoneName"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	SalesDeliveryState string `json:"SalesDeliveryState"`
	ShipperFacilityLicenseNumber *string `json:"ShipperFacilityLicenseNumber"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
	TransporterFacilityId *int `json:"TransporterFacilityId"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName *string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type ActiveReceipt struct {
	ArchivedDate *string `json:"ArchivedDate"`
	CaregiverLicenseNumber *string `json:"CaregiverLicenseNumber"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber"`
	Id int `json:"Id"`
	IdentificationMethod *string `json:"IdentificationMethod"`
	IsFinal bool `json:"IsFinal"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PatientRegistrationLocationId *int `json:"PatientRegistrationLocationId"`
	ReceiptNumber *string `json:"ReceiptNumber"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
}

type County struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

type DeliveriesCompleteRequest struct {
	AcceptedPackages []string `json:"AcceptedPackages"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime"`
	Id int `json:"Id"`
	PaymentType *string `json:"PaymentType"`
	ReturnedPackages []DeliveriesCompleteRequestReturnedPackage `json:"ReturnedPackages"`
}

type DeliveriesCompleteRequestReturnedPackage struct {
	Label string `json:"Label"`
	ReturnQuantityVerified int `json:"ReturnQuantityVerified"`
	ReturnReason string `json:"ReturnReason"`
	ReturnReasonNote string `json:"ReturnReasonNote"`
	ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure"`
}

type DeliveriesHubAcceptRequest struct {
	Id int `json:"Id"`
}

type DeliveriesHubDepartRequest struct {
	Id int `json:"Id"`
}

type DeliveriesHubRequest struct {
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	Id int `json:"Id"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions"`
	PlannedRoute string `json:"PlannedRoute"`
	TransporterFacilityId string `json:"TransporterFacilityId"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type DeliveriesHubVerifyIDRequest struct {
	Id int `json:"Id"`
	PaymentType string `json:"PaymentType"`
}

type DeliveriesRetailerDepartRequest struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId"`
}

type DeliveriesRetailerEndRequest struct {
	ActualArrivalDateTime string `json:"ActualArrivalDateTime"`
	Packages []DeliveriesRetailerEndRequestPackage `json:"Packages"`
	RetailerDeliveryId int `json:"RetailerDeliveryId"`
}

type DeliveriesRetailerEndRequestPackage struct {
	EndQuantity float64 `json:"EndQuantity"`
	EndUnitOfMeasure string `json:"EndUnitOfMeasure"`
	Label string `json:"Label"`
}

type DeliveriesRetailerRestockRequest struct {
	DateTime string `json:"DateTime"`
	Destinations *string `json:"Destinations"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	Packages []DeliveriesRetailerRestockRequestPackage `json:"Packages"`
	RetailerDeliveryId int `json:"RetailerDeliveryId"`
}

type DeliveriesRetailerRestockRequestPackage struct {
	PackageLabel string `json:"PackageLabel"`
	Quantity float64 `json:"Quantity"`
	RemoveCurrentPackage bool `json:"RemoveCurrentPackage"`
	TotalPrice float64 `json:"TotalPrice"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
}

type DeliveriesReturnReason struct {
	Name string `json:"Name"`
	RequiresImmatureWasteWeight bool `json:"RequiresImmatureWasteWeight"`
	RequiresMatureWasteWeight bool `json:"RequiresMatureWasteWeight"`
	RequiresNote bool `json:"RequiresNote"`
	RequiresWasteWeight bool `json:"RequiresWasteWeight"`
}

type DeliveryRetailer struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	AllowFullEdit bool `json:"AllowFullEdit"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	DateTime string `json:"DateTime"`
	Destinations []DeliveryRetailerDestination `json:"Destinations"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	FacilityLicenseNumber string `json:"FacilityLicenseNumber"`
	FacilityName string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Leg int `json:"Leg"`
	Packages []DeliveryRetailerPackage `json:"Packages"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	RestockDateTime *string `json:"RestockDateTime"`
	RetailerDeliveryNumber string `json:"RetailerDeliveryNumber"`
	RetailerDeliveryState string `json:"RetailerDeliveryState"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	TotalPriceSold int `json:"TotalPriceSold"`
	VehicleInfo string `json:"VehicleInfo"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type DeliveryRetailerDestination struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	ConsumerId string `json:"ConsumerId"`
	DeliveryNumber *string `json:"DeliveryNumber"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	FacilityLicenseNumber *string `json:"FacilityLicenseNumber"`
	FacilityName *string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientDeliveryZoneId *int `json:"RecipientDeliveryZoneId"`
	RecipientDeliveryZoneName *string `json:"RecipientDeliveryZoneName"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	SalesDeliveryState string `json:"SalesDeliveryState"`
	ShipperFacilityLicenseNumber *string `json:"ShipperFacilityLicenseNumber"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
	TransporterFacilityId *int `json:"TransporterFacilityId"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName *string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type DeliveryRetailerPackage struct {
	ArchivedDate *string `json:"ArchivedDate"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	IsOnInvestigationHold bool `json:"IsOnInvestigationHold"`
	IsOnInvestigationRecall bool `json:"IsOnInvestigationRecall"`
	IsOnRecall *bool `json:"IsOnRecall"`
	IsOnRecallCombined bool `json:"IsOnRecallCombined"`
	ItemServingSize *string `json:"ItemServingSize"`
	ItemStrainName *string `json:"ItemStrainName"`
	ItemSupplyDurationDays *int `json:"ItemSupplyDurationDays"`
	ItemUnitCbdContent *float64 `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentDose *float64 `json:"ItemUnitCbdContentDose"`
	ItemUnitCbdContentDoseUnitOfMeasureName *string `json:"ItemUnitCbdContentDoseUnitOfMeasureName"`
	ItemUnitCbdContentUnitOfMeasureName *string `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitCbdPercent *float64 `json:"ItemUnitCbdPercent"`
	ItemUnitQuantity *float64 `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName *string `json:"ItemUnitQuantityUnitOfMeasureName"`
	ItemUnitThcContent *float64 `json:"ItemUnitThcContent"`
	ItemUnitThcContentDose *float64 `json:"ItemUnitThcContentDose"`
	ItemUnitThcContentDoseUnitOfMeasureName *string `json:"ItemUnitThcContentDoseUnitOfMeasureName"`
	ItemUnitThcContentUnitOfMeasureName *string `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitThcPercent *float64 `json:"ItemUnitThcPercent"`
	ItemUnitVolume *float64 `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName *string `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight *float64 `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName *string `json:"ItemUnitWeightUnitOfMeasureName"`
	LastModified string `json:"LastModified"`
	PackageId int `json:"PackageId"`
	PackageLabel string `json:"PackageLabel"`
	ProductCategoryName *string `json:"ProductCategoryName"`
	ProductName *string `json:"ProductName"`
	Quantity float64 `json:"Quantity"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	RetailerDeliveryState *string `json:"RetailerDeliveryState"`
	TotalPrice float64 `json:"TotalPrice"`
	UnitOfMeasureAbbreviation *string `json:"UnitOfMeasureAbbreviation"`
	UnitOfMeasureName string `json:"UnitOfMeasureName"`
}

type InactiveDeliveriesRetailer struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	AllowFullEdit bool `json:"AllowFullEdit"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	DateTime string `json:"DateTime"`
	Destinations []interface{} `json:"Destinations"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	FacilityLicenseNumber string `json:"FacilityLicenseNumber"`
	FacilityName string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	Leg int `json:"Leg"`
	Packages []interface{} `json:"Packages"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	RestockDateTime *string `json:"RestockDateTime"`
	RetailerDeliveryNumber string `json:"RetailerDeliveryNumber"`
	RetailerDeliveryState string `json:"RetailerDeliveryState"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	TotalPriceSold int `json:"TotalPriceSold"`
	VehicleInfo string `json:"VehicleInfo"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate string `json:"VoidedDate"`
}

type InactiveDelivery struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	ConsumerId string `json:"ConsumerId"`
	DeliveryNumber *string `json:"DeliveryNumber"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	FacilityLicenseNumber *string `json:"FacilityLicenseNumber"`
	FacilityName *string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientDeliveryZoneId *int `json:"RecipientDeliveryZoneId"`
	RecipientDeliveryZoneName *string `json:"RecipientDeliveryZoneName"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	SalesDeliveryState string `json:"SalesDeliveryState"`
	ShipperFacilityLicenseNumber *string `json:"ShipperFacilityLicenseNumber"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
	TransporterFacilityId *int `json:"TransporterFacilityId"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName *string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type InactiveReceipt struct {
	ArchivedDate *string `json:"ArchivedDate"`
	CaregiverLicenseNumber *string `json:"CaregiverLicenseNumber"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber"`
	Id int `json:"Id"`
	IdentificationMethod *string `json:"IdentificationMethod"`
	IsFinal bool `json:"IsFinal"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PatientRegistrationLocationId *int `json:"PatientRegistrationLocationId"`
	ReceiptNumber *string `json:"ReceiptNumber"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
}

type PatientRegistrationLocation struct {
	Id int `json:"Id"`
	Name string `json:"Name"`
}

type ReceiptsExternalByExternalNumber struct {
	ArchivedDate *string `json:"ArchivedDate"`
	CaregiverLicenseNumber *string `json:"CaregiverLicenseNumber"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber"`
	Id int `json:"Id"`
	IdentificationMethod *string `json:"IdentificationMethod"`
	IsFinal bool `json:"IsFinal"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PatientRegistrationLocationId *int `json:"PatientRegistrationLocationId"`
	ReceiptNumber *string `json:"ReceiptNumber"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
}

type ReceiptsFinalizeRequest struct {
	Id int `json:"Id"`
}

type ReceiptsUnfinalizeRequest struct {
	Id int `json:"Id"`
}

type SalesDeliveriesRequest struct {
	ConsumerId *int `json:"ConsumerId"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	Transactions []SalesDeliveriesRequestTransaction `json:"Transactions"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
}

type SalesDeliveriesRequestTransaction struct {
	CityTax *string `json:"CityTax"`
	CountyTax *string `json:"CountyTax"`
	DiscountAmount *string `json:"DiscountAmount"`
	ExciseTax *string `json:"ExciseTax"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	MunicipalTax *string `json:"MunicipalTax"`
	PackageLabel string `json:"PackageLabel"`
	Price *string `json:"Price"`
	QrCodes *string `json:"QrCodes"`
	Quantity float64 `json:"Quantity"`
	SalesTax *string `json:"SalesTax"`
	SubTotal *string `json:"SubTotal"`
	TotalAmount float64 `json:"TotalAmount"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UnitThcContent *float64 `json:"UnitThcContent"`
	UnitThcContentUnitOfMeasure *string `json:"UnitThcContentUnitOfMeasure"`
	UnitThcPercent *float64 `json:"UnitThcPercent"`
	UnitWeight *float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasure *string `json:"UnitWeightUnitOfMeasure"`
}

type SalesDeliveriesRetailerRequest struct {
	ConsumerId *int `json:"ConsumerId"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	RetailerDeliveryId int `json:"RetailerDeliveryId"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	Transactions []SalesDeliveriesRetailerRequestTransaction `json:"Transactions"`
}

type SalesDeliveriesRetailerRequestTransaction struct {
	CityTax *string `json:"CityTax"`
	CountyTax *string `json:"CountyTax"`
	DiscountAmount *string `json:"DiscountAmount"`
	ExciseTax *string `json:"ExciseTax"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	MunicipalTax *string `json:"MunicipalTax"`
	PackageLabel string `json:"PackageLabel"`
	Price *string `json:"Price"`
	QrCodes *string `json:"QrCodes"`
	Quantity float64 `json:"Quantity"`
	SalesTax *string `json:"SalesTax"`
	SubTotal *string `json:"SubTotal"`
	TotalAmount float64 `json:"TotalAmount"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UnitThcContent *float64 `json:"UnitThcContent"`
	UnitThcContentUnitOfMeasure *string `json:"UnitThcContentUnitOfMeasure"`
	UnitThcPercent *float64 `json:"UnitThcPercent"`
	UnitWeight *float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasure *string `json:"UnitWeightUnitOfMeasure"`
}

type SalesDelivery struct {
	AcceptedDateTime *string `json:"AcceptedDateTime"`
	ActualArrivalDateTime *string `json:"ActualArrivalDateTime"`
	ActualDepartureDateTime *string `json:"ActualDepartureDateTime"`
	ActualReturnArrivalDateTime *string `json:"ActualReturnArrivalDateTime"`
	ActualReturnDepartureDateTime *string `json:"ActualReturnDepartureDateTime"`
	CompletedDateTime *string `json:"CompletedDateTime"`
	ConsumerId string `json:"ConsumerId"`
	DeliveryNumber *string `json:"DeliveryNumber"`
	Direction string `json:"Direction"`
	DriverEmployeeId string `json:"DriverEmployeeId"`
	DriverName string `json:"DriverName"`
	DriversLicenseNumber string `json:"DriversLicenseNumber"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime"`
	EstimatedReturnArrivalDateTime *string `json:"EstimatedReturnArrivalDateTime"`
	EstimatedReturnDepartureDateTime *string `json:"EstimatedReturnDepartureDateTime"`
	FacilityLicenseNumber *string `json:"FacilityLicenseNumber"`
	FacilityName *string `json:"FacilityName"`
	Id int `json:"Id"`
	LastModified string `json:"LastModified"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PlannedRoute string `json:"PlannedRoute"`
	RecipientAddressCity string `json:"RecipientAddressCity"`
	RecipientAddressCounty *string `json:"RecipientAddressCounty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode"`
	RecipientAddressState string `json:"RecipientAddressState"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2"`
	RecipientDeliveryZoneId *int `json:"RecipientDeliveryZoneId"`
	RecipientDeliveryZoneName *string `json:"RecipientDeliveryZoneName"`
	RecipientName *string `json:"RecipientName"`
	RecipientZoneId *int `json:"RecipientZoneId"`
	RecordedByUserName *string `json:"RecordedByUserName"`
	RecordedDateTime string `json:"RecordedDateTime"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	SalesDeliveryState string `json:"SalesDeliveryState"`
	ShipperFacilityLicenseNumber *string `json:"ShipperFacilityLicenseNumber"`
	TotalPackages int `json:"TotalPackages"`
	TotalPrice int `json:"TotalPrice"`
	Transactions []interface{} `json:"Transactions"`
	TransporterFacilityId *int `json:"TransporterFacilityId"`
	TransporterFacilityLicenseNumber *string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName *string `json:"TransporterFacilityName"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber"`
	VehicleMake string `json:"VehicleMake"`
	VehicleModel string `json:"VehicleModel"`
	VoidedDate *string `json:"VoidedDate"`
}

type SalesReceiptsRequest struct {
	CaregiverLicenseNumber *string `json:"CaregiverLicenseNumber"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber"`
	IdentificationMethod *string `json:"IdentificationMethod"`
	PatientLicenseNumber *string `json:"PatientLicenseNumber"`
	PatientRegistrationLocationId *int `json:"PatientRegistrationLocationId"`
	SalesCustomerType string `json:"SalesCustomerType"`
	SalesDateTime string `json:"SalesDateTime"`
	Transactions []SalesReceiptsRequestTransaction `json:"Transactions"`
}

type SalesReceiptsRequestTransaction struct {
	CityTax *string `json:"CityTax"`
	CountyTax *string `json:"CountyTax"`
	DiscountAmount *string `json:"DiscountAmount"`
	ExciseTax *string `json:"ExciseTax"`
	InvoiceNumber *string `json:"InvoiceNumber"`
	MunicipalTax *string `json:"MunicipalTax"`
	PackageLabel string `json:"PackageLabel"`
	Price *string `json:"Price"`
	QrCodes *string `json:"QrCodes"`
	Quantity float64 `json:"Quantity"`
	SalesTax *string `json:"SalesTax"`
	SubTotal *string `json:"SubTotal"`
	TotalAmount float64 `json:"TotalAmount"`
	UnitOfMeasure string `json:"UnitOfMeasure"`
	UnitThcContent *float64 `json:"UnitThcContent"`
	UnitThcContentUnitOfMeasure *string `json:"UnitThcContentUnitOfMeasure"`
	UnitThcPercent *float64 `json:"UnitThcPercent"`
	UnitWeight *float64 `json:"UnitWeight"`
	UnitWeightUnitOfMeasure *string `json:"UnitWeightUnitOfMeasure"`
}
