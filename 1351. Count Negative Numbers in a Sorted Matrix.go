func countNegatives(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

    lastNotNegIdx := cols - 1
    for j := 0; j < cols; j++ {
        if grid[0][j] < 0 {
            lastNotNegIdx = j - 1
            break
        }
    }

    result := cols - (lastNotNegIdx + 1)

    for i := 1; i < rows; i++ {
        for ; lastNotNegIdx >= 0; lastNotNegIdx-- {
            if grid[i][lastNotNegIdx] >= 0 {
                break
            }
        }
        result += cols - (lastNotNegIdx + 1)
    }

    return result
}
