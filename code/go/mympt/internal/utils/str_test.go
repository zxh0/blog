package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindCommonPrefixLen(t *testing.T) {
	require.Equal(t, 0, CommonPrefixLen("foo", "bar"))
	require.Equal(t, 2, CommonPrefixLen("bar", "baz"))
	require.Equal(t, 3, CommonPrefixLen("app", "apple"))
}

func TestToBin(t *testing.T) {
	require.Equal(t, "011001100110111101101111", ToBin("foo"))
	require.Equal(t, "011001100110111101110010", ToBin("for"))
	require.Equal(t, "foo", FromBin(ToBin("foo")))
}
