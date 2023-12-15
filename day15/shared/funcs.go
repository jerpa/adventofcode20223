package shared

func Hash15OnString(s string) int {
	x := 0
	for _, ch := range s {
		x = Hash15(x, int(ch))
	}
	return x
}
func Hash15(curr, ch int) int {
	curr += ch
	curr *= 17
	curr %= 256
	return curr
}
