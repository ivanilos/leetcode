func maxChunksToSorted(arr []int) int {
    sortedArr := make([]int, len(arr))
    copy(sortedArr, arr)

    slices.Sort(sortedArr)

    result := 0
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {
        freq[arr[i]]++
        if freq[arr[i]] == 0 {
            delete(freq, arr[i])
        }

        freq[sortedArr[i]]--
        if freq[sortedArr[i]] == 0 {
            delete(freq, sortedArr[i])
        }

        if len(freq) == 0 {
            result++
        }
    }

    return result
}
