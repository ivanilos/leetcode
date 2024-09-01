func construct2DArray(original []int, m int, n int) [][]int {
    sz := len(original)

    if sz != m * n {
        return [][]int{}
    }

    result := [][]int{}
    for j := 0; j < len(original); j++ {
        if j % n == 0 {
            result = append(result, []int{})
        }
        result[len(result) - 1] = append(result[len(result) - 1], original[j])
    }
    return result
}
