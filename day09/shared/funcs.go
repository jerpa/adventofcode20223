package shared

func AllZeroes(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
