func minimumDeletions(s string) int {
    prefixB := make([]int, len(s))
    suffixA := make([]int, len(s))

    prefixB[0] = getVal(s[0], 'b')
    for i := 1; i < len(s); i++ {
        prefixB[i] = prefixB[i - 1] + getVal(s[i], 'b')
    }

    suffixA[len(s) - 1] = getVal(s[len(s) - 1], 'a')
    for i := len(s) - 2; i >= 0; i-- {
        suffixA[i] = suffixA[i + 1] + getVal(s[i], 'a')
    }

    return solve(prefixB, suffixA)
}

func getVal(actual byte, target byte) int {
    if actual == target {
        return 1
    }
    return 0
}

func solve(prefixCost []int, suffixCost []int) int {
    result := min(suffixCost[0], prefixCost[len(prefixCost) - 1])

    for i := 0; i < len(prefixCost) - 1; i++ {
        result = min(result, prefixCost[i] + suffixCost[i + 1])
    }
    return result
}
