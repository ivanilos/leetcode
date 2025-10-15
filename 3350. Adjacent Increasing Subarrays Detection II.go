func maxIncreasingSubarrays(nums []int) int {
    prevMax := 0
    curMax := 0
    result := 0
    
    last := nums[0] + 1
    for i := 0; i < len(nums); i++ {
        if nums[i] > last {
            curMax++
        } else {
            result = max(result, max(prevMax / 2, min(prevMax, curMax)))

            prevMax = curMax
            curMax = 1
        }
        last = nums[i]

    }
    result = max(result, max(prevMax / 2, min(prevMax, curMax)))
    result = max(result, curMax / 2)

    return result
}
