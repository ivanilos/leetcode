func canCross(stones []int) bool {
    n := len(stones)

    stoneIdx := map[int]int{}
    for i, stone := range stones {
        stoneIdx[stone] = i
    }

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = -1
        }
    }

    result := solve(0, 0, stones, stoneIdx, dp)

    if result == 1 {
        return true
    }
    return false
}

func solve(prevIdx, curIdx int, stones []int, stoneIdx map[int]int, dp [][]int) int {
    if curIdx == len(stones) - 1 {
        return 1
    }

    if dp[prevIdx][curIdx] != -1 {
        return dp[prevIdx][curIdx]
    }

    dp[prevIdx][curIdx] = 0

    curJump := stones[curIdx] - stones[prevIdx]
    for nextJump := curJump - 1; nextJump <= curJump + 1; nextJump++ {
        nextPos := stones[curIdx] + nextJump
        if nextIdx, ok := stoneIdx[nextPos]; nextJump > 0 && ok {
            dp[prevIdx][curIdx] = max(dp[prevIdx][curIdx], solve(curIdx, nextIdx, stones, stoneIdx, dp))
        }
    }

    return dp[prevIdx][curIdx]
}
