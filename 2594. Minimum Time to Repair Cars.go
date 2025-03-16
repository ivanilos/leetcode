func repairCars(ranks []int, cars int) int64 {
    low := int64(1)
    high := int64(1e16)

    result := int64(0)
    for low <= high {
        mid := (low + high) / 2

        if canRepairAll(ranks, cars, mid) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}

func canRepairAll(ranks []int, cars int, maxTime int64) bool {
    total := int64(0)

    for _, rank := range ranks {
        q := float64(maxTime / int64(rank))
        total += int64(math.Sqrt(q))
    }
    return total >= int64(cars)
}
