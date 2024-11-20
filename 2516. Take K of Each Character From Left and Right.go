func takeCharacters(s string, k int) int {
    left := make([]int, 3)
    right := make([]int, 3)

    lastIdx := 0
    for lastIdx = 0; lastIdx < len(s); lastIdx++ {
        idx := s[lastIdx] - 'a'
        left[idx]++

        if hasNeeded(left, right, k) {
            break
        }
    }

    if lastIdx >= len(s) {
        return -1
    }

    result := lastIdx + 1
    rightIdx := len(s) - 1

    for ; lastIdx >= 0; lastIdx-- {
        idx := s[lastIdx] - 'a'
        left[idx]--

        for rightIdx >= lastIdx && !hasNeeded(left, right, k) {
            idx := s[rightIdx] - 'a'
            right[idx]++
            rightIdx--
        }

        if hasNeeded(left, right, k) {
            result = min(result, lastIdx + len(s) - 1 - rightIdx)
        }
    }

    return result
}

func hasNeeded(a []int, b []int, k int) bool {
    for i := 0; i < 3; i++ {
        if a[i] + b[i] < k {
            return false
        }
    }
    return true
}
