func islandPerimeter(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    result := 0

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == 1 {
                result += getPerimeter(grid, n, m, i, j)
            }
        }
    }
    return result
}

func getPerimeter(grid [][]int, n, m, i, j int) int {
    result := 0
    for dx := -1; dx <= 1; dx++ {
        if !isIn(i + dx, j, n, m) || grid[i + dx][j] == 0 {
            result++
        }
    }
    for dy := -1; dy <= 1; dy++ {
        if !isIn(i, j + dy, n, m) || grid[i][j + dy] == 0 {
            result++
        }
    }
    return result
}

func isIn(x, y, n, m int) bool {
    return 0 <= x && x < n && 0 <= y && y < m
}
