
export interface Harvest {
    ArchivedDate?: string;
    CurrentWeight?: number;
    DryingLocationId?: number;
    DryingLocationName?: string;
    DryingLocationTypeName?: string;
    DryingSublocationId?: number;
    DryingSublocationName?: string;
    FinishedDate?: string;
    HarvestStartDate?: string;
    HarvestType?: string;
    Id?: number;
    IsOnHold?: boolean;
    IsOnInvestigation?: boolean;
    IsOnInvestigationHold?: boolean;
    IsOnInvestigationRecall?: boolean;
    LabTestingState?: string;
    LabTestingStateDate?: string;
    LastModified?: string;
    Name?: string;
    PackageCount?: number;
    PatientLicenseNumber?: string;
    PlantCount?: number;
    SourceStrainCount?: number;
    SourceStrainNames?: string;
    Strains?: any[];
    TotalPackagedWeight?: number;
    TotalRestoredWeight?: number;
    TotalWasteWeight?: number;
    TotalWetWeight?: number;
    UnitOfWeightName?: string;
}
