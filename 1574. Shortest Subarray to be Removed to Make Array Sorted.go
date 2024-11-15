func findLengthOfShortestSubarray(arr []int) int {
    maxLen := max(solvePrefix(arr), solveSuffix(arr))

    return len(arr) - maxLen
}

func solvePrefix(arr []int) int {
    left := 0
    for i := 1; i < len(arr); i++ {
        if arr[i] >= arr[i - 1] {
            left++
        } else {
            break;
        }
    }

    result := left + 1
    right := len(arr)
    lastRightVal := arr[right - 1]

    for l := left; l >= 0; l-- {
        for right - 1 > l && arr[right - 1] >= arr[l] && arr[right - 1] <= lastRightVal {
            right--
            lastRightVal = arr[right]
        }

        result = max(result, l + 1 + len(arr) - right)
    }
    return result
}

func solveSuffix(arr []int) int {
    right := len(arr) - 1
    for i := len(arr) - 2; i >= 0; i-- {
        if arr[i] <= arr[i + 1] {
            right--
        } else {
            break;
        }
    }

    result := len(arr) - right
    left := -1
    lastLeftVal := arr[0]

    for r := right; r < len(arr); r++ {


        for left + 1 < r && arr[left + 1] <= arr[r] && arr[left + 1] >= lastLeftVal {
            left++
            lastLeftVal = arr[left]
        }
        result = max(result, left + 1 + len(arr) - r)
    }
    return result
}
