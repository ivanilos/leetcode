var INF = int(1e9)

func stoneGameII(piles []int) int {
    total := getTotal(piles)

    dp := make([][]int, len(piles) + 1)
    for i := 0; i < len(piles) + 1; i++ {
        dp[i] = make([]int, len(piles) + 1)
        for j := 0; j < len(piles) + 1; j++ {
            dp[i][j] = INF
        }
    }
    result := solve(0, 1, piles, dp)

    // A + B = total
    // A - B = result
    // A = (total + result) / 2
    return (result + total) / 2
}

func getTotal(piles []int) int {
    result := 0
    for _, pile := range piles {
        result += pile
    }
    return result
}

func solve(nextPile, M int, piles []int, dp [][]int) int {
    if nextPile >= len(piles) {
        return 0
    }
    if dp[nextPile][M] != INF {
        return dp[nextPile][M]
    }

    sum := piles[nextPile]
    result := sum - solve(nextPile + 1, M, piles, dp)

    for i := nextPile + 1; i < len(piles) && i < nextPile + 2 * M; i++ {
        sum += piles[i]
        nextM := max(M, i - nextPile + 1)
        nextM = min(len(piles), nextM)
        result = max(result, sum - solve(i + 1, nextM, piles, dp))
    }
    dp[nextPile][M] = result
    return result
}
