package shared

func DoPart(data []string, smudges int) int {
	sum := 0
	start := 0
	for i, v := range data {
		if v == "" {
			sum += FindMirrorValue(data[start:i], smudges)
			start = i + 1
		}
	}
	sum += FindMirrorValue(data[start:], smudges)

	return sum
}
func hasSmudgesRow(data []string, fromRow, smudges int) bool {
	sum := 0
	for i := 0; fromRow+i < len(data) && fromRow-1-i >= 0; i++ {
		for j := 0; j < len(data[0]); j++ {
			if data[fromRow+i][j] != data[fromRow-1-i][j] {
				sum++
				if sum <= smudges {
					continue
				}
				return false
			}
		}
	}
	return sum == smudges
}
func hasSmudgesCol(data []string, fromCol, smudges int) bool {
	sum := 0
	for i := 0; fromCol+i < len(data[0]) && fromCol-1-i >= 0; i++ {
		for j := 0; j < len(data); j++ {
			if data[j][fromCol+i] != data[j][fromCol-1-i] {
				sum++
				if sum <= smudges {
					continue
				}
				return false
			}
		}
	}
	return sum == smudges
}

func FindMirrorValue(data []string, smudges int) int {

	for row := 1; row < len(data); row++ {
		if hasSmudgesRow(data, row, smudges) {
			return row * 100
		}
	}
	for col := 1; col < len(data[0]); col++ {
		if hasSmudgesCol(data, col, smudges) {
			return col
		}
	}
	return 0
}
