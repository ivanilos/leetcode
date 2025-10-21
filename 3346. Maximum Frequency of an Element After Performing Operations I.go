func maxFrequency(nums []int, k int, numOperations int) int {
    freq := map[int]int{}

    for _, num := range nums {
        freq[num]++
    }

    numFreq := [][]int{}
    for key, val := range freq {
        numFreq = append(numFreq, []int{key, val})
    }

    slices.SortFunc(numFreq, func(a []int, b []int) int {
        return a[0] - b[0]
    })


    result := max(getFromExistingValue(numFreq, numOperations, k), getFromNewValue(numFreq, numOperations, k))

    return result
}

func getFromExistingValue(numFreq [][]int, numOperations, k int) int {
    l := 0
    r := -1
    extraFreq := 0

    result := 0
    for i := 0; i < len(numFreq); i++ {
        for numFreq[l][0] + k < numFreq[i][0] {
            extraFreq -= numFreq[l][1]
            l++
        }
        for r + 1 < len(numFreq) && numFreq[r + 1][0] - k <= numFreq[i][0] {
            extraFreq += numFreq[r + 1][1]
            r++
        }

        result = max(result, numFreq[i][1] + min(extraFreq - numFreq[i][1], numOperations))
    }

    return result
}

func getFromNewValue(numFreq [][]int, numOperations, k int) int {
    r := -1
    extraFreq := 0

    result := 0
    for i := 0; i < len(numFreq); i++ {
        for r + 1 < len(numFreq) && numFreq[r + 1][0] - 2 * k <= numFreq[i][0] {
            extraFreq += numFreq[r + 1][1]
            r++
        }

        result = max(result, min(extraFreq, numOperations))

        extraFreq -= numFreq[i][1]
    }

    return result
}
