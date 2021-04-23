package findPeakElement

func FindPeakElement(nums []int) int {
	// 由于题目假设nums[-1]=nums[n]=-∞。
	// 所以，我们从0开始往后遍历元素，如果某个元素大于其后面的元素，则该元素就是峰值元素。（但是它时O(n)，不符合题意）
	//
	//O(logN)一般考虑二分搜索。有如下规律：
	//
	//规律一：如果nums[i] > nums[i+1]，则在i之前一定存在峰值元素
	//
	//规律二：如果nums[i] < nums[i+1]，则在i+1之后一定存在峰值元素
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
