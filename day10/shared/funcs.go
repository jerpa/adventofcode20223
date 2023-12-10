package shared

import "errors"

type Direction string

const (
	Up    Direction = "U"
	Down  Direction = "D"
	Left  Direction = "L"
	Right Direction = "R"
)

func GetNeighbour(x, y int, data []string, direction Direction) (string, error) {
	if direction == Up && y > 0 {
		return string(data[y-1][x]), nil
	}
	if direction == Down && y < len(data)-1 {
		return string(data[y+1][x]), nil
	}
	if direction == Left && x > 0 {
		return string(data[y][x-1]), nil
	}
	if direction == Right && x < len(data[y])-1 {
		return string(data[y][x+1]), nil
	}
	return "", errors.New("out of bounds")
}
