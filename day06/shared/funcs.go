package shared

import (
	c "github.com/jerpa/adventofcode2023/common"
	"regexp"
)

type TAndD struct {
	Time     int
	Distance int
}

func GetTimesAndDistances(data []string) []TAndD {
	reg, _ := regexp.Compile(`\d+`)
	t := reg.FindAllString(data[0], -1)
	d := reg.FindAllString(data[1], -1)
	res := []TAndD{}
	for i := range t {
		res = append(res, TAndD{c.GetInt(t[i]), c.GetInt(d[i])})
	}

	return res
}
