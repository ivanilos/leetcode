func minimumEffort(tasks [][]int) int {
    slices.SortFunc(tasks, func(a, b []int) int {
        v1 := a[1] - a[0]
        v2 := b[1] - b[0]

        return v2 - v1
    })

    curEnergy := 0
    maxEnergy := 0
    for _, task := range tasks {
        need := max(task[0], task[1])
        if curEnergy < need {
            maxEnergy += need - curEnergy
            curEnergy += need - curEnergy
        }
        curEnergy -= task[0]
    }

    return maxEnergy
}
