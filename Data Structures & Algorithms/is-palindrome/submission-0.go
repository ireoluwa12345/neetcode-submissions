func isPalindrome(s string) bool {
    var builder strings.Builder
    
    for _, r := range s {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            builder.WriteRune(unicode.ToLower(r))
        }
    }
    
    cleanStr := builder.String()
    
    i, j := 0, len(cleanStr)-1
    for i < j {
        if cleanStr[i] != cleanStr[j] {
            return false
        }
        i++
        j--
    }
    return true
}