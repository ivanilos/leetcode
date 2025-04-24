func countCompleteSubarrays(nums []int) int {
    const MAX_NUM = 2000
    freq := make([]int, MAX_NUM + 1)

    distinctCount := getDistinctCount(nums)

    left := 0
    result := 0
    curDistinct := 0
    for right := 0; right < len(nums); right++ {
        freq[nums[right]]++

        if freq[nums[right]] == 1 {
           curDistinct++ 
        }

        for left <= right && curDistinct == distinctCount {
            result += len(nums) - right

            freq[nums[left]]--
            if freq[nums[left]] == 0 {
                curDistinct--
            }
            left++
        }
    }

    return result
}

func getDistinctCount(nums []int) int {
    result := map[int]bool{}

    for _, num := range nums {
        result[num] = true
    }
    return len(result)
}
