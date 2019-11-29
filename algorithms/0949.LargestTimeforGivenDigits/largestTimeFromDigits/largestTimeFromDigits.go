package largestTimeFromDigits

import (
	"math"
	"strconv"
)

func LargestTimeFromDigits(A []int) string {
	time := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}

			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				l := 6 - i - j - k
				hours := 10 * A[i] + A[j]
				mins := 10 *A[k] + A[l]
				if hours < 24 && mins < 60 {
					time = int(math.Max(float64(time), float64(hours * 60 + mins)))
				}
			}
		}
	}

	res := ""
	if time >= 0 {
		if time /60 >= 10 {
			res += strconv.Itoa(time / 60) + ":"
		} else {
			res += "0" + strconv.Itoa(time / 60) + ":"
		}
		if time%60 >= 10 {
			res += strconv.Itoa(time % 60)
		} else {
			res += "0" + strconv.Itoa(time % 60)
		}
	}
	return res
}