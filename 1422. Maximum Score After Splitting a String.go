func maxScore(s string) int {
    ones := 0
    for _, c := range s {
        if c == '1' {
            ones++
        }
    }

    result := 0
    prefZeroes := 0
    for i := 0; i < len(s) - 1; i++ {
        c := s[i]
        if c == '0' {
            prefZeroes++
        } else {
            ones--
        }
        result = max(result, ones + prefZeroes)
    }

    return result
}
