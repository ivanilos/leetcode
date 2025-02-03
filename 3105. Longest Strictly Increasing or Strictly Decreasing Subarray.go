func longestMonotonicSubarray(nums []int) int {
    return max(longestIncreasing(nums), longestDecreasing(nums))
}

func longestIncreasing(nums []int) int {
    result := 0

    cur := 0
    last := nums[0] - 1
    for _, num := range nums {
        if num > last {
            cur++
        } else {
            cur = 1
        }
        last = num

        result = max(result, cur)
    }
    return result
}

func longestDecreasing(nums []int) int {
    result := 0

    cur := 0
    last := nums[0] + 1
    for _, num := range nums {
        if num < last {
            cur++
        } else {
            cur = 1
        }
        last = num

        result = max(result, cur)
    }
    return result
}
