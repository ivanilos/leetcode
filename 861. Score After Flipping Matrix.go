func matrixScore(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])

    cols_ones := make([]int, m)
    for i := 0; i < n; i++ {
        swap := false
        if grid[i][0] == 0 {
            swap = true
        }

        for j := 0; j < m; j++ {
            if grid[i][j] == 1 && !swap || grid[i][j] == 0 && swap {
                cols_ones[j]++
            }
        }
    }

    result := 0
    base := 1
    for j := m - 1; j >= 0; j-- {
        result += base * max(cols_ones[j], n - cols_ones[j])
        base *= 2
    }
    return result
}
