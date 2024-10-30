func minimumMountainRemovals(nums []int) int {
    LISMap := getLISMap(nums)
    reversedNums := getReverse(nums)
    revLisMap := getLISMap(reversedNums)

    result := 0
    for i := 0; i < len(nums); i++ {
        if LISMap[i] > 1 && revLisMap[len(nums) - i - 1] > 1 {
            result = max(result, -1 + LISMap[i] + revLisMap[len(nums) - i - 1])
        }
    }

    if result >= 3 {
        return len(nums) - result
    }
    return -1
}

func getLISMap(nums []int) map[int]int {
    result := map[int]int{}

    sz := []int{}

    for i := 0; i < len(nums); i++ {
        pos := bsearch(nums[i], sz)

        if pos == len(sz) - 1 {
            result[i] = len(sz) + 1
            sz = append(sz, nums[i])
        } else if sz[pos + 1] >= nums[i] {
            result[i] = (pos + 1) + 1
            sz[pos + 1] = nums[i]
        }
    }
    return result
}

func bsearch(val int, v []int) int {
    low := 0
    high := len(v) - 1
    best := -1

    for low <= high {
        mid := (low + high) / 2

        if (v[mid] < val) {
            low = mid + 1
            best = mid
        } else {
            high = mid - 1
        }
    }
    return best
}

func getReverse(nums []int) []int {
    result := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        result[i] = nums[len(nums) - 1 - i]
    }
    return result
}
