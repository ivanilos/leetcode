func characterReplacement(s string, k int) int {
    freq := map[byte]int{}

    result := 0

    l := 0
    for r := 0; r < len(s); r++ {
        curChar := s[r]
        freq[curChar]++

        maxFreq := getMaxFreq(freq)

        for l <= r && curSize(l, r) > maxFreq + k {
            char := s[l]
            freq[char]--

            maxFreq = getMaxFreq(freq)
            l++
        }

        result = max(result, r - l + 1)
    }

    return result
}

func getMaxFreq(freq map[byte]int) int {
    result := 0

    for _, val := range freq {
        result = max(result, val)
    }

    return result
}

func curSize(l, r int) int {
    return r - l + 1
}
