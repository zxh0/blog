package utils

import (
	"encoding/json"
)

func ToPrettyJSON(v interface{}) string {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	return string(bytes)
}

func CommonPrefixLen(s1, s2 string) int {
	n1, n2 := len(s1), len(s2)
	i := 0
	for ; i < n1 && i < n2; i++ {
		if s1[i] != s2[i] {
			break
		}
	}
	return i
}
