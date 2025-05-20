func isZeroArray(nums []int, queries [][]int) bool {
    queryPrefSum := make([]int, len(nums) + 1)

    for _, query := range queries {
        l := query[0]
        r := query[1]

        queryPrefSum[l]++
        queryPrefSum[r + 1]--
    }

    curSum := 0
    for i := 0; i < len(nums); i++ {
        curSum += queryPrefSum[i]

        if curSum < nums[i] {
            return false
        }
    } 

    return true
}
