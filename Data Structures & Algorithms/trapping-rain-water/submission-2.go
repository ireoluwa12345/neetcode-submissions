func trap(height []int) int {
	response := 0
	if len(height) < 1 {
		return response
	}

	l, r := 0, len(height) - 1
	leftMax, rightMax := height[l], height[r]

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			response += leftMax - height[l]
		}else {
			r--
			rightMax = max(rightMax, height[r])
			response += rightMax - height[r]
		}
	}

	return response
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}