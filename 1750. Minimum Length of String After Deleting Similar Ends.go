func minimumLength(s string) int {
    firstIdx := 0
    lastIdx := len(s) - 1

    for firstIdx < lastIdx && s[firstIdx] == s[lastIdx] {
        val := s[firstIdx]
        for firstIdx <= lastIdx && s[firstIdx] == val {
            firstIdx++
        }
        for firstIdx <= lastIdx && s[lastIdx] == val {
            lastIdx--
        }
    }
    return lastIdx - firstIdx + 1
}
