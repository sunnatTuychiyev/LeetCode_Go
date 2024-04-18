func reverse(x int) int {
	var reversed int
	for x != 0 {
		digit := x % 10
		x /= 10
		reversed = reversed*10 + digit
	}

	if reversed < math.MinInt32 || reversed > math.MaxInt32 {
		return 0
	}

	return reversed
}