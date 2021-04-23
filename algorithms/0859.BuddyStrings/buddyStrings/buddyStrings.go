package buddyStrings

func BuddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) <= 1 {
		return false
	}
	if A == B {
		count := make([]int, 26)
		for _, v := range A {
			count[int(v)-'a']++
		}
		for _, v := range count {
			if v > 1 {
				return true
			}
		}
		return false
	} else {
		first, second := -1, -1
		for i, v := range A {
			if byte(v) != B[i] {
				if first == -1 {
					first = i
				} else if second == -1 {
					second = i
				} else {
					return false
				}
			}
		}
		return (second != -1 && A[first] == B[second]) && A[second] == B[first]
	}
}
