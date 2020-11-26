package tries

/*

Trie is an efficient information reTrieval data structure.

Using Trie, search complexities can be brought to optimal limit (key length).

If we store keys in binary search tree, a well balanced BST will need time proportional to M * log N, where M is maximum string length and N is number of keys in tree.

Using Trie, we can search the key in O(M) time. However the penalty is on Trie storage requirements.

Every node of Trie consists of multiple branches.

Each branch represents a possible character of keys.

We need to mark the last node of every key as end of word node.

A Trie node field isEndOfWord is used to distinguish the node as end of word node.

A simple structure to represent nodes of English alphabet can be as following,

*/

const (
	// AlphabetSize num of alphabet
	AlphabetSize = 24
)

// byteToInt rune 2 int
func byteToInt(ch rune) int {
	return int(ch) - int('a')
}

// IntToByte int 2 string
func IntToByte(ch int) string {
	return string(ch + int('a'))
}

// TrieNode Node of Trie
type TrieNode struct {
	children    []*TrieNode
	isEndOfWord bool
}

// newTrieNode create newTrieNode
func newTrieNode(isEndOfWord bool) *TrieNode {
	return &TrieNode{
		children:    make([]*TrieNode, AlphabetSize, AlphabetSize),
		isEndOfWord: isEndOfWord,
	}
}

// Trie tree
type Trie struct {
	Root *TrieNode
}

// NewTrie created NewTrie
func NewTrie() *Trie {
	return &Trie{
		Root: newTrieNode(false),
	}
}

// TrieInsertKey node to Trie
func (t *Trie) TrieInsertKey(key string) {
	root := t.Root
	for _, value := range key {
		index := byteToInt(value)
		if root.children[index] == nil {
			root.children[index] = newTrieNode(false)
		}
		root = root.children[index]
	}
	root.isEndOfWord = true
}

// TrieSearchKey key in Trie
func (t *Trie) TrieSearchKey(key string) bool {
	root := t.Root
	for _, value := range key {
		index := byteToInt(value)
		if index >= AlphabetSize || root.children[index] == nil {
			return false
		}
		root = root.children[index]
	}

	return root != nil && root.isEndOfWord
}

// TrieDeleteKey delete key
func (t *Trie) TrieDeleteKey(key string) {
	root := t.Root
	for _, value := range key {
		index := byteToInt(value)
		if index >= AlphabetSize || root.children[index] == nil {
			return
		}
		root = root.children[index]
	}
	root.isEndOfWord = false
}
