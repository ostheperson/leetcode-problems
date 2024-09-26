package problems

type TrieNode struct {
	children map[rune]*TrieNode
	count    int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		count:    0,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
		node.count++
	}
}

func (t *Trie) SumOfPrefixScores(word string) int {
	sum := 0
	node := t.root
	for _, char := range word {
		node = node.children[char]
		sum += node.count
	}
	return sum
}

func SumPrefixScores(words []string) []int {
	sum := make([]int, len(words))
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	for i, word := range words {
		sum[i] = trie.SumOfPrefixScores(word)
	}
	return sum
}
