func sortByBits(arr []int) []int {
    slices.SortFunc(arr, func(a int, b int) int {
        popA := popcount(a)
        popB := popcount(b)

        if popA == popB {
            return a - b
        } else {
            return popA - popB
        }
    })

    return arr
}

func popcount(a int) int {
    result := 0
    for a > 0 {
        result += a & 1
        a /= 2
    }
    return result
}
