func lexicalOrder(n int) []int {
    result := []int{}

    next := 1
    for i := 0; i < n; i++ {
        result = append(result, next)

        if next * 10 <= n {
            next *= 10
        } else {
            for next % 10 == 9 || next >= n {
                next /= 10
            }
            next++
        }
    }
    return result
}
