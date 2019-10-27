package corpFlightBookings

func CorpFlightBookings(bookings [][]int, n int) []int {
	ret := make([]int, n + 1)
	for _, v := range bookings {
		ret[v[0] - 1] += v[2]
		ret[v[1]] -= v[2]
	}
	for i := 1; i < n; i++ {
		ret[i] += ret[i - 1]
	}
	return ret[:n]
}