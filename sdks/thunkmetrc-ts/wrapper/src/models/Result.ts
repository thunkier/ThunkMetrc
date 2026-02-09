
export interface Result {
    ExpirationDateTime?: string;
    LabFacilityLicenseNumber?: string;
    LabFacilityName?: string;
    LabTestDetailRevokedDate?: string;
    LabTestResultDocumentFileId?: number;
    LabTestResultId?: number;
    OverallPassed?: boolean;
    PackageId?: number;
    ProductCategoryName?: string;
    ProductName?: string;
    ResultReleaseDateTime?: string;
    ResultReleased?: boolean;
    RevokedDate?: string;
    SourcePackageLabel?: string;
    TestComment?: string;
    TestInformationalOnly?: boolean;
    TestPassed?: boolean;
    TestPerformedDate?: string;
    TestResultLevel?: number;
    TestTypeName?: string;
}
