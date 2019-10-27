package addNegabinary

func AddNegabinary(arr1 []int, arr2 []int) []int {
	/**
	  1.基数为-2的数，进位是高一位加-1，借位是高一位减-1（即加1）
	    所以每一位加上（进位/借位）的结果有五种可能，-1、0、1、2、3；
	  2.这种数是无符号的，看来左边可以无限补零,右边的数总能借到位
	  3.最后要去掉先导零
	*/
	diff := len(arr1) - len(arr2)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			arr2 = append([]int{0}, arr2...)
		}
	} else {
		for i := 0; i < -diff; i++ {
			arr1 = append([]int{0}, arr1...)
		}
	}

	carry, res := 0, make([]int, len(arr1))
	for i := len(arr1) - 1; i >= 0; i-- {
		temp := arr1[i] + arr2[i] + carry
		if temp == 0 {
			res[i], carry = 0, 0
		} else if temp == 1 {
			res[i], carry = 1, 0
		} else if temp == 2 {
			res[i], carry = 0, -1
		} else if temp == 3 {
			res[i], carry = 1, -1
		} else if temp == -1 {
			res[i], carry = 1, 1
		}
	}

	for carry != 0 {
		if carry == -1 {
			res, carry = append([]int{1}, res...), 1
		} else if carry == 1 {
			res, carry = append([]int{1}, res...), 0
		}
	}
	for len(res) > 1 && res[0] == 0 {
		res = res[1:]
	}
	return res
}