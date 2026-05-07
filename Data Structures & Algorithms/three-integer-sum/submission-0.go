func threeSum(nums []int) [][]int {
	response := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		pointerOne, pointerTwo := i + 1, len(nums) - 1

		for pointerOne < pointerTwo {
			threeSum := nums[i] + nums[pointerOne] + nums[pointerTwo]

			if threeSum < 0 {
				pointerOne++
			}else if threeSum > 0 {
				pointerTwo--
			}else {
				response = append(response, []int{nums[i], nums[pointerOne], nums[pointerTwo]})
				pointerOne++

				for nums[pointerOne] == nums[pointerOne - 1] && pointerOne < pointerTwo {
					pointerOne++
				}
			}
		}
	}

	return response
}
