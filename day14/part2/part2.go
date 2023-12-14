package part2

import (
	"fmt"
	c "github.com/jerpa/adventofcode2023/common"
	"github.com/jerpa/adventofcode2023/day14/shared"
)

type info struct {
	sum      int
	checksum string
}

func Part2(data []string) int {

	checksums := []info{}
	for i := 0; i < 1000000000; i++ {
		data = shared.Tilt(data)
		data = shared.Rotate(data)
		data = shared.Tilt(data)
		data = shared.Rotate(data)
		data = shared.Tilt(data)
		data = shared.Rotate(data)
		data = shared.Tilt(data)
		data = shared.Rotate(data)

		checksums = append(checksums, info{sum: shared.CountWeight(data), checksum: fmt.Sprintf("%x", c.Checksum(data))})
		for j := 0; j < len(checksums)-1; j++ {
			if checksums[j].checksum == checksums[len(checksums)-1].checksum {
				l := 1000000000 - j - 1
				l %= i - j
				l += j
				return checksums[l].sum

			}
		}
	}

	return 0
}
