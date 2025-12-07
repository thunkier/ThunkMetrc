package models

type LabTest struct {
	Ids []float64 `json:"Ids"`
	Warnings *string `json:"Warnings"`
	CurrentPage int `json:"CurrentPage"`
	Data []LabTestData `json:"Data"`
	Page int `json:"Page"`
	PageSize int `json:"PageSize"`
	RecordsOnPage int `json:"RecordsOnPage"`
	Total int `json:"Total"`
	TotalPages int `json:"TotalPages"`
	TotalRecords int `json:"TotalRecords"`
	ExpirationDateTime *string `json:"ExpirationDateTime"`
	LabFacilityLicenseNumber string `json:"LabFacilityLicenseNumber"`
	LabFacilityName string `json:"LabFacilityName"`
	LabTestDetailRevokedDate *string `json:"LabTestDetailRevokedDate"`
	LabTestResultDocumentFileId *int `json:"LabTestResultDocumentFileId"`
	LabTestResultId int `json:"LabTestResultId"`
	OverallPassed bool `json:"OverallPassed"`
	PackageId int `json:"PackageId"`
	ProductCategoryName string `json:"ProductCategoryName"`
	ProductName string `json:"ProductName"`
	ResultReleaseDateTime string `json:"ResultReleaseDateTime"`
	ResultReleased bool `json:"ResultReleased"`
	RevokedDate *string `json:"RevokedDate"`
	SourcePackageLabel string `json:"SourcePackageLabel"`
	TestComment string `json:"TestComment"`
	TestInformationalOnly bool `json:"TestInformationalOnly"`
	TestPassed bool `json:"TestPassed"`
	TestPerformedDate string `json:"TestPerformedDate"`
	TestResultLevel float64 `json:"TestResultLevel"`
	TestTypeName string `json:"TestTypeName"`
	AlwaysPasses bool `json:"AlwaysPasses"`
	DependencyMode int `json:"DependencyMode"`
	Id int `json:"Id"`
	InformationalOnly bool `json:"InformationalOnly"`
	LabTestResultExpirationDays *int `json:"LabTestResultExpirationDays"`
	LabTestResultMaximum *string `json:"LabTestResultMaximum"`
	LabTestResultMinimum *string `json:"LabTestResultMinimum"`
	LabTestResultMode int `json:"LabTestResultMode"`
	MaxAllowedFailureCount int `json:"MaxAllowedFailureCount"`
	Name string `json:"Name"`
	RequiresTestResult bool `json:"RequiresTestResult"`
	ResearchAndDevelopment bool `json:"ResearchAndDevelopment"`
}

type LabTestData struct {
	Id int `json:"Id"`
	ItemCategories []LabTestDataItemCategory `json:"ItemCategories"`
	ItemCategoryCount int `json:"ItemCategoryCount"`
	LabTestTypeCount int `json:"LabTestTypeCount"`
	LabTestTypes []LabTestDataLabTestType `json:"LabTestTypes"`
	Name string `json:"Name"`
	RequiresAllFromLabTestBatch bool `json:"RequiresAllFromLabTestBatch"`
	ExpirationDateTime *string `json:"ExpirationDateTime"`
	LabFacilityLicenseNumber string `json:"LabFacilityLicenseNumber"`
	LabFacilityName string `json:"LabFacilityName"`
	LabTestDetailRevokedDate *string `json:"LabTestDetailRevokedDate"`
	LabTestResultDocumentFileId *int `json:"LabTestResultDocumentFileId"`
	LabTestResultId int `json:"LabTestResultId"`
	OverallPassed bool `json:"OverallPassed"`
	PackageId int `json:"PackageId"`
	ProductCategoryName string `json:"ProductCategoryName"`
	ProductName string `json:"ProductName"`
	ResultReleaseDateTime string `json:"ResultReleaseDateTime"`
	ResultReleased bool `json:"ResultReleased"`
	RevokedDate *string `json:"RevokedDate"`
	SourcePackageLabel string `json:"SourcePackageLabel"`
	TestComment string `json:"TestComment"`
	TestInformationalOnly bool `json:"TestInformationalOnly"`
	TestPassed bool `json:"TestPassed"`
	TestPerformedDate string `json:"TestPerformedDate"`
	TestResultLevel float64 `json:"TestResultLevel"`
	TestTypeName string `json:"TestTypeName"`
	AlwaysPasses bool `json:"AlwaysPasses"`
	DependencyMode int `json:"DependencyMode"`
	InformationalOnly bool `json:"InformationalOnly"`
	LabTestResultExpirationDays *int `json:"LabTestResultExpirationDays"`
	LabTestResultMaximum *string `json:"LabTestResultMaximum"`
	LabTestResultMinimum *string `json:"LabTestResultMinimum"`
	LabTestResultMode int `json:"LabTestResultMode"`
	MaxAllowedFailureCount int `json:"MaxAllowedFailureCount"`
	RequiresTestResult bool `json:"RequiresTestResult"`
	ResearchAndDevelopment bool `json:"ResearchAndDevelopment"`
}

type LabTestDataItemCategory struct {
	ProductCategory LabTestDataItemCategoryProductCategory `json:"ProductCategory"`
	ProductCategoryId string `json:"ProductCategoryId"`
	RequireLabTestBatch bool `json:"RequireLabTestBatch"`
}

type LabTestDataItemCategoryProductCategory struct {
	CanBeDecontaminated bool `json:"CanBeDecontaminated"`
	CanBeDestroyed bool `json:"CanBeDestroyed"`
	CanBeRemediated bool `json:"CanBeRemediated"`
	CanContainSeeds bool `json:"CanContainSeeds"`
	LabTestBatchNames []interface{} `json:"LabTestBatchNames"`
	Name string `json:"Name"`
	ProductCategoryType int `json:"ProductCategoryType"`
	QuantityType int `json:"QuantityType"`
	RequiresAdministrationMethod bool `json:"RequiresAdministrationMethod"`
	RequiresAllergens bool `json:"RequiresAllergens"`
	RequiresDescription bool `json:"RequiresDescription"`
	RequiresItemBrand bool `json:"RequiresItemBrand"`
	RequiresLabelPhotoDescription bool `json:"RequiresLabelPhotoDescription"`
	RequiresLabelPhotos int `json:"RequiresLabelPhotos"`
	RequiresNumberOfDoses bool `json:"RequiresNumberOfDoses"`
	RequiresPackagingPhotoDescription bool `json:"RequiresPackagingPhotoDescription"`
	RequiresPackagingPhotos int `json:"RequiresPackagingPhotos"`
	RequiresProductPDFDocuments int `json:"RequiresProductPDFDocuments"`
	RequiresProductPhotoDescription bool `json:"RequiresProductPhotoDescription"`
	RequiresProductPhotos int `json:"RequiresProductPhotos"`
	RequiresPublicIngredients bool `json:"RequiresPublicIngredients"`
	RequiresServingSize bool `json:"RequiresServingSize"`
	RequiresStrain bool `json:"RequiresStrain"`
	RequiresSupplyDurationDays bool `json:"RequiresSupplyDurationDays"`
	RequiresUnitCbdAContent bool `json:"RequiresUnitCbdAContent"`
	RequiresUnitCbdAContentDose bool `json:"RequiresUnitCbdAContentDose"`
	RequiresUnitCbdAPercent bool `json:"RequiresUnitCbdAPercent"`
	RequiresUnitCbdContent bool `json:"RequiresUnitCbdContent"`
	RequiresUnitCbdContentDose bool `json:"RequiresUnitCbdContentDose"`
	RequiresUnitCbdPercent bool `json:"RequiresUnitCbdPercent"`
	RequiresUnitThcAContent bool `json:"RequiresUnitThcAContent"`
	RequiresUnitThcAContentDose bool `json:"RequiresUnitThcAContentDose"`
	RequiresUnitThcAPercent bool `json:"RequiresUnitThcAPercent"`
	RequiresUnitThcContent bool `json:"RequiresUnitThcContent"`
	RequiresUnitThcContentDose bool `json:"RequiresUnitThcContentDose"`
	RequiresUnitThcPercent bool `json:"RequiresUnitThcPercent"`
	RequiresUnitVolume bool `json:"RequiresUnitVolume"`
	RequiresUnitWeight bool `json:"RequiresUnitWeight"`
}

type LabTestDataLabTestType struct {
	AlwaysPasses bool `json:"AlwaysPasses"`
	DependencyMode int `json:"DependencyMode"`
	Id int `json:"Id"`
	InformationalOnly bool `json:"InformationalOnly"`
	LabTestResultExpirationDays *int `json:"LabTestResultExpirationDays"`
	LabTestResultMaximum *string `json:"LabTestResultMaximum"`
	LabTestResultMinimum *string `json:"LabTestResultMinimum"`
	LabTestResultMode int `json:"LabTestResultMode"`
	MaxAllowedFailureCount int `json:"MaxAllowedFailureCount"`
	Name string `json:"Name"`
	RequiresTestResult bool `json:"RequiresTestResult"`
	ResearchAndDevelopment bool `json:"ResearchAndDevelopment"`
}

