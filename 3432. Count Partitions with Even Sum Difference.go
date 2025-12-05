func countPartitions(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }

    result := 0
    prefSum := 0
    for i := 0; i < len(nums) - 1; i++ {
        prefSum += nums[i]
        sum -= nums[i]

        if (prefSum - sum) % 2 == 0 {
            result++
        }
    }

    return result
}
