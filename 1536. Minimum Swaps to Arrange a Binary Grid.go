func minSwaps(grid [][]int) int {
    lastOneCol := make([]int, len(grid))

    for i, row := range grid {
        lastOneCol[i] = -1
        for j := len(row) - 1; j >= 0; j-- {
            if row[j] == 1 {
                lastOneCol[i] = j
                break
            }
        }
    }

    result := 0
    for i := 0; i < len(lastOneCol); i++ {
        chosen := -1
        for j := i; j < len(lastOneCol); j++ {
            if lastOneCol[j] <= i {
                chosen = j
                break
            }
        }

        if chosen == -1 {
            return -1
        }

        result += chosen - i
        for j := chosen; j > i; j-- {
            lastOneCol[j] = lastOneCol[j - 1]
        }
    }

    return result
}
