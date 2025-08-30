func minCut(s string) int {
    dp := make([]int, len(s) + 1)
    for i := 0; i < len(s); i++ {
        dp[i] = len(s)
    }

    isPalindrome := getIsPalindrome(s)

    dp[len(s)] = 0
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            if isPalindrome[i][j] {
                dp[i] = min(dp[i], 1 + dp[j + 1])
            }
        }
    }

    return dp[0] - 1
}

func getIsPalindrome(s string) [][]bool {
    isPalindrome := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        isPalindrome[i] = make([]bool, len(s))
    }

    for center := 0; center < len(s); center++ {
        for i, j := center, center; i >= 0 && j < len(s); i, j = i - 1, j + 1 {
            if s[i] == s[j] {
                isPalindrome[i][j] = true
            } else {
                break
            }
        }
    }

    for center := 0; center < len(s); center++ {
        for i, j := center, center + 1; i >= 0 && j < len(s); i, j = i - 1, j + 1 {
            if s[i] == s[j] {
                isPalindrome[i][j] = true
            } else {
                break
            }
        }
    }

    return isPalindrome
}
