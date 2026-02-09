
export interface CreatePackagesTestingRequestItem {
    ActualDate?: string;
    DecontaminateProduct?: boolean;
    DecontaminationDate?: string;
    DecontaminationSteps?: string;
    ExpirationDate?: string;
    Ingredients?: any[];
    IsDonation?: boolean;
    IsProductionBatch?: boolean;
    IsTradeSample?: boolean;
    Item?: string;
    LabTestStageId?: number;
    Location?: string;
    Note?: string;
    PatientLicenseNumber?: string;
    ProcessingJobTypeId?: number;
    ProductRequiresDecontamination?: boolean;
    ProductRequiresRemediation?: boolean;
    ProductionBatchNumber?: string;
    RemediateProduct?: boolean;
    RemediationDate?: string;
    RemediationMethodId?: number;
    RemediationSteps?: string;
    RequiredLabTestBatches?: any[];
    SellByDate?: string;
    Sublocation?: string;
    Tag?: string;
    UnitOfWeight?: string;
    UseByDate?: string;
}
