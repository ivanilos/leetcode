func dailyTemperatures(temperatures []int) []int {
    nextLarger := getNextLarger(temperatures)

    return getResult(temperatures, nextLarger)
}

func getNextLarger(temperatures []int) []int {
    nextLarger := make([]int, len(temperatures))

    nextLarger[len(temperatures) - 1] = len(temperatures)
    for i := len(temperatures) - 2; i >= 0; i-- {
        nextLarger[i] = i + 1
        for nextLarger[i] < len(temperatures) && temperatures[nextLarger[i]] <= temperatures[i] {
            nextLarger[i] = nextLarger[nextLarger[i]]
        }
    }

    return nextLarger
}

func getResult(temperatures []int, nextLarger []int) []int {
    result := make([]int, len(temperatures))
    for i := 0; i < len(temperatures); i++ {
        if nextLarger[i] == len(temperatures) {
            result[i] = 0
        } else {
            result[i] = nextLarger[i] - i
        }
    }

    return result
}
