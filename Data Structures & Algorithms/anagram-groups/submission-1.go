func groupAnagrams(strs []string) [][]string {
    // 1. Key is the frequency array, Value is a slice of strings
    anagramMap := make(map[[26]byte][]string)

    for _, s := range strs {
        var count [26]byte
        
        // 2. Count occurrences of each character
        for i := 0; i < len(s); i++ {
            count[s[i]-'a']++
        }

        // 3. Append the string to the corresponding frequency key
        anagramMap[count] = append(anagramMap[count], s)
    }

    // 4. Convert the map values into the required 2D slice format
    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }

    return result
}