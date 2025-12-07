package common

var Mocks = struct {
	Brands             string
	Empty              string
	ActiveItems        string
	ActiveHarvests     string
	VegetativePlants   string
	PackageTags        string
	PlantTags          string
	GenericList        string
	LocationTypes      string
	WasteMethods       string
	LabTestTypes       string
	ActivePackages     string
	TransferTypes      string
	ActiveFacilities   string
	ItemCategories     string
	Facilities         string
	ProcessingJobTypes string
	JobTypeCategories  string
	JobTypeAttributes  string
	Strains            string
	ActiveLocations    string
}{
	Brands:           `[{"Id": 12345, "Name": "Mock Brand"}]`,
	Empty:            `[]`,
	ActiveItems:      `[{"Name": "Mock Item", "UnitOfMeasureName": "Grams", "ProductCategory": {"Name": "Mock Category"}}]`,
	ActiveHarvests:   `[{"Id": 99999, "Name": "Mock Harvest"}]`,
	VegetativePlants: `[{"Label": "ABCDEF012345670000020202"}]`,
	PackageTags:      `[{"Label": "ABCDEF012345670000010011", "Tag": "ABCDEF012345670000010011"}, {"Label": "ABCDEF012345670000010012", "Tag": "ABCDEF012345670000010012"}]`,
	PlantTags:        `[{"Label": "ABCDEF012345670000020011", "Tag": "ABCDEF012345670000020011"}, {"Label": "ABCDEF012345670000020012", "Tag": "ABCDEF012345670000020012"}]`,
	GenericList:      `[{"Id": 1, "Name": "Mock Object"}]`,
	LocationTypes:    `[{"Name": "Default Location Type"}]`,
	WasteMethods:     `[{"Name": "Compost"}]`,
	LabTestTypes:     `[{"Name": "Potency"}]`,
	ActivePackages:   `[{"Label": "ABCDEF012345670000030001", "Quantity": 100.0, "UnitOfMeasureName": "Grams"}]`,
	TransferTypes:    `[{"Name": "Wholesale", "ForExternalIncomingShipments": true, "ForExternalOutgoingShipments": true}]`,
	ActiveFacilities: `[{"License": {"Number": "SF-SBX-1-4001"}, "Name": "Mock Facility B"}]`,
	ItemCategories: `[
		{ "Name": "Buds" },
		{ "Name": "Infused" },
		{ "Name": "Plants", "ProductCategoryType": "Plants" },
		{ "Name": "Vegetative Plant", "ProductCategoryType": "Plants" },
		{ "Name": "Immature Plant", "ProductCategoryType": "Plants" },
		{ "Name": "Seeds", "ProductCategoryType": "Plants" }
	]`,
	Facilities: `[
		{
			"FacilityId": 1,
			"Name": "Sandbox Medical Dispensary",
			"License": { "Number": "SBX-15-4001" },
			"FacilityType": {
				"CanGrowPlants": true,
				"CanCreateOpeningBalancePackages": true,
				"CanAssignLocationsToPackages": true,
				"CanUseComplianceLabel": true,
				"CanSellToPatients": true,
				"CanSellToConsumers": false,
				"CanManageItemBrands": true,
				"CanRecordProcessingJobs": true,
				"CanInfuseProducts": true,
				"RestrictHarvestPlantRestoreTimeHours": 24,
				"RestrictPlantBatchAdjustmentTimeHours": 24,
				"RestrictWholesalePriceEditDays": 10
			}
		}
	]`,
	ProcessingJobTypes: `[{"Name": "Aggregation"}]`,
	JobTypeCategories:  `[{"Name": "Concentrate"}]`,
	JobTypeAttributes:  `[{"Name": "Infused"}]`,
	Strains:            `[{"Name": "Spring Hill Kush"}]`,
	ActiveLocations:    `[{"Name": "Drying Room"}, {"Name": "Storage"}]`,
}
