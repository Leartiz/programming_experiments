package main

import (
	"fmt"
)

// https://leetcode.com/problems/implement-trie-prefix-tree/

type node struct {
	children map[byte]*node // ---> node*[26]
	isEnd    bool
}

type Trie struct {
	root *node
}

func Constructor() Trie {
	return Trie{
		root: &node{
			children: make(map[byte]*node),
			isEnd:    false, //  корневой узел
		},
	}
}

func (t *Trie) Insert(word string) {
	curNode := t.root
	for i := 0; i < len(word); i += 1 {

		tempNode, ok := curNode.children[word[i]]
		if !ok {
			tempNode = &node{
				children: make(map[byte]*node),
				isEnd:    false,
			}

			curNode.children[word[i]] = tempNode
		}

		// переход к следующими узлу
		curNode = tempNode
	}

	curNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curNode := t.root
	for i := 0; i < len(word); i += 1 {
		tempNode, ok := curNode.children[word[i]]
		if !ok {
			return false
		}

		curNode = tempNode
	}

	return curNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curNode := t.root
	for i := 0; i < len(prefix); i += 1 {
		tempNode, ok := curNode.children[prefix[i]]
		if !ok {
			return false
		}

		curNode = tempNode
	}

	return true
}

func main() {
	/*
		Input
		["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
		[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

		Output
		[null, null, true, false, true, null, true]
	*/
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}
