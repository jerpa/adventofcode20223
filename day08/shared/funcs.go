package shared

import "strings"

type Neighbour struct {
	Left  string
	Right string
}

func ParseData(data []string) map[string]Neighbour {
	res := map[string]Neighbour{}
	for _, v := range data {
		parts := strings.Split(v, " = (")
		n := strings.Split(strings.TrimSuffix(parts[1], ")"), ", ")
		res[parts[0]] = Neighbour{
			Left:  n[0],
			Right: n[1],
		}
	}
	return res
}
