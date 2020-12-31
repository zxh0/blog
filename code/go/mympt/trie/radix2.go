package trie

import (
	"mympt/internal/utils"
)

// radix=2^1
type PatriciaTrie struct {
	Root *RadixTrieNode
}

func NewPatriciaTrie() SearchTree {
	return &PatriciaTrie{Root: &RadixTrieNode{}}
}

func (node *PatriciaTrie) Put(key string, val uint) {
	node.Root.Put(utils.ToBin(key), val)
}
func (node *PatriciaTrie) Get(key string) uint {
	return node.Root.Get(utils.ToBin(key))
}
func (node *PatriciaTrie) ForEach(cb func(string, uint)) {
	node.Root.ForEach(func(k string, v uint) {
		cb(utils.FromBin(k), v)
	})
}
