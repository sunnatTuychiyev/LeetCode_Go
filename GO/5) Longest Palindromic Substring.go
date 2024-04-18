func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	var expandAroundCenter func(left, right int) string
	expandAroundCenter = func(left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}
		return s[left+1 : right]
	}

	var longest string
	for i := range s {
		// Check palindrome with odd length
		palindromeOdd := expandAroundCenter(i, i)
		// Check palindrome with even length
		palindromeEven := expandAroundCenter(i, i+1)
		// Update longest palindrome found so far
		if len(palindromeOdd) > len(longest) {
			longest = palindromeOdd
		}
		if len(palindromeEven) > len(longest) {
			longest = palindromeEven
		}
	}
	return longest
}
