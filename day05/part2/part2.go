package part2

import (
	"fmt"
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day05/shared"
	"regexp"
)

func Part2(data []string) int {
	// Brute force, execution time: 3m15s :)
	soils := [][]shared.SoilRange{}
	reg, _ := regexp.Compile(`\d+`)

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
	best := -1
	s := reg.FindAllString(data[0], -1)
	var p int
	o := 0
	for i := 0; i < len(s); i += 2 {
		f := c.GetInt(s[i])
		e := c.GetInt(s[i+1]) + f
		for x := f; x < e; x++ {
			p = x
			for j := range soils {
				for k := range soils[j] {
					if newSoil, ok := soils[j][k].NewSoil(p); ok {
						p = newSoil
						break
					}
				}
			}
			if best == -1 || p < best {
				best = p
			}
			o++
			if o == 1000000 {
				fmt.Print("*")
				o = 0
			}
		}
	}
	fmt.Println()
	return best
}
