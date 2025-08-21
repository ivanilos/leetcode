func removeBoxes(boxes []int) int {
    n := len(boxes)

    dp := make([][][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]int, n)
            for k := 0; k < n; k++ {
                dp[i][j][k] = -1
            }
        }
    }

    result := solve(0, n - 1, 0, boxes, dp)

    return result
}

func solve(l, r, leftLeftover int, boxes []int, dp [][][]int) int {
    if l > r {
        return 0
    }

    if dp[l][r][leftLeftover] != -1 {
        return dp[l][r][leftLeftover]
    }

    dp[l][r][leftLeftover] = (leftLeftover + 1) * (leftLeftover + 1) + solve(l + 1, r, 0, boxes, dp)

    for i := l + 1; i <= r; i++ {
        if boxes[i] == boxes[l] {
            candidateMax := solve(l + 1, i - 1, 0, boxes, dp) + solve(i, r, leftLeftover + 1, boxes, dp)
            dp[l][r][leftLeftover] = max(dp[l][r][leftLeftover], candidateMax)
        }
    }

    return dp[l][r][leftLeftover]
}
