func combinationSum(candidates []int, target int) [][]int {
    dp := make([][][][]int, len(candidates))
    for i := 0; i < len(candidates); i++ {
        dp[i] = make([][][]int, target + 1)
        for j := 0; j <= target; j++ {
            dp[i][j] = [][]int{}
        }
    }

    solve(0, target, candidates, dp)

    return dp[0][target]
}

func solve(pos int, target int, candidates []int, dp [][][][]int) [][]int {
    if target == 0 {
        return [][]int{
            []int{},
        }
    }

    if target < 0 || pos >= len(candidates) {
        return [][]int{}
    }

    if len(dp[pos][target]) > 0 {
        return dp[pos][target]
    }

    // dont choose elem
    without := solve(pos + 1, target, candidates, dp)

    // choose elem
    with := solve(pos, target - candidates[pos], candidates, dp)
    aux := slices.Clone(with)
    for i := 0; i < len(aux); i++ {
        aux[i] = append([]int{candidates[pos]}, with[i]...)
    }

    result := append(aux, without...)

    dp[pos][target] = result
    return result
}
