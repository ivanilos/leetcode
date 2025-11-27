const INF = int64(3e15)

func maxSubarraySum(nums []int, k int) int64 {
    prefSum := make([]int64, len(nums) + 1)
    prefSum[0] = 0
    for i, num := range nums {
        prefSum[i + 1] = prefSum[i] + int64(num)
    }

    result := -INF
    for i := 1; i <= k; i++ {
        result = max(result, solve(i, prefSum, k))
    }

    return result
}

func solve(start int, prefSum []int64, k int) int64 {
    result := -INF

    curSuf := -INF
    for i := start + k - 1; i < len(prefSum); i += k {
        curElem := prefSum[i] - prefSum[i - k]
        curSuf = max(curSuf + curElem, curElem)

        result = max(result, curSuf)
    }

    return result
}
