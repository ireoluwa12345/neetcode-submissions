func hasDuplicate(nums []int) bool {
	hashMap := make(map[int]int)
    for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			hashMap[nums[i]]++
			if hashMap[nums[i]] > 1 {
				return true
			}
		}else{
			hashMap[nums[i]] = 1
		}
	}
	return false
}
