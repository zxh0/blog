package utils

import (
	"encoding/json"
)

func ToPrettyJSON(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	return string(bytes)
}
