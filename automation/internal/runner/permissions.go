package runner

import (
	"strings"

	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

func ShouldRunGroup(group string, ft models.FacilityType) bool {
	lower := strings.ToLower(group)

	switch {
	case strings.Contains(lower, "plant"), strings.Contains(lower, "harvest"):
		return ft.CanGrowPlants

	case strings.Contains(lower, "strain"):
		return ft.CanUpdatePlantStrains || ft.CanReportStrainProperties

	case strings.Contains(lower, "additive"):
		return ft.CanInfuseProducts

	case strings.Contains(lower, "processing"):
		return ft.CanRecordProcessingJobs || ft.CanInfuseProducts

	case strings.Contains(lower, "transfer"):
		if strings.Contains(lower, "external") {
			return ft.EnableExternalIdentifier
		}
		return ft.CanTransferFromExternalFacilities

	case strings.Contains(lower, "package"):
		return ft.CanCreateOpeningBalancePackages || ft.CanCreateDerivedPackages || ft.CanSubmitPackagesForTesting

	case strings.Contains(lower, "lab") && strings.Contains(lower, "test"):
		return ft.CanTestPackages

	case strings.Contains(lower, "item"), strings.Contains(lower, "waste"):
		return true

	case strings.Contains(lower, "sublocation"):
		return ft.EnableSublocations

	case strings.Contains(lower, "location"):
		return ft.CanAssignLocationsToPackages

	case strings.Contains(lower, "delivery"):
		return ft.CanDeliverSalesToConsumers || ft.CanDeliverSalesToPatients || ft.RetailerDelivery

	case strings.Contains(lower, "sale"):
		return ft.CanSellToConsumers || ft.CanSellToPatients || ft.CanSellToExternalPatients

	case strings.Contains(lower, "patient"):
		return ft.CanHaveMemberPatients || ft.CanReportPatientCheckIns || ft.CanSellToPatients

	case strings.Contains(lower, "caregiver"):
		return ft.CanSellToCaregivers

	case strings.Contains(lower, "employee"):
		return true

	case strings.Contains(lower, "tag"):
		return ft.CanUseComplianceLabel || ft.CanTagPlantBatches
	}

	return true
}
