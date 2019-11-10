package isGoodArray

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

// 给你一个正整数数组 nums，你需要从中任选一些子集，然后将子集中每一个数乘以一个 任意整数，并求出他们的和。
//
// 假如该和结果为 1，那么原数组就是一个「好数组」，则返回 True；否则请返回 False。
func IsGoodArray(nums []int) bool {
	res := 0
	for _, v := range nums {
		res = gcd(res, v)
	}
	return res == 1
}