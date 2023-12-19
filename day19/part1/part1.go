package part1

import (
	"github.com/jerpa/adventofcode2023/day19/shared"
)

func Part1(data []string) int {
	workflows := map[string]shared.Workflow{}
	inParts := false
	sum := 0
	for _, v := range data {
		if v == "" {
			inParts = true
			continue
		}
		if !inParts {
			p := shared.ParseWorkflow(v)
			workflows[p.Name] = p
		} else {
			curr := "in"
			p := shared.ParseParts(v)
			for {
				if curr == "R" {
					break
				} else if curr == "A" {
					sum += p["x"] + p["m"] + p["a"] + p["s"]
					break
				}

				w := workflows[curr]
				curr = w.Else
				for _, r := range w.Rules {
					if r.True(p) {
						curr = r.Then
						break
					}
				}

			}
		}
	}

	return sum
}
