package part1

import c "github.com/jerpa/adventofcode2023/common"

func Part1(data []string) int {
	sum := 0
	for _, v := range data {
		sum += c.GetInt(v)
	}

	return sum
}
