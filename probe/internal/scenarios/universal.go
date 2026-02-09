package scenarios

import (
	"reflect"
	"strings"

	"github.com/thunkier/thunkmetrc/probe/internal/app"
	"github.com/thunkier/thunkmetrc/probe/internal/metrc/client"
)

type UniversalGatherWorkflow struct {
	Suffix          string
	ExcludePatterns []string
}

func (w *UniversalGatherWorkflow) Name() string {
	if w.Suffix != "" {
		return "Universal Data Gather " + w.Suffix
	}
	return "Universal Data Gather"
}

func (w *UniversalGatherWorkflow) ShouldRun(cfg app.Config) bool {
	return true
}

func (w *UniversalGatherWorkflow) Execute(cfg app.Config, f *client.Fetcher) (*app.WorkflowResult, error) {
	result := &app.WorkflowResult{
		DataGenerated:    make(map[string]int),
		SkippedEndpoints: make(map[string]string),
		EmptyEndpoints:   make([]string, 0),
	}

	val := reflect.ValueOf(f)
	typ := val.Type()

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		methodName := method.Name

		if !strings.HasPrefix(methodName, "Get") {
			continue
		}

		if methodName == "GetFacilities" {
			result.SkippedEndpoints[methodName] = "Reserved/Excluded"
			continue
		}

		allowlist := []string{
			"Types", "Reasons", "Phases", "States", "Statuses",
			"Waste", "CheckIns",
		}

		isAllowed := false
		for _, allow := range allowlist {
			if strings.Contains(methodName, allow) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			shouldExclude := false
			for _, pattern := range w.ExcludePatterns {
				if strings.Contains(methodName, pattern) {
					shouldExclude = true
					break
				}
			}
			if shouldExclude {
				result.SkippedEndpoints[methodName] = "Matched Exclude Pattern"
				continue
			}
		}

		methodVal := val.Method(i)
		methodType := methodVal.Type()

		var callArgs []reflect.Value
		shouldSkip := false

		if methodType.NumIn() == 0 {
			callArgs = []reflect.Value{}
		} else {
			argType := methodType.In(0)

			if argType.Kind() == reflect.Ptr {
				structType := argType.Elem()
				arg := reflect.New(structType)
				argElem := arg.Elem()

				shouldSkip = true
				for fIdx := 0; fIdx < structType.NumField(); fIdx++ {
					field := structType.Field(fIdx)
					fieldName := field.Name

					if fieldName == "LicenseNumber" {
						if fVal := argElem.Field(fIdx); fVal.CanSet() {
							fVal.SetString(cfg.FacilityLicenseNumber)
							shouldSkip = false
						}
						continue
					}
				}
				callArgs = []reflect.Value{arg}
			} else {
				allStrings := true
				for j := 0; j < methodType.NumIn(); j++ {
					if methodType.In(j).Kind() != reflect.String {
						allStrings = false
						break
					}
				}

				if allStrings {
					callArgs = make([]reflect.Value, methodType.NumIn())
					if methodType.NumIn() > 0 {
						callArgs[0] = reflect.ValueOf(cfg.FacilityLicenseNumber)
						for k := 1; k < methodType.NumIn(); k++ {
							callArgs[k] = reflect.ValueOf("")
						}
					}
				} else {
					result.SkippedEndpoints[methodName] = "Invalid Signature (Arg Type)"
					continue
				}
			}
		}

		if methodType.NumOut() != 2 {
			result.SkippedEndpoints[methodName] = "Invalid Signature (Returns)"
			continue
		}

		if strings.Contains(methodName, "ById") {
			result.SkippedEndpoints[methodName] = "Requires ID"
			continue
		}

		if shouldSkip {
			result.SkippedEndpoints[methodName] = "Requires Specific Params"
			continue
		}

		retvals := methodVal.Call(callArgs)

		var err error
		if len(retvals) > 0 {
			errVal := retvals[len(retvals)-1]
			if errVal.Kind() == reflect.Interface && !errVal.IsNil() {
				err = errVal.Interface().(error)
			}
		}

		if err != nil {

			cfg.Logger.Log("      -> Failed: %v", err)
		} else {
			count := 0
			if len(retvals) > 1 {
				resVal := retvals[0]
				if resVal.Kind() == reflect.Slice || resVal.Kind() == reflect.Array {
					count = resVal.Len()
					if count == 0 {
						result.EmptyEndpoints = append(result.EmptyEndpoints, methodName)
					}
				} else if resVal.Kind() != reflect.Invalid {
					isNil := false
					switch resVal.Kind() {
					case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface:
						if resVal.IsNil() {
							isNil = true
						}
					}

					if !isNil {
						count = 1
					}
				}
			}
			cfg.Logger.Log("      -> Success: Retrieved %d items", count)
			result.DataGenerated[methodName] = count
		}
	}

	return result, nil
}
