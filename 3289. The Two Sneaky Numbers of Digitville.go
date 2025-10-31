func getSneakyNumbers(nums []int) []int {
    result := []int{}
    found := make([]int, len(nums) - 2)

    for _, num := range nums {
        found[num]++

        if found[num] == 2 {
            result = append(result, num)
        }
    }

    return result
}
