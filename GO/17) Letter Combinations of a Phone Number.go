var letterMap = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	return generateCombinations(digits, 0, "")
}

func generateCombinations(digits string, index int, combination string) []string {
	if index == len(digits) {
		return []string{combination}
	}
	digit := digits[index]
	letters := letterMap[digit]
	var result []string
	for _, letter := range letters {
		result = append(result, generateCombinations(digits, index+1, combination+letter)...)
	}
	return result
}
