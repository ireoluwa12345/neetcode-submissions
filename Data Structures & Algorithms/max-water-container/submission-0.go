func maxArea(heights []int) int {
	maxArea := 0
	leftPointer, rightPointer := 0, len(heights) - 1

	for leftPointer < rightPointer {
		area := min(heights[leftPointer], heights[rightPointer]) * (rightPointer - leftPointer)
		if area > maxArea {
			maxArea = area
		}

		if heights[leftPointer] <= heights[rightPointer] {
			leftPointer++
		} else {
			rightPointer--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}