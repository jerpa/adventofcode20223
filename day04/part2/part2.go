package part2

import (
	"regexp"
	"strings"
)

func Part2(data []string) int {
	sum := 0
	reg := regexp.MustCompile(`\d+`)
	cc := map[int]int{}
	for i, v := range data {
		cc[i]++
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
				p++
			}
		}
		for x := 1; x <= p; x++ {
			cc[i+x] += cc[i]
		}
	}
	for i := range data {
		sum += cc[i]
	}

	return sum
}
