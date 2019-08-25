package trie

import (
	"fmt"
)

// Trie objects holds trie data structure
type Trie struct {
	root *Node
}

// NewTrie instantiates new Trie
func NewTrie() *Trie {
	return &Trie{root: NewNode(rune(0), nil)}
}

// Insert - inserts specified value under specified key to Trie
// insertion under same key overrides original value
func (tree *Trie) Insert(key string, value interface{}) {
	node := tree.root
	runes := []rune(key)
	level := uint(1)

	for _, r := range runes {
		if child, ok := node.children[r]; ok {
			node = child
		} else {
			new := NewNode(r, node)
			node.children[r] = new
			node = new
		}
		level++
	}
	node.data = value
	node.terminal = true
}

// Find - finds user value under specified key
func (tree *Trie) Find(key string) (value interface{}, found bool) {
	node, found := tree.findNodeByMask(key)

	if found && node.terminal {
		return node.data, true
	}

	return nil, false
}

// ContainsKey - returns true if specified key is present in structure
// ignores intermediate nonterminal node findings. It searches only for terminal values
func (tree *Trie) ContainsKey(key string) bool {
	_, found := tree.Find(key)
	return found
}

// DeleteKey - deletes key and all asociated user data
// returns true if deletion was successful
func (tree *Trie) DeleteKey(key string) bool {
	node, found := tree.findNodeByMask(key)
	if !found || node == nil || !node.terminal {
		return false
	}
	char := node.char
	node = node.parent
	for node.parent != nil && !node.terminal {
		char = node.char
		node = node.parent
	}

	delete(node.children, char)
	return true
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
