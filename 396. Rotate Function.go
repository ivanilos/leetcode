func maxRotateFunction(nums []int) int {
    prefixSum, suffixSum := getPrefixAndSuffixSums(nums)

    curResult := 0
    for i := 0; i < len(nums); i++ {
        curResult += nums[i] * i
    }

    result := curResult
    for i := 1; i < len(nums); i++ {
        curResult -= suffixSum[i]
        curResult += (len(nums) - 1) * nums[i - 1]
        
        if i - 2 >= 0 {
            curResult -= prefixSum[i - 2]
        }

        result = max(result, curResult)
    }

    return result
}

func getPrefixAndSuffixSums(nums []int) ([]int, []int) {
    prefixSum := make([]int, len(nums))
    prefixSum[0] = nums[0]

    for i := 1; i < len(nums); i++ {
        prefixSum[i] = nums[i] + prefixSum[i - 1]
    }

    suffixSum := make([]int, len(nums))
    suffixSum[len(nums) - 1] = nums[len(nums) - 1]

    for i := len(nums) - 2; i >= 0; i-- {
        suffixSum[i] = nums[i] + suffixSum[i + 1]
    }

    return prefixSum, suffixSum
}
