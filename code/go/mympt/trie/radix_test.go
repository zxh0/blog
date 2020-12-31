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

func TestRadix16Trie(t *testing.T) {
	trie := NewRadix16Trie()
	trie.Put("foo", 123) // 666f6f
	trie.Put("bar", 456) // 626172
	trie.Put("baz", 789) // 62617a
	require.Equal(t, uint(123), trie.Get("foo"))
	require.Equal(t, uint(456), trie.Get("bar"))
	require.Equal(t, uint(789), trie.Get("baz"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}

func TestPatriciaTrie(t *testing.T) {
	trie := NewPatriciaTrie()
	trie.Put("foo", 123) // 011001100110111101101111
	trie.Put("bar", 456) // 011000100110000101110010
	trie.Put("baz", 789) // 011000100110000101111010
	require.Equal(t, uint(123), trie.Get("foo"))
	require.Equal(t, uint(456), trie.Get("bar"))
	require.Equal(t, uint(789), trie.Get("baz"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}
