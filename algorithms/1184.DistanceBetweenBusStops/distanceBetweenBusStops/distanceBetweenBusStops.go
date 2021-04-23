package distanceBetweenBusStops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	sum := 0
	for i := start; i < destination; i++ {
		sum += distance[i]
	}

	ret, sum := sum, 0
	i := destination
	for i != start {
		sum += distance[i]
		i++
		if i == len(distance) {
			i = 0
		}
	}
	if sum != 0 && sum < ret {
		ret = sum
	}
	return ret
}
