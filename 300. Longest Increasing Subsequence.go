func lengthOfLIS(nums []int) int {
    lis := []int{}

    for _, num := range nums {
        pos := bsearch(lis, num)

        if pos == len(lis) - 1 {
            lis = append(lis, num)
        } else {
            lis[pos + 1] = num
        }
    }
    return len(lis)
}

func bsearch(lis []int, num int) int {
    low := 0
    high := len(lis) - 1
    best := -1

    for low <= high {
        mid := (low + high) / 2

        if lis[mid] >= num {
            high = mid - 1
        } else {
            best = mid
            low = mid + 1
        }
    }
    return best
}
