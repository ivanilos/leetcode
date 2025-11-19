func findFinalValue(nums []int, original int) int {
    numMap := map[int]bool{}

    for _, num := range nums {
        numMap[num] = true
    }

    result := original
    ok, _ := numMap[result]
    for ok {
        result *= 2

        ok, _ = numMap[result]
    }

    return result
}
