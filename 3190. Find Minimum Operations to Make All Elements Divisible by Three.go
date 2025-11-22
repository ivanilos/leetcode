func minimumOperations(nums []int) int {
    result := 0

    for _, num := range nums {
        remainder := num % 3
        result += min(remainder, 3 - remainder)
    }

    return result
}
