func gridGame(grid [][]int) int64 {
    cols := len(grid[0])

    if cols == 1 {
        return 0
    }

    bottomPrefix := make([]int64, cols)
    topSuffix := make([]int64, cols)

    bottomPrefix[0] = int64(grid[1][0])
    for i := 1; i < cols; i++ {
        bottomPrefix[i] = int64(grid[1][i]) + bottomPrefix[i - 1]
    }

    topSuffix[cols - 1] = int64(grid[0][cols - 1])
    for i := cols - 2; i >= 0; i-- {
        topSuffix[i] = int64(grid[0][i]) + topSuffix[i + 1]
    }

    result := min(topSuffix[1], bottomPrefix[cols - 2])
    for i := 2; i < cols; i++ {
        result = min(result, max(topSuffix[i], bottomPrefix[i - 2]))
    }

    return result
}
