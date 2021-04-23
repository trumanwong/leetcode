package MajorityChecker

type MajorityChecker struct {
	Arr []int
}

func Constructor(arr []int) MajorityChecker {
	checker := MajorityChecker{Arr: arr}
	return checker
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	m := make(map[int]int)
	for i := left; i <= right; i++ {
		m[this.Arr[i]]++
		if m[this.Arr[i]] >= threshold {
			return this.Arr[i]
		}
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
