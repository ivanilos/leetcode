func rangeAddQueries(n int, queries [][]int) [][]int {
    result := make([][]int, n)
    for i := 0; i < n; i++ {
        result[i] = make([]int, n)
    }

    for _, query := range queries {
        for row := query[0]; row <= query[2]; row++ {
            col1, col2 := query[1], query[3]
            result[row][col1]++

            if col2 + 1 < n {
                result[row][col2 + 1]--
            }
        }
    }

    for i := 0; i < n; i++ {
        curVal := 0
        for j := 0; j < n; j++ {
            curVal += result[i][j]
            result[i][j] = curVal
        }
    }

    return result
}
