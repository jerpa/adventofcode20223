package shared

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day10/shared"
)

type TrailData struct {
	Point c.Point
	Dir   shared.Direction
}

func Move(tiles []string, curr TrailData, trail map[TrailData]bool, energized map[c.Point]bool) map[c.Point]bool {
	for {
		trail[curr] = true
		if curr.Dir == shared.Right {
			curr.Point.X++
		} else if curr.Dir == shared.Left {
			curr.Point.X--
		} else if curr.Dir == shared.Up {
			curr.Point.Y--
		} else if curr.Dir == shared.Down {
			curr.Point.Y++
		}
		if curr.Point.X < 0 || curr.Point.Y < 0 || curr.Point.X >= len(tiles[0]) || curr.Point.Y >= len(tiles) {
			return energized
		}
		if curr.Dir == shared.Right {
			if tiles[curr.Point.Y][curr.Point.X] == '/' {
				curr.Dir = shared.Up
			} else if tiles[curr.Point.Y][curr.Point.X] == '\\' {
				curr.Dir = shared.Down
			}
		} else if curr.Dir == shared.Left {
			if tiles[curr.Point.Y][curr.Point.X] == '/' {
				curr.Dir = shared.Down
			} else if tiles[curr.Point.Y][curr.Point.X] == '\\' {
				curr.Dir = shared.Up
			}
		} else if curr.Dir == shared.Up {
			if tiles[curr.Point.Y][curr.Point.X] == '/' {
				curr.Dir = shared.Right
			} else if tiles[curr.Point.Y][curr.Point.X] == '\\' {
				curr.Dir = shared.Left
			}
		} else if curr.Dir == shared.Down {
			if tiles[curr.Point.Y][curr.Point.X] == '/' {
				curr.Dir = shared.Left
			} else if tiles[curr.Point.Y][curr.Point.X] == '\\' {
				curr.Dir = shared.Right
			}
		}
		energized[curr.Point] = true
		if _, ok := trail[curr]; ok {
			return energized
		}
		if (curr.Dir == shared.Right || curr.Dir == shared.Left) && tiles[curr.Point.Y][curr.Point.X] == '|' {
			energized = Move(tiles, TrailData{Dir: shared.Up, Point: c.Point{X: curr.Point.X, Y: curr.Point.Y}}, trail, energized)
			energized = Move(tiles, TrailData{Dir: shared.Down, Point: c.Point{X: curr.Point.X, Y: curr.Point.Y}}, trail, energized)
			return energized
		}
		if (curr.Dir == shared.Up || curr.Dir == shared.Down) && tiles[curr.Point.Y][curr.Point.X] == '-' {
			energized = Move(tiles, TrailData{Dir: shared.Left, Point: c.Point{X: curr.Point.X, Y: curr.Point.Y}}, trail, energized)
			energized = Move(tiles, TrailData{Dir: shared.Right, Point: c.Point{X: curr.Point.X, Y: curr.Point.Y}}, trail, energized)
			return energized
		}
	}
}
