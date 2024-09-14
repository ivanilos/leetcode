func longestSubarray(nums []int) int {
    maxi := 0
    result := 0
    curResult := 0

    for _, num := range nums {
        if num > maxi {
            curResult = 1
            result = 1
            maxi = num
        } else if num == maxi {
            curResult++
        } else {
            curResult = 0
        }
        result = max(result, curResult)
    }
    return result
}
