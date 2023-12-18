package part1

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day10/shared"
)

func Part1(data []string) int {
	graph := dijkstra.NewGraph()
	for y := range data {
		for x := range data[y] {
			for i := 1; i <= 3; i++ {
				if x+i < len(data[y]) {
					graph.AddMappedVertex(Vertex(shared.Right, c.Point{X: x, Y: y}, c.Point{X: x + i, Y: y}))
				}
				if x-i >= 0 {
					graph.AddMappedVertex(Vertex(shared.Left, c.Point{X: x, Y: y}, c.Point{X: x - i, Y: y}))
				}
				if y-i >= 0 {
					graph.AddMappedVertex(Vertex(shared.Up, c.Point{X: x, Y: y}, c.Point{X: x, Y: y - i}))
				}
				if y+i < len(data) {
					graph.AddMappedVertex(Vertex(shared.Down, c.Point{X: x, Y: y}, c.Point{X: x, Y: y + i}))
				}
			}
		}
	}
	for y := range data {
		for x := range data[y] {
			for i := 1; i <= 3; i++ {
				for j := 1; j <= 3; j++ {
					if x+i < len(data[y]) && y-j >= 0 {
						_ = graph.AddMappedArc(Vertex(shared.Down, c.Point{X: x, Y: y - j}, c.Point{X: x, Y: y}), Vertex(shared.Right, c.Point{X: x, Y: y}, c.Point{X: x + i, Y: y}), Cost(data, shared.Right, c.Point{X: x, Y: y}, c.Point{X: x + i, Y: y}))
					}
					if x+i < len(data[y]) && y+j < len(data) {
						_ = graph.AddMappedArc(Vertex(shared.Up, c.Point{X: x, Y: y + j}, c.Point{X: x, Y: y}), Vertex(shared.Right, c.Point{X: x, Y: y}, c.Point{X: x + i, Y: y}), Cost(data, shared.Right, c.Point{X: x, Y: y}, c.Point{X: x + i, Y: y}))
					}
					if x-i >= 0 && y-j >= 0 {
						_ = graph.AddMappedArc(Vertex(shared.Down, c.Point{X: x, Y: y - j}, c.Point{X: x, Y: y}), Vertex(shared.Left, c.Point{X: x, Y: y}, c.Point{X: x - i, Y: y}), Cost(data, shared.Left, c.Point{X: x, Y: y}, c.Point{X: x - i, Y: y}))
					}
					if x-i >= 0 && y+j < len(data) {
						_ = graph.AddMappedArc(Vertex(shared.Up, c.Point{X: x, Y: y + j}, c.Point{X: x, Y: y}), Vertex(shared.Left, c.Point{X: x, Y: y}, c.Point{X: x - i, Y: y}), Cost(data, shared.Left, c.Point{X: x, Y: y}, c.Point{X: x - i, Y: y}))
					}

					if y+i < len(data) && x-j >= 0 {
						_ = graph.AddMappedArc(Vertex(shared.Right, c.Point{X: x - j, Y: y}, c.Point{X: x, Y: y}), Vertex(shared.Down, c.Point{X: x, Y: y}, c.Point{X: x, Y: y + i}), Cost(data, shared.Down, c.Point{X: x, Y: y}, c.Point{X: x, Y: y + i}))
					}
					if y+i < len(data) && x+j < len(data[y]) {
						_ = graph.AddMappedArc(Vertex(shared.Left, c.Point{X: x + j, Y: y}, c.Point{X: x, Y: y}), Vertex(shared.Down, c.Point{X: x, Y: y}, c.Point{X: x, Y: y + i}), Cost(data, shared.Down, c.Point{X: x, Y: y}, c.Point{X: x, Y: y + i}))
					}
					if y-i >= 0 && x-j >= 0 {
						_ = graph.AddMappedArc(Vertex(shared.Right, c.Point{X: x - j, Y: y}, c.Point{X: x, Y: y}), Vertex(shared.Up, c.Point{X: x, Y: y}, c.Point{X: x, Y: y - i}), Cost(data, shared.Up, c.Point{X: x, Y: y}, c.Point{X: x, Y: y - i}))
					}
					if y-i >= 0 && x+j < len(data[y]) {
						_ = graph.AddMappedArc(Vertex(shared.Left, c.Point{X: x + j, Y: y}, c.Point{X: x, Y: y}), Vertex(shared.Up, c.Point{X: x, Y: y}, c.Point{X: x, Y: y - i}), Cost(data, shared.Up, c.Point{X: x, Y: y}, c.Point{X: x, Y: y - i}))
					}
				}
			}
		}
	}
	start := Vertex("", c.Point{X: 0, Y: 0}, c.Point{X: 0, Y: 0})
	end := Vertex("", c.Point{X: len(data[0]) - 1, Y: len(data) - 1}, c.Point{X: len(data[0]) - 1, Y: len(data) - 1})
	s := graph.AddMappedVertex(start)
	e := graph.AddMappedVertex(end)
	for i := 1; i <= 3; i++ {
		_ = graph.AddMappedArc(start, Vertex(shared.Right, c.Point{X: 0, Y: 0}, c.Point{X: i, Y: 0}), Cost(data, shared.Right, c.Point{X: 0, Y: 0}, c.Point{X: i, Y: 0}))
		_ = graph.AddMappedArc(start, Vertex(shared.Down, c.Point{X: 0, Y: 0}, c.Point{X: 0, Y: i}), Cost(data, shared.Down, c.Point{X: 0, Y: 0}, c.Point{X: 0, Y: i}))
		_ = graph.AddMappedArc(Vertex(shared.Right, c.Point{X: len(data[0]) - 1 - i, Y: len(data) - 1}, c.Point{X: len(data[0]) - 1, Y: len(data) - 1}), end, 0) //Cost(data, shared.Right, c.Point{X: len(data[0]) - 1 - i, Y: len(data) - 1}, c.Point{X: len(data[0]) - 1, Y: len(data) - 1}))
		_ = graph.AddMappedArc(Vertex(shared.Down, c.Point{X: len(data[0]) - 1, Y: len(data) - 1 - i}, c.Point{X: len(data[0]) - 1, Y: len(data) - 1}), end, 0)  //Cost(data, shared.Down, c.Point{X: len(data[0]) - 1, Y: len(data) - 1 - i}, c.Point{X: len(data[0]) - 1, Y: len(data) - 1}))
	}

	gr, _ := graph.Shortest(s, e)
	/*for _, v := range gr.Path {
		fmt.Println(graph.GetMapped(v))
	}*/
	return int(gr.Distance)
}
func Vertex(dir shared.Direction, from, to c.Point) string {
	if from.X == to.X && from.Y == to.Y {
		return fmt.Sprintf("%s", from)
	}
	return fmt.Sprintf("%s:%s:%s", dir, from, to)
}
func Cost(data []string, dir shared.Direction, from, to c.Point) int64 {
	sum := int64(0)
	for i := 1; i <= 3; i++ {
		if dir == shared.Right {
			if from.X+i <= to.X {
				sum += c.GetInt64(string(data[from.Y][from.X+i]))
			}
		} else if dir == shared.Left {
			if from.X-i >= to.X {
				sum += c.GetInt64(string(data[from.Y][from.X-i]))
			}
		} else if dir == shared.Up {
			if from.Y-i >= to.Y {
				sum += c.GetInt64(string(data[from.Y-i][from.X]))
			}
		} else if dir == shared.Down {
			if from.Y+i <= to.Y {
				sum += c.GetInt64(string(data[from.Y+i][from.X]))
			}
		}
	}
	return sum
}
