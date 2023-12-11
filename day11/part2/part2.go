package part2

import (
	"github.com/jerpa/adventofcode2023/day11/shared"
)

func Part2(data []string) int {
	return shared.CalculateDistances(data, 999999)
}
