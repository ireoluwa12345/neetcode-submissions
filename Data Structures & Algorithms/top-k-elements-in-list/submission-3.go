func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums) + 1)
	response := []int{}

	for _, num := range nums {
		count[num] += 1
	}

	for key, val := range count {
		freq[val] = append(freq[val], key)
	}

	for i := len(freq) - 1; i > 0; i-- {
		for _, val := range freq[i] {
			response = append(response, val)
			if len(response) == k {
				return response
			}
		}
	}

	return response
}
