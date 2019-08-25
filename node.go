package trie

// Node is internal data structure which hold one trie node
type Node struct {
	// char, which node represents
	char rune

	// child nodes
	children map[rune]*Node

	// parent node
	parent *Node

	// flag whether this node holds key value
	terminal bool

	// user data
	data interface{}
}

// NewNode create new instance of Node
func NewNode(char rune, parent *Node) *Node {
	return &Node{
		char:     char,
		parent:   parent,
		children: make(map[rune]*Node),
	}
}
