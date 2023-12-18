package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day10/shared"
	"github.com/twpayne/go-geom"
	"strings"
)

func Part1(data []string) int {

	x := float64(0)
	y := float64(0)
	coords := []geom.Coord{{0, 0}}
	for _, v := range data {
		s := strings.Split(v, " ")
		dir := shared.Direction(s[0])
		steps := c.GetInt(s[1])
		for i := 0; i < steps; i++ {
			switch dir {
			case shared.Up:
				y--
			case shared.Down:
				y++
			case shared.Left:
				x--
			case shared.Right:
				x++
			}
		}
		coords = append(coords, geom.Coord{x, y})
	}
	p := geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{coords})
	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return int(p.Area() - (p.Length() / 2) + 1 + p.Length())
}
