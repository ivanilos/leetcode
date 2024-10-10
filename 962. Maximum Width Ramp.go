func maxWidthRamp(nums []int) int {
    leftPossibleNums := [][]int{}

    result := 0
    for rightIdx := 0; rightIdx < len(nums); rightIdx++ {
        leftIdx := searchLeftmostSmaller(leftPossibleNums, nums[rightIdx])

        if leftIdx != -1 {
            result = max(result, rightIdx - leftIdx)
        }

        sz := len(leftPossibleNums)
        if sz == 0 || nums[rightIdx] < leftPossibleNums[sz - 1][0] {
            leftPossibleNums = append(leftPossibleNums, []int{nums[rightIdx], rightIdx})
        }
    }
    return result
}

func searchLeftmostSmaller(v [][]int, val int) int {
    left := 0
    right := len(v) - 1
    best := -1

    for left <= right {
        mid := (left + right) / 2

        if v[mid][0] > val {
            left = mid + 1
        } else {
            right = mid - 1
            best = mid
        }
    }
    if best != -1 {
        return v[best][1]
    }
    return -1
}
