func triangleNumber(nums []int) int {
    slices.Sort(nums)

    result := 0
    for i, c := range nums {
        result += countPairsWithSumLowerThan(c, nums[: i])
    }

    return result
}

func countPairsWithSumLowerThan(c int, nums []int) int {
    aIndex := 0
    result := 0
    for i, b := range nums {
        for aIndex + 1 < i && nums[aIndex] + b <= c {
            aIndex++
        }
        for aIndex - 1 >= 0 && nums[aIndex - 1] + b > c {
            aIndex--
        }

        if aIndex < i && nums[aIndex] + b > c {
            result += i - aIndex
        }
    }

    return result
}
