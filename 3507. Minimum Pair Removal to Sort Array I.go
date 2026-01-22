const INF = int(1e9)

func minimumPairRemoval(nums []int) int {
    result := 0

    for !nonDecreasing(nums) {
        result++

        leastSum := INF
        leastSumIdx := -1
        for i := 1; i < len(nums); i++ {
            if nums[i] + nums[i - 1] < leastSum {
                leastSum = nums[i] + nums[i - 1]
                leastSumIdx = i - 1
            }
        }

        nums[leastSumIdx] += nums[leastSumIdx + 1]
        copy(nums[leastSumIdx + 1 : ], nums[leastSumIdx + 2 :])
        nums = nums[: len(nums) - 1]
    }

    return result
}

func nonDecreasing(nums []int) bool {
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i - 1] {
            return false
        }
    }
    return true
}
