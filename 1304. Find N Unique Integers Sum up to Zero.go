func sumZero(n int) []int {
    if n == 2 {
        return []int{-1, 1}
    }
    result := []int{}

    sum := 0
    for i := 0; i < n - 1; i++ {
        sum += i
        result = append(result, i)
    }

    result = append(result, -sum)

    return result
}
