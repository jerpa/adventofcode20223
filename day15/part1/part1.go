package part1

import (
	"github.com/jerpa/adventofcode2023/day15/shared"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	for _, v := range strings.Split(data[0], ",") {
		sum += shared.Hash15OnString(v)
	}

	return sum
}
