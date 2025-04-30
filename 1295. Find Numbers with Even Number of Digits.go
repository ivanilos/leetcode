func findNumbers(nums []int) int {
    result := 0

    for _, num := range nums {
        str := strconv.Itoa(num)
        if len(str) % 2 == 0 {
            result++
        }
    }
    return result
}
