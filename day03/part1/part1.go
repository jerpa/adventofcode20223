package part1

import (
	"unicode"
)

func Part1(data []string) int {
	sum := 0

	for y, v := range data {
		start := -1
		stop := -1
		num := 0
		for x, c := range v {
			if unicode.IsDigit(c) {
				if start == -1 {
					start = x
				}
				num = num*10 + int(c-'0')
			}
			if x == len(data[y])-1 {
				x++
			}

			if x >= len(data[y]) || !unicode.IsDigit(c) {
				if start != -1 {
					stop = x - 1
					ok := false
					for i := start - 1; i <= stop+1; i++ {
						if hasPart(data, i, y-1) || hasPart(data, i, y+1) {
							ok = true
							break
						}
					}
					if ok || hasPart(data, start-1, y) || hasPart(data, stop+1, y) {
						//common.Print(start, stop, y, num)
						sum += num
					}
					start = -1
					stop = -1
					num = 0
				}
			}

		}

	}

	return sum
}
func hasPart(data []string, x, y int) bool {
	if x < 0 || y < 0 || y >= len(data) || x >= len(data[y]) {
		return false
	}
	c := data[y][x]
	if !unicode.IsDigit(rune(c)) && c != '.' {
		return true
	}

	return false
}
