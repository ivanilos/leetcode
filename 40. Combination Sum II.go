func combinationSum2(candidates []int, target int) [][]int {
    dp := make([][][][]int, len(candidates))

    sort.Ints(candidates)

    for i := 0; i < len(candidates); i++ {
        dp[i] = make([][][]int, target + 1)
    }

    for i := 0; i < len(candidates); i++ {
        dp[i][0] = [][]int{[]int{}}
    }
    if (candidates[0] <= target) {
        dp[0][candidates[0]] = [][]int{[]int{candidates[0]}}
    }

    for i := 1; i < len(candidates); i++ {
        for j := 1; j <= target; j++ {
            dp[i][j] = dp[i - 1][j]
            if j - candidates[i] >= 0 {
                newCombinations := add(candidates[i], dp[i - 1][j - candidates[i]])
                dp[i][j] = combine(dp[i][j], newCombinations)
            }
        }
    }

    return dp[len(candidates) - 1][target]
}

func add(val int, combinations [][]int) [][]int {
    result := [][]int{}
    for i := 0; i < len(combinations); i++ {
        cpy := make([]int, len(combinations[i]))
        copy(cpy, combinations[i])

        result = append(result, append(cpy, val))
    }
    return result
}

func combine(a [][]int, b [][]int) [][]int {
    mapa := map[int][]int{}

    for i := 0; i < len(a); i++ {
        mapa[getHash(a[i])] = a[i]
    }
    for i := 0; i < len(b); i++ {
        mapa[getHash(b[i])] = b[i]
    }

    result := [][]int{}
    for _, val := range(mapa) {
        result = append(result, val)
    }
    return result
}

func getHash(a []int) int {
    mod := int(1e9 + 9)
    result := 0
    for i := 0; i < len(a); i++ {
        result = (result * 53 + a[i]) % mod
    }
    return result
}
