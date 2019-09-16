package MyHashSet

type MyHashSet struct {
	HashSet []int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
	hashSet := make([]int, 1000000)

	for i := 0; i < len(hashSet); i++ {
		hashSet[i] = -1
	}
	return MyHashSet{HashSet: hashSet}
}


func (this *MyHashSet) Add(key int)  {
	this.HashSet[key] = 1
}


func (this *MyHashSet) Remove(key int)  {
	this.HashSet[key] = -1
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.HashSet[key] == -1 {
		return false
	}
	return true
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */