package trie

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"mympt/internal/utils"
)

type KV struct {
	Key string `json:"key"`
	Val uint   `json:"val"`
}

func TestRandomTree(t *testing.T) {
	randomTestSearchTree(t, NewTrie)
	randomTestSearchTree(t, NewRadixTrie)
	randomTestSearchTree(t, NewRadix16Trie)
	randomTestSearchTree(t, NewPatriciaTrie)
}

func randomTestSearchTree(t *testing.T, factory func() SearchTree) {
	for i := 0; i < 1000; i++ {
		kv := genRandomKV(10, 20)
		testTreeWithData(t, kv, factory())
	}
}

func testTreeWithData(t *testing.T, kvs []KV, tree SearchTree) {
	for _, kv := range kvs {
		tree.Put(kv.Key, kv.Val)
	}
	for _, kv := range kvs {
		if v := tree.Get(kv.Key); v != kv.Val {
			println("key:", kv.Key, ", expected val:", kv.Val, ", actual val:", v)
			println("kv:", utils.ToPrettyJSON(kvs))
			println("tree:", utils.ToPrettyJSON(tree))
			require.Equal(t, kv.Val, v, kv.Key)
		}
	}
}

func genRandomKV(maxKeyLen, kvCount int) (kvs []KV) {
	keys := make(map[string]uint, kvCount)
	kvs = make([]KV, kvCount)

	rand.Seed(time.Now().Unix())
	for i := 0; i < kvCount; i++ {
		keyLen := 1 + rand.Intn(maxKeyLen)
		key := randomStr(keyLen)
		if keys[key] == 0 {
			keys[key] = uint(i + 1)
		}
		kvs[i] = KV{
			Key: key,
			Val: keys[key],
		}
	}

	return
}

func randomStr(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := make([]byte, n)
	for j := 0; j < n; j++ {
		s[j] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(s)
}
