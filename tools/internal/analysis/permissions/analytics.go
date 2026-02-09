package permissions

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"

	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

func Run() {
	path := "../specs/source/captured/responses/Facilities/GetFacilities.json"
	data, err := os.ReadFile(path)
	if err != nil {
		path = "specs/source/captured/responses/Facilities/GetFacilities.json"
		data, err = os.ReadFile(path)
		if err != nil {
			panic(fmt.Sprintf("Could not read file: %v", err))
		}
	}

	var facilities []models.Facility
	if err := json.Unmarshal(data, &facilities); err != nil {
		panic(fmt.Sprintf("Could not parse JSON: %v", err))
	}

	fmt.Printf("Analyzing %d facilities...\n", len(facilities))

	totalFacilities := len(facilities)
	fieldCounts := make(map[string]int)
	fieldToFacilities := make(map[string][]string)

	for _, f := range facilities {
		val := reflect.ValueOf(f.FacilityType)
		typ := val.Type()

		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			value := val.Field(i)
			fieldName := field.Name

			isActive := false

			switch value.Kind() {
			case reflect.Bool:
				isActive = value.Bool()
			case reflect.Ptr:
				isActive = !value.IsNil()
			}

			if isActive {
				fieldCounts[fieldName]++
				fieldToFacilities[fieldName] = append(fieldToFacilities[fieldName], fmt.Sprintf("%s (%s)", f.Name, f.License.LicenseType))
			}
		}
	}

	var universal []string
	var unique []string
	var unused []string
	var common []string

	for field, count := range fieldCounts {
		switch count {
		case totalFacilities:
			universal = append(universal, field)
		case 1:
			unique = append(unique, field)
		default:
			common = append(common, field)
		}
	}

	val := reflect.ValueOf(models.FacilityType{})
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		name := typ.Field(i).Name
		if _, exists := fieldCounts[name]; !exists {
			unused = append(unused, name)
		}
	}

	sort.Strings(universal)
	sort.Strings(unique)
	sort.Strings(common)
	sort.Strings(unused)

	fmt.Println("\n=== UNIVERSAL PERMISSIONS (All Facilities Have These) ===")
	for _, f := range universal {
		fmt.Printf(" - %s\n", f)
	}

	fmt.Println("\n=== UNIQUE PERMISSIONS (Only 1 Facility Has These) ===")
	for _, f := range unique {
		owner := fieldToFacilities[f][0]
		fmt.Printf(" - %s: %s\n", f, owner)
	}

	fmt.Println("\n=== UNUSED PERMISSIONS (No Facilities Have These) ===")
	fmt.Printf("Total: %d unused fields (omitting list for brevity, but here are first 5):\n", len(unused))
	for i := 0; i < 5 && i < len(unused); i++ {
		fmt.Printf(" - %s\n", unused[i])
	}

	fmt.Println("\n=== COMMON PERMISSIONS (Shared by some, but not all) ===")
	for _, f := range common {
		count := fieldCounts[f]
		fmt.Printf(" - %s: %d facilities\n", f, count)
	}

	fmt.Println("\n=== FACILITY SIMILARITY GROUPS ===")

	type GroupData struct {
		Members      []string
		FacilityType models.FacilityType
	}
	groups := make(map[string]GroupData)

	for _, f := range facilities {
		sig := signature(f.FacilityType)
		if _, exists := groups[sig]; !exists {
			groups[sig] = GroupData{
				FacilityType: f.FacilityType,
			}
		}
		entry := groups[sig]
		entry.Members = append(entry.Members, fmt.Sprintf("%s (%s)", f.Name, f.License.LicenseType))
		groups[sig] = entry
	}

	type SortedGroup struct {
		Sig  string
		Data GroupData
	}
	var sortedGroups []SortedGroup
	for sig, data := range groups {
		sortedGroups = append(sortedGroups, SortedGroup{Sig: sig, Data: data})
	}
	sort.Slice(sortedGroups, func(i, j int) bool {
		//Sort by size (desc), then by signature (stable)
		if len(sortedGroups[i].Data.Members) != len(sortedGroups[j].Data.Members) {
			return len(sortedGroups[i].Data.Members) > len(sortedGroups[j].Data.Members)
		}
		return sortedGroups[i].Sig < sortedGroups[j].Sig
	})

	for i, g := range sortedGroups {
		fmt.Printf("Group %d (%d members):\n", i+1, len(g.Data.Members))
		for _, m := range g.Data.Members {
			fmt.Printf("  - %s\n", m)
		}

		fmt.Println("  Distinctive Capabilities:")
		val := reflect.ValueOf(g.Data.FacilityType)
		typ := val.Type()
		hasDistinctive := false

		var distinctFields []string

		for j := 0; j < val.NumField(); j++ {
			field := typ.Field(j)
			value := val.Field(j)
			fieldName := field.Name

			isActive := false
			switch value.Kind() {
			case reflect.Bool:
				isActive = value.Bool()
			case reflect.Ptr:
				isActive = !value.IsNil()
			}

			if isActive {
				if fieldCounts[fieldName] == totalFacilities {
					continue
				}

				output := fieldName
				if fieldCounts[fieldName] == len(g.Data.Members) {
					output += " (Exclusive)"
				}
				distinctFields = append(distinctFields, output)
			}
		}

		sort.Strings(distinctFields)
		for _, f := range distinctFields {
			fmt.Printf("   + %s\n", f)
			hasDistinctive = true
		}

		if !hasDistinctive {
			fmt.Println("   (None - Only Universal Permissions)")
		}
		fmt.Println()
	}
}

func signature(ft models.FacilityType) string {

	return fmt.Sprintf("%v", ft)
}
