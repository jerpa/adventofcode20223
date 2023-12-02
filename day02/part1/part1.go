package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"strings"
)

func Part1(data []string) int {
	sum := 0
	for _, v := range data {
		p := strings.Split(v, ":")
		id := c.GetInt(strings.Split(p[0], " ")[1])
		ok := true
	outer:
		for _, x := range strings.Split(p[1], ";") {
			x = strings.TrimSpace(x)
			for _, y := range strings.Split(x, ",") {
				y = strings.TrimSpace(y)
				h := strings.Split(y, " ")
				if (h[1] == "blue" && c.GetInt(h[0]) > 14) || (h[1] == "red" && c.GetInt(h[0]) > 12) || (h[1] == "green" && c.GetInt(h[0]) > 13) {
					ok = false
					break outer
				}

			}
		}
		if ok {
			sum += id
		}
	}

	return sum
}
