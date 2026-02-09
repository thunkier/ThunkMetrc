
export interface CreatePackagesPackagesRequestItem {
    ActualDate?: string;
    ExpirationDate?: string;
    Ingredients?: any[];
    IsDonation?: boolean;
    IsFinishedGood?: boolean;
    IsProductionBatch?: boolean;
    IsTradeSample?: boolean;
    Item?: string;
    LabTestStageId?: number;
    Location?: string;
    Note?: string;
    PatientLicenseNumber?: string;
    ProcessingJobTypeId?: number;
    ProductRequiresRemediation?: boolean;
    ProductionBatchNumber?: string;
    Quantity?: number;
    RequiredLabTestBatches?: boolean;
    SellByDate?: string;
    Sublocation?: string;
    Tag?: string;
    UnitOfMeasure?: string;
    UseByDate?: string;
    UseSameItem?: boolean;
}
