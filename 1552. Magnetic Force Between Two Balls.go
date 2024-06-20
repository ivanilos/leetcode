func maxDistance(position []int, m int) int {
    low := 0
    high := int(1e9)

    slices.Sort(position)

    result := low
    for low <= high {
        mid := (low + high) / 2

        last := position[0] - mid
        used := 0
        for i := 0; i < len(position); i++ {
            if position[i] - last >= mid {
                used++
                last = position[i]
            }
        }

        if used >= m {
            result = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return result
}
