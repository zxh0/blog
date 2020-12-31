package trie

import (
	"testing"

	"github.com/stretchr/testify/require"

	"mympt/internal/utils"
)

func TestSearchTrie(t *testing.T) {
	testSearchTrie(t, NewTrie())
	testSearchTrie(t, NewRadixTrie())
	testSearchTrie(t, NewPatriciaTrie())
	testSearchTrie(t, NewRadix16Trie())
}

func testSearchTrie(t *testing.T, tree SearchTree) {
	//tree.Put("", 123)
	tree.Put("C", 1972)
	tree.Put("C++", 1985)
	tree.Put("Python", 1989)
	tree.Put("Ruby", 1995)
	tree.Put("Java", 1995)
	tree.Put("JavaScript", 1995)
	tree.Put("C#", 2000)
	tree.Put("Java5", 2004)
	tree.Put("Scala", 2004)
	tree.Put("Groovy", 2007)
	tree.Put("Go", 2009)
	tree.Put("Rust", 2010)
	tree.Put("Swift", 2014)
	tree.ForEach(func(k string, v uint) {
		println(k, v)
	})
	println(utils.ToPrettyJSON(tree))

	//require.Equal(t, uint(123), tree.Get(""))
	require.Equal(t, uint(1972), tree.Get("C"))
	require.Equal(t, uint(1985), tree.Get("C++"))
	require.Equal(t, uint(1989), tree.Get("Python"))
	require.Equal(t, uint(1995), tree.Get("Ruby"))
	require.Equal(t, uint(1995), tree.Get("Java"))
	require.Equal(t, uint(1995), tree.Get("JavaScript"))
	require.Equal(t, uint(2000), tree.Get("C#"))
	require.Equal(t, uint(2004), tree.Get("Java5"))
	require.Equal(t, uint(2004), tree.Get("Scala"))
	require.Equal(t, uint(2007), tree.Get("Groovy"))
	require.Equal(t, uint(2009), tree.Get("Go"))
	require.Equal(t, uint(2010), tree.Get("Rust"))
	require.Equal(t, uint(2014), tree.Get("Swift"))
	require.Equal(t, uint(0), tree.Get("PHP"))
	require.Equal(t, uint(0), tree.Get("Ru"))
	require.Equal(t, uint(0), tree.Get("Rust2"))
}
