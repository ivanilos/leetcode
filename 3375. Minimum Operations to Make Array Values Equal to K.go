func minOperations(nums []int, k int) int {
    freq := map[int]int{}

    for _, num := range nums {
        if num < k {
            return -1
        }
        freq[num]++
    }

    result := 0
    
    for key, _ := range freq {
        if key != k {
            result++
        }
    }

    return result
}
