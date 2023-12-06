package part1

import (
	"github.com/jerpa/adventofcode2023/day06/shared"
)

func Part1(data []string) int {
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
