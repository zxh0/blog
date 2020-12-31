package trie

import (
	"mympt/internal/utils"
)

// radix=2^4
type Radix16Trie struct {
	Root *RadixTrieNode
}

func NewRadix16Trie() SearchTree {
	return &Radix16Trie{Root: &RadixTrieNode{}}
}

func (node *Radix16Trie) Put(key string, val uint) {
	node.Root.Put(utils.ToHex(key), val)
}
func (node *Radix16Trie) Get(key string) uint {
	return node.Root.Get(utils.ToHex(key))
}
func (node *Radix16Trie) ForEach(cb func(string, uint)) {
	node.Root.ForEach(func(k string, v uint) {
		cb(utils.FromHex(k), v)
	})
}
