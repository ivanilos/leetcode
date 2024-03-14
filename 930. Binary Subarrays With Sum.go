func numSubarraysWithSum(nums []int, goal int) int {
    freq := map[int]int{}

    freq[0] = 1

    curSum := 0
    result := 0
    for _, num := range nums {
        curSum += num
        result += freq[curSum - goal]
        freq[curSum] += 1
    }

    return result
}
