func longestSquareStreak(nums []int) int {
    numsMap := map[int]bool{}
    for _, num := range nums {
        numsMap[num] = true
    }

    result := 0
    memo := map[int]int{}
    for _, num := range nums {
        result = max(result, solve(num, numsMap, memo))
    }

    if result < 2 {
        return -1
    }
    return result
}

func solve(num int, numsMap map[int]bool, memo map[int]int) int {
    if _, ok := memo[num]; ok {
        return memo[num]
    }

    if _, ok := numsMap[num]; !ok {
        return 0
    }

    // allow overflow because numsMap >= 2
    return 1 + solve(num * num, numsMap, memo)
}
