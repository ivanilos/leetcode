func longestConsecutive(nums []int) int {
    mapa := map[int]bool{}
    for _, num := range nums {
        mapa[num] = true
    }

    dp := map[int]int{}

    result := 0
    for _, val := range nums {
        result = max(result, solve(val, mapa, dp))
    }
    return result
}

func solve(val int, mapa map[int]bool, dp map[int]int) int {
    if _, ok := mapa[val]; !ok {
        return 0
    }
    if _, ok := dp[val]; ok {
        return dp[val]
    }

    dp[val] = 1 + solve(val + 1, mapa, dp)
    return dp[val]
}
