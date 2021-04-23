package SumRange

type NumArray struct {
	Result []int
}

func Constructor(nums []int) NumArray {
	sum, result := 0, make([]int, len(nums))
	for i, v := range nums {
		result[i] = sum + v
		sum += v
	}
	numArray := NumArray{Result: result}
	return numArray
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.Result[j]
	} else {
		return this.Result[j] - this.Result[i-1]
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
