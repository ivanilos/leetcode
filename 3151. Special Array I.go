func isArraySpecial(nums []int) bool {
    lastParity := 1 - nums[0] % 2

    for _, num := range nums {
        if num % 2 == lastParity {
            return false
        }
        lastParity = num % 2
    }
    return true
}
