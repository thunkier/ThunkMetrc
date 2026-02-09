
export interface CreateProcessingJobPackagesRequestItem {
    ExpirationDate?: string;
    FinishDate?: string;
    FinishNote?: string;
    FinishProcessingJob?: boolean;
    IsFinishedGood?: boolean;
    Item?: string;
    JobName?: string;
    Location?: string;
    Note?: string;
    PackageDate?: string;
    PatientLicenseNumber?: string;
    ProductionBatchNumber?: string;
    Quantity?: number;
    SellByDate?: string;
    Sublocation?: string;
    Tag?: string;
    UnitOfMeasure?: string;
    UseByDate?: string;
    WasteCountQuantity?: number;
    WasteCountUnitOfMeasureName?: string;
    WasteVolumeQuantity?: number;
    WasteVolumeUnitOfMeasureName?: string;
    WasteWeightQuantity?: number;
    WasteWeightUnitOfMeasureName?: string;
}
