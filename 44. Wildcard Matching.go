func isMatch(s string, p string) bool {
    dp := make([][]int, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]int, len(p))
        for j := 0; j < len(p); j++ {
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
    if posS >= len(s) && posP >= len(p) {
        return 1
    }
    if posP >= len(p) {
        return 0
    }
    if posS >= len(s) {
        for i := posP; i < len(p); i++ {
            if p[i] != '*' {
                return 0
            }
        }
        return 1
    }

    if dp[posS][posP] != -1 {
        return dp[posS][posP]
    }

    dp[posS][posP] = 0
    if s[posS] == p[posP] || p[posP] == '?' {
        dp[posS][posP] = solve(posS + 1, posP + 1, s, p, dp)
    }
    if p[posP] == '*' {
        dp[posS][posP] = dp[posS][posP] | solve(posS, posP + 1, s, p, dp)
        dp[posS][posP] = dp[posS][posP] | solve(posS + 1, posP, s, p, dp)
    }

    return dp[posS][posP]
}
