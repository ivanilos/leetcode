func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := initDPTable(obstacleGrid)

    for i := 1; i < len(obstacleGrid); i++ {
        for j := 1; j < len(obstacleGrid[0]); j++ {
            if obstacleGrid[i][j] == 0 {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
            }
        }
    }
    return dp[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]
}

func initDPTable(obstacleGrid [][]int) [][]int {
    dp := make([][]int, len(obstacleGrid))
    for i := 0; i < len(obstacleGrid); i++ {
        dp[i] = make([]int, len(obstacleGrid[i]))
    }

    seenObstacle := false
    for i := 0; i < len(obstacleGrid); i++ {
        if obstacleGrid[i][0] == 1 {
            seenObstacle = true
        }
        if !seenObstacle {
            dp[i][0] = 1
        }
    }

    seenObstacle = false
    for j := 0; j < len(obstacleGrid[0]); j++ {
        if obstacleGrid[0][j] == 1 {
            seenObstacle = true
        }
        if !seenObstacle {
            dp[0][j] = 1
        }
    }

    return dp
}
