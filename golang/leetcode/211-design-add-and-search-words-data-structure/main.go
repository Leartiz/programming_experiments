package main

import (
	"fmt"
)

// https://leetcode.com/problems/design-add-and-search-words-data-structure/

// NOTE:
/*
	Design a data structure that supports adding new words and finding
	if a string matches any previously added string.

	Implement the WordDictionary class:
		- WordDictionary() Initializes the object.
		- void addWord(word) Adds word to the data structure,
			it can be matched later.
		- bool search(word) Returns true if there is any string
			in the data structure that matches word or false otherwise.
			word may contain dots '.' where dots can be matched with any letter.

	1 <= word.length <= 25
*/

type node struct {
	children map[byte]*node
	isEnd    bool
}

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &node{
			children: make(map[byte]*node),
			isEnd:    false,
		},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	curNode := wd.root // node <--- wd.root

	for i := 0; i < len(word); i += 1 {
		tempNode, ok := curNode.children[word[i]]
		if !ok {
			tempNode = &node{
				children: make(map[byte]*node),
				isEnd:    false,
			}

			// засунуть новый узел в "дерево"
			curNode.children[word[i]] = tempNode
		}

		// перейти к следующему
		curNode = tempNode
	}

	// последня буква
	curNode.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	return false
}

// -----------------------------------------------------------------------

func main() {
	/*
		Input
		["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
		[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]

		Output
		[null,null,null,null,false,true,true,true]
	*/
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad")) // false
	fmt.Println(wd.Search("bad")) // true
	fmt.Println(wd.Search(".ad")) // true
	fmt.Println(wd.Search("b..")) // true
}
