func partitionArray(nums []int, k int) int {
    result := 0

    slices.Sort(nums)

    last := nums[0]
    for _, num := range nums {
        if num - last > k {
            result++
            last = num
        }
    }

    return result + 1
}
