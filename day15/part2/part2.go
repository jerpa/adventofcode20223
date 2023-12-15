package part2

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day15/shared"
	"strings"
)

type Lens struct {
	Label string
	Value int
}

func Part2(data []string) int {
	sum := 0
	boxes := map[int][]Lens{}
	for _, s := range strings.Split(data[0], ",") {
		if strings.Index(s, "=") != -1 {
			d := strings.Split(s, "=")
			box := shared.Hash15OnString(d[0])
			if _, ok := boxes[box]; !ok {
				boxes[box] = []Lens{}
			}
			lens := Lens{
				Label: d[0],
				Value: c.GetInt(d[1]),
			}
			match := false
			for i := range boxes[box] {
				if boxes[box][i].Label == lens.Label {
					boxes[box][i] = lens
					match = true
					break
				}
			}
			if !match {
				boxes[box] = append(boxes[box], lens)
			}
		} else {
			d := strings.Split(s, "-")
			box := shared.Hash15OnString(d[0])
			if _, ok := boxes[box]; !ok {
				continue
			}

			for i := range boxes[box] {
				if boxes[box][i].Label == d[0] {
					boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					break
				}
			}
		}
	}
	for k, v := range boxes {
		for i, l := range v {
			sum += (k + 1) * (i + 1) * l.Value

		}
	}

	return sum
}
