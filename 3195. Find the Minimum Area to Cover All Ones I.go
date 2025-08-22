func minimumArea(grid [][]int) int {
    minRow := len(grid)
    maxRow := -1

    minCol := len(grid[0])
    maxCol := -1
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                minRow = min(minRow, i)
                maxRow = max(maxRow, i)
                minCol = min(minCol, j)
                maxCol = max(maxCol, j)
            }
        }
    }

    return (maxRow - minRow + 1) * (maxCol - minCol + 1)
}
