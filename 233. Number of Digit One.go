func countDigitOne(n int) int {
    str := strconv.Itoa(n)
    dp := make([][]int, len(str))
    for i := 0; i < len(str); i++ {
        dp[i] = make([]int, 2)
        for j := 0; j < 2; j++ {
            dp[i][j] = -1
        }
    }

    result := solve(0, 0, str, dp)

    return result
}

func solve(pos, isSmaller int, num string, dp [][]int) int {
    if pos >= len(num) {
        return 0
    }
    if dp[pos][isSmaller] != -1 {
        return dp[pos][isSmaller]
    }

    dp[pos][isSmaller] = 0

    posDigit := int(num[pos] - '0')
    for i := 0; i <= 9; i++ {
        add := 0
        if i == 1 {
            add = 1
        }

        if isSmaller == 1 {
            dp[pos][isSmaller] += add * getPossibilities(pos + 1, 1, num) + solve(pos + 1, 1, num, dp)
        } else {
            if i < posDigit {
                dp[pos][isSmaller] += add * getPossibilities(pos + 1, 1, num) + solve(pos + 1, 1, num, dp)
            } else {
                dp[pos][isSmaller] += add * getPossibilities(pos + 1, 0, num) + solve(pos + 1, 0, num, dp)
                break
            }
        }
    }

    return dp[pos][isSmaller]
}

func getPossibilities(pos, isSmaller int, num string) int {
    result := 1
    if isSmaller == 1 {
        for i := pos; i < len(num); i++ {
            result *= 10
        }
    } else {
        result, _ = strconv.Atoi(num[pos : len(num)])
        result++
    }

    return result
}
