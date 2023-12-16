package part2

import (
	c "github.com/jerpa/adventofcode2023/common"
	d "github.com/jerpa/adventofcode2023/day10/shared"
	"github.com/jerpa/adventofcode2023/day16/shared"
)

func Part2(data []string) int {
	best := 0

	for y := range data {
		energized := shared.Move(data, shared.TrailData{Dir: d.Right, Point: c.Point{X: -1, Y: y}}, map[shared.TrailData]bool{}, map[c.Point]bool{})
		if len(energized) > best {
			best = len(energized)
		}
		energized = shared.Move(data, shared.TrailData{Dir: d.Left, Point: c.Point{X: len(data[0]), Y: y}}, map[shared.TrailData]bool{}, map[c.Point]bool{})
		if len(energized) > best {
			best = len(energized)
		}
	}
	for x := range data[0] {
		energized := shared.Move(data, shared.TrailData{Dir: d.Down, Point: c.Point{X: x, Y: -1}}, map[shared.TrailData]bool{}, map[c.Point]bool{})
		if len(energized) > best {
			best = len(energized)
		}
		energized = shared.Move(data, shared.TrailData{Dir: d.Up, Point: c.Point{X: x, Y: len(data)}}, map[shared.TrailData]bool{}, map[c.Point]bool{})
		if len(energized) > best {
			best = len(energized)
		}
	}

	return best
}
