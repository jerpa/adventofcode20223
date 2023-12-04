package part1

import (
	"regexp"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	reg := regexp.MustCompile(`\d+`)
	for _, v := range data {
		r := strings.Split(v, ":")
		d := strings.Split(r[1], "|")
		w := reg.FindAllString(d[0], -1)
		m := reg.FindAllString(d[1], -1)
		wn := map[string]bool{}
		for _, x := range w {
			wn[x] = true
		}
		p := 0
		for _, x := range m {
			if _, ok := wn[x]; ok {
				if p == 0 {
					p = 1
				} else {
					p *= 2
				}
			}
		}
		sum += p
	}

	return sum
}
