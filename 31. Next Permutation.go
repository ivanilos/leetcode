func nextPermutation(nums []int)  {
    for i := len(nums) - 2; i >= 0; i-- {
        if nums[i] < nums[i + 1] {
            for j := len(nums) - 1; j >= 0; j-- {
                if nums[j] > nums[i] {
                    nums[i], nums[j] = nums[j], nums[i]
                    slices.Reverse(nums[i + 1 :])
                    return
                }
            }
        }
    }

    slices.Reverse(nums)
}
