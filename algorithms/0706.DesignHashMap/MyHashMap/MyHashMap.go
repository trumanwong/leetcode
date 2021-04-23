package MyHashMap

type MyHashMap struct {
	HashMap []int
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	hashMap := make([]int, 1000000)
	for i := 0; i < len(hashMap); i++ {
		hashMap[i] = -1
	}
	return MyHashMap{HashMap: hashMap}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.HashMap[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.HashMap[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.HashMap[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
