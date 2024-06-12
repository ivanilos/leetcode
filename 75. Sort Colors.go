func sortColors(nums []int)  {
    nextPosRed := 0
    nextPosBlue := len(nums) - 1
    for i := 0; i < len(nums); i++ {
        for nums[i] == 0 && nextPosRed < i || nums[i] == 2 && nextPosBlue > i {
            if nums[i] == 0 && nextPosRed < i {
                nums[nextPosRed], nums[i] = nums[i], nums[nextPosRed]
                nextPosRed++
            }
            if nums[i] == 2 && nextPosBlue > i {
                nums[nextPosBlue], nums[i] = nums[i], nums[nextPosBlue]
                nextPosBlue--
            }
        }
    }
}
