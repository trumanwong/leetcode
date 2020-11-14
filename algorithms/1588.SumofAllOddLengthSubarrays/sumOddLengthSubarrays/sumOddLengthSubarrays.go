package sumOddLengthSubarrays

func SumOddLengthSubarrays(arr []int) int {
	// left求法：对当前元素arr[i]来说，前面有i个元素，想要构成连续数组，可在arr[i]左面连续取0,1,2,...,i个元素，所以共有left=i+1种选择方法
	//right求法：arr[i]后有n-i-1个元素，想要构成连续数组，可在arr[i]右面连续取元素，共有right=n-i种选择方法(算上取0个元素)
	//left_odd, right_odd:为可选择数/2 (左面选奇数个元素右面选奇数个元素加上本身为奇数长度数组)
	//left_even, right_even:为(可选择数+1)/2 (左面选偶数个元素右面选偶数个元素加上本身为奇数长度数组)
	//sum:对每个arr[i]求它出现在奇数长度子数组中的次数，乘本身的值，最后求和
	res := 0
	for i := 0; i < len(arr); i++ {
		left, right := i + 1, len(arr) - i
		leftOdd, rightOdd := left / 2, right / 2
		leftEven, rightEven := (left + 1) / 2, (right + 1) / 2
		res += arr[i] * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return res
}