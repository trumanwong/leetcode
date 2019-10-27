package maskPII

func MaskPII(S string) string {
	s := []byte(S)
	isEmail := false
	if (s[0] >= 'a' && s[0] <= 'z') || (s[0] >= 'A' && s[0] <= 'Z') {
		isEmail = true
	}

	ret := []byte{}
	if isEmail {
		start := -1
		for i := 0; i < len(s); i++ {
			if s[i] == '@' {
				start = i
			}

			if s[i] >= 'A' && s[i] <= 'Z' {
				s[i] += 32
			}
		}
		ret = append(ret, s[0])
		ret = append(ret, []byte("*****")...)
		if start >= 2 {
			ret = append(ret, s[start - 1])
			ret = append(ret, s[start:len(s)]...)
		}
	} else {
		num := 0
		nums := []byte{}
		for _, v := range s {
			if v >= '0' && v <= '9' {
				num++
				nums = append(nums, v)
			}
		}

		if num == 10 {
			ret = append(ret, []byte("***-***-")...)
		} else {
			temp := ""
			for i := 0; i < num - 10; i++ {
				temp += "*"
			}
			ret = append(ret, []byte("+" + temp + "-***-***-")...)
		}
		ret = append(ret, nums[len(nums) - 4:len(nums)]...)
	}

	return string(ret)
}