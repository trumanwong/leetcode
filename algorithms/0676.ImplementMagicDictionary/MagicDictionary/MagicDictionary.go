package MagicDictionary

type MagicDictionary struct {
	words map[int][]string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	words := make(map[int][]string)
	return MagicDictionary{words: words}
}

/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string) {
	for _, v := range dict {
		this.words[len(v)] = append(this.words[len(v)], v)
	}
}

/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
	if _, ok := this.words[len(word)]; !ok {
		return false
	}

	diffCount := 0
	for _, w := range this.words[len(word)] {
		diffCount = 0
		for i := 0; i < len(word); i++ {
			if w[i] != word[i] {
				diffCount++
			}
			if diffCount > 1 {
				break
			}
		}
		if diffCount == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */
