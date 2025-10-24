func nextBeautifulNumber(n int) int {
    result := n + 1

    for !isGood(result) {
        result++
    }
    return result
}

func isGood(a int) bool {
    freq := map[int]int{}

    for a > 0 {
        freq[a % 10]++
        a /= 10
    }

    for key, val := range freq {
        if key != val {
            return false
        }
    }
    return true
}
