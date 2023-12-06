package main

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day06/part1"
	"github.com/jerpa/adventofcode2023/day06/part2"
	"time"
)

func main() {
	start := time.Now()

	c.Print("Part1: ", part1.Part1(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
	start = time.Now()
	c.Print("Part2: ", part2.Part2(c.GetData(Data)))
	c.Print("Took: ", time.Since(start).String())
}

var Data = `
Time:        42     89     91     89
Distance:   308   1170   1291   1467
`
