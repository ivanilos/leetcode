func maximumLength(s string) int {
    MAX_LEN := 50
    freq := make([][]int, 26)
    for i := 0; i < 26; i++ {
        freq[i] = make([]int, MAX_LEN + 1)
    }

    last := s[0]
    counter := 0
    for i := 0; i < len(s); i++ {
        if s[i] == last {
            counter++
        } else {
            freq[last - 'a'][counter]++
            counter = 1
            last = s[i]
        }
    }
    freq[last - 'a'][counter]++

    result := -1
    NEEDED_FREQ := 3
    for i := 0; i < 26; i++ {
        for j := MAX_LEN; j > 0; j-- {
            if freq[i][j] >= NEEDED_FREQ {
                result = max(result, j)
                break
            }
            if freq[i][j] > 0 {
                freq[i][j - 1] += freq[i][j] + 1
            }
        }
    }
    return result
}
