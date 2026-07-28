package rewritefixture

// Divide returns zero when count is zero.
func Divide(total, count int) int {
	if count == 0 {
		return 0
	}
	return total / count
}
