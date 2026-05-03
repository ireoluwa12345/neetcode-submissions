func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	longest := 0

	for _, num := range nums {
		numSet[num] = true
	}

	for _, num := range nums {
		if !numSet[num-1] {
			length := 1
			for numSet[num+length] {
				length++
			}

			longest = max(longest, length)
		}	
	}

	return longest
}