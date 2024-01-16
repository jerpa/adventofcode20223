package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
)

func Part1(data []string) int {
	m := map[c.Point]bool{}
outer:
	for y := range data {
		for x := range data[y] {
			if data[y][x] == 'S' {
				data[y] = data[y][:x] + "." + data[y][x+1:]
				m[c.Point{X: x, Y: y}] = true
				break outer
			}
		}
	}
	for i := 0; i < 64; i++ {
		n := map[c.Point]bool{}
		for k := range m {

			if k.X > 0 && data[k.Y][k.X-1] == '.' {
				n[c.Point{X: k.X - 1, Y: k.Y}] = true

			}
			if k.X < len(data[k.Y])-1 && data[k.Y][k.X+1] == '.' {
				n[c.Point{X: k.X + 1, Y: k.Y}] = true

			}
			if k.Y > 0 && data[k.Y-1][k.X] == '.' {
				n[c.Point{X: k.X, Y: k.Y - 1}] = true

			}
			if k.Y < len(data)-1 && data[k.Y+1][k.X] == '.' {
				n[c.Point{X: k.X, Y: k.Y + 1}] = true

			}
		}
		m = n
	}
	return len(m)
}
