func twoSum(numbers []int, target int) []int {
	ptr1, ptr2 := 0, len(numbers) - 1

	for ptr1 < ptr2 {
		currentSum := numbers[ptr1] + numbers[ptr2]
		if currentSum > target {
			ptr2--
		}else if currentSum < target {
			ptr1++
		}else{
			return []int{ptr1 + 1, ptr2 + 1}
		}
	}

	return []int{}
}
