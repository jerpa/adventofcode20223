package part2

import (
	"github.com/jerpa/adventofcode2023/day10/shared"
	"strings"
)

func Part2(data []string) int {
	y := 0
	x := 0
	m := map[int]map[int]shared.Direction{}
	for y = range data {
		x = strings.Index(data[y], "S")
		if x != -1 {
			break
		}
	}
	var dir shared.Direction
	s, err := shared.GetNeighbour(x, y, data, shared.Up)
	if err == nil && (s == "|" || s == "7" || s == "F") {
		dir = shared.Up
	}
	if dir == "" {
		s, err := shared.GetNeighbour(x, y, data, shared.Right)
		if err == nil && (s == "-" || s == "7" || s == "J") {
			dir = shared.Right
		}
	}
	if dir == "" {
		s, err := shared.GetNeighbour(x, y, data, shared.Down)
		if err == nil && (s == "|" || s == "L" || s == "J") {
			dir = shared.Down
		}
	}
	if dir == "" {
		s, err := shared.GetNeighbour(x, y, data, shared.Left)
		if err == nil && (s == "-" || s == "F" || s == "L") {
			dir = shared.Left
		}
	}

	startX := x
	startY := y

	lowestX := x
	if _, ok := m[y]; !ok {
		m[y] = map[int]shared.Direction{}
	}
	m[y][x] = dir
	for {
		if dir == shared.Up {
			y--
		} else if dir == shared.Right {
			x++
		} else if dir == shared.Down {
			y++
		} else if dir == shared.Left {
			x--
		}

		if x < lowestX {
			lowestX = x
		}
		if x == startX && y == startY {
			break
		}
		n := data[y][x]
		if _, ok := m[y]; !ok {
			m[y] = map[int]shared.Direction{}
		}
		if dir == shared.Left && n == 'L' {
			dir = shared.Up
		} else if dir == shared.Left && n == 'F' {
			dir = shared.Down
		} else if dir == shared.Right && n == '7' {
			dir = shared.Down
		} else if dir == shared.Right && n == 'J' {
			dir = shared.Up
		} else if dir == shared.Up && n == '7' {
			dir = shared.Left
		} else if dir == shared.Up && n == 'F' {
			dir = shared.Right
		} else if dir == shared.Down && n == 'L' {
			dir = shared.Right
		} else if dir == shared.Down && n == 'J' {
			dir = shared.Left
		}
		m[y][x] = dir

	}
	// Idea: Find the upper left corner of the loop and move clockwise. Then we always have enclosed tiles on the right hand side when we move upwards.
	sum := 0
	x = lowestX
	y = len(data)
	for k, v := range m {
		for r := range v {
			if r == x && k < y {
				y = k
			}
		}
	}
	startX = x
	startY = y
	if m[y][x] != shared.Right {
		// The map isn't in clockwise order, mirror it
		for k, v := range m {
			for r := range v {
				m[k][-r-10] = m[k][r]
				if m[k][r] == shared.Right {
					m[k][-r-10] = shared.Left
				} else if m[k][r] == shared.Left {
					m[k][-r-10] = shared.Right
				}
			}
		}
		x = -x - 10
		startX = x
	}
	last := shared.Direction("")
	for {
		if m[y][x] == shared.Up || (m[y][x] == shared.Left && last == shared.Up) {
			for xx := x + 1; xx < len(data[y]); xx++ {
				if _, ok := m[y][xx]; !ok {
					sum++
				} else {
					break
				}
			}
		}
		if m[y][x] == shared.Up {
			last = shared.Up
			y--
		} else if m[y][x] == shared.Right {
			last = shared.Right
			x++
		} else if m[y][x] == shared.Down {
			last = shared.Down
			y++
		} else {
			last = shared.Left
			x--
		}
		if x == startX && y == startY {
			break
		}
	}

	return sum
	// GAAH
}
