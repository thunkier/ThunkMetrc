
import type { SalesDeliveryRetailerDestinationsItem } from './SalesDeliveryRetailerDestinationsItem';
import type { SalesDeliveryRetailerPackagesItem } from './SalesDeliveryRetailerPackagesItem';
export interface DeliveryRetailer {
    AcceptedDateTime?: string;
    ActualDepartureDateTime?: string;
    AllowFullEdit?: boolean;
    CompletedDateTime?: string;
    DateTime?: string;
    Destinations?: SalesDeliveryRetailerDestinationsItem[];
    Direction?: string;
    DriverEmployeeId?: string;
    DriverName?: string;
    DriversLicenseNumber?: string;
    EstimatedDepartureDateTime?: string;
    FacilityLicenseNumber?: string;
    FacilityName?: string;
    Id?: number;
    LastModified?: string;
    Leg?: number;
    Packages?: SalesDeliveryRetailerPackagesItem[];
    RecordedByUserName?: string;
    RecordedDateTime?: string;
    RestockDateTime?: string;
    RetailerDeliveryNumber?: string;
    RetailerDeliveryState?: string;
    TotalPackages?: number;
    TotalPrice?: number;
    TotalPriceSold?: number;
    VehicleInfo?: string;
    VehicleLicensePlateNumber?: string;
    VehicleMake?: string;
    VehicleModel?: string;
    VoidedDate?: string;
}
