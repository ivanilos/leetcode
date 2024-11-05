func minChanges(s string) int {
    chunksSize := getChunkSizes(s)

    result := 0
    delta := 0
    for _, sz := range chunksSize {
        if (sz + delta) % 2 != 0 {
            result++
            delta = -1
        } else {
            delta = 0
        }
    }

    return result
}

func getChunkSizes(s string) []int {
    result := []int{}

    cur := 0
    lastChar := rune(s[0])
    for _, c := range s {
        if c == lastChar {
            cur++
        } else {
            result = append(result, cur)
            cur = 1
            lastChar = c
        }
    }
    result = append(result, cur)

    return result
}
