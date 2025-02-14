func isMatch(s string, p string) bool {
    dp := make([][]int, len(s) + 1)

    for i := 0; i <= len(s); i++ {
        dp[i] = make([]int, len(p) + 1)
        for j := 0; j <= len(p); j++ {
            dp[i][j] = -1
        }
    }

    result := solve(0, 0, s, p, dp)

    if result == 0 {
        return false
    }
    return true
}

func solve(posS, posP int, s, p string, dp [][]int) int {
    val, ok := baseCondition(posS, posP, s, p, dp)
    if ok {
        return val
    }

    dp[posS][posP] = 0
    if posP + 1 < len(p) && p[posP + 1] == '*' {
        dp[posS][posP] = max(dp[posS][posP], solve(posS, posP + 2, s, p, dp))
        if posS < len(s) && (s[posS] == p[posP] || p[posP] == '.') {
            dp[posS][posP] = max(dp[posS][posP], solve(posS + 1, posP, s, p, dp))
        }
    } else {
        if posS < len(s) && (s[posS] == p[posP] || p[posP] == '.') {
            dp[posS][posP] = max(dp[posS][posP], solve(posS + 1, posP + 1, s, p, dp))
        }
    }

    return dp[posS][posP]
}

func baseCondition(posS, posP int, s, p string, dp [][]int) (int, bool) {
    if posP >= len(p) {
        if posS >= len(s) {
            return 1, true
        }
        return 0, true
    }

    if dp[posS][posP] != -1 {
        return dp[posS][posP], true
    }

    return -1, false
}
