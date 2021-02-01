package leetcode208

import "testing"

func TestConstructort(t *testing.T) {
	trie := Constructor()
	// trie.Next = make([]*Trie, 26)
	trie.Insert("apple")

	if !trie.Search("apple") {
		t.Errorf("Example actually got false, but expect got true")
	}

	if trie.Search("app") {
		t.Errorf("Example actually got true, but expect got false")
	}

	if !trie.StartsWith("app") {
		t.Errorf("Example actually got false, but expect got true")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Errorf("Example actually got false, but expect got true")
	}
}
