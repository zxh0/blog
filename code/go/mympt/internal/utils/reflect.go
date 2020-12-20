package utils

import (
	"reflect"
	"sort"
)

func GetSortedKeys(i interface{}) []string {
	v := reflect.ValueOf(i)
	keys := make([]string, v.Len())
	for j, k := range v.MapKeys() {
		keys[j] = k.String()
	}

	sort.Strings(keys)
	return keys
}
