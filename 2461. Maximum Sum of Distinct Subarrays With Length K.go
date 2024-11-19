func maximumSubarraySum(nums []int, k int) int64 {
    freq := map[int]int{}

    result := int64(0)
    diffElems := 0
    curSum := int64(0)
    for i := 0; i < k; i++ {
        curSum += int64(nums[i])
        freq[nums[i]]++

        if freq[nums[i]] == 1 {
            diffElems++
        } else if freq[nums[i]] == 2{
            diffElems--
        }
    }

    if diffElems == k {
        result = curSum
    }

    for i := k; i < len(nums); i++ {
        curSum += int64(nums[i])
        freq[nums[i]]++

        if freq[nums[i]] == 1 {
            diffElems++
        } else if freq[nums[i]] == 2 {
            diffElems--
        }

        curSum -= int64(nums[i - k])
        freq[nums[i - k]]--
        if freq[nums[i - k]] == 1 {
            diffElems++
        } else if freq[nums[i - k]] == 0 {
            diffElems--
        }

        if diffElems == k {
            result = max(result, curSum)
        }
    }

    return result
}
