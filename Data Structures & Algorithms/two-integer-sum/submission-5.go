func twoSum(nums []int, target int) []int {
    listHash := make(map[int]int, len(nums))

    for i, num := range nums {
        difference := target - num
        
        if val, ok := listHash[difference]; ok {
            return []int{val, i}
        }
        
        listHash[num] = i
    }

    return []int{}
}