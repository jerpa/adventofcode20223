package shared

func Rotate(data []string) []string {
	res := []string{}
	for x := 0; x < len(data[0]); x++ {
		line := ""
		for y := len(data) - 1; y >= 0; y-- {
			line += string(data[y][x])
		}
		res = append(res, line)
	}
	return res
}
func Tilt(data []string) []string {
	for x := 0; x < len(data[0]); x++ {
		for y := 0; y < len(data); y++ {
			if data[y][x] == 'O' {
				var n int
				for n = y - 1; n >= 0 && data[n][x] == '.'; n-- {

				}
				n++

				if n == y {
					continue
				}
				data[n] = data[n][:x] + "O" + data[n][x+1:]
				data[y] = data[y][:x] + "." + data[y][x+1:]
			}
		}
	}
	return data
}
func CountWeight(data []string) int {
	sum := 0
	for y := range data {
		for x := range data[y] {
			if data[y][x] == 'O' {
				sum += len(data) - y
			}
		}
	}
	return sum
}
