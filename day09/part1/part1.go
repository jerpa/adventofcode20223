package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day09/shared"
	"regexp"
)

func Part1(data []string) int {
	sum := 0

	reg, _ := regexp.Compile(`[\-\d]+`)
	for _, v := range data {
		d := reg.FindAllString(v, -1)
		n := []int{}
		for _, x := range d {
			n = append(n, c.GetInt(x))
		}
		l := CalcX(n) + n[len(n)-1]

		sum += l
	}

	return sum
}
func CalcX(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if shared.AllZeroes(data) {
		return 0
	}

	n := []int{}
	for i := 1; i < len(data); i++ {
		n = append(n, data[i]-data[i-1])
	}
	if len(n) == 0 {
		return 0
	}
	x := CalcX(n)

	return x + n[len(n)-1]
}
