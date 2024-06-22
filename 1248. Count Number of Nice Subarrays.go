func numberOfSubarrays(nums []int, k int) int {
    oddInPrefix := 0
    oddInPrefixMap := map[int]int{}

    oddInPrefixMap[0] = 1
    result := 0
    for _, num := range(nums) {
        if num % 2 == 1 {
            oddInPrefix++
        }

        result += oddInPrefixMap[oddInPrefix - k]
        oddInPrefixMap[oddInPrefix] += 1
    }
    return result
}
