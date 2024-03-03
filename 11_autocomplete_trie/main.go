package main

import (
	"fmt"
	"os"
)

func main() {
	words := []string{"dog", "deer", "deal"}
	input := "de"
	trie := NewTrie()
	for _, word := range words {
		if err := trie.Insert(word); err != nil {
			fmt.Fprintf(os.Stderr, "Error inserting: %v\n", err)
		}
	}

	for _, completion := range trie.AllWithPrefix(input) {
		fmt.Println("=>", completion)
	}
}
