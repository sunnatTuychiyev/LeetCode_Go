func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1

	for left < right {
		minH := height[left]
		if height[right] < minH {
			minH = height[right]
		}

		area := minH * (right - left)
		if area > max {
			max = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
