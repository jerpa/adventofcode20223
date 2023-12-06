package part2

import (
	"github.com/jerpa/adventofcode2023/day06/shared"
	"strings"
)

func Part2(data []string) int {
	for i := range data {
		data[i] = strings.ReplaceAll(data[i], " ", "")
	}
	d := shared.GetTimesAndDistances(data)

	sum := 1
	for _, v := range d {
		ok := 0
		for t := 1; t < v.Time; t++ {
			if (v.Time-t)*t > v.Distance {
				ok++
			}
		}
		sum *= ok
	}

	return sum
}
