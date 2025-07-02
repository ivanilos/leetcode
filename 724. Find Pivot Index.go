func pivotIndex(nums []int) int {
    sum := getSum(nums)

    leftSum := 0
    for pivot := 0; pivot < len(nums); pivot++ {
        if leftSum == sum - leftSum - nums[pivot] {
            return pivot
        }

        leftSum += nums[pivot]
    }

    return -1
}

func getSum(nums []int) int {
    result := 0
    for _, num := range nums {
        result += num
    }

    return result
}
