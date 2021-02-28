package implementtrieprefixtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	// Trie trie = new Trie();
	// trie.insert("apple");
	// trie.search("apple");   // 返回 true
	// trie.search("app");     // 返回 false
	// trie.startsWith("app"); // 返回 true
	// trie.insert("app");
	// trie.search("app");     // 返回 true
	trie := Constructor()
	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, false, trie.Search("app"))
	assert.Equal(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	assert.Equal(t, true, trie.Search("app"))
}
