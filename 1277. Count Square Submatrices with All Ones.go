func countSquares(matrix [][]int) int {
    dp := make([][]int, len(matrix))

    for i := 0; i < len(matrix); i++ {
        dp[i] = make([]int, len(matrix[i]))
    }

    result := 0
    for i := 0; i < len(matrix); i++ {
        for j := 0; j < len(matrix[i]); j++ {
            dp[i][j] = matrix[i][j]

            if matrix[i][j] == 1 && i > 0 && j > 0 {
                dp[i][j] = 1 + min(dp[i - 1][j - 1], dp[i - 1][j], dp[i][j - 1])
            }
            result += dp[i][j]
        }
    }
    return result
}
