func constructTransformedArray(nums []int) []int {
    sz := len(nums)
    result := make([]int, sz)

    for i := 0; i < sz; i++ {
        pos := (((i + nums[i]) % sz) + sz) % sz
        result[i] = nums[pos]
    }

    return result
}
