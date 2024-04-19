func intToRoman(num int) string {
	var res string
	switch {
	case num >= 1000:
		res = GreaterThousand(num)
	case num >= 500:
		res = GreaterFiveHundred(num)
	case num >= 100:
		res = GreaterHundred(num)
	case num >= 50:
		res = GreaterFifty(num)
	case num >= 10:
		res = GreaterTen(num)
	case num >= 5:
		res = GreaterFive(num)
	default:
		res = GreaterOne(num)
	}

	return res
}

func GreaterThousand(num int) string {
	var res string
	for num >= 1000 {
		res += "M"
		num -= 1000
	}

	res += GreaterFiveHundred(num)

	return res
}

func GreaterFiveHundred(num int) string {
	var res string
	if num/900 == 1 {
		res += "CM"
		num -= 900
	} else if num >= 500 {
		res += "D"
		num -= 500
	}

	res += GreaterHundred(num)

	return res
}

func GreaterHundred(num int) string {
	var res string
	if num/400 == 1 {
		res += "CD"
		num -= 400
	} else if num >= 100 {
		for num >= 100 {
			res += "C"
			num -= 100
		}
	}

	res += GreaterFifty(num)

	return res
}

func GreaterFifty(num int) string {
	var res string
	if num/90 == 1 {
		res += "XC"
		num -= 90
	} else if num >= 50 {
		res += "L"
		num -= 50
	}

	res += GreaterTen(num)

	return res
}

func GreaterTen(num int) string {
	var res string
	if num/40 == 1 {
		res += "XL"
		num -= 40
	} else if num >= 10 {
		for num >= 10 {
			res += "X"
			num -= 10
		}
	}

	res += GreaterFive(num)

	return res
}

func GreaterFive(num int) string {
	var res string
	if num == 9 {
		res += "IX"
		num -= 9
	} else if num >= 5 {
		res += "V"
		num -= 5
	}

	res += GreaterOne(num)

	return res
}

func GreaterOne(num int) string {
	var res string
	if num == 4 {
		res += "IV"
		num -= 4
	} else {
		for num > 0 {
			res += "I"
			num -= 1
		}
	}