func divideString(s string, k int, fill byte) []string {
    result := []string{}

    for i := 0; i < len(s); i += k {
        result = append(result, s[i : min(i + k, len(s))])
    }

    last := []byte(result[len(result) - 1])
    for i := 0; i < (k - len(s) % k) % k; i++ {
        last = append(last, fill)
    }
    result[len(result) - 1] = string(last)

    return result
}
