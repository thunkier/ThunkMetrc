
export interface PatientsStatus {
    Active?: boolean;
    ConcentrateOuncesAllowed?: number;
    ConcentrateOuncesAvailable?: number;
    ConcentrateOuncesPurchased?: number;
    FlowerOuncesAllowed?: number;
    FlowerOuncesAvailable?: number;
    FlowerOuncesPurchased?: number;
    IdentificationMethod?: string;
    InfusedOuncesAllowed?: number;
    InfusedOuncesAvailable?: number;
    InfusedOuncesPurchased?: number;
    PatientLicenseExpirationDate?: string;
    PatientLicenseNumber?: string;
    PatientRegistrationStartDate?: string;
    PurchaseAmountDays?: number;
    RegistrationNumber?: string;
    ThcOuncesAllowed?: number;
    ThcOuncesAvailable?: number;
    ThcOuncesPurchased?: number;
}
