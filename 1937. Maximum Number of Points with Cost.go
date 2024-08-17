func maxPoints(points [][]int) int64 {
    dp := make([][]int64, len(points))
    for i := 0; i < len(points); i++ {
        dp[i] = make([]int64, len(points[i]))
    }

    for j := 0; j < len(points[0]); j++ {
        dp[0][j] = int64(points[0][j])
    }

    for i := 1; i < len(points); i++ {
        leftMax := make([]int64, len(points[i]))
        rightMax := make([]int64, len(points[i]))

        leftMax[0] = dp[i - 1][0]
        for j := 1; j < len(points[i]); j++ {
            leftMax[j] = max(leftMax[j - 1] - 1, dp[i - 1][j])
        }

        rightMax[len(points[i]) - 1] = dp[i - 1][len(points[i]) - 1]
        for j := len(points[i]) - 2; j >= 0; j-- {
           rightMax[j] = max(rightMax[j + 1] - 1, dp[i - 1][j])
        }

        for j := 0; j < len(points[i]); j++ {
            dp[i][j] = int64(points[i][j]) + max(leftMax[j], rightMax[j])
        }
    }

    result := int64(0)
    for j := 0; j < len(points[0]); j++ {
        result = max(result, dp[len(points) - 1][j])
    }
    return result
}

func abs(x int64) int64 {
    if x < 0 {
        return -x
    }
    return x
}
