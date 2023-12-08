package part2

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day08/shared"
	"strings"
)

func Part2(data []string) int {
	m := shared.ParseData(data[2:])

	curr := []string{}
	for k := range m {
		if k[2] == 'A' {
			curr = append(curr, k)
		}
	}
	lcms := []int{}
	for i := range curr {
		sum := 0
	outer:
		for {
			for _, pos := range strings.TrimSpace(data[0]) {
				if pos == 'L' {
					curr[i] = m[curr[i]].Left
				} else {
					curr[i] = m[curr[i]].Right
				}

				sum++
				if curr[i][2] == 'Z' {
					break outer
				}
			}
		}
		lcms = append(lcms, sum)
	}

	return c.LCM(lcms...)
}
