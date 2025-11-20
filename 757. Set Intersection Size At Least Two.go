func intersectionSizeTwo(intervals [][]int) int {
    slices.SortFunc(intervals, func(a []int, b []int) int{
        return a[1] - b[1]
    })

    result := map[int]bool{}

    for _, interval := range intervals {
        alreadyIn := 0

        for key, _ := range result {
            if interval[0] <= key && key <= interval[1] {
                alreadyIn++
            }
        }

        nextNum := interval[1]
        for alreadyIn < 2 {
            if _, ok := result[nextNum]; !ok {
                alreadyIn++
                result[nextNum] = true
            }
            nextNum--
        }
    }

    return len(result)
}
