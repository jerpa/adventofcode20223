package part2

import (
	c "github.com/jerpa/adventofcode2023/common"
	"strings"
)

func Part2(data []string) int {
	sum := 0
	for _, v := range data {
		data := map[string]int{}

		for _, x := range strings.Split(strings.Split(v, ":")[1], ";") {
			x = strings.TrimSpace(x)
			for _, y := range strings.Split(x, ",") {
				y = strings.TrimSpace(y)
				h := strings.Split(y, " ")

				m := c.GetInt(h[0])
				if _, ok := data[h[1]]; !ok || m > data[h[1]] {
					data[h[1]] = m
				}
			}
		}
		sum += data["blue"] * data["red"] * data["green"]
	}

	return sum
}
