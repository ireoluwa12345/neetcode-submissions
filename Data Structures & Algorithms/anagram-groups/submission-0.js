class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        let map = new Map()

        for (let i = 0; i < strs.length; i++) {
            let sorted_data = strs[i].split("").sort().join("")
            if (!map.has(sorted_data)) {
                map.set(sorted_data, [])
            }

            map.get(sorted_data).push(strs[i]);
        }

        return Array.from(map.values())
    }
}
