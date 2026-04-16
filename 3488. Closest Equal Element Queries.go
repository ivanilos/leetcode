func solveQueries(nums []int, queries []int) []int {
    valToPos := map[int][]int{}
    indexInValToPos := map[int]int{}

    for i, num := range nums {
        indexInValToPos[i] = len(valToPos[num])
        valToPos[num] = append(valToPos[num], i)    
    }

    result := make([]int, len(queries))
    for i, queryIdx := range queries {
        queryVal := nums[queryIdx]

        if len(valToPos[queryVal]) == 1 {
            result[i] = -1
        } else {
            result[i] = getResult(queryIdx, queryVal, valToPos, indexInValToPos, nums)
        }
    }

    return result
}

func getResult(queryIdx, queryVal int, valToPos map[int][]int, indexInValToPos map[int]int, nums []int) int {
    sz := len(valToPos[queryVal])

    curIdx := indexInValToPos[queryIdx]
    previousIdx := (curIdx - 1 + sz) % sz
    nextIdx := (curIdx + 1) % sz

    prevDist := abs(valToPos[queryVal][curIdx] - valToPos[queryVal][previousIdx])
    goingToPrev := min(len(nums) - prevDist, prevDist)

    nextDist := abs(valToPos[queryVal][curIdx] - valToPos[queryVal][nextIdx])
    goingToNext := min(len(nums) - nextDist, nextDist)

    return min(goingToPrev, goingToNext)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
