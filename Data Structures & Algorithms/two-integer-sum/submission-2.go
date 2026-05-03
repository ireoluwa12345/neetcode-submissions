func twoSum(nums []int, target int) []int {
    listHash := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        listHash[nums[i]] = i;
    }

    for i := 0; i < len(nums); i++ {
        difference := target - nums[i]
        val, ok := listHash[difference]

        if ok && val != i {
            return []int{i, val}
        }
    }

    return []int{}
}
