func resultsArray(nums []int, k int) []int {
    result := make([]int, len(nums) - k + 1)
    for i := 0; i < len(nums) - k + 1; i++ {
        result[i] = -1
    }

    curConsec := 1
    for i := 1; i < k; i++ {
        if nums[i] == 1 + nums[i - 1] {
            curConsec++
        } else {
            curConsec = 1
        }
        fmt.Println(i, curConsec)
    }
    if curConsec >= k {
        result[0] = nums[k - 1]
    }

    for i := k; i < len(nums); i++ {
        if nums[i] == 1 + nums[i - 1] {
            curConsec++
        } else {
            curConsec = 1
        }

        if curConsec >= k {
            result[i - k + 1] = nums[i]
        } 
    }

    return result
}
