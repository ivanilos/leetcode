func tupleSameProduct(nums []int) int {
    freq := map[int]int{}

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            mult := nums[i] * nums[j]
            freq[mult]++
        }
    }

    result := 0
    for _, val := range freq {
        result += (val * (val - 1)) / 2
    }
    return result * 8
}
