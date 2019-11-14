package WordDictionary

type WordDictionary struct {
	words map[int][]string
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	words := make(map[int][]string)
	return WordDictionary{words: words}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	this.words[len(word)] = append(this.words[len(word)], word)
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if _, ok := this.words[len(word)]; !ok {
		return false
	}
	for _, w := range this.words[len(word)] {
		i, judge := 0, true
		for i < len(w) {
			if word[i] != '.' && w[i] != word[i] {
				judge = false
				break
			}
			i++
		}
		if judge {
			return true
		}
	}
	return false
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */