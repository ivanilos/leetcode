var INF = int(1e9)

func cherryPickup(grid [][]int) int {
    if grid[0][0] == -1 {
        return 0
    }

    n := len(grid)
    m := len(grid[0])

    dp := make([][][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]int, m)
            for k := 0; k < m; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    result := solve(0, 0, 0, n, m, grid, dp)

    return max(0, result)
}

func solve(row1, row2, col, n, m int, grid [][]int, dp [][][]int) int {
    col2 := row1 + col - row2
    if row1 >= n || row2 >= n || col >= m || col2 >= m {
        return -INF
    }
    if row1 == n - 1 && row2 == n - 1 && col == m - 1 {
        return grid[row1][col]
    }
    if dp[row1][row2][col] != -1 {
        return dp[row1][row2][col]
    }

    dp[row1][row2][col] = 0
    if grid[row1][col] == -1 || grid[row2][col2] == -1 {
        dp[row1][row2][col] = -INF
        return -INF
    }

    curCherries := grid[row1][col]
    if row1 != row2 {
        curCherries += grid[row2][col2]
    }

    dp[row1][row2][col] = curCherries + solve(row1 + 1, row2 + 1, col, n, m, grid, dp)
    dp[row1][row2][col] = max(dp[row1][row2][col], curCherries + solve(row1, row2 + 1, col + 1, n, m, grid, dp))
    dp[row1][row2][col] = max(dp[row1][row2][col], curCherries + solve(row1, row2, col + 1, n, m, grid, dp))
    dp[row1][row2][col] = max(dp[row1][row2][col], curCherries + solve(row1 + 1, row2, col, n, m, grid, dp))

    return dp[row1][row2][col]
}
