func specialArray(nums []int) int {
    maxVal := 1000
    counter := make([]int, maxVal + 1)

    for _, num := range(nums) {
        counter[num]++
    }

    result := -1
    bigger := 0
    for i := maxVal; i >= 0; i-- {
        bigger += counter[i]
        if bigger == i {
            result = bigger
            break
        }
    }

    return result
}
