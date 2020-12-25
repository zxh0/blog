package trie

import (
	"mympt/internal/utils"
)

// radix=2^8
type RadixTrieNode struct {
	Prefix string                    `json:"prefix,omitempty"`
	Val    uint                      `json:"val,omitempty"`
	Kids   map[string]*RadixTrieNode `json:"kids,omitempty"`
}

func NewRadixTrie() SearchTree {
	return &RadixTrieNode{}
}

func (node *RadixTrieNode) Put(key string, val uint) {
	if node.Prefix == "" && node.Val == 0 && node.Kids == nil {
		node.Prefix, node.Val = key, val
		return
	}

	if node.Prefix == key {
		node.Val = val
		return
	}

	n := utils.CommonPrefixLen(node.Prefix, key)
	if len(node.Prefix) > n {
		node.split(n)
	}
	if len(key) > n {
		kid := node.findOrCreateKid(key[n : n+1])
		kid.Put(key[n+1:], val)
	} else {
		node.Val = val
	}
}

func (node *RadixTrieNode) findOrCreateKid(char string) *RadixTrieNode {
	if node.Kids == nil {
		node.Kids = make(map[string]*RadixTrieNode)
	}
	if node.Kids[char] == nil {
		node.Kids[char] = &RadixTrieNode{}
	}
	return node.Kids[char]
}

func (node *RadixTrieNode) split(n int) {
	node.Kids = map[string]*RadixTrieNode{
		node.Prefix[n : n+1]: {
			Prefix: node.Prefix[n+1:],
			Val:    node.Val,
			Kids:   node.Kids,
		},
	}
	node.Prefix = node.Prefix[:n]
	node.Val = 0
}

func (node *RadixTrieNode) Get(key string) uint {
	if key == node.Prefix {
		return node.Val
	}
	if node.Kids != nil {
		n := utils.CommonPrefixLen(node.Prefix, key)
		if n == len(node.Prefix) {
			if kid, found := node.Kids[key[n:n+1]]; found {
				return kid.Get(key[n+1:])
			}
		}
	}
	return 0
}

func (node *RadixTrieNode) ForEach(cb func(string, uint)) {
	if node.Val > 0 {
		cb(node.Prefix, node.Val)
	}

	for _, char := range utils.GetSortedKeys(node.Kids) {
		kid := node.Kids[char]
		kid.ForEach(func(key string, val uint) {
			cb(node.Prefix+char+key, val)
		})
	}
}
