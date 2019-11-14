package Trie

import "strings"

type Trie struct {
	words map[int][]string
}


/** Initialize your data structure here. */
func Constructor() Trie {
	words := make(map[int][]string)
	return Trie{words: words}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	this.words[len(word)] = append(this.words[len(word)], word)
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if _, ok := this.words[len(word)]; !ok {
		return false
	}
	for _, w := range this.words[len(word)] {
		if w == word {
			return true
		}
	}
	return false
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for k, ws := range this.words {
		if k < len(prefix) {
			continue
		}
		for _, w := range ws {
			if strings.HasPrefix(w, prefix) {
				return true
			}
		}
	}
	return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */