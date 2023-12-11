package shared

import (
	c "github.com/jerpa/adventofcode2023/common"
	"slices"
)

func CalculateDistances(data []string, addForEmpty int) int {
	emptyCols := map[int]bool{}
	emptyRows := map[int]bool{}
	for y := range data {
		empty := true
		for x := range data[y] {
			if data[y][x] != '.' {
				empty = false
				break
			}
		}
		if empty {
			emptyRows[y] = true
		}
	}
	for x := range data[0] {
		empty := true
		for y := range data {
			if data[y][x] != '.' {
				empty = false
				break
			}
		}
		if empty {
			emptyCols[x] = true
		}
	}
	galaxies := []c.Point{}
	for y := range data {
		for x := range data[y] {
			if data[y][x] == '#' {
				galaxies = append(galaxies, c.Point{X: x, Y: y})
			}
		}
	}
	sum := 0
	for i, s := range galaxies {
		for _, g := range galaxies[i+1:] {
			d := s.ManhattanFrom(g)
			for y := slices.Min([]int{s.Y, g.Y}) + 1; y < slices.Max([]int{s.Y, g.Y}); y++ {
				if _, ok := emptyRows[y]; ok {
					d += addForEmpty
				}
			}
			for x := slices.Min([]int{s.X, g.X}) + 1; x < slices.Max([]int{s.X, g.X}); x++ {
				if _, ok := emptyCols[x]; ok {
					d += addForEmpty
				}
			}
			sum += d
		}

	}

	return sum
}
