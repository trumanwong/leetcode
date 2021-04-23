package partitionLabels

func PartitionLabels(S string) []int {
	res := make([]int, 0)
	if len(S) == 0 {
		return res
	}

	cache := make([]int, 26)
	for i := 0; i < len(S); i++ {
		cache[int(S[i]-'a')] = i
	}

	for i := 0; i < len(S); {
		index := cache[int(S[i]-'a')]
		for j := i + 1; j < index && j < len(S); j++ {
			if cache[int(S[j]-'a')] > index {
				index = cache[int(S[j]-'a')]
			}
		}
		res = append(res, index-i+1)
		i = index + 1
	}
	return res
}
