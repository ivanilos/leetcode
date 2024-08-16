func maxDistance(arrays [][]int) int {
    curMin := getMin(arrays[0])
    curMax := getMax(arrays[0])

    result := 0
    for i := 1; i < len(arrays); i++ {
        newMin := getMin(arrays[i])
        newMax := getMax(arrays[i])

        result = max(result, abs(newMax - curMin))
        result = max(result, abs(curMax - newMin))
        curMin = min(curMin, newMin)
        curMax = max(curMax, newMax)
    }
    return result
}

func getMin(array []int) int {
    result := array[0]
    for i := 1; i < len(array); i++ {
        result = min(result, array[i])
    }
    return result
}

func getMax(array []int) int {
    result := array[0]
    for i := 1; i < len(array); i++ {
        result = max(result, array[i])
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
