package RandomizedSet

import "math/rand"

type RandomizedSet struct {
	numsMap map[int]int
	nums []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	nums := make([]int, 0)
	numsMap := make(map[int]int)
	return RandomizedSet{nums: nums, numsMap: numsMap}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numsMap[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.numsMap[val] = 1
	return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numsMap[val]; !ok {
		return false
	}
	delete(this.numsMap, val)
	i := 0
	for this.nums[i] != val {
		i++
	}
	this.nums = append(this.nums[0:i], this.nums[i + 1:]...)
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 0 {
		return -1
	}
	index := rand.Int() % len(this.nums)
	return this.nums[index]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */