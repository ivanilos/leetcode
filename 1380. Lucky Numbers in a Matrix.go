func luckyNumbers (matrix [][]int) []int {
    INF := int(1e9)
    rows := len(matrix)
    cols := len(matrix[0])

    minInRow := make([]int, rows)
    maxInCol := make([]int, cols)

    for i := 0; i < rows; i++ {
        minInRow[i] = INF
    }
    for j := 0; j < cols; j++ {
        maxInCol[j] = -INF
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            minInRow[i] = min(minInRow[i], matrix[i][j])
            maxInCol[j] = max(maxInCol[j], matrix[i][j])
        }
    }

    result := []int{}
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if matrix[i][j] == minInRow[i] && matrix[i][j] == maxInCol[j] {
                result = append(result, matrix[i][j])
            }
        }
    }
    return result
}
