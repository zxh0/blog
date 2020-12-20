package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSortedKeys(t *testing.T) {
	keys := GetSortedKeys(map[string]int{
		"foo": 123,
		"bar": 345,
		"baz": 456,
	})
	require.Equal(t, []string{"bar", "baz", "foo"}, keys)
}
