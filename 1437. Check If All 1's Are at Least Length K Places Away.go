func kLengthApart(nums []int, k int) bool {
    lastIdx := -k - 1
    for i, num := range nums {
        if num == 1 {
            if i - lastIdx - 1 < k {
                return false
            }
            lastIdx = i
        }
    }

    return true
}
