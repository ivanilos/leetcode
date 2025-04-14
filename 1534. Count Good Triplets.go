func countGoodTriplets(arr []int, a int, b int, c int) int {
    n := len(arr)

    result := 0
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            for k := j + 1; k < n; k++ {
                if isGood(arr[i], arr[j], a) && isGood(arr[j], arr[k], b) && isGood(arr[i], arr[k], c) {
                    result++
                }
            }
        }
    }

    return result
}

func isGood(a, b, maxDiff int) bool {
    return abs(a - b) <= maxDiff
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
