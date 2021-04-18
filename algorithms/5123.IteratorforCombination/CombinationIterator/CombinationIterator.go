package CombinationIterator

type CombinationIterator struct {
	characters []byte
	indexes []int
	cursor int
	pointer int
	start int
	end int
	combinationLength int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
	indexes := make([]int, 0)
	for i := 0; i < combinationLength; i++ {
		indexes = append(indexes, i)
	}
	return CombinationIterator{
		characters: []byte(characters),
		indexes: indexes,
		start: 0,
		cursor: 1,
		end: combinationLength - 1,
		pointer: combinationLength - 1,
		combinationLength: combinationLength,
	}
}


func (this *CombinationIterator) Next() string {
	res := make([]byte, 0)
	if len(this.indexes) < this.combinationLength {
		return string(res)
	}
	for i := 0; i < len(this.indexes); i++ {
		res = append(res, this.characters[this.indexes[i]])
	}
	this.pointer++
	if this.pointer >= len(this.characters) {
		this.end++
		this.indexes = make([]int, 0)
		this.cursor++
		if this.end < len(this.characters) {
			for i := this.start; i <= this.end; i++ {
				if i != this.start && i < this.cursor {
					continue
				}
				this.indexes = append(this.indexes, i)
			}
		}
		if this.end >= len(this.characters) || len(this.indexes) < this.combinationLength {
			this.start++
			this.end = this.start + this.combinationLength - 1
			this.cursor = this.start + 1
			this.indexes = make([]int, 0)
			if this.end < len(this.characters) {
				for i := this.start; i <= this.end; i++ {
					this.indexes = append(this.indexes, i)
				}
			}
		}

		this.pointer = this.end
	} else if this.pointer < len(this.characters) {
		this.indexes = this.indexes[:len(this.indexes) - 1]
		this.indexes = append(this.indexes, this.pointer)
	} else {
		this.indexes = make([]int, 0)
	}
	return string(res)
}


func (this *CombinationIterator) HasNext() bool {
	if len(this.indexes) < this.combinationLength {
		return false
	}
	return true
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */