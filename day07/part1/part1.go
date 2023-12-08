package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day07/shared"
	"sort"
	"strings"
)

func Part1(data []string) int {
	d := []shared.Hand{}
	sum := 0
	for _, v := range data {
		s := strings.Split(v, " ")
		h := shared.Hand{
			Cards: []shared.Card{
				shared.Card(s[0][0]),
				shared.Card(s[0][1]),
				shared.Card(s[0][2]),
				shared.Card(s[0][3]),
				shared.Card(s[0][4]),
			},
			Bid: c.GetInt(s[1]),
		}
		h.SetType(false)
		d = append(d, h)
	}

	sort.Slice(d, func(i, j int) bool {
		if d[i].Type == d[j].Type {
			for k := range d[i].Cards {
				if d[i].Cards[k].Value() != d[j].Cards[k].Value() {
					return d[i].Cards[k].Value() > d[j].Cards[k].Value()
				}
			}
			return d[i].Bid > d[j].Bid
		}
		return d[i].Type > d[j].Type
	})
	for i, v := range d {
		//c.Print(i, v)
		sum += v.Bid * (len(d) - i)
	}
	return sum
}
