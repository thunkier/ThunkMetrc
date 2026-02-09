package scenarios

import (
	"github.com/thunkier/thunkmetrc/probe/internal/app"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

type Scenario struct {
	Name        string
	Description string
	Matcher     func(f models.Facility) bool
	Steps       []app.DataWorkflow
}

func (s *Scenario) Matches(f models.Facility) bool {
	if s.Matcher != nil {
		return s.Matcher(f)
	}
	return false
}

var ScenarioRegistry = []Scenario{
	{
		Name:        "Medical Grow Site",
		Description: "Unique permissions including PatientAffiliation requirements",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.PackagesRequirePatientAffiliation &&
				f.FacilityType.PlantsRequirePatientAffiliation
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Med Grow)",
				ExcludePatterns: []string{
					"Sales", "Retail", "Processing", "LabTest", "Additives", "Employees",
				},
			},
		},
	},

	{
		Name:        "Research Certificate",
		Description: "Research facility with Grow capabilities",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanGrowPlants &&
				f.FacilityType.CanPackageVegetativePlants &&
				f.FacilityType.CanTrackMotherPlantsAsGrowthPhase &&
				f.FacilityType.CanTrackVegetativePlants &&
				f.FacilityType.CanUpdatePlantStrains &&
				f.FacilityType.IsRetail &&
				f.FacilityType.RestrictHarvestPlantRestoreTimeHours != nil &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip &&
				!f.FacilityType.CanTagPlantBatches &&
				!f.FacilityType.CanReportHarvestSchedules
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Research)",
				ExcludePatterns: []string{
					"Sales", "Retail",
				},
			},
		},
	},

	{
		Name:        "Medical Processor",
		Description: "Medical-only Processor",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanInfuseProducts &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRecordProcessingJobs &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanReturnPackageQuantity &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsMedical
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Med Processor)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "ODA",
		Description: "ODA Facility",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanCreateOpeningBalancePackages &&
				f.FacilityType.CanTestPackages &&
				f.FacilityType.IsMedical &&
				f.FacilityType.IsRetail &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip &&
				!f.FacilityType.CanRequestProductRemediation
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(ODA)",
				ExcludePatterns: []string{
					"Harvest", "Processing",
				},
			},
		},
	},

	{
		Name:        "Hemp Grower Certificate",
		Description: "Hemp Grower (No actual CanGrowPlants permission in sandbox)",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip &&
				!f.FacilityType.CanGrowPlants &&
				!f.FacilityType.CanInfuseProducts &&
				!f.FacilityType.CanRecordProcessingJobs
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Hemp Grower)",
				ExcludePatterns: []string{
					"Sales", "Retail", "Processing", "LabTest",
				},
			},
		},
	},

	{
		Name:        "Recreational Wholesaler",
		Description: "Wholesaler",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsRetail &&
				!f.FacilityType.CanRecordProcessingJobs &&
				!f.FacilityType.CanInfuseProducts &&
				!f.FacilityType.CanGrowPlants
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Rec Wholesaler)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Sales", "Retail", "Patient", "Caregiver", "LabTest",
				},
			},
		},
	},

	{
		Name:        "Hemp Handler Certificate",
		Description: "Hemp Handler (Processor-like)",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanInfuseProducts &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRecordProcessingJobs &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Hemp Handler)",
				ExcludePatterns: []string{
					"Harvest", "Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "Recreational Processor",
		Description: "Recreational Processor",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanInfuseProducts &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRecordProcessingJobs &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanReturnPackageQuantity &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsRetail &&
				!f.FacilityType.CanGrowPlants
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Rec Processor)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "Warm Springs Producer",
		Description: "Producer",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanGrowPlants &&
				f.FacilityType.CanPackageVegetativePlants &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanReportHarvestSchedules &&
				f.FacilityType.CanTrackMotherPlantsAsGrowthPhase &&
				f.FacilityType.CanTrackVegetativePlants &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUpdatePlantStrains &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsRetail &&
				f.FacilityType.RestrictHarvestPlantRestoreTimeHours != nil &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip &&
				!f.FacilityType.CanTagPlantBatches
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Warm Springs)",
				ExcludePatterns: []string{
					"Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "Cow Creek Producer",
		Description: "Producer with Exclusive TagPlantBatches",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanTagPlantBatches
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Cow Creek)",
				ExcludePatterns: []string{
					"Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "Recreational Producer",
		Description: "Generic Rec Producer",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanCreateTradeSamplePackages &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanGrowPlants &&
				f.FacilityType.CanInfuseProducts &&
				f.FacilityType.CanPackageVegetativePlants &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanRecordProcessingJobs &&
				f.FacilityType.CanRemediatePackagesWithFailedLabResults &&
				f.FacilityType.CanReportHarvestSchedules &&
				f.FacilityType.CanReturnPackageQuantity &&
				f.FacilityType.CanTrackMotherPlantsAsGrowthPhase &&
				f.FacilityType.CanTrackVegetativePlants &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.CanUpdatePlantStrains &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsRetail &&
				f.FacilityType.RestrictHarvestPlantRestoreTimeHours != nil
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Rec Producer)",
				ExcludePatterns: []string{
					"Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "OLCC",
		Description: "Generic OLCC",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanAssignLocationsToPackages &&
				f.FacilityType.CanCreateOpeningBalancePackages &&
				f.FacilityType.CanRequestProductRemediation &&
				f.FacilityType.CanTestPackages &&
				f.FacilityType.CanUpdateLocationsOnPackages &&
				f.FacilityType.IsMedical &&
				f.FacilityType.IsRetail &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(OLCC)",
				ExcludePatterns: []string{
					"Sales", "Retail",
				},
			},
		},
	},

	{
		Name:        "Laboratory",
		Description: "Testing Laboratory",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanCreateProcessValidationPackages &&
				f.FacilityType.CanTestPackages &&
				f.FacilityType.IsMedical &&
				f.FacilityType.IsRetail &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Lab)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Processing", "Sales", "Retail", "Patient", "Caregiver",
				},
			},
		},
	},

	{
		Name:        "Medical Dispensary",
		Description: "Generic Medical Dispensary",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.AdvancedSales &&
				f.FacilityType.CanDownloadProductLabel &&
				f.FacilityType.CanReceiveAssociateProductLabel &&
				f.FacilityType.CanSellToCaregivers &&
				f.FacilityType.CanSellToPatients &&
				f.FacilityType.CanUseComplianceLabel &&
				f.FacilityType.IsMedical &&
				f.FacilityType.SalesDeliveryAllowAddress &&
				f.FacilityType.SalesDeliveryAllowCity &&
				f.FacilityType.SalesDeliveryAllowCounty &&
				f.FacilityType.SalesDeliveryAllowPlannedRoute &&
				f.FacilityType.SalesDeliveryAllowState &&
				f.FacilityType.SalesDeliveryAllowZip
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Med Dispensary)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Processing", "LabTest",
				},
			},
		},
	},

	{
		Name:        "Recreational Retailer",
		Description: "Recreational Retailer",
		Matcher: func(f models.Facility) bool {
			return f.FacilityType.CanDeliverSalesToConsumers &&
				f.FacilityType.CanDeliverSalesToPatients &&
				f.FacilityType.CanSellToConsumers &&
				f.FacilityType.SalesDeliveryRequirePatientNumber &&
				f.FacilityType.SalesRequireCaregiverNumber &&
				f.FacilityType.SalesRequireCaregiverPatientNumber &&
				f.FacilityType.SalesRequirePatientNumber
		},
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{
				Suffix: "(Rec Retail)",
				ExcludePatterns: []string{
					"Harvest", "Plant", "Processing",
				},
			},
		},
	},

	{
		Name:        "Universal Fallback",
		Description: "Fallback for unknown facility types",
		Matcher:     func(f models.Facility) bool { return true },
		Steps: []app.DataWorkflow{
			&UniversalGatherWorkflow{Suffix: "(Universal)"},
		},
	},
}
