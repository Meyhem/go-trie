package radixtree

type Node struct {
	char     rune
	children map[rune]*Node

	data interface{}
}

func NewNode(char rune) *Node {
	return &Node{
		char:     char,
		children: make(map[rune]*Node),
	}
}
