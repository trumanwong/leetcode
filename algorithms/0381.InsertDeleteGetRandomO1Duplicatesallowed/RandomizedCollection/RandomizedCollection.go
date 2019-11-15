package RandomizedCollection

import "math/rand"

type RandomizedCollection struct {
	numsMap map[int]int
	nums []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	nums := make([]int, 0)
	numsMap := make(map[int]int)
	return RandomizedCollection{nums: nums, numsMap: numsMap}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	ret := true
	this.nums = append(this.nums, val)
	if _, ok := this.numsMap[val]; ok {
		ret = false
	}
	this.numsMap[val]++
	return ret
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	if _, ok := this.numsMap[val]; ok {
		this.numsMap[val]--
		if this.numsMap[val] == 0 {
			delete(this.numsMap, val)
		}
		i := 0
		for i < len(this.nums) && this.nums[i] != val {
			i++
		}
		this.nums = append(this.nums[0:i], this.nums[i + 1:]...)
		return true
	}
	return false
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	if len(this.nums) == 0 {
		return -1
	}
	index := rand.Int() % len(this.nums)
	return this.nums[index]
}


/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */