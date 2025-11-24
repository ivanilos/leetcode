func prefixesDivBy5(nums []int) []bool {
    result := make([]bool, len(nums))

    curMod := 0
    for i := 0; i < len(nums); i++ {
        curMod = (2 * curMod + nums[i]) % 5

        if curMod == 0 {
            result[i] = true
        }
    }

    return result
}
