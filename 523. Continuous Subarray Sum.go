func checkSubarraySum(nums []int, k int) bool {
    prefix := 0
    prefixMap := map[int]int{}
    prefixMap[0] = -1

    for i := 0; i < len(nums); i++ {
        num := nums[i]
        prefix = (prefix + num) % k
        if last, ok := prefixMap[prefix]; ok {
            if i - last > 1 {
                return true
            }
        } else {
            prefixMap[prefix] = i
        }
    }
    return false
}
