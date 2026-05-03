func productExceptSelf(nums []int) []int {
	response := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++ {
		response[i] = prefix
		prefix = prefix * nums[i];
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		response[i] = response[i] * postfix
		postfix = postfix * nums[i]
	}

	return response
}
