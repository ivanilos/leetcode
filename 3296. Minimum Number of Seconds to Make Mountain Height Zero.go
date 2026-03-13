func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
    low := int64(0)
    high := int64(1e16)

    best := high

    prefSum := getPrefSum(mountainHeight)

    for low <= high {
        mid := (low + high) / 2

        if can(mountainHeight, workerTimes, prefSum, mid) {
            best = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    return best
}

func getPrefSum(n int) []int64 {
    prefSum := make([]int64, n + 1)

    prefSum[0] = 0
    for i := 1; i <= n; i++ {
        prefSum[i] = int64(i) + prefSum[i - 1]
    }
    return prefSum
}

func can(mountainHeight int, workerTimes []int, prefSum []int64, maxTime int64) bool {
    reduced := int64(0)
    for _, workerTime := range workerTimes {
        reduced += search(maxTime / int64(workerTime), prefSum)
    }

    return reduced >= int64(mountainHeight)
}

func search(val int64, v []int64) int64 {
    low := int64(0);
    high := int64(len(v) - 1)

    best := int64(0)
    for low <= high {
        mid := (low + high) / 2

        if v[mid] <= val {
            best = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return best
}
