func hasIncreasingSubarrays(nums []int, k int) bool {
    goodStartingIdx := map[int]bool{}

    curSize := 0
    last := nums[0] - 1
    for i := 0; i < len(nums); i++ {
        if nums[i] > last {
            curSize++
        } else {
            curSize = 1 
        }
        last = nums[i]

        if curSize >= k {
            goodStartingIdx[i - k + 1] = true
        }
    }

    for i := 0; i < len(nums); i++ {
        if goodStartingIdx[i] && goodStartingIdx[i + k] {
            return true
        }
    }

    return false
}
