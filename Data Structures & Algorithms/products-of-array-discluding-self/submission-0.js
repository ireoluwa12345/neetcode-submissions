class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    productExceptSelf(nums) {
        let resultsArray = new Array(nums.length)
        for (let i = 0; i < nums.length; i++) {
            let currentProduct = 1
            for (let j = 0; j < nums.length; j++){
                if (i != j) {
                    currentProduct = currentProduct * nums[j]
                }
            }
            resultsArray[i] = currentProduct
        }

        return resultsArray;
    }
}
