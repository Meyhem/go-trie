package trie

import "fmt"

type Node struct {
	char     rune
	children map[rune]*Node
	level    uint

	terminal bool
	data     interface{}
}

func (n Node) String() string {
	return fmt.Sprintf("'%s' - level: %d", string(n.char), n.level)
}

func NewNode(char rune, level uint) *Node {
	return &Node{
		char:     char,
		level:    level,
		children: make(map[rune]*Node),
	}
}
