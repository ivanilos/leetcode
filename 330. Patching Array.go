func minPatches(nums []int, n int) int {
    result := 0
    lastMade := 0

    for i := 0; i < len(nums) && lastMade < n; i++ {
        for nums[i] > lastMade + 1 && lastMade < n {
            // patch with lastMade + 1
            result++
            lastMade = lastMade + (lastMade + 1)
        }
        lastMade += nums[i]
        fmt.Println(i, nums[i], lastMade, result)
    }

    for lastMade < n {
        // patch with lastMade + 1
        result++
        lastMade = lastMade + (lastMade + 1)
    }
    return result
}
