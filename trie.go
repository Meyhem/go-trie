package trie

import (
	"fmt"
)

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: NewNode(rune(0), 0)}
}

func (tree *Trie) Insert(key string, value interface{}) {
	node := tree.root
	runes := []rune(key)
	level := uint(1)

	for _, r := range runes {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			new := NewNode(r, level)
			node.children[r] = new
			node = new
		}
		level++
	}
	node.data = value
	node.terminal = true
}

func (tree *Trie) FindByKey(key string) (value interface{}, found bool) {
	node, found := tree.findNodeByMask(key)

	if found && node.terminal {
		return node.data, true
	}

	return nil, false
}

func (tree *Trie) ContainsKey(key string) bool {
	_, found := tree.FindByKey(key)
	return found
}

func (tree *Trie) findNodeByMask(mask string) (n *Node, found bool) {
	node := tree.root
	runes := []rune(mask)

	for _, r := range runes {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			return nil, false
		}
	}

	return node, true
}

func (tree *Trie) dump(node *Node) {
	if node == nil {
		node = tree.root
	}

	fmt.Println(node)
	for _, v := range node.children {
		tree.dump(v)
	}
}
