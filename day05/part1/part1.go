package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day05/shared"
	"regexp"
)

func Part1(data []string) int {
	seeds := []int{}
	soils := [][]shared.SoilRange{}
	reg, _ := regexp.Compile(`\d+`)
	s := reg.FindAllString(data[0], -1)
	for _, v := range s {
		seeds = append(seeds, c.GetInt(v))
	}
	pos := -1

	for _, v := range data[1:] {
		if v == "" {
			pos++
			soils = append(soils, []shared.SoilRange{})
			continue
		}
		s := reg.FindAllString(v, -1)
		if len(s) == 0 {
			continue
		}
		soils[pos] = append(soils[pos], shared.SoilRange{
			OutStart: c.GetInt(s[0]),
			InStart:  c.GetInt(s[1]),
			InEnd:    c.GetInt(s[2]) + c.GetInt(s[1]) - 1,
		})
	}
	for i := range seeds {
		for j := range soils {
			for k := range soils[j] {
				if newSoil, ok := soils[j][k].NewSoil(seeds[i]); ok {
					seeds[i] = newSoil
					break
				}
			}
		}
	}
	return c.MinInt(seeds)
}
