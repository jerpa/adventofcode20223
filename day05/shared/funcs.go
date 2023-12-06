package shared

func DoNothing() {
}

type SoilRange struct {
	InStart  int
	InEnd    int
	OutStart int
}

func (s *SoilRange) NewSoil(i int) (int, bool) {
	if i >= s.InStart && i <= s.InEnd {
		return s.OutStart + (i - s.InStart), true
	}
	return -1, false
}
