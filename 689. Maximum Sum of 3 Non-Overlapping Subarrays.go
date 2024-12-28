func maxSumOfThreeSubarrays(nums []int, k int) []int {
    prefSizeK := make([]int, len(nums) - k + 1)

    for i := 0; i < k; i++ {
        prefSizeK[0] += nums[i]
    }

    for i := 1; i < len(nums) - k + 1; i++ {
        prefSizeK[i] = prefSizeK[i - 1] + nums[i + k - 1] - nums[i - 1]
    }

    dp, dpIdx := initDP(prefSizeK)

    solve(0, 3, k, prefSizeK, dp, dpIdx)

    result := []int{}
    pos := 0
    need := 3

    for need > 0 {
        result = append(result, dpIdx[pos][need])
        pos = dpIdx[pos][need] + k
        need--
    }

    return result
}

func initDP(v []int) ([][]int, [][]int) {
    dp := make([][]int, len(v))
    for i := 0; i < len(v); i++ {
        dp[i] = make([]int, 4)

        for j := 0; j < 4; j++ {
            dp[i][j] = -int(1e9)
        }
    }

    dpIdx := make([][]int, len(v))
    for i := 0; i < len(v); i++ {
        dpIdx[i] = make([]int, 4) 
    }

    return dp, dpIdx
}

func solve(pos, need, k int, v []int, dp, dpIdx [][]int) int {
    if pos >= len(v) || need == 0 {
        if need == 0 {
            return 0
        }
        return -int(1e9)
    }

    if dp[pos][need] >= 0 {
        return dp[pos][need]
    } 

    // take val
    result := v[pos] + solve(pos + k, need - 1, k, v, dp, dpIdx)

    // dont take val
    result2 := solve(pos + 1, need, k, v, dp, dpIdx)

    if result >= result2 {
        dp[pos][need] = result
        dpIdx[pos][need] = pos 
    } else if pos + 1 < len(v) {
        dp[pos][need] = result2
        dpIdx[pos][need] = dpIdx[pos + 1][need]
    }

    return dp[pos][need]
}
