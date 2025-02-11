func longestOnes(nums []int, k int) int {
    left := 0
    result := 0
    flippedZeroes := 0
    curOnes := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 {
            flippedZeroes++
        } else {
            curOnes++
        }

        for left <= right && flippedZeroes > k {
            if nums[left] == 0 {
                flippedZeroes--
            } else {
                curOnes--
            }
            left++
        }

        result = max(result, curOnes + flippedZeroes)
    }

    return result
}
