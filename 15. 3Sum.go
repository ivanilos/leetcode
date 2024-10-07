func threeSum(nums []int) [][]int {
    sort.Ints(nums)

    result := [][]int{}

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }

        leftIdx := i + 1
        rightIdx := len(nums) - 1

        for leftIdx < rightIdx {
            curSum := nums[i] + nums[leftIdx] + nums[rightIdx]
            if curSum < 0 {
                leftIdx++
            } else if curSum > 0 {
                rightIdx--
            } else {
                result = append(result, []int{nums[i], nums[leftIdx], nums[rightIdx]})

                leftIdx++
                for leftIdx < rightIdx && nums[leftIdx] == nums[leftIdx - 1] {
                    leftIdx++
                }
            }
        }
    }

    return result
}
