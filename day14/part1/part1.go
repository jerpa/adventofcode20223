package part1

import "github.com/jerpa/adventofcode2023/day14/shared"

func Part1(data []string) int {
	data = shared.Tilt(data)
	return shared.CountWeight(data)
}
