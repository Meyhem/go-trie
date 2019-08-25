package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tree := NewTrie()
	tree.Insert("123", "heyhou")
	tree.Insert("1234", "foo")
	tree.Insert("132", "bar")

	val, found := tree.FindByKey("123")

	if !found {
		t.Error("there should be value under key '123'")
	}

	if val != "heyhou" {
		t.Error("tree found wrong value under key '123'")
	}

	val, found = tree.FindByKey("1234")

	if !found {
		t.Error("there should be value under key '1234'")
	}

	if val != "foo" {
		t.Error("tree found wrong value under key '1234'")
	}

	val, found = tree.FindByKey("badkey")

	if found {
		t.Error("there shouldn't be value under key 'badkey'")
	}

	if val != nil {
		t.Error("tree should find nil value under key 'badkey'")
	}

	val, found = tree.FindByKey("12")

	if found {
		t.Error("there shouldn't be value under sub key '12'")
	}

	if val != nil {
		t.Error("tree should find nil value under sub key '12'")
	}

	if !tree.ContainsKey("123") {
		t.Error("tree should contain key '123'")
	}

	if tree.ContainsKey("12") {
		t.Error("tree shouldn't contain key '1234'")
	}
}
