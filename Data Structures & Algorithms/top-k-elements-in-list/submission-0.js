class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        let map = new Map()

        nums.forEach((num) => {
            if (!map.has(num)) {
                map.set(num, 0)
            }

            map.set(num, map.get(num) + 1)
        })


        const mapEntries = [...map.entries()];

        let sorted_data = mapEntries.sort((a, b) => {
            return b[1] - a[1]
        })

        console.log(sorted_data) 

        return sorted_data.slice(0, k).map((pair) => pair[0])
    }
}
