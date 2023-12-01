package part2

import (
	c "github.com/jerpa/adventofcode2023/common"
	"regexp"
)

func Part2(data []string) int {
	r, err := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, v := range data {
		n := []int{}
		for i := 0; i < len(v); i++ {
			s := r.FindAllString(v[i:], -1)
			if len(s) > 0 {
				n = append(n, getVal(s[0])...)
			}
		}

		sum += n[0]*10 + n[len(n)-1]
	}

	return sum
}
func getVal(s string) []int {
	switch s {
	case "one":
		return []int{1}
	case "two":
		return []int{2}
	case "three":
		return []int{3}
	case "four":
		return []int{4}
	case "five":
		return []int{5}
	case "six":
		return []int{6}
	case "seven":
		return []int{7}
	case "eight":
		return []int{8}
	case "nine":
		return []int{9}
	default:
		return []int{c.GetInt(s)}
	}
	return []int{0}
}
