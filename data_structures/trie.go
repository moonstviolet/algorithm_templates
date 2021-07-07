package data_structures

type TrieNode struct {
	exist bool
	next  [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{},
	}
}

func (t *Trie) Insert(s string) {
	ptr := t.root
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if ptr.next[idx] == nil {
			ptr.next[idx] = &TrieNode{}
		}
		ptr = ptr.next[idx]
	}
	ptr.exist = true
}

func (t *Trie) Search(s string) bool {
	ptr := t.root
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if ptr.next[idx] == nil {
			return false
		}
		ptr = ptr.next[idx]
	}
	return ptr.exist
}
