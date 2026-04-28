const INF = int(1e9)

func minOperations(grid [][]int, x int) int {
    nums := []int{}

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            nums = append(nums, grid[i][j])
        }
    }

    slices.Sort(nums)

    rightSum := 0
    residue := map[int]bool{}
    for _, val := range nums {
        residue[val % x] = true
        rightSum += val
    }

    if len(residue) > 1 {
        return -1
    }

    result := INF
    n := len(nums)

    leftSum := 0
    for i, val := range nums {
        curResult := rightSum - (n - i) * val + (val * i) - leftSum
        result = min(result, curResult / x)

        rightSum -= val
        leftSum += val
    }

    return result
}
