package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	endOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		endOfWord: false,
	}
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
		root: NewTrieNode(),
	}
}


func (this *Trie) Insert(word string)  {
    current := this.root
	for i := 0; i < len(word); i++ {
		_, occur := current.children[word[i]]
		if !occur {
			current.children[word[i]] = NewTrieNode()
		}
		current = current.children[word[i]]
	}
	current.endOfWord = true
}


func (this *Trie) Search(word string) bool {
    current := this.root
	for i := 0; i < len(word); i++ {
		_, occur := current.children[word[i]]
		if !occur {
			return false
		}
		current = current.children[word[i]]
	}
	return current.endOfWord
}


func (this *Trie) StartsWith(prefix string) bool {
    current := this.root

	for i := 0; i < len(prefix); i++ {
		_, occur := current.children[prefix[i]]
		if !occur {
			return false
		}
		current = current.children[prefix[i]]
	}
	return true
}

func main() {
	trie := Constructor();
	trie.Insert("apple");
	fmt.Println(trie.Search("apple")); 
	fmt.Println(trie)  // return True
	trie.Search("app");     // return False
	trie.StartsWith("app"); // return True
	trie.Insert("app");
	trie.Search("app");     // return True

}