package sortColors

func swap(arr []int, x int, y int) {
	t := arr[x]
	arr[x] = arr[y]
	arr[y] = t
}

// Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.
//
//Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
//
//Note: You are not suppose to use the library's sort function for this problem.
//
//Example:
//
//Input: [2,0,2,1,1,0]
//Output: [0,0,1,1,2,2]
//Follow up:
//
//A rather straight forward solution is a two-pass algorithm using counting sort.
//First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
//Could you come up with a one-pass algorithm using only constant space?
func SortColors(nums []int) {
	L := -1
	R := len(nums)
	current := 0
	mid := 1
	// 将1作为中间值， 小的放前面， 大的放后面
	for current < R {
		if nums[current] < mid {
			L = L + 1
			swap(nums, L, current)
			current = current + 1
		} else if nums[current] > mid {
			R = R - 1
			swap(nums, R, current)
		} else {
			current++
		}
	}
}
