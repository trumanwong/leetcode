package distributeCandies

func DistributeCandies(candies int, num_people int) []int {
	ret := make([]int, num_people)
	index, step := 0, 1
	for candies > 0 {
		if index >= num_people {
			index = 0
		}

		if candies - step <= 0 {
			ret[index] += candies
		} else {
			ret[index] += step
		}
		candies -= step
		index++
		step++
	}
	return ret
}