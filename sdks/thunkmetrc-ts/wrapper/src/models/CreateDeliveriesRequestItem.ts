
export interface CreateDeliveriesRequestItem {
    ConsumerId?: number;
    DriverEmployeeId?: string;
    DriverName?: string;
    DriversLicenseNumber?: string;
    EstimatedArrivalDateTime?: string;
    EstimatedDepartureDateTime?: string;
    PatientLicenseNumber?: string;
    PhoneNumberForQuestions?: string;
    PlannedRoute?: string;
    RecipientAddressCity?: string;
    RecipientAddressCounty?: string;
    RecipientAddressPostalCode?: string;
    RecipientAddressState?: string;
    RecipientAddressStreet1?: string;
    RecipientAddressStreet2?: string;
    RecipientName?: string;
    RecipientZoneId?: number;
    SalesCustomerType?: string;
    SalesDateTime?: string;
    Transactions?: any[];
    TransporterFacilityLicenseNumber?: string;
    VehicleLicensePlateNumber?: string;
    VehicleMake?: string;
    VehicleModel?: string;
}
