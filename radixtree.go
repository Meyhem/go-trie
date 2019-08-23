package trie

import (
	"fmt"
)

type RadixTree struct {
	root *Node
}

func NewRadixTree() *RadixTree {
	return &RadixTree{root: NewNode(rune(0))}
}

func (tree *RadixTree) Insert(key string, value interface{}) {
	node := tree.root
	runes := []rune(key)

	for _, r := range runes {
		fmt.Println("kek", node)
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			new := NewNode(r)
			node.children[r] = new
			node = new
		}
	}
	node.data = value
}

func (tree *RadixTree) dump(node *Node) {
	if node == nil {
		node = tree.root
	}

	fmt.Println(node)
	for _, v := range node.children {
		tree.dump(v)
	}
}

func (tree *RadixTree) findNode(key string) {

}
