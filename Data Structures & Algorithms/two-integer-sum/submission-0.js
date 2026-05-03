class Solution {
    /**
     * @param {number[]} nums
     * @param {number} target
     * @return {number[]}
     */
    twoSum(nums, target) {
        let map = {};
        for (let i = 0; i < nums.length; i++){
            let difference = target - nums[i]
            if (map[nums[i]] != undefined) {
                return [map[nums[i]], i]
            }else{
                map[difference] = i
            }
        }

        return []
    }
}
