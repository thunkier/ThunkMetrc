package util

import (
	"encoding/json"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomHexString(length int) string {
	hexCharset := "0123456789abcdef"
	b := make([]byte, length)
	for i := range b {
		b[i] = hexCharset[seededRand.Intn(len(hexCharset))]
	}
	return string(b)
}

func RandomizeConflictingNames(jsonBody string, conflicts []string) string {
	isConflict := func(val string) bool {
		for _, c := range conflicts {
			if c == val {
				return true
			}
		}
		return false
	}

	var data interface{}
	if err := json.Unmarshal([]byte(jsonBody), &data); err != nil {
		return jsonBody
	}

	var walk func(v interface{}) interface{}
	walk = func(v interface{}) interface{} {
		switch val := v.(type) {
		case map[string]interface{}:
			newMap := make(map[string]interface{})
			for k, v2 := range val {
				if k == "Name" || k == "Tag" || k == "Label" {
					if s, ok := v2.(string); ok && isConflict(s) {
						suffix := GenerateRandomString(5)
						if k == "Tag" {
							newMap[k] = s + "-" + suffix
						} else {
							newMap[k] = s + " " + suffix
						}
						continue
					}
				}
				newMap[k] = walk(v2)
			}
			return newMap
		case []interface{}:
			newArr := make([]interface{}, len(val))
			for i, v2 := range val {
				newArr[i] = walk(v2)
			}
			return newArr
		default:
			return val
		}
	}

	newData := walk(data)
	out, err := json.MarshalIndent(newData, "", "  ")
	if err != nil {
		return jsonBody
	}
	return string(out)
}
