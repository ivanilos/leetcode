const INF int = int(1e9)

func smallestRange(nums [][]int) []int {
    result := []int{-INF, INF}

    startIdx := make([]int, len(nums))
    iterations := getTotalIterations(nums)

    for it := 0; it < iterations; it++ {
        mini, maxi := getValues(startIdx, nums)

        if maxi - mini < result[1] - result[0] {
            result = []int{mini, maxi}
        }

        for i := 0; i < len(nums); i++ {
            if nums[i][startIdx[i]] == mini && startIdx[i] + 1 < len(nums[i]) {
                startIdx[i]++
                break
            }
        }
    }

    return result
}

func getTotalIterations(nums [][]int) int {
    result := 0

    for i := 0; i < len(nums); i++ {
        result += len(nums[i]) - 1
    }
    return result + 1
}

func getValues(idxs []int, nums [][]int) (int, int) {
    mini := INF
    maxi := -INF

    for i := 0; i < len(nums); i++ {
        mini = min(mini, nums[i][idxs[i]])
        maxi = max(maxi, nums[i][idxs[i]])
    }
    return mini, maxi
}
