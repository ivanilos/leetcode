func maxUniqueSplit(s string) int {
    result := 0
    for mask := 0; mask < 1 << len(s); mask++ {
        result = max(result, splitValue(s, mask))
    }
    return result
}

func splitValue(s string, mask int) int {
    exists := map[string]bool{}

    start := 0
    result := 0
    for i := 0; i < len(s); i++ {
        bit := (mask >> i) & 1

        if bit == 1 {
            result++

            str := s[start : i + 1]
            if _, ok := exists[str]; ok {
                return 0
            }

            exists[str] = true
            start = i + 1
        }
    }

    return result
}
