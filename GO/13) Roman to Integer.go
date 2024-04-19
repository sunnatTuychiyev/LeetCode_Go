func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0
	for i := len(s) - 1; i >= 0; i-- {
		value := romanValues[s[i]]
		if value >= prev {
			result += value
		} else {
			result -= value
		}
		prev = value
	}

	return result
}
