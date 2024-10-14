func minimumTotal(triangle [][]int) int {
    dp := make([][]int, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([]int, len(triangle))
    }

    dp[0][0] = triangle[0][0]
    for i := 1; i < len(triangle); i++ {
        row := i % 2

        dp[row][0] = triangle[i][0] + dp[1 - row][0]
        for j := 1; j < i; j++ {
            dp[row][j] = triangle[i][j] + min(dp[1 - row][j - 1], dp[1 - row][j])
        }
        dp[row][i] = triangle[i][i] + dp[1 - row][i - 1]
    }

    lastRow := (len(triangle) - 1) % 2
    result := dp[lastRow][0]
    for j := 1; j < len(triangle); j++ {
        result = min(result, dp[lastRow][j])
    }
    return result
}
