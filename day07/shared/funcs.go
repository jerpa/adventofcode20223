package shared

import "sort"

type Hand struct {
	Cards []Card
	Bid   int
	Type  Type
}
type Card rune

type Type int

var order = []Card{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var jokerOrder = []Card{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

const (
	FiveOfAKind  Type = 10
	FourOfAKind  Type = 9
	FullHouse    Type = 8
	ThreeOfAKind Type = 7
	TwoPair      Type = 6
	OnePair      Type = 5
	HighCard     Type = 4
)

func (c Card) Value() int {
	for i, v := range order {
		if v == c {
			return len(order) - i
		}
	}
	return -1
}
func (c Card) ValueWithJoker() int {
	for i, v := range jokerOrder {
		if v == c {
			return len(jokerOrder) - i
		}
	}
	return -1
}
func (c Card) String() string {
	return string(c)
}

func (h *Hand) SetType(withJoker bool) {
	gm := map[Card]int{}
	for _, v := range h.Cards {
		gm[v]++
	}
	jokerValue := 0
	if withJoker {
		jokerValue = gm['J']
		/*for k := range gm {
			if k != 'J' {
				gm[k] += gm['J']
			}
		}*/
		if jokerValue == 5 {
			h.Type = FiveOfAKind
			return
		}
		delete(gm, 'J')
	}
	g := []int{}
	for _, v := range gm {
		g = append(g, v)
	}
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	if withJoker {
		g[0] += jokerValue
	}
	if len(g) == 1 {
		h.Type = FiveOfAKind
		return
	}
	if len(g) == 2 {
		for _, v := range g {
			if v == 4 {
				h.Type = FourOfAKind
				return
			}
			if v == 3 {
				h.Type = FullHouse
				return
			}
		}
	}
	if len(g) == 3 {
		for _, v := range g {
			if v == 3 {
				h.Type = ThreeOfAKind
				return
			}
			if v == 2 {
				h.Type = TwoPair
				return
			}
		}
	}
	if len(g) == 4 {
		h.Type = OnePair
		return
	}
	h.Type = HighCard
}
