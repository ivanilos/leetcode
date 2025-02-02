func check(nums []int) bool {
    start := 0
    sz := len(nums)
    for i := 1; i < sz; i++ {
        if nums[i] < nums[i - 1] {
            start = i
            break
        }
    }

    for i := 0; i < sz - 1; i++ {
        if nums[(i + start) % sz] > nums[(i + 1 + start) % sz] {
            return false
        }
    }

    return true
}
