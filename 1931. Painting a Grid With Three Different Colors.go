const MOD = int(1e9 + 7)

func colorTheGrid(m int, n int) int {
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 1 << (2 * m)) // 4 ^ m

        for j := 0; j < (1 << (2 * m)); j++ {
            dp[i][j] = -1
        }
    }

    validMasks := getValidMasks(m)

    return solve(n - 1, 0, m, validMasks, dp)
}

func getValidMasks(m int) []int {
    result := []int{}

    for i := 0; i < (1 << (2 * m)); i++ {
        if isValidMask(i, m) {
            result = append(result, i)
        }
    }

    return result
}

func isValidMask(mask, m int) bool {
    prev := 0
    for i := 0; i < m; i++ {
        val := (mask >> (2 * i)) & 3

        if val == prev || val == 0 {
            return false
        }

        prev = val
    }
    return true
}

func solve(row, mask, m int, validMasks []int, dp [][]int) int {
    if row <= -1 {
        return 1
    }
    if dp[row][mask] != -1 {
        return dp[row][mask]
    }

    result := 0
    for _, newMask := range validMasks {
        if !isAdjacentWithSameColor(mask, newMask, m) {
            result += solve(row - 1, newMask, m, validMasks, dp)
            result %= MOD
        }
    }

    dp[row][mask] = result
    return result
}

func isAdjacentWithSameColor(mask, newMask, m int) bool {
    for i := 0; i < m; i++ {
        prev := (mask >> (2 * i)) & 3
        cur := (newMask >> (2 * i)) & 3

        if cur == prev {
            return true
        }
    }
    return false
}
