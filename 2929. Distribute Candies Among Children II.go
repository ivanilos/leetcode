func distributeCandies(n int, limit int) int64 {
    result := int64(0)
    for first := 0; first <= limit && first <= n; first++ {
        left := n - first

        if left > 2 * limit {
            continue
        }

        if left <= limit {
            result += int64(left + 1)
        } else {
            result += int64(limit - (left - limit) + 1)
        }
    }

    return result
}
