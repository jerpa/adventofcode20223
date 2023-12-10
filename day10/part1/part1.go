package part1

import (
	"github.com/jerpa/adventofcode2023/day10/shared"
	"strings"
)

func Part1(data []string) int {
	y := 0
	x := 0
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
	if dir == "" {
		s, err := shared.GetNeighbour(x, y, data, shared.Right)
		if err == nil && (s == "-" || s == "7" || s == "J") {
			dir = shared.Right
		}
	}
	startX := x
	startY := y
	steps := 0

	for {
		if dir == shared.Up {
			y--
		} else if dir == shared.Down {
			y++
		} else if dir == shared.Left {
			x--
		} else if dir == shared.Right {
			x++
		}
		steps++
		if x == startX && y == startY {
			break
		}
		n := data[y][x]
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
	}

	sum := steps / 2

	return sum
}
