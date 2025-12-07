package wrapper

type ProcessingJobsCreateAdjustV1RequestItem struct {
	CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
	Packages []ProcessingJobsCreateAdjustV1RequestItem_Packages `json:"Packages,omitempty"`
	Id int `json:"Id,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	AdjustmentNote string `json:"AdjustmentNote,omitempty"`
}

type ProcessingJobsCreateAdjustV1RequestItem_Packages struct {
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type ProcessingJobsCreateAdjustV2RequestItem struct {
	VolumeUnitOfMeasureName string `json:"VolumeUnitOfMeasureName,omitempty"`
	WeightUnitOfMeasureName string `json:"WeightUnitOfMeasureName,omitempty"`
	Packages []ProcessingJobsCreateAdjustV2RequestItem_Packages `json:"Packages,omitempty"`
	Id int `json:"Id,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	AdjustmentNote string `json:"AdjustmentNote,omitempty"`
	CountUnitOfMeasureName string `json:"CountUnitOfMeasureName,omitempty"`
}

type ProcessingJobsCreateAdjustV2RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateJobtypesV1RequestItem struct {
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	Category string `json:"Category,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
}

type ProcessingJobsCreateJobtypesV2RequestItem struct {
	Category string `json:"Category,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ProcessingJobsCreateStartV1RequestItem struct {
	CountUnitOfMeasure string `json:"CountUnitOfMeasure,omitempty"`
	VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure,omitempty"`
	WeightUnitOfMeasure string `json:"WeightUnitOfMeasure,omitempty"`
	Packages []ProcessingJobsCreateStartV1RequestItem_Packages `json:"Packages,omitempty"`
	StartDate string `json:"StartDate,omitempty"`
	JobName string `json:"JobName,omitempty"`
	JobType string `json:"JobType,omitempty"`
}

type ProcessingJobsCreateStartV1RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateStartV2RequestItem struct {
	WeightUnitOfMeasure string `json:"WeightUnitOfMeasure,omitempty"`
	Packages []ProcessingJobsCreateStartV2RequestItem_Packages `json:"Packages,omitempty"`
	StartDate string `json:"StartDate,omitempty"`
	JobName string `json:"JobName,omitempty"`
	JobType string `json:"JobType,omitempty"`
	CountUnitOfMeasure string `json:"CountUnitOfMeasure,omitempty"`
	VolumeUnitOfMeasure string `json:"VolumeUnitOfMeasure,omitempty"`
}

type ProcessingJobsCreateStartV2RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type ProcessingJobsCreatepackagesV1RequestItem struct {
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	PackageDate string `json:"PackageDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Location string `json:"Location,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	JobName string `json:"JobName,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	WasteCountQuantity string `json:"WasteCountQuantity,omitempty"`
	WasteWeightQuantity string `json:"WasteWeightQuantity,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	WasteVolumeQuantity string `json:"WasteVolumeQuantity,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	FinishProcessingJob bool `json:"FinishProcessingJob,omitempty"`
	Item string `json:"Item,omitempty"`
	Note string `json:"Note,omitempty"`
}

type ProcessingJobsCreatepackagesV2RequestItem struct {
	Note string `json:"Note,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	Item string `json:"Item,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	WasteWeightQuantity string `json:"WasteWeightQuantity,omitempty"`
	WasteVolumeQuantity string `json:"WasteVolumeQuantity,omitempty"`
	FinishProcessingJob bool `json:"FinishProcessingJob,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	JobName string `json:"JobName,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	WasteCountQuantity string `json:"WasteCountQuantity,omitempty"`
	Tag string `json:"Tag,omitempty"`
	PackageDate string `json:"PackageDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
}

type ProcessingJobsUpdateFinishV1RequestItem struct {
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	Id int `json:"Id,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	TotalCountWaste string `json:"TotalCountWaste,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	TotalVolumeWaste string `json:"TotalVolumeWaste,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	TotalWeightWaste int `json:"TotalWeightWaste,omitempty"`
}

type ProcessingJobsUpdateFinishV2RequestItem struct {
	TotalVolumeWaste string `json:"TotalVolumeWaste,omitempty"`
	TotalCountWaste string `json:"TotalCountWaste,omitempty"`
	Id int `json:"Id,omitempty"`
	FinishDate string `json:"FinishDate,omitempty"`
	WasteVolumeUnitOfMeasureName string `json:"WasteVolumeUnitOfMeasureName,omitempty"`
	TotalWeightWaste int `json:"TotalWeightWaste,omitempty"`
	WasteWeightUnitOfMeasureName string `json:"WasteWeightUnitOfMeasureName,omitempty"`
	FinishNote string `json:"FinishNote,omitempty"`
	WasteCountUnitOfMeasureName string `json:"WasteCountUnitOfMeasureName,omitempty"`
}

type ProcessingJobsUpdateJobtypesV1RequestItem struct {
	CategoryName string `json:"CategoryName,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
	Attributes []string `json:"Attributes,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ProcessingJobsUpdateJobtypesV2RequestItem struct {
	Attributes []string `json:"Attributes,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	CategoryName string `json:"CategoryName,omitempty"`
	ProcessingSteps string `json:"ProcessingSteps,omitempty"`
}

type ProcessingJobsUpdateUnfinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type ProcessingJobsUpdateUnfinishV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SublocationsCreateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
}

type SublocationsUpdateV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type SalesCreateDeliveryV1RequestItem struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	Transactions []SalesCreateDeliveryV1RequestItem_Transactions `json:"Transactions,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type SalesCreateDeliveryV1RequestItem_Transactions struct {
	SubTotal string `json:"SubTotal,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	Price string `json:"Price,omitempty"`
}

type SalesCreateDeliveryV2RequestItem struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	Transactions []SalesCreateDeliveryV2RequestItem_Transactions `json:"Transactions,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
}

type SalesCreateDeliveryV2RequestItem_Transactions struct {
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	Price string `json:"Price,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem struct {
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Packages []SalesCreateDeliveryRetailerV1RequestItem_Packages `json:"Packages,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	Destinations []SalesCreateDeliveryRetailerV1RequestItem_Destinations `json:"Destinations,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Packages struct {
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Destinations struct {
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transactions []SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
}

type SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions struct {
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem struct {
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	Packages []SalesCreateDeliveryRetailerV2RequestItem_Packages `json:"Packages,omitempty"`
	Destinations []SalesCreateDeliveryRetailerV2RequestItem_Destinations `json:"Destinations,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Packages struct {
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Destinations struct {
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	Transactions []SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
}

type SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
}

type SalesCreateDeliveryRetailerDepartV1RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerDepartV2RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerEndV1RequestItem struct {
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
	Packages []SalesCreateDeliveryRetailerEndV1RequestItem_Packages `json:"Packages,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerEndV1RequestItem_Packages struct {
	Label string `json:"Label,omitempty"`
	EndQuantity int `json:"EndQuantity,omitempty"`
	EndUnitOfMeasure string `json:"EndUnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryRetailerEndV2RequestItem struct {
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
	Packages []SalesCreateDeliveryRetailerEndV2RequestItem_Packages `json:"Packages,omitempty"`
}

type SalesCreateDeliveryRetailerEndV2RequestItem_Packages struct {
	EndQuantity int `json:"EndQuantity,omitempty"`
	EndUnitOfMeasure string `json:"EndUnitOfMeasure,omitempty"`
	Label string `json:"Label,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV1RequestItem struct {
	DateTime string `json:"DateTime,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Destinations string `json:"Destinations,omitempty"`
	Packages []SalesCreateDeliveryRetailerRestockV1RequestItem_Packages `json:"Packages,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV1RequestItem_Packages struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV2RequestItem struct {
	Destinations string `json:"Destinations,omitempty"`
	Packages []SalesCreateDeliveryRetailerRestockV2RequestItem_Packages `json:"Packages,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
}

type SalesCreateDeliveryRetailerRestockV2RequestItem_Packages struct {
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	RemoveCurrentPackage bool `json:"RemoveCurrentPackage,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV1RequestItem struct {
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	Transactions []SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions `json:"Transactions,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions struct {
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV2RequestItem struct {
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RetailerDeliveryId int `json:"RetailerDeliveryId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	RecipientZoneId int `json:"RecipientZoneId,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	Transactions []SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions `json:"Transactions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
}

type SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions struct {
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Price string `json:"Price,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesCreateReceiptV1RequestItem struct {
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Transactions []SalesCreateReceiptV1RequestItem_Transactions `json:"Transactions,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
}

type SalesCreateReceiptV1RequestItem_Transactions struct {
	CountyTax string `json:"CountyTax,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Price string `json:"Price,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
}

type SalesCreateReceiptV2RequestItem struct {
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Transactions []SalesCreateReceiptV2RequestItem_Transactions `json:"Transactions,omitempty"`
}

type SalesCreateReceiptV2RequestItem_Transactions struct {
	ExciseTax string `json:"ExciseTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
}

type SalesCreateTransactionByDateV1RequestItem struct {
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Price string `json:"Price,omitempty"`
}

type SalesUpdateDeliveryV1RequestItem struct {
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	Id int `json:"Id,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	Transactions []SalesUpdateDeliveryV1RequestItem_Transactions `json:"Transactions,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
}

type SalesUpdateDeliveryV1RequestItem_Transactions struct {
	SalesTax string `json:"SalesTax,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Price string `json:"Price,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
}

type SalesUpdateDeliveryV2RequestItem struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	Id int `json:"Id,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ConsumerId int `json:"ConsumerId,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	Transactions []SalesUpdateDeliveryV2RequestItem_Transactions `json:"Transactions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
}

type SalesUpdateDeliveryV2RequestItem_Transactions struct {
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
}

type SalesUpdateDeliveryCompleteV1RequestItem struct {
	PaymentType string `json:"PaymentType,omitempty"`
	AcceptedPackages []string `json:"AcceptedPackages,omitempty"`
	ReturnedPackages []SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages `json:"ReturnedPackages,omitempty"`
	Id int `json:"Id,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
}

type SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages struct {
	ReturnQuantityVerified int `json:"ReturnQuantityVerified,omitempty"`
	ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure,omitempty"`
	ReturnReason string `json:"ReturnReason,omitempty"`
	ReturnReasonNote string `json:"ReturnReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
}

type SalesUpdateDeliveryCompleteV2RequestItem struct {
	PaymentType string `json:"PaymentType,omitempty"`
	AcceptedPackages []string `json:"AcceptedPackages,omitempty"`
	ReturnedPackages []SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages `json:"ReturnedPackages,omitempty"`
	Id int `json:"Id,omitempty"`
	ActualArrivalDateTime string `json:"ActualArrivalDateTime,omitempty"`
}

type SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages struct {
	ReturnReason string `json:"ReturnReason,omitempty"`
	ReturnReasonNote string `json:"ReturnReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
	ReturnQuantityVerified int `json:"ReturnQuantityVerified,omitempty"`
	ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure,omitempty"`
}

type SalesUpdateDeliveryHubV1RequestItem struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	Id int `json:"Id,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterFacilityId string `json:"TransporterFacilityId,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type SalesUpdateDeliveryHubV2RequestItem struct {
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Id int `json:"Id,omitempty"`
	TransporterFacilityId string `json:"TransporterFacilityId,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type SalesUpdateDeliveryHubAcceptV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubAcceptV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubDepartV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubDepartV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryHubVerifyIdV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	PaymentType string `json:"PaymentType,omitempty"`
}

type SalesUpdateDeliveryHubVerifyIdV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	PaymentType string `json:"PaymentType,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem struct {
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Packages []SalesUpdateDeliveryRetailerV1RequestItem_Packages `json:"Packages,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	Destinations []SalesUpdateDeliveryRetailerV1RequestItem_Destinations `json:"Destinations,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Id int `json:"Id,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Destinations struct {
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	Transactions []SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Id int `json:"Id,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	DriverEmployeeId int `json:"DriverEmployeeId,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions struct {
	SalesTax string `json:"SalesTax,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Price string `json:"Price,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
}

type SalesUpdateDeliveryRetailerV1RequestItem_Packages struct {
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem struct {
	DriverName string `json:"DriverName,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Destinations []SalesUpdateDeliveryRetailerV2RequestItem_Destinations `json:"Destinations,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Packages []SalesUpdateDeliveryRetailerV2RequestItem_Packages `json:"Packages,omitempty"`
	DriverEmployeeId string `json:"DriverEmployeeId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Id int `json:"Id,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Destinations struct {
	RecipientAddressCounty string `json:"RecipientAddressCounty,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	Transactions []SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions `json:"Transactions,omitempty"`
	RecipientAddressState string `json:"RecipientAddressState,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	DriverEmployeeId int `json:"DriverEmployeeId,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	RecipientZoneId string `json:"RecipientZoneId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	RecipientAddressStreet1 string `json:"RecipientAddressStreet1,omitempty"`
	ConsumerId string `json:"ConsumerId,omitempty"`
	Id int `json:"Id,omitempty"`
	RecipientAddressCity string `json:"RecipientAddressCity,omitempty"`
	RecipientAddressPostalCode string `json:"RecipientAddressPostalCode,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	RecipientName string `json:"RecipientName,omitempty"`
	RecipientAddressStreet2 string `json:"RecipientAddressStreet2,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions struct {
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Price string `json:"Price,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
}

type SalesUpdateDeliveryRetailerV2RequestItem_Packages struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalPrice float64 `json:"TotalPrice,omitempty"`
	DateTime string `json:"DateTime,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesUpdateReceiptV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Transactions []SalesUpdateReceiptV1RequestItem_Transactions `json:"Transactions,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Id int `json:"Id,omitempty"`
}

type SalesUpdateReceiptV1RequestItem_Transactions struct {
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	Price string `json:"Price,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type SalesUpdateReceiptV2RequestItem struct {
	Transactions []SalesUpdateReceiptV2RequestItem_Transactions `json:"Transactions,omitempty"`
	IdentificationMethod string `json:"IdentificationMethod,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CaregiverLicenseNumber string `json:"CaregiverLicenseNumber,omitempty"`
	PatientRegistrationLocationId int `json:"PatientRegistrationLocationId,omitempty"`
	Id int `json:"Id,omitempty"`
	SalesDateTime string `json:"SalesDateTime,omitempty"`
	ExternalReceiptNumber string `json:"ExternalReceiptNumber,omitempty"`
	SalesCustomerType string `json:"SalesCustomerType,omitempty"`
}

type SalesUpdateReceiptV2RequestItem_Transactions struct {
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	Price string `json:"Price,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
}

type SalesUpdateReceiptFinalizeV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateReceiptUnfinalizeV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type SalesUpdateTransactionByDateV1RequestItem struct {
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	ExciseTax string `json:"ExciseTax,omitempty"`
	DiscountAmount string `json:"DiscountAmount,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	Price string `json:"Price,omitempty"`
	CityTax string `json:"CityTax,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SubTotal string `json:"SubTotal,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	MunicipalTax string `json:"MunicipalTax,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	TotalAmount float64 `json:"TotalAmount,omitempty"`
	QrCodes string `json:"QrCodes,omitempty"`
	CountyTax string `json:"CountyTax,omitempty"`
	SalesTax string `json:"SalesTax,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
}

type StrainsCreateV1RequestItem struct {
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type StrainsCreateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
}

type StrainsCreateUpdateV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
}

type StrainsUpdateV2RequestItem struct {
	SativaPercentage float64 `json:"SativaPercentage,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	TestingStatus string `json:"TestingStatus,omitempty"`
	ThcLevel float64 `json:"ThcLevel,omitempty"`
	CbdLevel float64 `json:"CbdLevel,omitempty"`
	IndicaPercentage float64 `json:"IndicaPercentage,omitempty"`
}

type AdditivesTemplatesCreateV2RequestItem struct {
	ActiveIngredients []AdditivesTemplatesCreateV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	Note string `json:"Note,omitempty"`
	Name string `json:"Name,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
	RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
}

type AdditivesTemplatesCreateV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type AdditivesTemplatesUpdateV2RequestItem struct {
	AdditiveType string `json:"AdditiveType,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	Name string `json:"Name,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	RestrictiveEntryIntervalQuantityDescription string `json:"RestrictiveEntryIntervalQuantityDescription,omitempty"`
	Id int `json:"Id,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	RestrictiveEntryIntervalTimeDescription string `json:"RestrictiveEntryIntervalTimeDescription,omitempty"`
	ActiveIngredients []AdditivesTemplatesUpdateV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
}

type AdditivesTemplatesUpdateV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type HarvestsCreateFinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsCreatePackageV1RequestItem struct {
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	Tag string `json:"Tag,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	Note string `json:"Note,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	Item string `json:"Item,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Ingredients []HarvestsCreatePackageV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	Location string `json:"Location,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type HarvestsCreatePackageV1RequestItem_Ingredients struct {
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
}

type HarvestsCreatePackageV2RequestItem struct {
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Ingredients []HarvestsCreatePackageV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	Tag string `json:"Tag,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	Note string `json:"Note,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type HarvestsCreatePackageV2RequestItem_Ingredients struct {
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
}

type HarvestsCreatePackageTestingV1RequestItem struct {
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	Ingredients []HarvestsCreatePackageTestingV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	Location string `json:"Location,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Item string `json:"Item,omitempty"`
	Note string `json:"Note,omitempty"`
}

type HarvestsCreatePackageTestingV1RequestItem_Ingredients struct {
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
}

type HarvestsCreatePackageTestingV2RequestItem struct {
	Item string `json:"Item,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ProductRequiresDecontamination bool `json:"ProductRequiresDecontamination,omitempty"`
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	Note string `json:"Note,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Location string `json:"Location,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Ingredients []HarvestsCreatePackageTestingV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	RemediationMethodId int `json:"RemediationMethodId,omitempty"`
	DecontaminateProduct bool `json:"DecontaminateProduct,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	RequiredLabTestBatches []interface{} `json:"RequiredLabTestBatches,omitempty"`
	Tag string `json:"Tag,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
	RemediateProduct bool `json:"RemediateProduct,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type HarvestsCreatePackageTestingV2RequestItem_Ingredients struct {
	HarvestId int `json:"HarvestId,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
}

type HarvestsCreateRemoveWasteV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	WasteType string `json:"WasteType,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsCreateUnfinishV1RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type HarvestsCreateWasteV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	WasteType string `json:"WasteType,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type HarvestsUpdateFinishV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type HarvestsUpdateLocationV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
}

type HarvestsUpdateMoveV1RequestItem struct {
	HarvestName string `json:"HarvestName,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type HarvestsUpdateRenameV1RequestItem struct {
	OldName string `json:"OldName,omitempty"`
	NewName string `json:"NewName,omitempty"`
	Id int `json:"Id,omitempty"`
}

type HarvestsUpdateRenameV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	OldName string `json:"OldName,omitempty"`
	NewName string `json:"NewName,omitempty"`
}

type HarvestsUpdateRestoreHarvestedPlantsV2RequestItem struct {
	HarvestId int `json:"HarvestId,omitempty"`
	PlantTags []string `json:"PlantTags,omitempty"`
}

type HarvestsUpdateUnfinishV2RequestItem struct {
	Id int `json:"Id,omitempty"`
}

type LabTestsCreateRecordV1RequestItem struct {
	Results []LabTestsCreateRecordV1RequestItem_Results `json:"Results,omitempty"`
	Label string `json:"Label,omitempty"`
	ResultDate string `json:"ResultDate,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsCreateRecordV1RequestItem_Results struct {
	Passed bool `json:"Passed,omitempty"`
	Notes string `json:"Notes,omitempty"`
	LabTestTypeName string `json:"LabTestTypeName,omitempty"`
	Quantity float64 `json:"Quantity,omitempty"`
}

type LabTestsCreateRecordV2RequestItem struct {
	ResultDate string `json:"ResultDate,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
	Results []LabTestsCreateRecordV2RequestItem_Results `json:"Results,omitempty"`
	Label string `json:"Label,omitempty"`
}

type LabTestsCreateRecordV2RequestItem_Results struct {
	Passed bool `json:"Passed,omitempty"`
	Notes string `json:"Notes,omitempty"`
	LabTestTypeName string `json:"LabTestTypeName,omitempty"`
	Quantity float64 `json:"Quantity,omitempty"`
}

type LabTestsUpdateLabtestdocumentV1RequestItem struct {
	LabTestResultId int `json:"LabTestResultId,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsUpdateLabtestdocumentV2RequestItem struct {
	LabTestResultId int `json:"LabTestResultId,omitempty"`
	DocumentFileName string `json:"DocumentFileName,omitempty"`
	DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
}

type LabTestsUpdateResultReleaseV1RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type LabTestsUpdateResultReleaseV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type LocationsCreateV1RequestItem struct {
	LocationTypeName string `json:"LocationTypeName,omitempty"`
	Name string `json:"Name,omitempty"`
}

type LocationsCreateV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
}

type LocationsCreateUpdateV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	LocationTypeName string `json:"LocationTypeName,omitempty"`
}

type LocationsUpdateV2RequestItem struct {
	LocationTypeName string `json:"LocationTypeName,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantsCreateAdditivesV1RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	PlantLabels []string `json:"PlantLabels,omitempty"`
}

type PlantsCreateAdditivesV1RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantsCreateAdditivesV2RequestItem struct {
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	PlantLabels []string `json:"PlantLabels,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
}

type PlantsCreateAdditivesV2RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantsCreateAdditivesBylocationV1RequestItem struct {
	ActiveIngredients []PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
}

type PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantsCreateAdditivesBylocationV2RequestItem struct {
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	ActiveIngredients []PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
}

type PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients struct {
	Percentage float64 `json:"Percentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem struct {
	SublocationName string `json:"SublocationName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
	Volume string `json:"Volume,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
}

type PlantsCreateAdditivesUsingtemplateV2RequestItem struct {
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	PlantLabels []string `json:"PlantLabels,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
	Volume string `json:"Volume,omitempty"`
}

type PlantsCreateChangegrowthphasesV1RequestItem struct {
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
}

type PlantsCreateHarvestplantsV1RequestItem struct {
	HarvestName string `json:"HarvestName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Plant string `json:"Plant,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
}

type PlantsCreateManicureV2RequestItem struct {
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	Plant string `json:"Plant,omitempty"`
}

type PlantsCreateManicureplantsV1RequestItem struct {
	DryingLocation string `json:"DryingLocation,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	Plant string `json:"Plant,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PlantsCreateMoveplantsV1RequestItem struct {
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
}

type PlantsCreatePlantbatchPackageV1RequestItem struct {
	Item string `json:"Item,omitempty"`
	Location string `json:"Location,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Count int `json:"Count,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PackageTag string `json:"PackageTag,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
}

type PlantsCreatePlantbatchPackageV2RequestItem struct {
	PlantLabel string `json:"PlantLabel,omitempty"`
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Count int `json:"Count,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PackageTag string `json:"PackageTag,omitempty"`
	Location string `json:"Location,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PlantsCreatePlantingsV1RequestItem struct {
	SublocationName string `json:"SublocationName,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PlantsCreatePlantingsV2RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantLabel string `json:"PlantLabel,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PlantsCreateWasteV1RequestItem struct {
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	PlantLabels []interface{} `json:"PlantLabels,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
	Note string `json:"Note,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
}

type PlantsCreateWasteV2RequestItem struct {
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	PlantLabels []interface{} `json:"PlantLabels,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	Note string `json:"Note,omitempty"`
}

type PlantsUpdateAdjustV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	AdjustCount int `json:"AdjustCount,omitempty"`
	AdjustReason string `json:"AdjustReason,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
}

type PlantsUpdateGrowthphaseV2RequestItem struct {
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
}

type PlantsUpdateHarvestV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Plant string `json:"Plant,omitempty"`
	Weight float64 `json:"Weight,omitempty"`
	UnitOfWeight string `json:"UnitOfWeight,omitempty"`
	DryingLocation string `json:"DryingLocation,omitempty"`
	DryingSublocation string `json:"DryingSublocation,omitempty"`
	HarvestName string `json:"HarvestName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PlantsUpdateLocationV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PlantsUpdateMergeV2RequestItem struct {
	SourcePlantGroupLabel string `json:"SourcePlantGroupLabel,omitempty"`
	MergeDate string `json:"MergeDate,omitempty"`
	TargetPlantGroupLabel string `json:"TargetPlantGroupLabel,omitempty"`
}

type PlantsUpdateSplitV2RequestItem struct {
	SplitDate string `json:"SplitDate,omitempty"`
	SourcePlantLabel string `json:"SourcePlantLabel,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	TagLabel string `json:"TagLabel,omitempty"`
	StrainLabel string `json:"StrainLabel,omitempty"`
}

type PlantsUpdateStrainV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	StrainId int `json:"StrainId,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	Id int `json:"Id,omitempty"`
}

type PlantsUpdateTagV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Label string `json:"Label,omitempty"`
	TagId int `json:"TagId,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	ReplaceDate string `json:"ReplaceDate,omitempty"`
}

type RetailIdCreateAssociateV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	QrUrls []string `json:"QrUrls,omitempty"`
}

type RetailIdCreateGenerateV2Request struct {
	Quantity int `json:"Quantity,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
}

type RetailIdCreateMergeV2Request struct {
	PackageLabels []string `json:"packageLabels,omitempty"`
}

type RetailIdCreatePackageInfoV2Request struct {
	PackageLabels []string `json:"packageLabels,omitempty"`
}

type PatientsCreateV2RequestItem struct {
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
}

type PatientsCreateAddV1RequestItem struct {
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
}

type PatientsCreateUpdateV1RequestItem struct {
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	NewLicenseNumber string `json:"NewLicenseNumber,omitempty"`
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PatientsUpdateV2RequestItem struct {
	LicenseEffectiveEndDate string `json:"LicenseEffectiveEndDate,omitempty"`
	RecommendedSmokableQuantity int `json:"RecommendedSmokableQuantity,omitempty"`
	InfusedOuncesAllowed int `json:"InfusedOuncesAllowed,omitempty"`
	RecommendedPlants int `json:"RecommendedPlants,omitempty"`
	HasSalesLimitExemption bool `json:"HasSalesLimitExemption,omitempty"`
	LicenseNumber string `json:"LicenseNumber,omitempty"`
	ThcOuncesAllowed int `json:"ThcOuncesAllowed,omitempty"`
	ConcentrateOuncesAllowed int `json:"ConcentrateOuncesAllowed,omitempty"`
	MaxFlowerThcPercentAllowed int `json:"MaxFlowerThcPercentAllowed,omitempty"`
	NewLicenseNumber string `json:"NewLicenseNumber,omitempty"`
	FlowerOuncesAllowed int `json:"FlowerOuncesAllowed,omitempty"`
	MaxConcentrateThcPercentAllowed int `json:"MaxConcentrateThcPercentAllowed,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	LicenseEffectiveStartDate string `json:"LicenseEffectiveStartDate,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem struct {
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Destinations []TransfersCreateExternalIncomingV1RequestItem_Destinations `json:"Destinations,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations struct {
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Packages []TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages struct {
	SellByDate string `json:"SellByDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
}

type TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem struct {
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	Destinations []TransfersCreateExternalIncomingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations struct {
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	Packages []TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	Transporters []TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages struct {
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters struct {
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	TransporterDetails []TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
}

type TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem struct {
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	Name string `json:"Name,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	Destinations []TransfersCreateTemplatesV1RequestItem_Destinations `json:"Destinations,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations struct {
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersCreateTemplatesV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateTemplatesV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations_Transporters struct {
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type TransfersCreateTemplatesV1RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	Destinations []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	Name string `json:"Name,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations struct {
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	Transporters []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterDetails []TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
}

type TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem struct {
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	Destinations []TransfersUpdateExternalIncomingV1RequestItem_Destinations `json:"Destinations,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	ShipperName string `json:"ShipperName,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	TransferId int `json:"TransferId,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations struct {
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	Transporters []TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters struct {
	DriverName string `json:"DriverName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
}

type TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages struct {
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem struct {
	ShipperName string `json:"ShipperName,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	ShipperMainPhoneNumber string `json:"ShipperMainPhoneNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	ShipperLicenseNumber string `json:"ShipperLicenseNumber,omitempty"`
	ShipperAddress1 string `json:"ShipperAddress1,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	ShipperAddressCity string `json:"ShipperAddressCity,omitempty"`
	ShipperAddressState string `json:"ShipperAddressState,omitempty"`
	TransferId int `json:"TransferId,omitempty"`
	ShipperAddress2 string `json:"ShipperAddress2,omitempty"`
	ShipperAddressPostalCode string `json:"ShipperAddressPostalCode,omitempty"`
	Destinations []TransfersUpdateExternalIncomingV2RequestItem_Destinations `json:"Destinations,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations struct {
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	GrossUnitOfWeightId int `json:"GrossUnitOfWeightId,omitempty"`
	Transporters []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	Packages []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters struct {
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	TransporterDetails []TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
}

type TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages struct {
	UseByDate string `json:"UseByDate,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	PackagedDate string `json:"PackagedDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
	ItemName string `json:"ItemName,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	TransferTemplateId int `json:"TransferTemplateId,omitempty"`
	Destinations []TransfersUpdateTemplatesV1RequestItem_Destinations `json:"Destinations,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	Name string `json:"Name,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations struct {
	Transporters []TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	Packages []TransfersUpdateTemplatesV1RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters struct {
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	TransporterDetails string `json:"TransporterDetails,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type TransfersUpdateTemplatesV1RequestItem_Destinations_Packages struct {
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	TransferTemplateId int `json:"TransferTemplateId,omitempty"`
	Destinations []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations `json:"Destinations,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations struct {
	Packages []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages `json:"Packages,omitempty"`
	PlannedRoute string `json:"PlannedRoute,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransferDestinationId int `json:"TransferDestinationId,omitempty"`
	RecipientLicenseNumber string `json:"RecipientLicenseNumber,omitempty"`
	InvoiceNumber string `json:"InvoiceNumber,omitempty"`
	TransferTypeName string `json:"TransferTypeName,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	Transporters []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters `json:"Transporters,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters struct {
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	PhoneNumberForQuestions string `json:"PhoneNumberForQuestions,omitempty"`
	EstimatedArrivalDateTime string `json:"EstimatedArrivalDateTime,omitempty"`
	TransporterDetails []TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails `json:"TransporterDetails,omitempty"`
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	IsLayover bool `json:"IsLayover,omitempty"`
	EstimatedDepartureDateTime string `json:"EstimatedDepartureDateTime,omitempty"`
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails struct {
	DriverOccupationalLicenseNumber string `json:"DriverOccupationalLicenseNumber,omitempty"`
	DriverName string `json:"DriverName,omitempty"`
	DriverLicenseNumber string `json:"DriverLicenseNumber,omitempty"`
	DriverLayoverLeg string `json:"DriverLayoverLeg,omitempty"`
	VehicleMake string `json:"VehicleMake,omitempty"`
	VehicleModel string `json:"VehicleModel,omitempty"`
	VehicleLicensePlateNumber string `json:"VehicleLicensePlateNumber,omitempty"`
}

type TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	WholesalePrice string `json:"WholesalePrice,omitempty"`
	GrossWeight float64 `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightName string `json:"GrossUnitOfWeightName,omitempty"`
}

type TransportersCreateDriverV2RequestItem struct {
	Name string `json:"Name,omitempty"`
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	EmployeeId string `json:"EmployeeId,omitempty"`
}

type TransportersCreateVehicleV2RequestItem struct {
	Model string `json:"Model,omitempty"`
	LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
	Make string `json:"Make,omitempty"`
}

type TransportersUpdateDriverV2RequestItem struct {
	DriversLicenseNumber string `json:"DriversLicenseNumber,omitempty"`
	EmployeeId string `json:"EmployeeId,omitempty"`
	Id string `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type TransportersUpdateVehicleV2RequestItem struct {
	Id string `json:"Id,omitempty"`
	Make string `json:"Make,omitempty"`
	Model string `json:"Model,omitempty"`
	LicensePlateNumber string `json:"LicensePlateNumber,omitempty"`
}

type ItemsCreateV1RequestItem struct {
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	Strain string `json:"Strain,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Name string `json:"Name,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	Description string `json:"Description,omitempty"`
}

type ItemsCreateV2RequestItem struct {
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	Name string `json:"Name,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	Strain string `json:"Strain,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	Description string `json:"Description,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
}

type ItemsCreateBrandV2RequestItem struct {
	Name string `json:"Name,omitempty"`
}

type ItemsCreateFileV2RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreatePhotoV1RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreatePhotoV2RequestItem struct {
	FileName string `json:"FileName,omitempty"`
	EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
}

type ItemsCreateUpdateV1RequestItem struct {
	Strain string `json:"Strain,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	Name string `json:"Name,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	Id int `json:"Id,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	ItemBrand string `json:"ItemBrand,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	Description string `json:"Description,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
}

type ItemsUpdateV2RequestItem struct {
	ItemBrand string `json:"ItemBrand,omitempty"`
	AdministrationMethod string `json:"AdministrationMethod,omitempty"`
	UnitThcPercent float64 `json:"UnitThcPercent,omitempty"`
	UnitThcContent float64 `json:"UnitThcContent,omitempty"`
	ProductPhotoDescription string `json:"ProductPhotoDescription,omitempty"`
	UnitThcAcontent float64 `json:"UnitThcAContent,omitempty"`
	UnitCbdPercent float64 `json:"UnitCbdPercent,omitempty"`
	ProcessingJobTypeName string `json:"ProcessingJobTypeName,omitempty"`
	UnitThcContentDose float64 `json:"UnitThcContentDose,omitempty"`
	Allergens string `json:"Allergens,omitempty"`
	UnitThcAcontentUnitOfMeasure string `json:"UnitThcAContentUnitOfMeasure,omitempty"`
	UnitThcAcontentDoseUnitOfMeasure string `json:"UnitThcAContentDoseUnitOfMeasure,omitempty"`
	UnitVolume float64 `json:"UnitVolume,omitempty"`
	LabelPhotoDescription string `json:"LabelPhotoDescription,omitempty"`
	NumberOfDoses string `json:"NumberOfDoses,omitempty"`
	UnitCbdContent float64 `json:"UnitCbdContent,omitempty"`
	UnitCbdAcontentDoseUnitOfMeasure string `json:"UnitCbdAContentDoseUnitOfMeasure,omitempty"`
	ProductPdffileSystemIds string `json:"ProductPDFFileSystemIds,omitempty"`
	Description string `json:"Description,omitempty"`
	LabelImageFileSystemIds string `json:"LabelImageFileSystemIds,omitempty"`
	UnitThcAcontentDose float64 `json:"UnitThcAContentDose,omitempty"`
	ItemCategory string `json:"ItemCategory,omitempty"`
	UnitWeight float64 `json:"UnitWeight,omitempty"`
	ProcessingJobCategoryName string `json:"ProcessingJobCategoryName,omitempty"`
	UnitThcContentUnitOfMeasure string `json:"UnitThcContentUnitOfMeasure,omitempty"`
	UnitCbdApercent float64 `json:"UnitCbdAPercent,omitempty"`
	ProductImageFileSystemIds string `json:"ProductImageFileSystemIds,omitempty"`
	UnitCbdAcontentUnitOfMeasure string `json:"UnitCbdAContentUnitOfMeasure,omitempty"`
	ItemIngredients string `json:"ItemIngredients,omitempty"`
	PublicIngredients string `json:"PublicIngredients,omitempty"`
	UnitCbdContentDoseUnitOfMeasure string `json:"UnitCbdContentDoseUnitOfMeasure,omitempty"`
	UnitWeightUnitOfMeasure string `json:"UnitWeightUnitOfMeasure,omitempty"`
	UnitVolumeUnitOfMeasure string `json:"UnitVolumeUnitOfMeasure,omitempty"`
	Name string `json:"Name,omitempty"`
	GlobalProductName string `json:"GlobalProductName,omitempty"`
	UnitCbdContentDose float64 `json:"UnitCbdContentDose,omitempty"`
	UnitThcApercent float64 `json:"UnitThcAPercent,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Strain string `json:"Strain,omitempty"`
	UnitCbdAcontentDose float64 `json:"UnitCbdAContentDose,omitempty"`
	ServingSize string `json:"ServingSize,omitempty"`
	UnitCbdContentUnitOfMeasure string `json:"UnitCbdContentUnitOfMeasure,omitempty"`
	PackagingPhotoDescription string `json:"PackagingPhotoDescription,omitempty"`
	UnitCbdAcontent float64 `json:"UnitCbdAContent,omitempty"`
	Id int `json:"Id,omitempty"`
	SupplyDurationDays int `json:"SupplyDurationDays,omitempty"`
	PackagingImageFileSystemIds string `json:"PackagingImageFileSystemIds,omitempty"`
	UnitThcContentDoseUnitOfMeasure string `json:"UnitThcContentDoseUnitOfMeasure,omitempty"`
}

type ItemsUpdateBrandV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PackagesCreateV1RequestItem struct {
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Note string `json:"Note,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Tag string `json:"Tag,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
	Item string `json:"Item,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	Ingredients []PackagesCreateV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
}

type PackagesCreateV1RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Item string `json:"Item,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	Location string `json:"Location,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
	Ingredients []PackagesCreateV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
}

type PackagesCreateV2RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateAdjustV1RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesCreateAdjustV2RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesCreateChangeItemV1RequestItem struct {
	Label string `json:"Label,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PackagesCreateChangeLocationV1RequestItem struct {
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	MoveDate string `json:"MoveDate,omitempty"`
}

type PackagesCreateFinishV1RequestItem struct {
	Label string `json:"Label,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
}

type PackagesCreatePlantingsV1RequestItem struct {
	LocationName string `json:"LocationName,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
	PlantedDate string `json:"PlantedDate,omitempty"`
	UnpackagedDate string `json:"UnpackagedDate,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	PackageAdjustmentAmount int `json:"PackageAdjustmentAmount,omitempty"`
	PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
}

type PackagesCreatePlantingsV2RequestItem struct {
	PlantBatchType string `json:"PlantBatchType,omitempty"`
	PlantCount int `json:"PlantCount,omitempty"`
	UnpackagedDate string `json:"UnpackagedDate,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	PackageAdjustmentUnitOfMeasureName string `json:"PackageAdjustmentUnitOfMeasureName,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	PlantedDate string `json:"PlantedDate,omitempty"`
	PackageAdjustmentAmount int `json:"PackageAdjustmentAmount,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	LocationName string `json:"LocationName,omitempty"`
	SublocationName string `json:"SublocationName,omitempty"`
	StrainName string `json:"StrainName,omitempty"`
}

type PackagesCreateRemediateV1RequestItem struct {
	RemediationSteps string `json:"RemediationSteps,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	RemediationMethodName string `json:"RemediationMethodName,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
}

type PackagesCreateTestingV1RequestItem struct {
	Ingredients []PackagesCreateTestingV1RequestItem_Ingredients `json:"Ingredients,omitempty"`
	RequiredLabTestBatches bool `json:"RequiredLabTestBatches,omitempty"`
	Tag string `json:"Tag,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	Location string `json:"Location,omitempty"`
	Note string `json:"Note,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Item string `json:"Item,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PackagesCreateTestingV1RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateTestingV2RequestItem struct {
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Ingredients []PackagesCreateTestingV2RequestItem_Ingredients `json:"Ingredients,omitempty"`
	Note string `json:"Note,omitempty"`
	Tag string `json:"Tag,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	ProductRequiresRemediation bool `json:"ProductRequiresRemediation,omitempty"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	ProductionBatchNumber string `json:"ProductionBatchNumber,omitempty"`
	UseSameItem bool `json:"UseSameItem,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	LabTestStageId int `json:"LabTestStageId,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	Location string `json:"Location,omitempty"`
	IsProductionBatch bool `json:"IsProductionBatch,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	ProcessingJobTypeId int `json:"ProcessingJobTypeId,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PackagesCreateTestingV2RequestItem_Ingredients struct {
	Package string `json:"Package,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
}

type PackagesCreateUnfinishV1RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateAdjustV2RequestItem struct {
	Quantity int `json:"Quantity,omitempty"`
	UnitOfMeasure string `json:"UnitOfMeasure,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateChangeNoteV1RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Note string `json:"Note,omitempty"`
}

type PackagesUpdateDecontaminateV2RequestItem struct {
	DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
	PackageLabel string `json:"PackageLabel,omitempty"`
	DecontaminationMethodName string `json:"DecontaminationMethodName,omitempty"`
	DecontaminationDate string `json:"DecontaminationDate,omitempty"`
}

type PackagesUpdateDonationFlagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateDonationUnflagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateExternalidV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	ExternalId string `json:"ExternalId,omitempty"`
}

type PackagesUpdateFinishV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateItemV2RequestItem struct {
	Item string `json:"Item,omitempty"`
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateLabTestRequiredV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
}

type PackagesUpdateLocationV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	MoveDate string `json:"MoveDate,omitempty"`
}

type PackagesUpdateNoteV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	Note string `json:"Note,omitempty"`
}

type PackagesUpdateRemediateV2RequestItem struct {
	PackageLabel string `json:"PackageLabel,omitempty"`
	RemediationMethodName string `json:"RemediationMethodName,omitempty"`
	RemediationDate string `json:"RemediationDate,omitempty"`
	RemediationSteps string `json:"RemediationSteps,omitempty"`
}

type PackagesUpdateTradesampleFlagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateTradesampleUnflagV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateUnfinishV2RequestItem struct {
	Label string `json:"Label,omitempty"`
}

type PackagesUpdateUsebydateV2RequestItem struct {
	Label string `json:"Label,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
}

type PatientCheckInsCreateV1RequestItem struct {
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PatientCheckInsCreateV2RequestItem struct {
	CheckInDate string `json:"CheckInDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
}

type PatientCheckInsUpdateV1RequestItem struct {
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	Id int `json:"Id,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
}

type PatientCheckInsUpdateV2RequestItem struct {
	CheckInLocationId int `json:"CheckInLocationId,omitempty"`
	RegistrationStartDate string `json:"RegistrationStartDate,omitempty"`
	RegistrationExpiryDate string `json:"RegistrationExpiryDate,omitempty"`
	CheckInDate string `json:"CheckInDate,omitempty"`
	Id int `json:"Id,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PlantBatchesCreateAdditivesV1RequestItem struct {
	AdditiveType string `json:"AdditiveType,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ActiveIngredients []PlantBatchesCreateAdditivesV1RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
}

type PlantBatchesCreateAdditivesV1RequestItem_ActiveIngredients struct {
	Name string `json:"Name,omitempty"`
	Percentage float64 `json:"Percentage,omitempty"`
}

type PlantBatchesCreateAdditivesV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	EpaRegistrationNumber string `json:"EpaRegistrationNumber,omitempty"`
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	ProductTradeName string `json:"ProductTradeName,omitempty"`
	ActiveIngredients []PlantBatchesCreateAdditivesV2RequestItem_ActiveIngredients `json:"ActiveIngredients,omitempty"`
	ProductSupplier string `json:"ProductSupplier,omitempty"`
	ApplicationDevice string `json:"ApplicationDevice,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	AdditiveType string `json:"AdditiveType,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
}

type PlantBatchesCreateAdditivesV2RequestItem_ActiveIngredients struct {
	Percentage float64 `json:"Percentage,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesCreateAdditivesUsingtemplateV2RequestItem struct {
	TotalAmountUnitOfMeasure string `json:"TotalAmountUnitOfMeasure,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	AdditivesTemplateName string `json:"AdditivesTemplateName,omitempty"`
	Rate string `json:"Rate,omitempty"`
	Volume string `json:"Volume,omitempty"`
	TotalAmountApplied int `json:"TotalAmountApplied,omitempty"`
}

type PlantBatchesCreateAdjustV1RequestItem struct {
	ReasonNote string `json:"ReasonNote,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
}

type PlantBatchesCreateAdjustV2RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	Quantity int `json:"Quantity,omitempty"`
	AdjustmentReason string `json:"AdjustmentReason,omitempty"`
	AdjustmentDate string `json:"AdjustmentDate,omitempty"`
	ReasonNote string `json:"ReasonNote,omitempty"`
}

type PlantBatchesCreateChangegrowthphaseV1RequestItem struct {
	NewSublocation string `json:"NewSublocation,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	Count int `json:"Count,omitempty"`
	StartingTag string `json:"StartingTag,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Name string `json:"Name,omitempty"`
	CountPerPlant string `json:"CountPerPlant,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewLocation string `json:"NewLocation,omitempty"`
}

type PlantBatchesCreateGrowthphaseV2RequestItem struct {
	NewLocation string `json:"NewLocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	GrowthPhase string `json:"GrowthPhase,omitempty"`
	NewSublocation string `json:"NewSublocation,omitempty"`
	CountPerPlant string `json:"CountPerPlant,omitempty"`
	Count int `json:"Count,omitempty"`
	StartingTag string `json:"StartingTag,omitempty"`
	GrowthDate string `json:"GrowthDate,omitempty"`
	Name string `json:"Name,omitempty"`
}

type PlantBatchesCreatePackageV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Location string `json:"Location,omitempty"`
	Tag string `json:"Tag,omitempty"`
	Note string `json:"Note,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	Count int `json:"Count,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Item string `json:"Item,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PlantBatchesCreatePackageFrommotherplantV1RequestItem struct {
	Location string `json:"Location,omitempty"`
	Item string `json:"Item,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	Id int `json:"Id,omitempty"`
	Tag string `json:"Tag,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Note string `json:"Note,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Count int `json:"Count,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
}

type PlantBatchesCreatePackageFrommotherplantV2RequestItem struct {
	Tag string `json:"Tag,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Count int `json:"Count,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Note string `json:"Note,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
	Id int `json:"Id,omitempty"`
	Item string `json:"Item,omitempty"`
}

type PlantBatchesCreatePlantingsV2RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	SourcePlantBatches string `json:"SourcePlantBatches,omitempty"`
	Strain string `json:"Strain,omitempty"`
	Name string `json:"Name,omitempty"`
	Type string `json:"Type,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
}

type PlantBatchesCreateSplitV1RequestItem struct {
	GroupName string `json:"GroupName,omitempty"`
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Strain string `json:"Strain,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
}

type PlantBatchesCreateSplitV2RequestItem struct {
	Count int `json:"Count,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Strain string `json:"Strain,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	GroupName string `json:"GroupName,omitempty"`
}

type PlantBatchesCreateWasteV1RequestItem struct {
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
	Note string `json:"Note,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
}

type PlantBatchesCreateWasteV2RequestItem struct {
	Note string `json:"Note,omitempty"`
	WasteDate string `json:"WasteDate,omitempty"`
	PlantBatchName string `json:"PlantBatchName,omitempty"`
	WasteMethodName string `json:"WasteMethodName,omitempty"`
	MixedMaterial string `json:"MixedMaterial,omitempty"`
	WasteWeight float64 `json:"WasteWeight,omitempty"`
	UnitOfMeasureName string `json:"UnitOfMeasureName,omitempty"`
	ReasonName string `json:"ReasonName,omitempty"`
}

type PlantBatchesCreatepackagesV1RequestItem struct {
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	IsDonation bool `json:"IsDonation,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
	Note string `json:"Note,omitempty"`
	SellByDate string `json:"SellByDate,omitempty"`
	Tag string `json:"Tag,omitempty"`
	PlantBatch string `json:"PlantBatch,omitempty"`
	Id int `json:"Id,omitempty"`
	Count int `json:"Count,omitempty"`
	Item string `json:"Item,omitempty"`
	ActualDate string `json:"ActualDate,omitempty"`
	IsTradeSample bool `json:"IsTradeSample,omitempty"`
	ExpirationDate string `json:"ExpirationDate,omitempty"`
	UseByDate string `json:"UseByDate,omitempty"`
}

type PlantBatchesCreateplantingsV1RequestItem struct {
	ActualDate string `json:"ActualDate,omitempty"`
	Count int `json:"Count,omitempty"`
	PatientLicenseNumber string `json:"PatientLicenseNumber,omitempty"`
	Location string `json:"Location,omitempty"`
	SourcePlantBatches string `json:"SourcePlantBatches,omitempty"`
	Name string `json:"Name,omitempty"`
	Type string `json:"Type,omitempty"`
	Strain string `json:"Strain,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PlantBatchesUpdateLocationV2RequestItem struct {
	MoveDate string `json:"MoveDate,omitempty"`
	Name string `json:"Name,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PlantBatchesUpdateMoveplantbatchesV1RequestItem struct {
	MoveDate string `json:"MoveDate,omitempty"`
	Name string `json:"Name,omitempty"`
	Location string `json:"Location,omitempty"`
	Sublocation string `json:"Sublocation,omitempty"`
}

type PlantBatchesUpdateNameV2RequestItem struct {
	NewGroup string `json:"NewGroup,omitempty"`
	Id int `json:"Id,omitempty"`
	Group string `json:"Group,omitempty"`
}

type PlantBatchesUpdateStrainV2RequestItem struct {
	StrainName string `json:"StrainName,omitempty"`
	Id int `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
	StrainId int `json:"StrainId,omitempty"`
}

type PlantBatchesUpdateTagV2RequestItem struct {
	Id int `json:"Id,omitempty"`
	Group string `json:"Group,omitempty"`
	TagId int `json:"TagId,omitempty"`
	NewTag string `json:"NewTag,omitempty"`
	ReplaceDate string `json:"ReplaceDate,omitempty"`
}

