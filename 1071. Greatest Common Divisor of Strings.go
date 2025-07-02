func gcdOfStrings(str1 string, str2 string) string {
    if len(str1) > len(str2) {
        str1, str2 = str2, str1
    }

    result := ""
    for end := 0; end < len(str1); end++ {
        cand := str1[: end + 1]

        if isPeriod(str1, cand) && isPeriod(str2, cand) {
            result = str1[: end + 1]
        }
    }

    return result
}

func isPeriod(str string, cand string) bool {
    if len(str) % len(cand) != 0 {
        return false
    }

    for i := 0; i < len(str); i++ {
        if str[i] != cand[i % len(cand)] {
            return false
        }
    }

    return true
}
