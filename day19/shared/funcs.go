package shared

import (
	c "github.com/jerpa/adventofcode2023/common"
	"strings"
)

type Rule struct {
	Variable string
	Below    bool
	Value    int
	Then     string
}

func (r Rule) True(p Parts) bool {
	if r.Below {
		return p[r.Variable] < r.Value
	} else {
		return p[r.Variable] > r.Value
	}
}

type Workflow struct {
	Rules []Rule
	Name  string
	Else  string
}

func ParseWorkflow(w string) Workflow {
	w = strings.TrimSuffix(w, "}")
	x := strings.Split(w, "{")
	res := Workflow{
		Name: strings.TrimSpace(x[0]),
	}
	w = strings.TrimPrefix(x[1], "{")
	for _, v := range strings.Split(w, ",") {
		if !strings.Contains(v, ":") {
			res.Else = v
		} else {
			p := Rule{}
			r := strings.Split(v, ":")
			p.Then = r[1]
			if strings.Contains(r[0], "<") {
				p.Below = true
				z := strings.Split(r[0], "<")
				p.Variable = z[0]
				p.Value = c.GetInt(z[1])
			} else {
				p.Below = false
				z := strings.Split(r[0], ">")
				p.Variable = z[0]
				p.Value = c.GetInt(z[1])
			}

			res.Rules = append(res.Rules, p)

		}
	}
	return res
}

type Parts map[string]int

func ParseParts(p string) Parts {
	res := make(Parts)
	p = strings.Trim(p, "{")
	p = strings.Trim(p, "}")
	for _, v := range strings.Split(p, ",") {
		x := strings.Split(v, "=")
		res[x[0]] = c.GetInt(x[1])
	}

	return res
}

func DoNothing() {
}
