package compare

func Larger(a int, b int) int {
	if a < b { // Oops! This comparison is backward!
		return a
	} else {
		return b
	}
}
