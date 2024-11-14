func minimizedMaximum(n int, quantities []int) int {
    low := 1
    high := slices.Max(quantities)


    result := high
    for low <= high {
        mid := (low + high) / 2

        neededStores := 0
        for i := 0; i < len(quantities); i++ {
            neededStores += (quantities[i] + mid - 1) / mid
        }

        if neededStores <= n {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}
