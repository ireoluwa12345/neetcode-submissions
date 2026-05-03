func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[byte]int)
	mapT := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        mapS[s[i]]++
        mapT[t[i]]++
    }
    
    for char, count := range mapS {
        if mapT[char] != count {
            return false
        }
    }

    return true
}
