func minMirrorPairDistance(nums []int) int {
    result := len(nums)

    numToIdx := map[int]int{}
    for i := len(nums) - 1; i >= 0; i-- {
        num := nums[i]
        if idx, ok := numToIdx[getMirror(num)]; ok {
            result = min(result, idx - i)
        }

        numToIdx[num] = i
    }

    if result == len(nums) {
        return -1
    }
    return result
}

func getMirror(num int) int {
    result := 0

    for num > 0 {
        digit := num % 10
        num /= 10

        result = 10 * result + digit
    }

    return result
}
