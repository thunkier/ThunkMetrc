
export interface PlantBatch {
    DestroyedCount?: number;
    Id?: number;
    IsOnHold?: boolean;
    IsOnInvestigation?: boolean;
    IsOnInvestigationHold?: boolean;
    IsOnInvestigationRecall?: boolean;
    LastModified?: string;
    LocationId?: number;
    LocationName?: string;
    LocationTypeName?: string;
    MultiPlantBatch?: boolean;
    Name?: string;
    PackagedCount?: number;
    PatientLicenseNumber?: string;
    PlantBatchTypeId?: number;
    PlantBatchTypeName?: string;
    PlantedDate?: string;
    SourcePackageId?: number;
    SourcePackageLabel?: string;
    SourcePlantBatchIds?: any[];
    SourcePlantBatchNames?: string;
    SourcePlantId?: number;
    SourcePlantLabel?: string;
    StrainId?: number;
    StrainName?: string;
    SublocationId?: number;
    SublocationName?: string;
    TrackedCount?: number;
    UntrackedCount?: number;
}
