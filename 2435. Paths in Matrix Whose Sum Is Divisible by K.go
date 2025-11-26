const MOD = int(1e9 + 7)

func numberOfPaths(grid [][]int, k int) int {
    rows := len(grid)
    cols := len(grid[0])

    dp := make([][][]int, rows)
    for i := 0; i < rows; i++ {
        dp[i] = make([][]int, cols)
        for j := 0; j < cols; j++ {
            dp[i][j] = make([]int, k)
        }
    }

    remSum := (grid[0][0]) % k
    dp[0][0][remSum] = 1
    for i := 1; i < rows; i++ {
        remSum = (remSum + grid[i][0]) % k
        dp[i][0][remSum] = 1
    }

    remSum = (grid[0][0]) % k
    for j := 1; j < cols; j++ {
        remSum = (remSum + grid[0][j]) % k
        dp[0][j][remSum] = 1
    }

    for i := 1; i < rows; i++ {
        for j := 1; j < cols; j++ {
            for rem := 0; rem < k; rem++ {
                lastRem := (((rem - grid[i][j]) % k) + k) % k
                dp[i][j][rem] = (dp[i][j][rem] +  dp[i][j - 1][lastRem] + dp[i - 1][j][lastRem]) % MOD
            }
        }
    }

    return dp[rows - 1][cols - 1][0]
}
