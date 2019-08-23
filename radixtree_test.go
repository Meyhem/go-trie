package radixtree

import (
	"testing"
)

func TestRadixTree(t *testing.T) {
	tree := NewRadixTree()
	tree.Insert("kek", "vole")
	tree.Insert("keke", "kokot")
	tree.Insert("keke", "boľšoj")
	tree.dump(nil)
}
