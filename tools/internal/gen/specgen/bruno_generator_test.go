package specgen

import "testing"

func TestGetBrunoFileName(t *testing.T) {
	tests := []struct {
		opName    string
		groupName string
		expected  string
	}{
		{
			opName:    "CreateBatch",
			groupName: "Plants",
			expected:  "CreatePlantsBatch",
		},
		{
			opName:    "UpdateLocation",
			groupName: "Facilities",
			expected:  "UpdateFacilitiesLocation",
		},
		{
			opName:    "GetById",
			groupName: "Packages",
			expected:  "GetPackagesById",
		},
		{
			opName:    "DeleteItem",
			groupName: "Sales",
			expected:  "DeleteSalesItem",
		},
		{
			opName:    "StartProcessing",
			groupName: "Plant Batches",
			expected:  "StartPlantBatchesProcessing",
		},
		{
			opName:    "CustomOp",
			groupName: "Admin",
			expected:  "AdminCustomOp",
		},
		{
			opName:    "Create",
			groupName: "Transfers",
			expected:  "CreateTransfers",
		},
		{
			opName:    "AdjustPackage",
			groupName: "Packages",
			expected:  "AdjustPackagesPackage",
		},
	}

	for _, tt := range tests {
		t.Run(tt.opName+"_"+tt.groupName, func(t *testing.T) {
			got := getBrunoFileName(tt.opName, tt.groupName)
			if got != tt.expected {
				t.Errorf("getBrunoFileName(%q, %q) = %v; want %v", tt.opName, tt.groupName, got, tt.expected)
			}
		})
	}
}
