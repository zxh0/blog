package trie

//type IterationCallback = func(key string, val int)

type SearchTree interface {
	Put(key string, val uint)
	Get(key string) uint
	ForEach(cb func(key string, val uint))
}
