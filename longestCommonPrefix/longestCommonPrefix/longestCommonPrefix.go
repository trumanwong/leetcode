package longestCommonPrefix

func LongestCommonPrefix(strs []string) string {
	//dif 表示比较过程中出现字符不相同赋值为0  index代表此下标之前的字符都相同
	dif := 0
	ret := ""
	var i, j int
	strSize := len(strs)
	if strSize == 0 {
		return ""
	} else if strSize == 1 {
		// 只有一个字符串，直接原字符串返回
		return strs[0]
	}
	//获取字符串中长度最小的例数
	columnMin := len(strs[0])
	for i = 1; i < strSize; i++ {
		if len(strs[i]) < columnMin {
			columnMin = len(strs[i])
		}
	}
	if columnMin == 0 {
		return ""
	}
	for i = 0; i < columnMin; i++ {
		// 每次都得重新初始化
		dif = 1
		for j = 0; j < strSize - 1; j++ {
			if strs[j][i] == strs[j + 1][i] {
				continue
			} else {
				dif = 0
				if i == 0 {
					return ""
				}
				break;
			}
		}
		if j == strSize - 1 || dif == 1 {
			ret = ret + string(strs[0][i])
		}
	}
	return ret
}
