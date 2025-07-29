func smallestSubarrays(nums []int) []int {
    nextOn := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        nextOn[i] = make([]int, 31)
    }

    for i := len(nums) - 1; i >= 0; i-- {
        num := nums[i]

        for j := 0; j < 31; j++ {
            nextOn[i][j] = len(nums)
            if ((num >> j) & 1) == 1 {
                nextOn[i][j] = i
            } else if i + 1 < len(nums) {
                nextOn[i][j] = nextOn[i + 1][j]
            }
        }
    }


    result := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        result[i] = 1
        for j := 0; j < 31; j++ {
            if nextOn[i][j] < len(nums) {
                result[i] = max(result[i], nextOn[i][j] - i + 1)
            }
        }
    }

    return result
}
