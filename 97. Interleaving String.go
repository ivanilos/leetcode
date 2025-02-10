func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) {
        return false
    }
    
    dp := make([][]bool, 2)
    for i := 0; i < 2; i++ {
        dp[i] = make([]bool, len(s2) + 1)
    }

    dp[0][len(s2)] = true
    for j := len(s2) - 1; j >= 0; j-- {
        if s2[j] == s3[j + len(s1)] {
            dp[0][j] = dp[0][j + 1]
        }
    }

    nextDpTableRow := 1
    for row := len(s1) - 1; row >= 0; row-- {
        if s1[row] == s3[row + len(s2)] {
            dp[nextDpTableRow][len(s2)] = true
        } else {
            dp[nextDpTableRow][len(s2)] = false
        }

        for j := len(s2) - 1; j >= 0; j-- {
            dp[nextDpTableRow][j] = false
            if s2[j] == s3[row + j] {
                dp[nextDpTableRow][j] = dp[nextDpTableRow][j] || dp[nextDpTableRow][j + 1]
            }
            if s1[row] == s3[row + j] {
                dp[nextDpTableRow][j] = dp[nextDpTableRow][j] || dp[1 - nextDpTableRow][j]
            }
        }

        nextDpTableRow = 1 - nextDpTableRow
    }

    return dp[1 - nextDpTableRow][0]
}
