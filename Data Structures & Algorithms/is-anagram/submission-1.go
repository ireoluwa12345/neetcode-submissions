func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[rune]int)
	mapT := make(map[rune]int)

	for _, char := range s {
        mapS[char]++
    }
    for _, char := range t {
        mapT[char]++
    }

    // Compare counts
    for char, count := range mapS {
        if mapT[char] != count {
            return false
        }
    }

    return true
}
