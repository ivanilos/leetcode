func combinationSum3(k int, n int) [][]int {
    result := [][]int{}

    cur := []int{}
    solve(1, k, n, cur, &result)

    return result
}

func solve(next, k, n int, cur []int, result *[][]int) {
    if n == 0 && k == 0 {
        cpy := make([]int, len(cur))
        copy(cpy, cur)
        *result = append(*result, cpy)
        return
    }
    if next > 9 {
        return
    }
    if k < 0 {
        return
    }
    if n < 0 {
        return
    }

    // take number
    cur = append(cur, next)
    solve(next + 1, k - 1, n - next, cur, result)
    cur = cur[0 : len(cur) - 1]

    // dont take
    solve(next + 1, k, n, cur, result)
}
