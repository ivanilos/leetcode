func flowerGame(n int, m int) int64 {
    result := int64(0)

    for x := 1; x <= n; x++ {
        if x % 2 == 0 {
            result += int64((m + 1) / 2)
        } else {
            result += int64(m / 2)
        }
    }
    return result
}
