package part1

import (
	c "github.com/jerpa/adventofcode2023/common"
	"regexp"
)

func Part1(data []string) int {
	r, err := regexp.Compile(`\d`)
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, v := range data {
		s := r.FindAllString(v, -1)
		sum += c.GetInt(s[0])*10 + c.GetInt(s[len(s)-1])
	}

	return sum
}
