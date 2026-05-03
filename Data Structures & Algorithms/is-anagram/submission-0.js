class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
        if (s.length !== t.length) return false;

        let mapS = {};
        let mapT = {};

        for (let i = 0; i < s.length; i++) {
            mapS[s[i]] = (mapS[s[i]] || 0) + 1;
            mapT[t[i]] = (mapT[t[i]] || 0) + 1;
        }

        for (let char in mapS) {
            if (mapS[char] !== mapT[char]) {
                return false;
            }
        }

        return true;
    }
}
