package trie

import (
	"testing"

	"github.com/stretchr/testify/require"

	"mympt/internal/utils"
)

func TestRadixTrie(t *testing.T) {
	trie := NewRadixTrie()
	trie.Put("Java", 1995)
	trie.Put("JavaScript", 1996)
	trie.Put("Groovy", 2007)
	trie.Put("Golang", 2012)
	require.Equal(t, uint(1995), trie.Get("Java"))
	require.Equal(t, uint(1996), trie.Get("JavaScript"))
	require.Equal(t, uint(2007), trie.Get("Groovy"))
	require.Equal(t, uint(2012), trie.Get("Golang"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}
