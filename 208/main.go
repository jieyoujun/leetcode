package implementtrieprefixtree

type Trie struct {
	next [26]*Trie // 每个字母下面都可能接26个字母
	end  bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, letter := range word {
		letter -= 'a'
		if node.next[letter] == nil {
			node.next[letter] = &Trie{}
		}
		node = node.next[letter]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, letter := range word {
		letter -= 'a'
		if node.next[letter] == nil {
			return false
		}
		node = node.next[letter]
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, letter := range prefix {
		letter -= 'a'
		if node.next[letter] == nil {
			return false
		}
		node = node.next[letter]
	}
	return true
}
