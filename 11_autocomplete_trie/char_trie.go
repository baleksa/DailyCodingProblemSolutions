package main

type Trie *trieNode

type trieNode struct {
	hasVal   bool
	children map[rune]*trieNode
}

func NewTrie() *trieNode {
	return &trieNode{false, map[rune]*trieNode{}}
}

func (trie *trieNode) Insert(s string) error {
	node := trie
	for _, r := range s {
		if next, ok := node.children[r]; ok == true {
			node = next
		} else {
			node.children[r] = NewTrie()
			node = node.children[r]
		}
	}

	node.hasVal = true
	return nil
}

func (trie *trieNode) Contains(s string) bool {
	node := trie
	for _, r := range s {
		if next, ok := node.children[r]; ok == true {
			node = next
		} else {
			return false
		}
	}
	return node.hasVal
}

func (trie *trieNode) Find(s string) *trieNode {
	node := trie
	for _, r := range s {
		if next, ok := node.children[r]; ok == true {
			node = next
		} else {
			return nil
		}
	}
	return node
}

func (trie *trieNode) AllWithPrefix(s string) []string {
	start := trie.Find(s)
	if start == nil {
		return nil
	}

	words := []string{}

	var traverse func(string, *trieNode)
	traverse = func(str string, tr *trieNode) {
		if tr.hasVal {
			words = append(words, str)
		}
		for r, node := range tr.children {
			traverse(str+string(r), node)
		}
	}

	traverse(s, start)

	return words
}
