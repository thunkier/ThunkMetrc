package util

import (
	"encoding/json"
	"fmt"
	"strings"
)

func HasData(body string) bool {
	body = strings.TrimSpace(body)
	if body == "" {
		return false
	}

	if strings.HasPrefix(body, "[") {
		var arr []interface{}
		if err := json.Unmarshal([]byte(body), &arr); err == nil && len(arr) > 0 {
			return true
		}
	}

	if strings.HasPrefix(body, "{") {
		var obj map[string]interface{}
		if err := json.Unmarshal([]byte(body), &obj); err == nil {
			if data, ok := obj["Data"]; ok {
				if dataArr, ok := data.([]interface{}); ok && len(dataArr) > 0 {
					return true
				}
			}
		}
	}

	return false
}

func ParseList(body string, target interface{}) error {
	var wrapper struct {
		Data json.RawMessage `json:"Data"`
	}
	if err := json.Unmarshal([]byte(body), &wrapper); err == nil && len(wrapper.Data) > 0 {
		return json.Unmarshal(wrapper.Data, target)
	}
	return json.Unmarshal([]byte(body), target)
}

func TransformBody(body string, modifier func(map[string]interface{}) bool) (string, error) {
	var wrapper struct {
		Data []map[string]interface{} `json:"Data"`
	}
	if err := json.Unmarshal([]byte(body), &wrapper); err == nil && wrapper.Data != nil {
		var newData []map[string]interface{}
		changed := false
		for _, item := range wrapper.Data {
			if modifier(item) {
				newData = append(newData, item)
			}
			changed = true
		}
		if changed {
			wrapper.Data = newData
			bytes, err := json.MarshalIndent(wrapper, "", "  ")
			return string(bytes), err
		}
		return body, nil
	}

	var arrReq []map[string]interface{}
	if err := json.Unmarshal([]byte(body), &arrReq); err == nil {
		var newData []map[string]interface{}
		for _, item := range arrReq {
			if modifier(item) {
				newData = append(newData, item)
			}
		}
		bytes, err := json.MarshalIndent(newData, "", "  ")
		return string(bytes), err
	}

	var objReq map[string]interface{}
	if err := json.Unmarshal([]byte(body), &objReq); err == nil {
		if !modifier(objReq) {
			return "[]", nil
		}
		bytes, err := json.MarshalIndent(objReq, "", "  ")
		return string(bytes), err
	}
	return body, nil
}

func ExtractID(body string) string {
	var resp struct {
		Ids []int `json:"Ids"`
		Id  int   `json:"Id"`
	}
	if err := json.Unmarshal([]byte(body), &resp); err == nil {
		if len(resp.Ids) > 0 {
			return fmt.Sprintf("%d", resp.Ids[0])
		}
		if resp.Id != 0 {
			return fmt.Sprintf("%d", resp.Id)
		}
	}
	return ""
}

func ExtractNames(jsonBody string) []string {
	var names []string
	jsonBody = strings.TrimSpace(jsonBody)

	extractFromMap := func(m map[string]interface{}) []string {
		var found []string
		addIfStr := func(key string) {
			if v, ok := m[key].(string); ok && v != "" {
				found = append(found, v)
			}
		}
		keys := []string{
			"Name", "Tag", "Label", "PackageLabel", "LicenseNumber",
			"JobName", "ProcessingJobType", "PlantBatchName",
			"LicensePlateNumber", "DriverName", "DriversLicenseNumber", "EmployeeId",
		}
		for _, k := range keys {
			addIfStr(k)
		}

		if v, ok := m["PlantLabels"].([]interface{}); ok {
			for _, item := range v {
				if s, ok := item.(string); ok && s != "" {
					found = append(found, s)
				}
			}
		}
		return found
	}

	var arr []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonBody), &arr); err == nil {
		for _, item := range arr {
			names = append(names, extractFromMap(item)...)
		}
		return names
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(jsonBody), &obj); err == nil {
		names = append(names, extractFromMap(obj)...)
	}

	return names
}

func ExtractExistingEntities(jsonBody string) map[string]int {
	result := make(map[string]int)
	type Entity struct {
		Id    int    `json:"Id"`
		Name  string `json:"Name"`
		Label string `json:"Label"`
		Tag   string `json:"Tag"`
	}
	var arr []Entity
	if err := json.Unmarshal([]byte(jsonBody), &arr); err == nil {
		for _, item := range arr {
			if item.Id == 0 {
				continue
			}
			if item.Name != "" {
				result[item.Name] = item.Id
			}
			if item.Label != "" {
				result[item.Label] = item.Id
			}
			if item.Tag != "" {
				result[item.Tag] = item.Id
			}
		}
	}
	return result
}
