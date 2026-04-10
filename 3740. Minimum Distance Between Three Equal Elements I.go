const INF = int(1e9)

func minimumDistance(nums []int) int {
    result := INF

    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                if nums[i] == nums[j] && nums[j] == nums[k] {
                    result = min(result, (j - i) + (k - j) + (k - i))
                }
            }
        }
    }

    if result == INF {
        return -1
    }
    return result
}
