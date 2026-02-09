package models

type Batch struct {
	Id int `json:"Id"`
	ItemCategories []BatchItemCategory `json:"ItemCategories"`
	ItemCategoryCount int `json:"ItemCategoryCount"`
	LabTestTypeCount int `json:"LabTestTypeCount"`
	LabTestTypes []BatchLabTestType `json:"LabTestTypes"`
	Name string `json:"Name"`
	RequiresAllFromLabTestBatch bool `json:"RequiresAllFromLabTestBatch"`
}

type BatchItemCategory struct {
	ProductCategory BatchItemCategoryProductCategory `json:"ProductCategory"`
	ProductCategoryId string `json:"ProductCategoryId"`
	RequireLabTestBatch bool `json:"RequireLabTestBatch"`
}

type BatchItemCategoryProductCategory struct {
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

type BatchLabTestType struct {
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

type LabTestDocumentRequest struct {
	DocumentFileBase64 string `json:"DocumentFileBase64"`
	DocumentFileName string `json:"DocumentFileName"`
	LabTestResultId int `json:"LabTestResultId"`
}

type LabTestsType struct {
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

type RecordRequest struct {
	DocumentFileBase64 string `json:"DocumentFileBase64"`
	DocumentFileName string `json:"DocumentFileName"`
	Label string `json:"Label"`
	ResultDate string `json:"ResultDate"`
	Results []RecordRequestResult `json:"Results"`
}

type RecordRequestResult struct {
	LabTestTypeName string `json:"LabTestTypeName"`
	Notes string `json:"Notes"`
	Passed bool `json:"Passed"`
	Quantity float64 `json:"Quantity"`
}

type Result struct {
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
}

type ResultsReleaseRequest struct {
	PackageLabel string `json:"PackageLabel"`
}
