package trie

import (
	"mympt/internal/utils"
)

// 256-way
type TrieNode struct {
	Val  uint                 `json:"val,omitempty"`
	Kids map[string]*TrieNode `json:"kids,omitempty"`
}

func NewTrie() SearchTree {
	return &TrieNode{}
}

func (node *TrieNode) Put(key string, val uint) {
	if key == "" {
		node.Val = val
	} else {
		kid := node.findOrCreateKid(key[:1])
		kid.Put(key[1:], val)
	}
}

func (node *TrieNode) findOrCreateKid(char string) *TrieNode {
	if node.Kids == nil {
		node.Kids = make(map[string]*TrieNode)
	}
	if node.Kids[char] == nil {
		node.Kids[char] = &TrieNode{}
	}
	return node.Kids[char]
}

func (node *TrieNode) Get(key string) uint {
	if key == "" {
		return node.Val
	}
	if kid, found := node.Kids[key[:1]]; found {
		return kid.Get(key[1:])
	}
	return 0
}

func (node *TrieNode) ForEach(cb func(string, uint)) {
	if node.Val > 0 {
		cb("", node.Val)
	}
	for _, char := range utils.GetSortedKeys(node.Kids) {
		kid := node.Kids[char]
		kid.ForEach(func(key string, val uint) {
			cb(char+key, val)
		})
	}
}
