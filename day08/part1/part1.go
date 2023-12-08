package part1

import (
	"github.com/jerpa/adventofcode2023/day08/shared"
	"strings"
)

func Part1(data []string) int {
	m := shared.ParseData(data[2:])

	sum := 0
	curr := "AAA"
	for {
		for _, pos := range strings.TrimSpace(data[0]) {
			if pos == 'L' {
				curr = m[curr].Left
			} else {
				curr = m[curr].Right
			}
			sum++

			if curr == "ZZZ" {
				return sum
			}
		}
	}

}
