func repeatLimitedString(s string, repeatLimit int) string {
    freq := make([]int, 26)

    for _, c := range s {
        idx := c - 'a'
        freq[idx]++
    }

    result := []rune{}
    solve(25, 25, 0, repeatLimit, freq, &result)

    return string(result)
}

func solve(next, last, usedLast, repeatLimit int, freq []int, result *[]rune) {
    if next < 0 {
        return
    }

    if freq[next] == 0 {
        for i := next - 1; i >= 0; i-- {
            if freq[i] != 0 {
                solve(next - 1, last, usedLast, repeatLimit, freq, result)
                break
            }
        }
        return 
    }

    if next == last && usedLast == repeatLimit {
        solve(next - 1, last, usedLast, repeatLimit, freq, result)
        return 
    }

    *result = append(*result, rune(next + 'a'))
    freq[next]--
    if (next == last) {
        usedLast++
    } else {
        usedLast = 1
        last = next
    }

    solve(25, last, usedLast, repeatLimit, freq, result)
    return 
}
