package trie

import (
	"testing"

	"github.com/stretchr/testify/require"

	"mympt/internal/utils"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Put("foo", 123)
	trie.Put("bar", 456)
	trie.Put("baz", 789)
	require.Equal(t, uint(123), trie.Get("foo"))
	require.Equal(t, uint(456), trie.Get("bar"))
	require.Equal(t, uint(789), trie.Get("baz"))
	trie.ForEach(func(key string, val uint) {
		println(key, ":", val)
	})
	println(utils.ToPrettyJSON(trie))
}
