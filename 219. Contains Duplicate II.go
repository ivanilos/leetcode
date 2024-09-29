func containsNearbyDuplicate(nums []int, k int) bool {
    mapa := map[int]int{}
    for i := 0; i <= k && i < len(nums); i++ {
        mapa[nums[i]]++
        if mapa[nums[i]] >= 2 {
            return true
        }
    }

    for i := k + 1; i < len(nums); i++ {
        mapa[nums[i - (k + 1)]]--
        mapa[nums[i]]++
        if mapa[nums[i]] >= 2 {
            return true
        }
    }
    return false
}
