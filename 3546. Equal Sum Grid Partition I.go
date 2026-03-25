func canPartitionGrid(grid [][]int) bool {
    rows := len(grid)
    cols := len(grid[0])

    rowSum := make([]int64, rows)
    colSum := make([]int64, cols)

    totalSum := int64(0)
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            totalSum += int64(grid[i][j])
            rowSum[i] += int64(grid[i][j])
            colSum[j] += int64(grid[i][j])
        }
    }

    curSum := totalSum
    for i := 0; i < rows - 1; i++ {
        curSum -= rowSum[i]

        if curSum * 2 == totalSum {
            return true
        }
    }

    curSum = totalSum
    for j := 0; j < cols - 1; j++ {
        curSum -= colSum[j]

        if curSum * 2 == totalSum {
            return true
        }
    }

    return false
}
