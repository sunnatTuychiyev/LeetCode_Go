func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	index, step := 0, 1

	for _, char := range s {
		rows[index] += string(char)
		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}
		index += step
	}

	result := ""
	for _, row := range rows {
		result += row
	}
	return result
}