package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	d "github.com/jerpa/adventofcode2023/day10/shared"
	"github.com/jerpa/adventofcode2023/day16/shared"
)

func Part1(data []string) int {
	energized := map[c.Point]bool{}

	energized = shared.Move(data, shared.TrailData{Dir: d.Right, Point: c.Point{X: -1, Y: 0}}, map[shared.TrailData]bool{}, energized)

	return len(energized)
}
