func minimumAbsDifference(arr []int) [][]int {
    slices.Sort(arr)

    result := [][]int{}
    minDiff := arr[1] - arr[0]
    for i := 1; i < len(arr); i++ {
        if arr[i] - arr[i - 1] == minDiff {
            result = append(result, []int{arr[i - 1], arr[i]})
        } else if arr[i] - arr[i - 1] < minDiff {
            minDiff = arr[i] - arr[i - 1]
            result = [][]int{}
            result = append(result, []int{arr[i - 1], arr[i]})
        }
    }

    return result
}
